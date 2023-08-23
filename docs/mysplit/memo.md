# memo
コマンドラインツールの作成をする。要件は以下のとおり。
ターミナルで使えるsplitを自前で実装する。
* go公式のパッケージ、自作のパッケージのみ使用可能
* オプションは'-l, -n, -b'が必須

goのカルチャーからテスト駆動でやりたいけど無理ぽそう。

## コマンドライン引数
Goのデフォルトで用意されているflagパッケージを使用する。
(https://pkg.go.dev/flag)</br>
難点はオプションの競合をどうするか。
現時点でfunc (*FlagSet) xxxFunc()を使って非nilエラーが返せるからここでなんとかできないか。

どうしてもダメそうならos.Argsで入力とってきて、自前でパースする。

### -l
lineの数を指定する。デフォルトは1000行。

### -n
ファイルの分割数を指定する。デフォルトは2。

### -b
ファイルの分割サイズを指定する。デフォルトは1MB。
バイトサイズの扱いはこの辺でいけそう。
~~https://pkg.go.dev/github.com/inhies/go-bytesize~~（デフォルトじゃなかった）

### -a
サフィックスを指定する。デフォルトは'xaa'。
Xは第2引数で指定した文字列がつく。デフォルトは'x'。
intの値でaがつく数が変わる。
```
Ex) default -> xaa, xab, xac, ...
    5  ->  xaaaaa, xaaaab, xaaaac, ...
```
## ファイルの入出力
ioパッケージを使用する。
エラー時にファイルを閉じるかどうかドキュメント見てじっそうすること。
## テストについて
テストはGoのデフォルトで用意されているtestingパッケージを使用する。

### コマンドライン引数のテストについて
flagパッケージを使用すると、テストの際にコマンドライン引数を指定することができない。
テストにはflag.commandLine.setを使用し境界値分析を行う。
他にテスト方法知らない。。。
