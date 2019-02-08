/*
package calc : calcuration package
*/
package calc

import (
  "errors"
)

/*
isEven 入力した値が偶数か否かを判定する
  args num:数値
  returns true:偶数/false:奇数
*/
func isEven(num int) bool {
  return (num % 2 == 0)
}

/*
 Calc 引数をもとに四則演算を行う
  args 
       left:左辺値
       right:右辺値
       operation:演算子
  returs
       answer:計算結果
       err:エラーinterface 
*/
func Calc(left, right int, operation string) (answer int, err error) {
  // やはり, アセンブリのadd-sub-mul-divがしっくりくるね
  if operation == "add" {
    return left + right, nil
  } else if operation == "sub" {
    return left - right, nil
  } else if operation == "mul" {
    return left * right, nil
  } else if operation == "div" {
    // +ゼロ割+は悪の手
    // Segmentation Faultになりますよ
    if right == 0 {
      return 0, errors.New("divided by zero")
    }
    return left / right, nil
  }
  // 異常な演算子ですよ(実装してないだけでは?)
  return 0, errors.New("invalid operator")
}