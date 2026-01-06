package main

/*
#cgo darwin LDFLAGS: ${SRCDIR}/../rust/target/release/librust2go_sample.a
#include <stdint.h>

extern int32_t add(int32_t a, int32_t b);
*/
import "C"

func rustAdd(a, b int32) int32 {
	return int32(C.add(C.int32_t(a), C.int32_t(b)))
}
