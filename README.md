# webcron
------------

一个定时任务管理器，基于Go语言和beego框架开发。用于统一管理项目中的定时任务，提供可视化配置界面、执行日志记录、邮件通知等功能，无需依赖*unix下的crontab服务。

## 项目背景


## 功能特点

* 统一管理多种定时任务。
* 秒级定时器，使用crontab的时间表达式。
* 可随时暂停任务。
* 记录每次任务的执行结果。
* 执行结果邮件通知。

## 界面截图


## 安装说明

系统需要安装Go和MySQL。

获取源码

	$ go get github.com/SuperDsj/webcron
	
打开配置文件 conf/app.conf，修改相关配置。
	

创建数据库webcron，再导入install.sql

	$ mysql -u username -p -D webcron < install.sql

运行
	
	$ ./webcron
	或
	$ nohup ./webcron 2>&1 > error.log &
	设为后台运行

访问： 

http://localhost:8000

帐号：admin
密码：admin888
