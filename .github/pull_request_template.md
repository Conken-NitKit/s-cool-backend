## 概要

## デバッグ項目

- [ ] (実装に不具合がないことを確認するために行った項目を記載してください 1)
- [ ] (実装に不具合がないことを確認するために行った項目を記載してください 2)

## スクリーンショット

| Before | After |
| :-: | :-: |
| | |

## Self-Checking points 🚨

レビューを依頼する前に必ず確認してほしいポイントです。

### 共通（命名)
- [ ] `data`, `info`, `value` などの汎用的で抽象度の高い変数名を使っていない [参考](https://neos21.net/blog/2020/01/28-01.html)
- [ ] 配列（コレクション）は末尾に `s` をつけて複数形にするか `xxxList` と命名することで配列であることを明確にしている [参考](https://teratail.com/questions/161176)
- [ ] マジックナンバーは極力存在せず、定数を用いて表現されてい [参考](https://twitter.com/program_shiba/status/1483378634975072260)
- [ ] パッと見で何をやってるかわからないような処理の結果を **説明変数** している [参考](https://wb-hp.com/blog/2020/11/09/explanatory-variable.html)
- [ ] 複雑な条件式の結果を **要約変数** を適用している [参考](https://twitter.com/hakuto00/status/1362608154840760320)

### 共通(処理系)
- [ ] **早期リターン（ガード節）** を適用することで条件分岐の簡略化している [参考](https://zenn.dev/media_engine/articles/early_return)
- [ ] if文では「調査対象」を左側に、「比較対象」を右側に配置している [参考](https://twitter.com/yuuuma_11/status/1347374986160340992/photo/2)
- [ ] 特定の値のパターンによって分岐する場合は `is/else文` ではなく、`switch文` を使う [参考](https://blog.senseshare.jp/if-switch.html)

### 共通（コメント系）
- [ ] コメントは **説明変数** でも伝わらない場合にのみ適用している [参考](https://kitsune.blog/comment)
- [ ] コメントには適切なアノテーションコメントが記載されている [参考](https://qiita.com/taka-kawa/items/673716d77795c937d422)

### 共通(レイアウト系)
- [ ] インデントが整っており見やすいコードになっている
- [ ] 変数や定数が、なるべく使う箇所の近くで宣言されている [参考](https://twitter.com/program_shiba/status/1484360099799859200)

### 共通(その他)
- [ ] フォーマット差分などは PR 上に一言 `インラインコメント` を付けて、レビュワーが省エネできるように [参考](https://docs.github.com/ja/pull-requests/collaborating-with-pull-requests/reviewing-changes-in-pull-requests/commenting-on-a-pull-request)