### 安装模块
go get github.com/go-sql-driver/mysql
go get github.com/go-xorm/xorm

go mod vendor


#### 数据库
create database gin_study DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

grant all on gin_study.* to "gin_study"@"%" identified by 'av123456';

flush privileges;


git的tag功能
git 下打标签其实有2种情况

轻量级的：它其实是一个独立的分支,或者说是一个不可变的分支.指向特定提交对象的引用
带附注的：实际上是存储在仓库中的一个独立对象，它有自身的校验和信息，包含着标签的名字，标签说明，标签本身也允许使用 GNU Privacy Guard (GPG) 来签署或验证,电子邮件地址和日期，一般我们都建议使用含附注型的标签，以便保留相关信息
所以我们推荐使用第二种标签形式

Top
创建tag
git tag -a V1.2 -m 'release 1.2'
上面的命令我们成功创建了本地一个版本 V1.2 ,并且添加了附注信息 'release 1.2'

Top
查看tag
git tag
要显示附注信息,我们需要用 show 指令来查看

git show V1.2
但是目前这个标签仅仅是提交到了本地git仓库.如何同步到远程代码库

git push origin --tags
如果刚刚同步上去,你缺发现一个致命bug ,需要重新打版本,现在还为时不晚.

git tag -d V1.2
到这一步我们只是删除了本地 V1.2的版本,可是线上V1.2的版本还是存在,如何办?这时我们可以推送的空的同名版本到线下,达到删除线上版本的目标:

git push origin :refs/tags/V1.2
如何获取远程版本?

git fetch origin tag V1.2
这样我们可以精准拉取指定的某一个版本.适用于运维同学部署指定版本.