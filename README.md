# s-cool-backend

## dockerコマンド
### コンテナの起動
```sh
$ docker-compose up -d
```

### コンテナの停止
```sh
$ docker-compose down
```
### キャッシュクリア
```sh
$ docker-compose down -v
$ docker-compose build --no-cache
$ docker-compose up -
```

## Commit/Push について

### - commit接頭辞
```
fix　　　：　バグ修正
add　　　：　新規（ファイル）機能追加
update　：　機能修正（バグではない）
remove　：　削除（ファイル）
```

### commitメッセージの例
```
add: ボタンのコンポーネントを作成(#100)
fix: ボタンの表示バグの修正(#200)

接頭辞: 変更内容の概要(#issue番号)
```

## Branch/PR について

### - 常設ブランチ
```
develop
- 本番環境 PR用branchはここから
```
### - 新規のブランチ
```
feature/#100_feature_button
fix/#200_fix_button
hotfix/fix_button_bug

命名規則
1.接頭辞は以下の通りとする.
  機能追加: feature
  修正: fix
  バグの修正: hotfix
2.issueが存在する場合は接頭辞に続けて #100 のようにissue番号を含める.(複数人での開発の際は極力issueを立てた後branchを立てるようにする).
3.最後に要約を記載 
```
