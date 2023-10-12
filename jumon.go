package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: base64 [-e ファイル] [-d 文字列]")

		return
	}

	switch args[0] {
	case "-e":
		encode(args[1])
		break
	case "-d":
		decode(args[1])
		break
	default:
		fmt.Println("Invalid argument")
		return
	}
}

func encode(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	encoded := base64.StdEncoding.EncodeToString(data)

	fmt.Println(encoded)

	return encoded
}

func decode(str string) []byte {
	data, err := base64.StdEncoding.DecodeString(str)

	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	// fmt.Println(string(data))

	err = os.WriteFile("output.bin", data, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return data
}
