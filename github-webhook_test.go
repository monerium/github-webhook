package main

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func TestHex(t *testing.T) {
	fmt.Println("Simple")
	sha := "a4e652fe34f2be526dd7ae298dfc2807dfef34c5"
	fmt.Println(sha)
	fmt.Println(len(sha))
	fmt.Println(hex.DecodeString(sha))
	fmt.Println("Hello")
	sha = strings.TrimPrefix("sha1=a4e652fe34f2be526dd7ae298dfc2807dfef34c5", "sha1=")
	fmt.Println(sha)
	fmt.Println(len(sha))
	fmt.Println(hex.DecodeString(sha))
}
