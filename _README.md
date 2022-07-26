# W3bstream Phase1 进度

## phase1

|                              | 说明                       | 编码 | unit test | bench test | integrate test |
| :---                         | :---                       | :--- | :---      | :----      | :---           |
| kit/model & modelgen         | db model 代码生成          | done | done      | done       | done           |
| kit/enum & enumgen           | IntStriner                 | done | done      | done       | done           |
| kit/statuserr & statusxgen   | StatusCode, Source, Msg... | done | done      | done       | done           |
| conf/log                     | 日志 logrus                | done | done      | -          | done           |
| conf/env                     | 配置parser,loader          | done | done      | -          | done           |
| conf/postgres                | 数据库连接配置             | done | done      | -          | done           |
| conf/mqtt                    | mqtt连接配置和基本操作     | done | done      | -          | done           |
| http/operator                | http get/post/...Outputer  | done | done      | -          | -              |
| http/transport               | http router/handler        | done | done      | -          | -              |
| http/transformer/form        | -                          | done | done      | -          | -              |
| http/transformer/json        | -                          | done | done      | -          | -              |
| http/transformer/binary      | -                          | done | done      | -          | -              |
| http/transformer/plain       | -                          | done | done      | -          | -              |
| http/middlewares/cors        | -                          | done | done      | -          | -              |
| http/middlewares/auth        | jwt                        | done | done      | -          | -              |
| http/statuserr               | http 状态码 错误           | done | -         | -          | -              |
| base/validator               | 数据校验                   | done | -         | -          | -              |
| base/context                 | context transmit           | done | done      | done       | done           |
| base/app                     | 服务脚手架                 | done | done      | -          | done           |
| other                        | -                          | -    | -         | -          | -              |

## phase2

|               | 说明  | 编码 | unit test | bench test | integrate test |
| :---          | :---  | :--- | :---      | :----      | :---           |
| applet create |       | done | -         | x          | -              |
| applet remove |       | done | -         | x          | -              |
| conf/ipfs     |       |      | -         | x          | -              |
| conf/graphql  |       |      | -         | x          | -              |

## 周计划(07.25~07.29)

1. 测试applet已完成功能
2. applet deploy

* conf/ipfs
* conf/graphql

3. 完善一下基础模块功能

## 代码仓库

1. [脚手架](https://github.com/saitofun/qkit)
2. [W3bstream backend](https://github.com/iotexproject/w3bstream)
