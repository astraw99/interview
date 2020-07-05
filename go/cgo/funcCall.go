package main

/*
#include <stdio.h>
#include <stdlib.h>

void out(char* in) {
  printf("%s\n", in);
}
*/
import "C"

import (
	"unsafe"
)

func main() {
	cstr := C.CString("go is cool")
	C.out(cstr) //ok
	C.free(unsafe.Pointer(cstr))
}
