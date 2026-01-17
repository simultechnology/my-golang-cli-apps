package main

import "fmt"

// 値渡し: データがコピーされる
func doubleValue(val int) {
	val = val * 2
	fmt.Printf("  [関数内] doubleValue の中で2倍にしました: %d\n", val)
}

// ポインタ渡し: 住所を受け取る
func doublePointer(ptr *int) {
	// *ptr で「その住所の中身」にアクセスして書き換え
	*ptr = *ptr * 2
	fmt.Printf("  [関数内] doublePointer の中で2倍にしました: %d\n", *ptr)
}

func main() {
	println("--- 1. 値渡し (Pass by Value) の実験 ---")
	score := 100
	fmt.Printf("呼び出し前: %d\n", score)

	doubleValue(score) // scoreの「コピー」が渡される

	fmt.Printf("呼び出し後: %d (変わっていません！)\n", score)

	println("\n--- 2. ポインタ渡し (Pass by Pointer) の実験 ---")
	fmt.Printf("呼び出し前: %d\n", score)

	doublePointer(&score) // scoreの「住所」を渡す

	fmt.Printf("呼び出し後: %d (変わりました！)\n", score)
}
