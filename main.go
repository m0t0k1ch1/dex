package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strings"
)

const (
	hexPrefix = "0x"
)

func main() {
	os.Exit(run())
}

func run() int {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("no input")
		return 1
	}

	src := args[1]
	isHex := strings.HasPrefix(src, hexPrefix)

	var result string
	var err error

	switch isHex {
	case true:
		result, err = hex2Decimal(src)
	case false:
		result, err = decimal2Hex(src)
	}
	if err != nil {
		fmt.Println(err)
		return 1
	}

	fmt.Println(result)
	return 0
}

func hex2Decimal(src string) (string, error) {
	i, ok := new(big.Int).SetString(strings.TrimPrefix(src, hexPrefix), 16)
	if !ok {
		return "", fmt.Errorf("hex -> decimal : can not parse")
	}

	return i.String(), nil
}

func decimal2Hex(src string) (string, error) {
	i, ok := new(big.Int).SetString(src, 10)
	if !ok {
		return "", fmt.Errorf("decimal -> hex : can not parse")
	}

	return hexPrefix + hex.EncodeToString(i.Bytes()), nil
}
