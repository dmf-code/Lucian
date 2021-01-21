# rain_dog_backend

落水狗Go后台

### gin官方文档
```shell script
https://gin-gonic.com/zh-cn/docs/examples/
```

### 起步
首先进入 `internal->app` 路径下

第一次部署项目需要先初始化数据

```shell script
go run main.go Migration
```

这个命令会创建数据表和迁移基础框架数据，默认会创建一个管理员账号 `admin` : `123456`

```shell script
go run main.go
```

这样子就能够自动安装所需要的包了

> 注意这里面会出现被墙的问题，导致无法下载包的问题这里面可以设置代理解决
> set http_proxy=127.0.0.1:1080 （这个是我的小飞机的代理端口）


### Restful接口 命名推荐如下

|   HTTP方法   |   URI   |  动作    |
| ----         | ----               | ----    |
|   GET        |    /photos	        |   index   |
|   GET        |    /photos/:photo  |   show   |
|   PUT        |    /photos/:photo  |   update   |
|   POST       |    /photos         |   store |
|   DELETE     |    /photos/:photo  |   destroy   |


### Casbin 权限介绍
```text
EnforceSafe 检查权限 (bool, error)
```

`model.conf` 文件的 `matchers` 的用途，编写不同的规则语句进行判断。这里有一堆内置函数。

> 官方匹配函数文档 https://casbin.org/docs/zh-CN/function

### 前后分离权限思考
因为项目是前后端分离的，所以其实我们全部权限一开始其实是写好的。然后为什么还要请求接口去渲染呢？这里是因为
不同的角色拥有不同的权限，如果不请求接口进行渲染的话，我们就不知道哪些是他应该渲染出来的菜单。所以这里还是
需要请求接口的形式，让前端动态渲染界面。


### .env文件初始配置
```text
SECRET_KEY=YOURSECRETKEYGOESHERE # 密码盐字段
USER=root # 数据库用户
PASSWORD=root # 数据库密码
IP=192.168.3.9 # 数据库ip
PORT=9003 # 数据库端口
DBNAME=rain_dog # 数据库名称
TOKEN_EXPIRE_TIME=1024 # token过期时间，单位小时
INIT_ADMIN_TABLE=true # 第一次填写true初始化表数据，之后改为false
```

### 参考项目
```text
https://github.com/it234/goapp
```

## 感谢以下框架的开源支持

- [Gin] - [https://gin-gonic.com/](https://gin-gonic.com/)
- [GORM] - [http://gorm.io/](http://gorm.io/)
- [Casbin] - [https://casbin.org/](https://casbin.org/)
