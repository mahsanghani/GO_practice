package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)

	fmt.Println(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()
	fmt.Println(uEnc)
	fmt.Println(string(uDec))
}