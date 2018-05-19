## app端接口

###1. 客户端App调起H5游戏

    http://xxx.xxx.xxx?app_version=xxx&addr=xxxxx&type=xxxx&name=xxxxxx&...

必填参数：

 - [app_version] 应用版本号
 - [addr]        钱包地址        //用户ID，唯一标识用户
 - [type]        钱包类型
 - [name]        钱包用户昵称

###2. 游戏结束调用客户端App

使用deeplink，参数追加在url后面

    coinpay://com.x.wallet/game?result=xxx

必填参数：

 - [result] 个人游戏结果

## 后台服务端接口

###1. 开始游戏，获取糖果

url:api/game/start
method: post
参数:

    {
        "game_id": xx      //必填，游戏类型id，app会传过去
        "game_room_id": xx //必填，游戏每一期id
        "players": [       //必填
            {
                "name": "用户名"
                "addr": "用户地址" //必填
            }
        ]
    }

返回:

    {
        "code" : 0         // 状态码 0表示成功
        "message" : ""     // 消息
        "value": [         // 业务返回
            {
                "candy_count": xxx //可分配糖果数量，**字符串类型**
                "candy_type" : xxx //糖果类型 candy_type < 10000 表示是糖果
                "candy_label": xxx //糖果标签
            },
            {
                "candy_count": xxx //可分配糖果数量，**字符串类型**
                "candy_type" : xxx //糖果类型 candy_type >= 10000 表示是钻石
                "candy_label": xxx //糖果标签
            }
        ]
    }

###2. 结束游戏调用后台服务器
method: post
url:api/game/over
参数:

    {
        "game_id":xx            //同上
        "game_room_id":xx       //同上
        "players":[             //同上
            {
                "name": "xx"
                "addr": "xx"
                "score":"xx"    //获取到的糖果数量，**字符串类型**
            }
        ]
    }

返回:

    {
        "code" : 0         // 状态码 0表示成功
        "message" : ""     // 消息
        "value": {         // 业务返回

        }
    }