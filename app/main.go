/*
Copyright © 2025 Masaki Noguchi <mnologics@gmail.com>
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	/* 	問題
	Dog API からランダムに画像 URL を取得し、その URL を出力するサブコマンドrandomを実装しましょう。
	API のエンドポイントはGET https://dog.ceo/api/breeds/image/randomです。
	*/
	url := "https://dog.ceo/api/breeds/image/random"
	massage := fetchUrl(url)
	fmt.Println(massage)

}
