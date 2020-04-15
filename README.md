# rain_dog_backend

落水狗Go后台

### gin官方文档
```shell script
https://gin-gonic.com/zh-cn/docs/examples/
```

### 起步
首先进入 `rainDog->src->blog` 路径下

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

### 前后分离权限思考
因为项目是前后端分离的，所以我认为前端界面菜单可以直接写死，只需要使用前后端规定好的菜单路由就可以了。

### 参考项目
```text
https://github.com/it234/goapp
```

## 感谢以下框架的开源支持

- [Gin] - [https://gin-gonic.com/](https://gin-gonic.com/)
- [GORM] - [http://gorm.io/](http://gorm.io/)
- [Casbin] - [https://casbin.org/](https://casbin.org/)
