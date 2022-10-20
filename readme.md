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

- 什么是数据库事务
  - 一个由多个数据库操作组成的不可分割的工作单元
- 为什么需要数据库事务
  - 保证工作单元的可靠和一致性
  - 保证同一工作的单元被并发访问时，彼此独立，互不影响
- 事务具备哪些特性
  - 原子性
  - 一致性
  - 隔离性
  - 持久性
- 数据库事务的隔离级别
  - 读未提交
  - 读已提交
  - 可重复读
  - serilization
- 如何在golang中优雅地实现事务
- 如何测试golang中的事务
  - the best way to make sure that our transaction works well is to run it with several concurrent go routines.

##### 第七章 DB transaction lock & How to handle deadlock

- 如何测试死锁问题

- 如何排查数据库死锁
- 如何处理数据库死锁

##### 第八章 如何避免死锁

- 如何从业务逻辑层面避免死锁

##### 第九章 事务隔离级别及读现象

##### 第十章 github action

- workflow
- job
- step
- action
- runner

##### 第十一章 实现restful api

- 路由机制

- 参数校验

- 参数绑定

- 结果响应

  - 正常结果
  - 异常结果

- web服务

- api单元测试

  - 表驱动测试

  - mock数据访问层
  - 保证测试覆盖率（test case齐全）

##### 第十二章 配置加载

- viper
- 配置文件场景

##### 第十三章 Mock测试

- gomock
- 打桩(stub)
- 校验结果

##### 第十四章 自定义参数校验

- 定义validator.Func类型函数并实现校验逻辑
- 将上述函数注册到gin Validator中
- 将自定义校验器应用到特定参数上

##### 常用第三方库

- [golang-migrate](https://github.com/golang-migrate/migrate)
- [Gorm](https://gorm.io/)
- [sqlx](https://github.com/jmoiron/sqlx) 

- [sqlc](https://sqlc.dev/)
- [testify](https://github.com/stretchr/testify)

##### 最佳实践总结

- 数据库设计最佳实践
  - 表名用复数
  - 结构体名用单数
  - 使用数据库迁移工具管理数据库变更(migration)
- 测试最佳实践
  - 优先使用随机数据进行测试
  - 各测试用例之间保持独立，避免相互影响
- 通过main_test.go初始化测试环境和相关依赖
  
  - 同一个包只能有一个主测试(MainTest)
  
- 事务处理最佳实践

  https://blog.devgenius.io/go-golang-clean-architecture-repositories-vs-transactions-9b3b7c953463

- 小技巧
  - 编译时校验接口实现如：var _ Querier = (*Queries)(nil)
  - 尽量让构建过程自动化(make)
  - 闭包应用：Go lacks support for generics type, closure is often used when we want to get the result from a callback function, because the callback function itself doesn’t know the exact type of the result it should return.