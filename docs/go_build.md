# LuneScript から変換した go のコードをビルドする

`lnsc` で go に変換したコードを、単体の go のプログラムとしてビルドする手順。

## 1. 必要なファイル

| ファイル | 作り方 |
|---|---|
| 変換した go のコード (`xxx.go`) | `lnsc xxx.lns SAVE -langGo --main xxx` |
| `main.go` | `lnsc xxx mkmain main.go` |
| `go.mod` | `lnsc xxx mkmain main.go --gomod <参照先>` で `main.go` と一緒に出力できる (既にある場合は上書きしない) |

変換したコードは、LuneScript のランタイム (`github.com/ifritJP/LuneScript/src/lune/base/runtime_go`) を import する。
ランタイムのソースや `lune` ディレクトリを手元にコピーする必要は無い。

## 2. 手順

```sh
lnsc hello.lns SAVE -langGo --main hello --app myapp
lnsc hello mkmain main.go --gomod /path/to/LuneScript/src --app myapp
go mod tidy
go build -o hello .
```

- `--app` は go のモジュール名になる。`main.go` が別ディレクトリのパッケージを import する場合も、この名前を使う。
  省略すると lnsc の既定の名前 (`lnsc`) になるので、指定すること。
- `go mod tidy` で、ランタイムが依存するモジュール (`gopher-lua` など) を取得する。

## 3. `--gomod` の参照先

| 指定 | 出力する go.mod | 用途 |
|---|---|---|
| LuneScript のソースの `src` ディレクトリ | `require github.com/ifritJP/LuneScript/src v0.0.0` と、そのディレクトリへの `replace` | 使っている lnsc と同じ版のランタイムでビルドする。**推奨** |
| バージョン (`master` など) | `require github.com/ifritJP/LuneScript/src master` | GitHub からランタイムを取得する (`go mod tidy` で疑似バージョンに解決される) |

変換したコードは、変換した lnsc と同じ版のランタイムを前提とする。
バージョンを指定する場合は、使っている lnsc と同じ版のランタイムが GitHub にあることを確認すること
(2026-10-17 時点で、`master` は 2024-01-06 の版に解決される)。

なお、go のモジュールは `src` ディレクトリにあるので、go のバージョンとして使うには `src/v1.6.1` のような形のタグが必要になる
(リポジトリ直下の `v1.6.0` などのタグは、go のモジュールのバージョンとしては使えない)。

## 4. cgo と Lua

| ビルド | 必要なもの | 備考 |
|---|---|---|
| 通常 (`go build`) | cgo と C の Lua 5.3 (`pkg-config` で `lua-5.3` が見つかること) | Lua 5.3 と同じ動作になる |
| `go build -tags gopherlua` | 不要 (`CGO_ENABLED=0` でもビルドできる) | Lua の部分を gopher-lua (Lua 5.1 相当) で動かす。実数の文字列化などが Lua 5.1 の形式になるので、Lua 5.3 と同じ出力にするには、変換時に `--int2str need0` を指定する |

Lua の機能 (`Luaval`、`_load`、`__luago` など) を使わないプログラムでも、ランタイムは Lua の VM を持つので、どちらかが必要になる。
