# What's GitHub Flow?
 - ブランチ作成
 - commit & push
 - PullRequest & Review
 - Merge
の流れを取るもの

## ブランチ作成
`$ git branch {作業ブランチ名}`

ここで, {作業ブランチ名}は誰が見てもそのブランチの意図がわかるように命名すべき

## commit & push
 - commit
 - pushする前に, masterブランチで`$ git pull --rebase origin master`をして, 作業ブランチで`$ git rebase -i master`を実行すべき(リモートのmasterブランチの内容でのconflictを防ぐため)
 - push origin {作業ブランチ名}

## PullRequest & Review
originにマージできるかを確認

## Merge
OKならoriginにマージ(squashがいいよ)
