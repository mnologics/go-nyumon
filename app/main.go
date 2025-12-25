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
	msg := result["message"]
	list := msg.([]interface{})
	var output string
	for i, v := range list {
		output += v.(string) + "\n"
		if i >= 9 {
			// break
		}
	}
	output = output[:len(output)-1] // 最後の改行を削除
	return output
}

func byteToJson(body []byte) map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil
	}
	return result
}

func main() { // Station10
	/*
		問題
		Station9 で犬種を指定して画像を取得できるようにしました。 この Station では犬種一覧を取得できるようにしましょう。

		API のエンドポイントはGET https://dog.ceo/api/breeds/list です。 ※全件取得すると多すぎるので、最初の 10件だけ出力してください。

		仕様
		コマンドはgo run ./app/main.go breed-listとする。
		1種類ごとに改行して出力する。
		例）

		$ go run ./app/main.go breed-list

		affenpinscher
		african
		akita
		....
	*/

	args := os.Args
	if len(args) < 2 || args[1] != "breed-list" {
		fmt.Println("Usage: go run ./app/main.go breed-list")
		os.Exit(1)
		return
	}
	url := "https://dog.ceo/api/breeds/list"
	massage := fetchUrl(url)
	if massage == "" {
		fmt.Println("犬種一覧の取得に失敗しました")
		os.Exit(1)
		return
	}
	fmt.Println(massage)
}
