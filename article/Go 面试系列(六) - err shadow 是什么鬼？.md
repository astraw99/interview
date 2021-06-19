在日常工作中，我们经常使用 `err != nil` 来判断程序或函数是否报错，或者使用 `defer {recover = err}` 来判断是否有 `panic` 严重错误，但稍不注意，很容易掉进 `err shadow` 的陷阱。

## 1. 变量作用域

```go
package main

import "fmt"

func main() {
	x := 100
	func() {
		x := 200 // x will shadow outer x
		fmt.Println(x)
	}()

	fmt.Println(x)
}
```

输出如下：
```go
200
100
```

**结果分析：**
`x` 变量在 `func` 里面打印为 `200`，在外层打印为 `100`，这就是变量的作用域(`variable scope`)。`func` 里面的变量 `x` 是一个新变量，只不过与外层 `x` 重名了(`variable redeclaration`)，此时里层 `x` 的作用域仅限于 `func {} block`，而外层 `x` 的作用域则是 `main {} block`，此时里层变量 `x` 发生了 `variable shadowing`，外层 `x` 不受影响，依然是 `100`。

改一下写法：
```go
package main

import "fmt"

func main() {
	x := 100
	func() {
		x = 200 // x will override outer x
		fmt.Println(x)
	}()

	fmt.Println(x)
}
```

输出如下：
```go
200
200
```
此时，`func` 里面的变量 `x` 仅仅是覆盖了外层 `x`，并没有定义新的变量，所以内外层输出都是 `200`。

## 2. err shadow - 无名 error
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("func err1:", test1())
}

func test1() error {
	var err error

	defer func() {
		fmt.Println("defer err1:", err)
	}()

	if _, err := os.Open("xxx"); err != nil {
		return err
	}

	return nil
}
```
输出如下：
```go
defer err1: <nil>
func err1: open xxx: no such file or directory
```

**结果分析：**
`func test1` 首先定义了 `var err error` 变量，但下面的 `os.Open` 报错使用 `err :=` 被局部 `err shadow` 了，虽然显式使用了 `return err` 返回错误，但由于 `test1() error` 返回参数是无名的(`unnamed variable`)，导致 `defer` 中 `err` 获取不到被 `err shadow` 的错误 `err`，取的仍然是外层初始化 `var err error` 值，所以输出为 `err1: <nil>`。

只需要将第 `19` 行改一下，即可避免 `err shadow`：
```go
if _, err = os.Open("xxx"); err != nil {
		return err
	}
```
输出如下：
```go
defer err1: open xxx: no such file or directory
func err1: open xxx: no such file or directory
```

## 3. err shadow - 有名 error
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("func err2:", test2())
}

func test2() (err error) {

	defer func() {
		fmt.Println("defer err2:", err)
	}()

	if _, err := os.Open("xxx"); err != nil {
		return // return without err will compilation error
	}

	return
}
```

上面的 `test2` 运行会有编译报错，这是 `go compiler` 在编译时做了 `variable shadowing` 检查，发现有就直接编译报错。修改一下即可：

```go
func main() {
	fmt.Println("func err3:", test3())
}

func test3() (err error) {

	defer func() {
		fmt.Println("defer err3:", err)
	}()

	if _, err := os.Open("xxx"); err != nil {
		return err
	}

	return
}
```

输出如下：
```go
defer err3: open xxx: no such file or directory
func err3: open xxx: no such file or directory
```

## 4. 嵌套 err shadow
```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println("func err4:", test4())
}

func test4() (err error) {

	defer func() {
		fmt.Println("defer err4:", err)
	}()

	if _, err := os.Open("xxx"); err == nil {
		if err := json.Unmarshal([]byte("{}"), &struct{}{}); err == nil {
			fmt.Println("OK")
		}
	}

	return
}
```

输出如下：
```go
defer err4: <nil>
func err4: <nil>
```

**结果分析：**
`func test4()` 是一个有名返回 `err error`，则函数初始化时会 `var err error` 定义对应的有名变量(`named variable`)，但下面的 `os.Open` 或 `json.Unmarshal` 都使用了 `err :=` 重定义 `err` 变量，造成了 `err shadow`，因此在函数退出时，外层 `err` 依然是 `nil`，`defer` 获取也就是 `nil`。

改一下写法即可：
```go
func main() {
	fmt.Println("func err5:", test5())
}

func test5() (err error) {

	defer func() {
		fmt.Println("defer err5:", err)
	}()

	if _, err = os.Open("xxx"); err == nil {
		if err = json.Unmarshal([]byte("{}"), &struct{}{}); err == nil {
			fmt.Println("OK")
		}
	}

	return
}
```

输出如下：
```go
defer err5: open xxx: no such file or directory
func err5: open xxx: no such file or directory
```

## 5. 小结
本文通过几个实例，分析了在实际工作中很容易出现的 `err shadow` 问题，究其本质原因主要是变量作用域引起的，在官方文档中提到：`An identifier declared in a block may be redeclared in an inner block. While the identifier of the inner declaration is in scope, it denotes the entity declared by the inner declaration`.

另外，在函数返回值命名方面，我们需要考虑无名、有名参数的情况，在保证代码逻辑正确的情况下，建议使用工具 `go linter` 或 `go vet` 来检测编译器没检测到的 `variable shadowing`，避免踩到坑。