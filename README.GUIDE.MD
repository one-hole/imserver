1. ClientManager
2. Client
3. Room ?

package main

func execManager() {
  根据参数（来选择性和 Redis 或 RabbitMQ）创建连接、并监听

  -> package source RunRedis
  -> package source RunRabbit
}

package sockets

func (ClientManager) Exec() {
  这里去监听事件
  1: 新的连接接入
  2: 连接释放
  3: 待推送的消息事件
}

package models

func Init() {
  1. 创建数据库连接
  2. 抽象持久化对象
}

1. 黑白名单功能
2. 连接 Token 有效期的功能（用户续费的时候很管用）
3. Token 鉴权的功能（比如客户订购了 CSGO 但是没有订购 Dota2）

package api

file: api/ws/base.go

1. 校验 IP 白名单
2. 校验 Token 权限
3. 分发长连接到不同的 Manager

package sockets

func ServeWS() {
  1. 升级短连接为长连接
  2. 生成连接抽象对象
  3. 注册连接对象到相关的 Manager
}