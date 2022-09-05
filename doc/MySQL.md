# MySQL

MySQL 和 Navicat 安装完成

安装目录：C:\Program Files\MySQL\MySQL Server 8.0\bin

账号和密码：

root密码：12345678



## 连接MySQL

https://github.com/go-sql-driver/mysql



## 终端中使用 MySQL

```text
mysql -u root -p
```

 

gorm文档：https://gorm.io/zh_CN/docs/index.html



创建表：

```mysql
create table Users (
`id` int unsigned auto_increment,
`nick_name` varchar(100) not null,
`password` varchar(100) not null,
primary key (`id`)
)engine=InnoDB default charset=utf8mb4;
```

说明：gorm 读取的使用结构体定义称「NickName」，读取的时候是「nick_name」



查看 warning：

> show warnings

用 utf8 需要使用 utf8mb4



清空数据库表：

> truncate table 表名;

