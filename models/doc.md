### Models

#### Config

```
Config 表控制全局行为、比如 开启|关闭 白名单、黑名单、租户永久试用等
```

#### Tenant

```
租户、Client 背后实际的承载方

Name
Email (Uniq)
Phone
PasswordDigest
ReceivePort
ServerCount
```

#### Server

```
Client 实例的本质

ID
Address
Enable
```