#include <julia.h>

void init() {
	jl_init("/Users/david_neumann/builds/julia/usr/lib");
	JL_SET_STACK_BASE;
}
