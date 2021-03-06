# ch03

## 3.2 バージョン管理

```bash
$ make version
```

## 3.3 効率的なI/O処理

### 出力のバッファリング

```bash
$ make buffering
```

### 複数の出力先に一度に書き込む

```bash
$ make multi-writer
$ echo -n foo | ./multi-writer/multi-writer
```

## 3.4 乱数を扱う

```bash
$ make rand
```

## 3.5 人間が扱いやすい形式の数値

```bash
$ make humanize
$ humanize/main humanize/main
humanize/main: 2.6 MB
```

## 3.6 Goから外部コマンドを実行する

```bash
$ make exec
$ cat Makefile | exec/exec
```

## 3.8 シグナルを扱う

```bash
# デフォルトのシグナル挙動
$ make no-signal

# シグナルを自前ハンドリング
$ make handle-signal

# 独自のシグナルを定義してハンドリング
$ make timer-signal
```

## 3.9 goroutineの停止

```bash
$ make stop-goroutine
```
