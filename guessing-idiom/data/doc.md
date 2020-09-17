## 通用参数

> 通用请求   

| name | type | intro | example |
| ---- | ---- | ----- | ------- |

> 通用返回值    

| name | type | intro |
| ---- | ---- | ----- |
| code | int | 状态码：0 成功，其他为错误，参照错误码小结 |
| data | mixed | 接口正文 |
| msg | string | 错误附加信息 |

## 生成关卡数据

> GET: /gm/games/create

注：每次最多生成 20 条   

## 获取关卡数据

> GET: /game/stage/get    

| name | type | intro | example |
| ---- | ---- | ----- | ------- |
| level | string | 关卡等级 | 1 |

``` json
{
    "code": 0,
    "data": {
        "content": "[{\"idiom\":\"招摇撞骗\",\"mean\":\"撞骗：寻机...的注意。\",\"pos\":\"0,20,30,5\",\"vis\":\"0111\"}]",
        "level": 1
    },
    "msg": ""
}

```

| name | type | intro |
| ---- | ---- | ----- |
| content | string | 关卡数据 |
| level | int | 关卡等级 |


