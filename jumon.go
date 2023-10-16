package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/atotto/clipboard"
	"golang.org/x/term"
)

var encodeFilename string
var outputFilename string
var useClipboard bool

func init() {
	// testing.Init()
	flag.StringVar(&encodeFilename, "e", "", "エンコードするファイルを指定する")
	flag.StringVar(&outputFilename, "o", "", "デコードのアウトプットファイル名を指定する")
	flag.BoolVar(&useClipboard, "c", false, "エンコード結果をクリップボードに出力する")
	flag.Parse()
}

func main() {

	reader := os.Stdin

	if term.IsTerminal(int(reader.Fd())) && encodeFilename == "" {
		flag.PrintDefaults()
		return
	}

	if !term.IsTerminal(int(reader.Fd())) {
		byte, err := io.ReadAll(reader)

		if err != nil {
			fmt.Println(err)
			return
		}

		if len(byte) > 0 {
			data, err := decode(string(byte))

			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			if outputFilename == "" {
				outputFilename = "output.bin"
			}

			err = os.WriteFile(outputFilename, data, 0644)

			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		return
	}

	if encodeFilename != "" {

		if !fileExists(encodeFilename) {
			fmt.Printf("%s が利用できません\n", encodeFilename)
			return
		}

		data, err := encode(encodeFilename)

		if err != nil {
			fmt.Println(err)
			return
		}

		if useClipboard {
			clipboard.WriteAll(data)
		} else {
			fmt.Println(data)
		}

		return
	}
}

func encode(filename string) (string, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(data)

	return encoded, nil
}

func decode(str string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(str)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
