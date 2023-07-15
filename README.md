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
    {
        "do_time": 0.0,
        "pose": "track"
    }
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
        {
            "do_time": 0.0,
            "pose": "track"
        }
    ]
}
```

- 次の文章が物忘れ
- これはstateが`forget`なのでtakubo.LineNumberをインクリメントしない

```json
{
    "text": "だれとおばあさんがすんでいたんだっけ",
    "state": "forget",
    "behavior":[
        {
            "do_time": 2.0,
            "pose": "look-up"
        },
        {
            "do_time": 0.0,
            "pose": "track"
        }
    ]
}
```

## 物忘れ

### request

```json
{
    "word": "おじいさん",
    "line_number": 2,
    "title": "momotaro"
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
        {
            "do_time": 5.0,
            "pose": "look-up"
        },
        {
            "do_time": 0.0,
            "pose": "track"
        }
    ]
}
```

不正解の場合

```json
{
    "text": "そうだっけ？",
    "behavior":[
        {
            "do_time": 1.0,
            "pose": "look-up"
        },
        {
            "do_time": 0.0,
            "pose": "track"
        }
    ]
}
```

## port番号

```go
const (
 BackendIP     = "0.0.0.0"
 DynamixelIP   = "0.0.0.0"
 WizWebIP      = "0.0.0.0"
 BackendPort   = "3303"
 DynamixelPort = "5550"
 WizWebPort    = "5551"
)
```
