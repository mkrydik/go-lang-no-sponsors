# go-lang-no-sponsors

Go Lang のスポンサーの提供でお送りします (ダジャレ)


## 使い方

このプロジェクトは以下の環境で動作確認しました。

- macOS Monterey v12.4 : Go v1.19 : Terminal.app

次のようにインストールすることで、環境変数 `GOBIN` もしくは `GOPATH` で指定されたディレクトリ (通常は `${HOME}/go/bin/`) に実行ファイルが生成され、コマンドが実行できるようになります。

```bash
# コマンドをインストールして実行する
$ git clone https://github.com/mkrydik/go-lang-no-sponsors.git && cd "$(basename "$_" .git)"
$ go install
$ go-lang-no-sponsors  # ${HOME}/go/bin/go-lang-no-sponsors

# ソースを直接実行する
$ go run ./main.go

# バイナリをビルドして実行する
$ go build
$ ./go-lang-no-sponsors
```

- 実行例 (ターミナルウィンドウの幅・高さによって罫線や空行の表示が変わります)

```bash
$ go-lang-no-sponsors
+-------------------------------------------------+
|                                                 |
|                                                 |
|                                                 |
|                     提   供                     |
|                                                 |
|                                                 |
|                                                 |
|                     mkrydik                     |
|                                                 |
|                                                 |
|                                                 |
+-------------------------------------------------+
```

ターミナルに提供スーパーが出力されます。


## コーディング関連コマンド

```bash
# フォーマット
$ go fmt

# テスト
$ go test
```
