// Go provides built-in support for base64 encoding/decoding
package main

import b64 "encoding/base64"
import "fmt"

func main() {

	data := "abc123!?$*$()'-=@~"

	// Go supports both standard and URL-compatible base64
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
