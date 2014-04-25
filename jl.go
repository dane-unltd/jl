package jl

/*
#cgo CFLAGS: -Wno-strict-aliasing -fno-omit-frame-pointer -I/Users/david_neumann/builds/julia/src -I/Users/david_neumann/builds/julia/src/support
#cgo LDFLAGS: -L/Users/david_neumann/builds/julia/usr/lib -ljulia
#include <stdlib.h>
#include "jl.h"
*/
import "C"

func init(path string) {
	C.init()
}
