package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

// import "fmt"

func Print(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs, (*C.FILE)(C.stdout))
}

func main() {
	Print("abc")
	print("abc")
}
