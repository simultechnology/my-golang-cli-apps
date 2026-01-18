# My Golang CLI Apps

Go言語の練習用CLIアプリケーション集

## プロジェクト一覧

### english-words-viewer

英単語データファイルを読み込んで表示するCLIツール

**機能:**
- テキストファイルから英単語と日付を解析
- 正規表現を使った空白区切りのデータ分割
- 単語リストの先頭10件を表示

**使い方:**

```bash
cd english-words-viewer
go run main.go
```

### pointer-basics

Go言語のポインタ、スライス、配列の挙動を学ぶための実験コード集
- 値渡し vs ポインタ渡し
- スライスと配列の内部構造の違い
- メソッドレシーバの使い分け

**使い方:**
```bash
cd pointer-basics/01-basics
# 各ファイルを個別に実行
go run main.go
```

## 必要環境

- Go 1.x以上

## ライセンス

練習用プロジェクト
