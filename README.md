# practiceRepository
存放goweb练习所用文件仓库
bubble清单
一个基于gin+gorm开发的练手小项目。

前端页面基于vue和ElementUI开发,前端源码地址：bubble_frontend。

使用指南
配置MySQL
在你的数据库中执行以下命令，创建本项目所用的数据库：
CREATE DATABASE goweb02 DEFAULT CHARSET=utf8mb4;
在go_web_app03/conf/config.ini文件中按如下提示配置数据库连接信息。
port = 6060
release = false

[mysql]
user = 你的数据库用户名
password = 你的数据库密码
host = 你的数据库host地址
port = 你的数据库端口
db = goweb02
编译
go build
执行
Mac/Unix：
./go_web_app03 conf/config.ini

Windows:
go_web_app03.exe conf/config.ini
启动之后，使用浏览器打开http://127.0.0.1:6060/即可。
