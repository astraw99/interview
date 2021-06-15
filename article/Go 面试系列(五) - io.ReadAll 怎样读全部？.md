在进行本地 `file` 文件内容读取，或进行 `HTTP` 网络接口通信的时候，我们经常使用 `io.ReadAll` 来读取远程接口返回的 `resp.Body`，但接口返回数据量有大有小，`io.ReadAll` 是怎样完成全部数据的读取的？

带着此疑问，让我们走近 `io.ReadAll` 源码一探究竟：

## 1. Demo 读取文件内容
```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 读取文件内容
	fileInfo, err := os.Open("./abc.go")
	if err != nil {
		panic(err)
	}

	contentBytes, err := io.ReadAll(fileInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(contentBytes))
}
```

此时读取的 `io stream` 大小并不知道，`io.ReadAll` 使用什么策略读取全部数据呢？滑动窗口？线性/指数递增读取？`Talk is cheap. Show me the code`.

## 2. io.ReadAll Code
`go1.16/src/io/io.go#L626`
```go
// ReadAll reads from r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Read
// as an error to be reported.
func ReadAll(r Reader) ([]byte, error) {
	b := make([]byte, 0, 512)
	for {
		if len(b) == cap(b) {
			// Add more capacity (let append pick how much).
			b = append(b, 0)[:len(b)]
		}
		//println(cap(b))
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == EOF {
				err = nil
			}
			return b, err
		}
	}
}
```
**源码解析：**
从上面源码可以看到，使用 `make` 先默认申请 `cap = 512` 的 `[]byte`，然后进入 `for` 循环迭代，直到数据全部读取完成。`for` 循环中，首先通过 `len(b) == cap(b)` 判断 `b` 的容量是否满了，如果已经满了，使用 `append(b, 0)` 追加一个元素，此时会发生什么呢？

我们知道，一个 `slice` 容量不够了需要扩容，但扩容机制是怎样的呢？继续 `Show me the code`.

## 3. slice 扩容机制
`go1.16/src/runtime/slice.go#L125`
```go
// growslice handles slice growth during append.
// It is passed the slice element type, the old slice, and the desired new minimum capacity,
// and it returns a new slice with at least that capacity, with the old data
// copied into it.
// The new slice's length is set to the old slice's length,
// NOT to the new requested capacity.
// This is for codegen convenience. The old slice's length is used immediately
// to calculate where to write new values during an append.
// TODO: When the old backend is gone, reconsider this decision.
// The SSA backend might prefer the new length or to return only ptr/cap and save stack space.
func growslice(et *_type, old slice, cap int) slice {
	...

	newcap := old.cap
	doublecap := newcap + newcap
	//println("newcap: ", newcap)
	//println("cap: ", cap)
	if cap > doublecap {
		newcap = cap
	} else {
		if old.cap < 1024 {
			newcap = doublecap
		} else {
			// Check 0 < newcap to detect overflow
			// and prevent an infinite loop.
			for 0 < newcap && newcap < cap {
				newcap += newcap / 4
			}
			// Set newcap to the requested cap when
			// the newcap calculation overflowed.
			if newcap <= 0 {
				newcap = cap
			}
		}
	}
...
}
```
**源码解析：**
从上面源码可以看到，`slice` 扩容算法为：
1). 当需要的容量(`cap`)超过原切片容量的两倍(`doublecap`)时，会使用需要的容量作为新容量(`newcap`)；
2). 当原切片容量 `< 1024` 时，新切片的容量(`newcap`)会直接翻倍(`doublecap`)；
3). 当原切片容量 `>= 1024` 时，会按原切片容量反复地增加 `1/4`，直到新容量(`newcap`)超过所需要的容量；

举例说明：
在上面 `io.ReadAll` 源码中，初始 `slice cap = 512`，后面扩容将会：
```go
512
1024(doublecap)
1280(1024 + 1024/4)
1600(1280 + 1280/4)
2000(1600 + 1600/4)
...
```

