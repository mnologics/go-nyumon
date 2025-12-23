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
	if result["status"] != "success" {
		return ""
	}
	return result["message"].(string)
}

func byteToJson(body []byte) map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil
	}
	return result
}

func main() { // Station9
	/*
			問題
		Station8 でランダムに画像を取得するサブコマンドを実装しました。 この Station で犬種を指定して画像を取得できるようにしましょう。

		仕様
		コマンドはdog-cli breed {犬種名}とする。
		犬種が指定されていない場合はエラーとする。
		API から取得できない犬種や無効な文字列の場合はエラーとする
		エラーメッセージ
		犬種が指定されていない場合はError: 犬種が指定されていませんと出力する。 例）
		$ go run ./app/main.go breed

		Error: 犬種が指定されていません
		API から取得できない犬種や無効な文字列の場合はError: 無効な犬種ですと出力する。 例）
		$ go run ./app/main.go breed invalid

		Error: 無効な犬種です
	*/
	// API のエンドポイントはGET https://dog.ceo/api/breed/{犬種名}/images/randomです。

	args := os.Args
	if len(args) < 3 || args[1] != "breed" {
		fmt.Println("Error: 犬種が指定されていません")
		os.Exit(1)
		return
	}
	breed := args[2]
	url := fmt.Sprintf("https://dog.ceo/api/breed/%s/images/random", breed)
	massage := fetchUrl(url)
	if massage == "" {
		fmt.Println("Error: 無効な犬種です")
		os.Exit(1)
		return
	}
	fmt.Println(massage)
}
