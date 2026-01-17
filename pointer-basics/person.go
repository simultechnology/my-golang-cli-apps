package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 値レシーバ: メソッドが呼ばれると struct のデータがコピーされる
// 内部で書き換えても、コピーが変わるだけで元データには影響しない
func (p Person) HaveBirthdayValue() {
	p.Age++
	fmt.Printf("  [メソッド内] %s は誕生日を迎えました (値レシーバ): %d歳\n", p.Name, p.Age)
}

// ポインタレシーバ: struct の住所を受け取る
// 元のデータを書き換えることができる
func (p *Person) HaveBirthdayPointer() {
	p.Age++
	fmt.Printf("  [メソッド内] %s は誕生日を迎えました (ポインタレシーバ): %d歳\n", p.Name, p.Age)
}

func main() {
	gameCharacter := Person{Name: "Mario", Age: 30}
	fmt.Printf("初期状態: %+v\n", gameCharacter)

	fmt.Println("\n--- 1. 値レシーバの実験 ---")
	gameCharacter.HaveBirthdayValue()
	fmt.Printf("呼び出し後: %+v (歳をとっていません！)\n", gameCharacter)

	fmt.Println("\n--- 2. ポインタレシーバの実験 ---")
	// Go言語の便利機能: ポインタレシーバでも、通常の変数から直接呼び出せます
	// コンパイラが自動的に (&gameCharacter).HaveBirthdayPointer() と解釈してくれます
	gameCharacter.HaveBirthdayPointer()
	fmt.Printf("呼び出し後: %+v (歳をとりました！)\n", gameCharacter)
}
