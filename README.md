# kinaco

kinaco create unique IDs using 3 or 4 letter words.

## Usage

```
result := kinaco.Generate()
// result should be "hoge-fuga-zoo"
```

```
result := kinaco.Delimiter("/").Num(4).Generate()
// result should be "hoge/fuga/zoo/four"
```
