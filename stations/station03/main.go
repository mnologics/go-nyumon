package main

import "fmt"

func main() {
	//books := []Book{} // 値は自由に定義してください
	books := []Book{
		{
			Title:      "Go言語入門",
			Author:     "山田太郎",
			Price:      2800,
			Categories: []string{"プログラミング", "Go"},
		},
		{
			Title:      "Pythonスタートブック",
			Author:     "鈴木一郎",
			Price:      2500,
			Categories: []string{"プログラミング", "Python"},
		},
		{
			Title:      "JavaScript本格入門",
			Author:     "田中花子",
			Price:      3000,
			Categories: []string{"プログラミング", "JavaScript"},
		},
	}	

	totalPrice := Total(books)
	fmt.Println("合計", totalPrice, "円")

	authorBooksMap := ToMap(books)
	authorName := "田中花子"
	fmt.Println("著者: ", authorName, ", タイトル: ", authorBooksMap[authorName].Title)
}

// 構造体 Book を定義 ---
type Book struct {
	Title      string
	Author     string
	Price      int
	Categories []string
}

// フィールド
// Title ... 文字列、本のタイトル
// Author ... 文字列、著者
// Price ... 整数、価格
// Categories ... 文字列のスライス、カテゴリ

// --------------------

// Bookの配列を引数とし、本の合計価格を戻り値とする関数 `Total` を実装する
func Total(books []Book) int {
	// books 配列の各要素の Price フィールドを合計して戻り値とする
	sum := 0;
	for _, book := range books {
		sum += book.Price
	}
	return sum
}
	
// 3. キーを「著者名 (Author)」、値を構造体 Book とするマップを戻り値とする関数 `ToMap` を実装する
func ToMap(books []Book) map[string]Book {
	// books 配列の各要素をマップに格納して戻り値とする
	authorBooksMap := make(map[string]Book)
	for _, book := range books {
		authorBooksMap[book.Author] = book
	}
	return authorBooksMap
}

// --- IGNORE ---
// func Total() int { return 0 }
// func ToMap() map[string]Book { return nil }
// --- IGNORE ---
