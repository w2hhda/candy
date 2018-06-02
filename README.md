
文档
基础参数：
{
    "app_version": "xxxx",
    "imei":"xxxxx",
    "os_version":"xxxx",
    ...//具体业务参数
}

基础返回：
{
    "code": 0 //0 代表成功
    "message": "success"
    "value" : //具体业务返回
}

========================================
### 1. 提交游戏结果
url： /api/game/record
参数： 基础参数
name：用户名
addr：地址
candy_type: 糖果类型,
count: 糖果数量
game_id: 游戏ID

{
    "imei":"xxxxxxxx",
    "addr":"0xfb0a596ec522791b99a7162fee7891a43186745D", //改这个
    "name":"dada",
    "game_id":1,
    "candy_type":1,
    "count":100
}

返回：
{
    "code": 0,
    "message": "success",
    "value": ""
}

=======17:51============================


### 1.糖果列表
url： /api/candy/list
参数： 基础参数
      page_number：代表上一页的页码，第一次取传0
      addrs ：地址数组
{
    "page_number": 0,
    "addrs":["0xfb0a596ec522791b99a7162fee7891a43186745R"]
}
返回：
{
    "code": 0,
    "message": "success",
    "value": {
        "page_number": 1,
        "page_size": 10,
        "total_page": 1,
        "list": [
            {
                "id": 7,
                "all_count": "1000000",
                "remaining_count": "1000000",
                "token_addr": "asdfasd",
                "candy_label": "ETH",
                "candy_type": 1,
                "rate": 11
            },
            {
                "id": 8,
                "all_count": "100",
                "remaining_count": "100",
                "token_addr": "12312312",
                "candy_label": "EOS",
                "candy_type": 10000,
                "rate": 13
            },
            {
                "id": 9,
                "all_count": "1000",
                "remaining_count": "1000",
                "token_addr": "sfasdfas",
                "candy_label": "BTC",
                "candy_type": 2,
                "rate": 12
            },
            {
                "id": 10,
                "all_count": "111",
                "remaining_count": "111",
                "token_addr": "sdfsdf",
                "candy_label": "NEO",
                "candy_type": 10001,
                "rate": 14
            }
        ]
    }
}

### 2.游戏入口
url： /api/candy/index
参数： 基础参数
{
    "code": 0,
    "message": "success",
    "value": {
        "all_candy_count": "1.001211E+06",
        "all_game_list": [
            {
                "id": 1,
                "url": "https://beego.me/docs/mvc/model/orm.md",
                "sort": 1,
                "status": 1,
                "icon": "https://beego.me/static/img/beego_purple.png",
                "name": "jump"
            }
        ]
    }
}

### 3.排行榜
url：/api/rank
参数： 基础参数
      page_number：代表上一页的页码，第一次取传0
{
    "code": 0,
    "message": "success",
    "value": {
        "page_number": 0,
        "page_size": 10,
        "total_page": 1,
        "List": [
            {
                "addr": "0xfb0a596ec522791b99a7162fee7891a43186745b",
                "count": "1321",
                "value": "1.5753E+04"
            },
            {
                "addr": "0xfb0a596ec522791b99a7162fee7891a43183453",
                "count": "222",
                "value": "2.442E+03"
            }
        ]
    }
}

### 4.账单
url：/api/record/list
参数：基础参数
     page_number：代表上一页的页码，第一次取传0
     addrs ：地址数组
{
    "code": 0,
    "message": "success",
    "value": {
        "page_number": 1,
        "page_size": 10,
        "total_page": 1,
        "List": [
            {
                "id": 1,
                "addr": "0xfb0a596ec522791b99a7162fee7891a43186745b",
                "count": "100",
                "candy": {
                    "id": 7,
                    "all_count": "1000000",
                    "remaining_count": "1000000",
                    "token_addr": "asdfasd",
                    "candy_label": "ETH",
                    "candy_type": 1,
                    "rate": 11
                },
                "create_at": "2222"
            }
        ]
    }
}

### 5. 我的糖果
url：/api/token/list
参数：基础参数
     addrs ：地址数组
{
    "code": 0,
    "message": "success",
    "value": [
        {
            "addr": [
                "0xfb0a596ec522791b99a7162fee7891a43186745b"
            ],
            "label": "BTC",
            "count": "1.222E+03",
            "icon": "",
            "rate": 12
        },
        {
            "addr": [
                "0xfb0a596ec522791b99a7162fee7891a43186745b",
                "0xfb0a596ec522791b99a7162fee7891a43183453"
            ],
            "label": "ETH",
            "count": "3.21E+02",
            "icon": "",
            "rate": 11
        }
    ]
}

