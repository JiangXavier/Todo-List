# Todo-List
### 项目主要功能
- 用户注册登录 ( jwt-go鉴权 )
- 新增/删除/修改/查询 备忘录
- 存储每条备忘录的浏览次数
- 分页功能
### 项目主要依赖
- Gin
- Gorm
- mysql
- redis
- ini
- jwt-go
- logrus
### 项目结构
```shell
TodoList/
├── api
├── cache
├── conf
├── middleware
├── model
├── pkg
│   └──  util
├── routes
├── serializer
└── service
```
- api : 用于定义接口函数
- cache : 放置redis缓存
- conf : 用于存储配置文件
- middleware : 应用中间件
- model : 应用数据库模型
- pkg/util : 工具函数
- routes : 路由逻辑处理
- serializer : 将数据序列化为 json 的函数
- service : 接口函数的实现
### 配置文件
#### 本地运行服务以及mysql数据库
**conf/config.ini**
```ini
# debug开发模式,release生产模式
[service]
AppMode = debug
HttpPort = :3000
# 运行端口号 3000端口

[mysql]
Db = mysql
DbHost =
# mysql ip地址
DbPort = 
# mysql 端口号
DbUser = 
# mysql 用户名
DbPassWord = 
# mysql 密码
DbName = 
# mysql 名字
```
#### redis
**cache/common.go**
```go
func ConnRedis() {
	rd := redis.NewClient(&redis.Options{
		Addr: "", // url
		Password: "",
		DB:1,   // 1号数据库
	})
	result, err := rd.Ping().Result()
	if err != nil {
		fmt.Println("ping err :",err)
		return
	}
	fmt.Println(result)
}
```
### 项目运行
#### 下载依赖
`go mod tidy`
#### 运行
`go run main.go`
### 备注
- `mysql`是存储主要数据
- `redis`用来存储备忘录的浏览次数
