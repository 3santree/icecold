# Icecold

Obfuscate shellcode and output an executable for you.

This project is a POC of the last method explained in this [blog](https://www.bordergate.co.uk/shellcode-obfuscation/)

## Usage

Defaulty, it will generate a `a.exe` that pops up calc in `output` folder

```cmd
cd icecold
go run .
.\output\a.exe
```

To test your shellcode, replace the shellcode in `main.go`

```golang
func main() {
	shellcode, _ := hex.DecodeString("505152535657556A605A6863616C6354594883EC2865488B32488B7618488B761048AD488B30488B7E3003573C8B5C17288B741F204801FE8B541F240FB72C178D5202AD813C0757696E4575EF8B741F1C4801FE8B34AE4801F799FFD74883C4305D5F5E5B5A5958C3")
	// Put Your shellcode below and comment out the line above
	//shellcode := []byte{<Your Shellcode>}
```

## Detection

Using common windows/x64/meterpreter/reverse_tcp shellcode from msf, the executable we generated was detected by 3/71 vendors according to virustotal as of Auguest 2023.

![2023-08-12_13-05](https://github.com/3santree/icecold/assets/37735352/6cb7097e-8a19-4c79-8389-bc468c3eed61)

