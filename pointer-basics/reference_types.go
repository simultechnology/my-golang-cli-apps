package main

import "fmt"

func modifySlice(s []string) {
	if len(s) > 0 {
		s[0] = "Modified" // スライスの中身はポインタ経由で参照されているので書き換わる
		fmt.Println("  [関数内] スライスの要素を書き換えました")
	}
}

func modifyMap(m map[string]int) {
	m["score"] = 999 // マップも参照型なので書き換わる
	fmt.Println("  [関数内] マップの値を書き換えました")
}

func modifyArray(a [3]string) {
	a[0] = "Modified"
	fmt.Println("  [関数内] 配列の要素を書き換えました")
}

func main() {
	println("--- 1. スライス (Reference Type) ---")
	// スライスは内部に「ポインタ」を持っているので、コピーしても参照先は同じ
	mySlice := []string{"Original", "Data"}
	fmt.Printf("呼び出し前: %v\n", mySlice)
	modifySlice(mySlice)
	fmt.Printf("呼び出し後: %v (値渡しなのに中身が変わる！)\n", mySlice)

	println("\n--- 2. 配列 (Value Type) ---")
	// 配列 (長さ固定) は値そのものなので、コピーされる
	myArray := [3]string{"Original", "Data", "End"}
	fmt.Printf("呼び出し前: %v\n", myArray)
	modifyArray(myArray)
	fmt.Printf("呼び出し後: %v (変わりません)\n", myArray)

	println("\n--- 3. マップ (Reference Type) ---")
	// マップも内部等はポインタのようなもの
	myMap := map[string]int{"score": 10}
	fmt.Printf("呼び出し前: %v\n", myMap)
	modifyMap(myMap)
	fmt.Printf("呼び出し後: %v (変わります！)\n", myMap)

	println("\n★ まとめ: スライスとマップは `*` を付けなくてもポインタのような挙動をします")
}
