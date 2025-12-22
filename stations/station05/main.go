package main

import (
	"fmt"
	"os"
)

func main() {
	sampleFile := "./sample.txt"
	CheckFileExist(sampleFile)

	fmt.Println() // 改行のため

	notExistFile := "./abc.txt"
	CheckFileExist(notExistFile)

	fmt.Println() // 改行のため
}

// CheckFileExist は、指定されたパスのファイルが存在するかどうかを確認し、結果を出力します。
func CheckFileExist(path string) {
	// ここにファイルの存在を確認するコードを実装してください
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Print("ファイルが見つかりませんでした")
	} else {
		fmt.Print("ファイルが見つかりました")
	}
}
