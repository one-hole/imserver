### How to start project

* Dev env
  ```
  GO_ENV=debug go run main.go
  ```
  
* Docker env for debug
  ```
  docker run -e GO_ENV=release -v "/$(pwd)/config/config_docker.yml:/root/config/config.yml" -p 8000:8000 7064805/imserver:latest
  ```

* Docker env for release
  ```
  docker run -e GO_ENV=release -v "/$(pwd)/config/config_docker.yml:/root/config/config.yml" -v "/$(pwd)/logs:/root/logs" -p 8000:8000 7064805/imserver:latest
  ```

### How to customize your log format

* [gin-logrus](https://github.com/w-zengtao/gin-logrus/blob/master/logger.go) rewrite function `Logger()`
* More details in [logrus](https://github.com/sirupsen/logrus)
---

运行环境依赖如下

* Redis
* MySQL
* RabbitMQ

---


使用到的三方库如下

* [Redis](https://github.com/go-redis/redis)
* [Gorm](https://github.com/jinzhu/gorm)
* [RabbitMQ](https://github.com/streadway/amqp)
* [JSON](https://github.com/tidwall/gjson)
* [WebSocket](https://github.com/gorilla/websocket)
* ~~[YAML](https://github.com/go-yaml/yaml)~~
* [Gin](https://github.com/gin-gonic/gin)
* [Viper](https://github.com/spf13/viper)
* [Logrus](https://github.com/sirupsen/logru)

----

正在处理

* [x] 迁移到 Gin 上 (Down)
* 希望有 Web Rest APi 管理 ([Gin](https://github.com/gin-gonic/gin)) （撸了一个开头...）
* 希望能基于上述 APi 加入可视化管理

下一步

* 分布式 - 主（Master）、从（Slave）、调度（Schedule）、注册（Register) 真他妈越走越偏了
* Docker 打包的时候传入一些环境变量 & Monut 一下 ConfigFile