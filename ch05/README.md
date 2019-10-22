# ch05

## 5.1 動的な型の判別

```bash
$ make classical-type-assertion
```

## 5.2 reflectパッケージ

```bash
# reflect.Value
$ make reflect-value-of

# reflect.Type
$ make reflect-type
```

## 5.3 reflectの利用例 ハマらないためのレシピ集

```bash
# 値を動的に生成する
$ make dynamic-create
```

```bash
# 見える範囲・見えない範囲
$ make reflect-scope-pkg

# Setできる値
$ make reflect-can-set

# reflectでinterfaceを満たしているかどうかの確認
$ make reflect-type-implements

# 動的なselect文の構築
$ cd select-ch; ./select-ch
$ reflect-select-ch/reflect-select-ch reflect-select-ch/aaa.txt reflect-select-ch/bbb.txt reflect-select-ch/ccc.txt
```

## 5.4 reflectのパフォーマンスまとめ

```bash
# reflectと型アサーションの比較
$ make reflect-access-performance

# reflectによるソート
$ make reflect-sort-performance
```
