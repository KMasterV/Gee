## 项目介绍

本项目是基于gin框架搭建，gorm驱动mysql数据库，实现用户注册，登录的实例，仅供学习分享，感谢阅读。

## 运行环境

确保本地环境已安装golang语言1.16以上版本，go get 项目依赖，并配置config/application.yml文件以适应本地数据库环境

```
# 拉取代码
git clone -b master https://github.com/KMasterV/Gee defined_name
# 进入项目目录
cd defined_name
# 安装项目依赖
go get
# 配置yml文件
datasource:
  host: localhost
  port: 3306
  database: databaseName
  username: root
  password: root
```

## 启动项目

```
go run main.go router.go
```

