# 简介
一个Gin + Gorm + Mysql + bootstrap实现的博客，采用MVC架构。

## 规范包及目录
```
|-controller

|-dao

|-router

|-model

|-assets

|-templates
````

## 使用指南

### 下载
```
git clone https://github.com/youke0721/GinTodo.git
```

### 配置MySQL
1. 在你的数据库中执行以下命令，创建本项目所用的数据库：

```
CREATE DATABASE golang_blog DEFAULT CHARSET=utf8mb4;
```

2. 在`dao.go`文件中配置数据库连接信息。

```
dsn := "你的数据库用户名:你的数据库密码6@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
```

### 编译

```
go run main.go
```
启动之后，使用浏览器打开<http://127.0.0.1:8080/>即可。
![1716998003713](https://github.com/youke0721/GinBlog/assets/118320030/5be4a9a8-43f7-4121-b175-7b3e64c762bd)


   



