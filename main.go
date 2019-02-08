package main

// こんな感じでまとめてインポートすることも可能
import (
  "fmt"
  "../go_review/message"
)

// main関数があるものが最初によばれるよ
// `go run`で呼ぶファイルはこのmain関数がないと許されないよ
// そんなことをしたらセグ木の下に埋めてもらって構わないよ!
func main() {
  fmt.Println(message.GetHelloMessage("hoge"))
}