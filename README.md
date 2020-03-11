wsoui: The OUI database for Go [![Build Status](https://travis-ci.org/hnw/wsoui.svg?branch=master)](https://travis-ci.org/hnw/wsoui)
=====================

[English version](./README-en.md)

## 説明

`wsoui` はMACアドレスからベンダー名を検索するGo製のライブラリです。

MACアドレスの先頭三桁はOUI (Organizationally Unique Identifier) と呼ばれ、IEEEがベンダーに切り売りしています。逆に言うと、MACアドレスの先頭三桁を見ればどこのベンダーが作ったNICか（もしくはBlueToothデバイスか）を知ることができます。

本ライブラリではOUIのデータをGoのmapとして持ち、検索用の関数から利用しています。

OUIの元データとしては、WireSharkが加工して配布している[Wireshark manufacturer database](https://gitlab.com/wireshark/wireshark/raw/master/manuf)を利用しています。

## 特徴

- 省バイナリサイズ・省メモリ
  - 本ライブラリは低スペックマシンでも動作するよう実装されています。実際、私はこのライブラリを使ったプログラムを64MBメモリ・32MBストレージのLinuxマシン上で動かしています。
- 最新のOUIデータが反映されている
  - 元データからの反映はTravis CIで週に1回自動的に行われるので、常に最新のOUIデータが利用できます。

## 利用方法

### `func LookUp(mac string) (string, error)`

この関数は引数としてMACアドレス `mac` を取り、対応するベンダー名の略称（`"Cisco"` のような表記、最大8文字）を返します。

`mac` は先頭3オクテット分のみを利用しますので、文字列全体がMACアドレスのフォーマットに従っている必要はありません。オクテットの区切り文字は `-` または `:` である必要があります。もしくは区切り文字無しでも構いません。

与えられた `mac` に対応するベンダー名が見つからなかった場合は空文字列と非nil値の`error`を返します。

## サンプル

``` go
package main

import (
	"fmt"
	"github.com/hnw/wsoui"
)

func main() {
	abbr, _ := wsoui.LookUp("cc-20-e8")
	fmt.Print(abbr) // Apple
}
```

## ライセンス

ouidata.go のみGPLとします。元ファイルの[Wireshark manufacturer database](https://code.wireshark.org/review/gitweb?p=wireshark.git;a=blob_plain;f=manuf)がGPLに従うと考えられるためです。

ouidata.go以外のファイルはMITとします。

したがって、本ライブラリをリンクした実行バイナリはGPLに従います。
