package message

/*
GetHelloMessage メッセージを取得する
  args name:名前
  return 挨拶文
*/
func GetHelloMessage(name string) string {
  // こんな感じで返り値にはセミコロンは不必要
  // Rustに似てるね
  return "Hello " + name
}