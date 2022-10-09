### Backend master class



##### 第一章 数据库设计

- 工具：

  [dbdigram.io]: https://dbdiagram.io

- 表结构

  - accounts
  - entries
  - transfers

##### 第二章 环境搭建

- 安装docker
- 配置postgresql容器
- 安装数据库可视化工具
- 执行表结构设计语句

##### 第三章 database migration

- 安装golang-migrate
- 生成migration文件
- 使用MakeFile简化项目构建

##### 第四章 CRUD代码生成（sqlc）

- 常用数据库操作库对比：database/sql，Gorm，Sqlx，Sqlc
- 安装Sqlc
- 使用sqlc init创建配置文件并修改
- 编写sql query语句
- 生成CRUD代码：sqlc generate

##### 第五章 CRUD单元测试

- Test Create Account
- Generate random data
- Test Get Account
- Test Update Account
- Test Delete Account
- Test List Accounts

##### 第六章 数据库事务实现

##### 第七章 DB transaction lock & How to handle deadlock

##### 第八章 如何避免死锁

##### 第九章 事务隔离级别及读现象

##### 第十章 github action

##### 第十一章 实现restful api

##### 第十二章 配置加载

##### 第十三章 Mock测试





##### 最佳实践总结

- 数据库设计最佳实践
  - 表名用复数
  - 结构体名用单数
- 测试最佳实践
  - 优先使用随机数据进行测试
  - 各测试用例之间保持独立，互不影响

- 事务处理最佳实践

  https://blog.devgenius.io/go-golang-clean-architecture-repositories-vs-transactions-9b3b7c953463