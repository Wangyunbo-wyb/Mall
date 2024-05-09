# Mall
-------------------------

### 主要功能
+ 用户注册登录(jwt-go)
+ 用户基本信息修改，解绑定邮箱，修改密码
+ 商品的发布，浏览，各个商品的浏览次数，以及部分种类商品的排行等
+ 购物车的加入，删除，浏览等
+ 订单的创建，删除，支付等
+ 地址的增加，删除，修改等
+ 添加ELK体系，方便日志查看和管理
+ 商品秒杀功能还未完善（只完成了消息队列的初始化）


### 项目结构
```
gin-mall
├── api             # 用于定义接口函数，也就是controller的作用
├── cmd             # 程序入口
├── conf            # 配置文件
├── doc             # 文档
├── middleware      # 中间件
├── model           # 数据库模型
├── pkg
│  ├── e            # 错误码
│  └── util         # 工具函数
├── repository
│  ├── cache        # Redis缓存
│  ├── db           # 持久层的mysql
│  │  ├── dao       # dao层，对db进行操作
│  │  └── model     # 定义mysql的模型
│  ├── es           # ElasticSearch，形成elk体系
│  └── mq           # 放置各种mq，kafka，rabbitmq等等
├── routes          # 路由逻辑处理
├── serializer      # 将数据序列化为 json 的函数，便于返回给前端
└── service         # 接口函数的实现

```

### 
