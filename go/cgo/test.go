package main

/*
#include <stdlib.h>
// 注释必须紧挨 import
*/
import (
	"C" // 必须单独引用
)

import (
	"fmt"
	"unsafe"
)

func main() {
	cs := C.CString("my Cgo wangcheng")

	fmt.Printf("%T\n", cs)
	fmt.Println(string(*cs))

	C.free(unsafe.Pointer(cs))

}
