运行环境依赖如下

* Redis
* RabbitMQ

---
User Story 如下

*  可以针对某个连接直接推送
*  可以针对某些连接直接推送
*  可以针对某个房间直接推送
*  可以针对某个类房间直接推送

Client 对于某个「客户端」连接的抽象

ClientManger 管理所有的连接 & 连接的释放
* 推送的数据源应该来自于 RabbitMQ & Redis
* 连接可以加入、可以离开某个房间
* 消息是针对房间进行推送
* 房间的建立时机

---


使用到的三方库如下

* [Redis](https://github.com/go-redis/redis)
* [RabbitMQ](https://github.com/streadway/amqp)
* [JSON](https://github.com/tidwall/gjson)
* [WebSocket](https://github.com/gorilla/websocket)
* [YAML](https://github.com/go-yaml/yaml)

----

正在处理

* [x] 迁移到 Gin 上 (Down)
* 希望有 Web Rest APi 管理 ([Gin](https://github.com/gin-gonic/gin)) （撸了一个开头...）
* 希望能基于上述 APi 加入可视化管理

下一步

* 分布式 - 主（Master）、从（Slave）、调度（Schedule）、注册（Register) 真他妈越走越偏了
* Docker 打包的时候传入一些环境变量 & Monut 一下 ConfigFile