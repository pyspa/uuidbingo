# UUID Bingo

uuidを指定して、uuidでビンゴができる楽しいゲームです。忘年会の余興としてお使いください。


## 遊び方

サーバーを立ち上げて、 `http://localhost:8080/141468b0-b0b5-466b-b0a9-0cbd57317c4f` のようなURLをブラウザーで開くと、ビンゴシートを得ることができます。お手持ちのUUID生成器を使ってUUIDを生成し、UUIDが一致したときに、そのマスをチェックしてくだしさい。縦横斜め、いずれかのラインのすべてのマスがチェックされると、あなたの勝ちです。

ビンゴのマスで作られるUUIDは、[NewSha1関数](https://pkg.go.dev/github.com/google/uuid#NewSHA1) を使った、Version 5 UUIDです。これに合わせたUUID生成器を使用してください。


## サーバーの始め方

```shellscript
go build
```

して、

```shellscript
./uuidbingo
```

してください。8080ポートでサーバーが起動します。

