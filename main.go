package main

/*
#include <stdio.h>

void printint(int v) {
    printf("int: %d\n", v);
}
*/
import "C"

func main() {
	v := 42
	C.printint(C.int(v))
}
