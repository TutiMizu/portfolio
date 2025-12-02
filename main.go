package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. ルート設定: "/" というURLに来たら helloHandler を動かす
	http.HandleFunc("/", helloHandler)

	fmt.Println("サーバーが起動しました！")
	fmt.Println("ブラウザで http://localhost:8080 にアクセスしてください")

	// 2. サーバー起動: ポート8080番で待ち受ける
	// (プログラムはここで止まって、ずっとアクセスを待ち続けます)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
	}
}

// ブラウザに返事をする係（ハンドラー関数）
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// w (Writer) に文字を書き込むと、それがブラウザに届く
	fmt.Fprintf(w, "Hello, Backend World! from Docker")
}
