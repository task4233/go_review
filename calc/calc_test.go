package calc

import (
  "testing"
)

// isEven()のテスト関数
func Test_isEven(t *testing.T) {
  // テストケースの構造体
  type testCase struct {
    num    int  // 入力値
    expect bool // 期待値:偶数か否か
  }

  // テストリストの生成
  // Arrayで宣言してどんどんいれてくよー
  // javascriptの
  // const hoge = []
  // ;['hoge'].foreach(...)
  // みたいなかんじですね!
  testlist := []testCase{
    // 2 -> true(Even)
    testCase{ num: 2, expect: true },
    // 3 -> false(Odd)
    testCase{ num: 3, expect: false },
  }

  // testlistのループ
  for _, data := range testlist {
    // 関数を呼び出して結果と期待値を比較するよ
    result := isEven(data.num)
    if result != data.expect {
      // 結果と期待値が同じでない=>テスト失敗, 残念でした!
      t.Errorf("Test Failed! expect: %v, result: %v", data.expect, result)
    }   
  }
}

// Calcのテスト関数
func Test_Calc(t *testing.T) {
  // テストケースの構造体
  type testCase struct {
    left               int    // 左辺値
    right              int    // 右辺値
    operation          string // 演算子
    expectAnswer       int    // 期待値:演算結果
    expectErrEqNil     bool   // 期待値:err==nil? この変数名すごい良いと思いませんか? 私は思いました
  }

  // テストリストの生成
  testlist := []testCase{
    // ====
    // True
    // ====
    // 10 + 5 = 15
    testCase{ left: 10, right: 5, operation: "add", expectAnswer: 10+5, expectErrEqNil: true },
    // 10 - 5 = 5
    testCase{ left: 10, right: 5, operation: "sub", expectAnswer: 10-5, expectErrEqNil: true },
    // 10 * 5 = 50
    testCase{ left: 10, right: 5, operation: "mul", expectAnswer: 10*5, expectErrEqNil: true },
    // 10 / 5 = 2
    testCase{ left: 10, right: 5, operation: "div", expectAnswer: 10/5, expectErrEqNil: true },

    // ====
    // Err
    // ====
    // 演算子が異常=> invalid operator
    // よく考えたら?は三項演算子だよね
    testCase{ left: 10, right: 5, operation: "?", expectAnswer: 0, expectErrEqNil: false },
    // 10/0 => divided by zero
    testCase{ left: 10, right: 0, operation: "/", expectAnswer: 0, expectErrEqNil: false },
  }

  // テストケースを回す
  for _, data := range testlist {
    
    answer, err := Calc(data.left, data.right, data.operation)
    if (err == nil) != data.expectErrEqNil {
      // エラーの判定結果と期待値が異なる=>test failed...
      t.Errorf("Test failed! err: %s", err)
    } else if answer != data.expectAnswer {
      // 演算結果と期待値が異なる=>test failed!
      t.Errorf("Test failed(Result)! %d %s %d = %d", data.left, data.operation, data.right, answer)
    }
  }
}