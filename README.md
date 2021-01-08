#### Golang版本
- 1.12.4
### 用到的组件
- go get -u github.com/gin-gonic/gin
- go get github.com/jinzhu/gorm
- go get github.com/go-sql-driver/mysql
- go get -u github.com/go-redis/redis
- go get  github.com/bitly/go-simplejson
- go get github.com/gin-contrib/cors
- go get github.com/nfnt/resize
- go get github.com/skip2/go-qrcode

### 需要配置的环境变量：
- MYSQL_PASSWORD：数据库的密码
- ALIYUN_KEYID:阿里云短信的keyid
- ALIYUN_SECRET：阿里云短信的secret
- STATIC_PATH:静态文件存放的目录
- VERSION_NUM:版本号，例如：1.2
- RECOMMEND_PLAN:推荐的训练计划id