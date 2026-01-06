package main

/*
#cgo darwin LDFLAGS: ${SRCDIR}/../rust/target/release/librust2go_sample.a
#include <stdint.h>

extern int32_t add(int32_t a, int32_t b);
*/
import "C"

import "fmt"

func main() {
	result := C.add(2, 3)
	fmt.Println("Result from Rust:", result)
}
