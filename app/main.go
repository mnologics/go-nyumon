/*
Copyright © 2025 Masaki Noguchi <mnologics@gmail.com>
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetchUrl(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	result := byteToJson(body)
	return result
}

func byteToJson(body []byte) string {
	var result map[string]interface{}
	err := json.Unmarshal(body, &result)
	if err != nil {
		return ""
	}
	return result["message"].(string)
}

func main() {
	// Station8

	// 引数を受け取る処理を追加
	/* オプション引数仕様
	引数名はimagesとする。
	例）dog-cli random --images 2
	短縮系: -i
	オプション引数が指定されていない場合やオプション引数のデフォルト値は1とする。
	*/
	args := os.Args
	images := 1
	if len(args) >= 4 && (args[2] == "--images" || args[2] == "-i") {
		fmt.Sscanf(args[3], "%d", &images)
	}

	for i := 0; i < images; i++ {
		/* 	問題
		Dog API からランダムに画像 URL を取得し、その URL を出力するサブコマンドrandomを実装しましょう。
		API のエンドポイントはGET https://dog.ceo/api/breeds/image/randomです。
		*/
		url := "https://dog.ceo/api/breeds/image/random"
		massage := fetchUrl(url)
		fmt.Println(massage)
	}

}