实际扩容 `cap` 是这样的吗？让我们验证一下：
```go
before newcap:  1024
-after newcap:  1024
before newcap:  1280
-after newcap:  1280
before newcap:  1600
-after newcap:  1792
before newcap:  2240
-after newcap:  2304
```
奇怪？发现 `after newcap` 并没有按照上面预想的值扩容，仔细挖代码，发现除了按照上面 `slice cap `扩容外，还对内存分配进行了“对齐”：

`go1.16/src/runtime/slice.go#L198`
```go
    println("before newcap: ", newcap)

	var overflow bool
	var lenmem, newlenmem, capmem uintptr
	// Specialize for common values of et.size.
	// For 1 we don't need any division/multiplication.
	// For sys.PtrSize, compiler will optimize division/multiplication into a shift by a constant.
	// For powers of 2, use a variable shift.
	switch {
	...
	case isPowerOfTwo(et.size):
		var shift uintptr
		if sys.PtrSize == 8 {
			// Mask shift for better code generation.
			shift = uintptr(sys.Ctz64(uint64(et.size))) & 63
		} else {
			shift = uintptr(sys.Ctz32(uint32(et.size))) & 31
		}
		lenmem = uintptr(old.len) << shift
		newlenmem = uintptr(cap) << shift
		capmem = roundupsize(uintptr(newcap) << shift) // 进入到内存块(memory block)分配
		overflow = uintptr(newcap) > (maxAlloc >> shift)
		newcap = int(capmem >> shift)
	...
	}

	println("after newcap: ", newcap)
```


进入到内存块(`memory block`)分配：
`go1.16/src/runtime/msize.go#L13`
```go
// Returns size of the memory block that mallocgc will allocate if you ask for the size.
func roundupsize(size uintptr) uintptr {
	if size < _MaxSmallSize {
		if size <= smallSizeMax-8 {
			return uintptr(class_to_size[size_to_class8[divRoundUp(size, smallSizeDiv)]])
		} else {
			return uintptr(class_to_size[size_to_class128[divRoundUp(size-smallSizeMax, largeSizeDiv)]])
		}
	}
	if size+_PageSize < size {
		return size
	}
	return alignUp(size, _PageSize)
}
```

获取 `spanClass` 对应的 `size`：
`go1.16/src/runtime/sizeclasses.go#L84`
```go
const (
	_NumSizeClasses = 68
)

var class_to_size = [_NumSizeClasses]uint16{0, 8, 16, 24, 32, 48, 64, 80, 96, 112, 128, 
144, 160, 176, 192, 208, 224, 240, 256, 288, 320, 352, 384, 416, 448, 480, 512, 576, 640, 
704, 768, 896, 1024, 1152, 1280, 1408, 1536, 1792, 2048, 2304, 2688, 3072, 3200, 3456, 
4096, 4864, 5376, 6144, 6528, 6784, 6912, 8192, 9472, 9728, 10240, 10880, 12288, 13568, 
14336, 16384, 18432, 19072, 20480, 21760, 24576, 27264, 28672, 32768}
```

从上面 `68` 类 `spanClass` 可以看到，我们想要分配 `1600` 被对齐到了 `1792`，`2240` 被对齐到了 `2304`，符合下面的验证结果：
```go
before newcap:  1024
-after newcap:  1024
before newcap:  1280
-after newcap:  1280
before newcap:  1600
-after newcap:  1792
before newcap:  2240
-after newcap:  2304
```

## 4. 小结
从上面的源码分析可以看到，`io.ReadAll` 通过使用 `slice append` 自动扩容 + 内存对齐机制，使用增加的容量来实现对 `io stream` 的全部读取。`slice append` 扩容算法为：
1). 当需要的容量(`cap`)超过原切片容量的两倍(`doublecap`)时，会使用需要的容量作为新容量(`newcap`)；
2). 当原切片容量 `< 1024` 时，新切片的容量(`newcap`)会直接翻倍(`doublecap`)；
3). 当原切片容量 `>= 1024` 时，会按原切片容量反复地增加 `1/4`，直到新容量(`newcap`)超过所需要的容量；

后面将会有更多系列文章，解读内存分配、`GC` 机制、`GPM` 调度、面试系列、`K8s` 系列、`etcd` 系列等，如有错误恳请指正。最后，祝大家端午节快乐～
