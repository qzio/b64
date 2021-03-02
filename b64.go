package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var decode bool
var nostrip bool

func main() {
	flag.BoolVar(&decode, "d", false, "decode")
	flag.BoolVar(&nostrip, "nostrip", false, "nostrip")
	flag.Parse()

	if !isPipe() {
		fmt.Println("only piped input for now")
		os.Exit(1)
	}

	buf, err := io.ReadAll(os.Stdin)
	check(err)
	buf = bytes.TrimSuffix(buf, []byte("\n"))

	var res string
	if decode {
		res, err = decodeStr(string(buf))
		check(err)
	} else {
		res = encode(buf)
	}
	fmt.Printf("\n%s\n", res)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func decodeStr(str string) (string, error) {
	if !nostrip {
		if l := len(str) % 4; l > 0 {
			str += strings.Repeat("=", 4-l)
		}
	}
	b, err := base64.URLEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func encode(buf []byte) string {
	res := ""
	if !nostrip {
		res = strings.TrimRight(base64.URLEncoding.EncodeToString(buf), "=") + "\n"
	} else {
		res = base64.URLEncoding.EncodeToString(buf) + "\n"
	}
	return res
}

func isPipe() bool {
	info, err := os.Stdin.Stat()
	check(err)
	return (info.Mode() & os.ModeCharDevice) == 0
}
