# hitcomic-local-server
> HITcomic show local ticker server.



## Staff API (Cert API)

GET `/staff/:key`

Response

```json
Content-Type: "application/json"

{
    "result":"success",
    "quantity": 2, //有 2 次入场记录
    "times": 3, //剩餘次數
    "records":[
        {
            "time": "2018-06-18 09:29:20",
            "result": 0,
            "picture": "4BCA7104-B36F-415B-977D-1886AE149230.jpg"
        },
        {
            "time": "2018-06-18 12:21:30",
            "result": 0,
            "picture": "4BCA7104-B36F-415B-977D-1886AE149230.jpg"
        }
    ]
}
```



POST `/staff/:key`

Request

```json
Content-Type: "multipart/form-data" //http 內容類型發生變化

... 文件上傳 bit流


```

