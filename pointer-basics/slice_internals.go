package main

import "fmt"

func appendElement(s []int) {
	// ここで s は main関数の slice の「コピー」を受け取っています。
	// (ポインタ、長さ、容量) の3点セットがコピーされています。

	// append で要素を追加します
	s = append(s, 999)

	fmt.Printf("  [関数内] appendしました。長さ: %d, 容量: %d, 先頭アドレス: %p\n", len(s), cap(s), s)
}

func main() {
	fmt.Println("--- 1. スライスの正体 (Slice Header) ---")
	// 長さ2, 容量2のスライス
	slice := []int{10, 20}

	fmt.Printf("変数 slice のアドレス (&slice): %p (3点セットがある場所)\n", &slice)
	fmt.Printf("中身の配列の先頭アドレス (slice ): %p (実際のデータがある場所)\n", slice)

	fmt.Println("\n--- 2. append の落とし穴 ---")
	fmt.Printf("呼び出し前: 長さ: %d, 容量: %d, 中身: %v\n", len(slice), cap(slice), slice)

	// 関数にスライスを渡して append してもらう
	appendElement(slice)

	fmt.Printf("呼び出し後: 長さ: %d, 容量: %d, 中身: %v\n", len(slice), cap(slice), slice)

	fmt.Println("\n★ 解説:")
	fmt.Println("  呼び出し後も中身が変わっていない！？")
	fmt.Println("  理由: 関数内で append して「長さ」や「容量」、そして場合によっては「ポインタ」まで変わりましたが、")
	fmt.Println("       呼び出し元の slice 変数 (3点セット) は更新されなかったからです。")
	fmt.Println("  正解: s = append(s, val) のように、戻り値を受け取る必要があります。")
}
