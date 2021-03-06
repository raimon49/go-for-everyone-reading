# ch06

## 6.2 testingパッケージ入門

```bash
$ make test
```

## 6.3 ベンチマーク入門

```bash
$ make bench
```

## 6.4 テストの実践的なテクニック

### reflect.DeepEqualを使う

```bash
$ make reflect
```

### Race Detectorを使って競合状態を検出する

```bash
# Data Raceが検出される
$ make race

# Data Raceが解消され検出されない
$ make fixed-data-race
```

### TestMainによるテストの制御

```bash
$ make setup
```

### Build Constraintsを利用したテスト対象の指定

```bash
$ make integration
```

### テストにおける変数または手続きの書き換え

```bash
$ make gists
```

### インターフェースを使ったモック

```bash
$ make gists_interface
```

### net/http/httptestパッケージを使ったサーバーのテスト

```bash
$ make server_test
```

### テストカバレッジの出力

```bash
$ make coverprofile
```
