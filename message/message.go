package message

/*
GetHelloMessage メッセージを取得する
  args name:名前
  return 挨拶文
*/
func GetHelloMessage(name string) string {
  return "Hello " + name
}