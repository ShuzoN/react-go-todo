### 開発環境用のgo

localでgoが必要です(editor補完用)

```sh $ make install/local
```


```sh
// host
$ export GOPATH=~/go >> ~/.zshrc
$ export GOROOT=/usr/local/go/bin/go >> ~/.zshrc
```

### 開発環境のserver

サーバの立ち上げ

```sh
$ make server
```

調子が悪い時は以下

```sh
$ make server/down server
```

サービスごとのrebuildは以下

```sh
$ make web/rebuild
```

```sh
$ mysql/rebuild mysql/setup
```

### ports / hosts


#### serviceが使うports

```sh
$ cat Makefile | head -n 5
```

#### hosts

mac/linux(docker host)の `/etc/hosts` に以下をたすこと

```
# headphonista
127.0.0.1 headphonista.mu.dev
```

serverの挙動は [http://localhost](https://localhost) へgo
go監視ツール realizeの管理画面は [http://localhost:3000](https://localhost:3000)

### vscodeユーザの人向け

`./vscode/settings.json` をvscodeの `settings.json` に追記する
