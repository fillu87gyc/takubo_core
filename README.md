# たくぼーこあ

## これはなに

- たくぼーのコアになる部分です

## 話し始める前

`GET /v2/detect`

### request

```json
{
    "recog": "ももたろー"
}
```

### response

wake wordだった場合

```json
{
  "text": "いまからね、桃太郎をね、話すよ",
  "state": "talking",
  "behavior":[
    "track"
  ]
}
```

違う場合(なにもしない)

```json
{
    "text": "",
    "state": "detect"
}
```

## 発話中

`GET /v2/talking`

### request

```json
{
    "title": "momotaro",
    "line_number": 2
}
```

### response

- 通常文の発話
- stateがtalkingのレスポンスが帰ってきた場合にのみtakubo.LineNumberをインクリメント

```json
{
    "text": "どんぶらこ、どんぶらこと",
    "state": "talking",
    "behavior":[
        "track"
    ]
}
```

- 次の文章が物忘れ
- これはstateが`forget`なのでtakubo.LineNumberをインクリメントしない

```json
{
    "text": "だれとおばあさんがすんでいたんだっけ",
    "state": "forget",
    "behavior": [
        "think",
        "track"
    ]
}
```

## 物忘れ

### request

```json
{
    "title": "momotaro",
    "line_number": 2
}
```

### response

- 正解の場合
- state == `talking`なのでtakubo.LineNumberをインクリメント

```json
{
    "text": "そうだ、それそれー",
    "state": "talking",
    "behavior":[
        ["look-up", 5.0],
        ["track", 0]
    ]
}
```

不正解の場合

```json
{
    "text": "そうだっけ？",
    "behavior":[
        ["look-up", 1.0],
        ["track", 0]
    ]
}
```
