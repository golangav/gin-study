### 安装模块
go get github.com/go-sql-driver/mysql
go get github.com/go-xorm/xorm

go mod vendor


#### 数据库
create database gin_study DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

grant all on gin_study.* to "gin_study"@"%" identified by 'av123456';

flush privileges;