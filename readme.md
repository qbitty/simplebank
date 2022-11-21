

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

##### 第二十四章 构建最小化的golang app镜像 -- 多阶段构建

- 基于golang镜像构建golang app二进制可执行文件

- 基于最小化系统镜像及上一步的构建结果构建最终镜像

- 注意：

  - 国内构建时可能无法下载相应的golang module，可通过在Dockerfile中声明ENV GOPROXY https://goproxy.cn,direct解决

    ```dockerfile
    ENV GOPROXY https://goproxy.cn,direct
    ```

  - 多阶段构建时使用如下格式引用特定阶段的构建结果

    ```dockerfile
    COPY --from=builder /app/main .
    ```

##### 第二十五章 使用docker network连接两个独立的容器

- 创建docker 网络
- 运行docker容器时通过--network <networ name>指定网络名称
- 运行docker容器时注意指定数据源环境变量

##### 第二十六章 使用docker compose管理服务

- 定义docker-compose.yaml
- 修改Dockerfile以适配docker compose
- 注意服务之间的依赖关系
##### 第二十七章 创建免费的AWS账号

##### 第二十八章 通过github action自动构建及推送docker镜像到AWS ECR

- 创建镜像仓库
- 创建用户github-ci并记下对应的access-id和access-key
- 创建deployment用户组
- 给用户组授权授权
- 将github-ci关联到deloyment用户组
- 在github仓库的设置中的secrets下的actions中添加AWS_ACCESS_KEY_ID和AWS_SECRET_ACCESS_KEY（对应github-ci的取值）

##### 第二十九章 通过AWS RDS中创建生产级数据库

##### 第三十章 使用aws secrets manager管理生产环境密钥

- 生成随机token_key

  ```sh
  openssl rand -hex 64 | head -c 32
  ```

- 创建aws 密钥

- 安装aws cli

- 配置本地aws cli认证

- 授予用户组访问secret manager service的权限

- 解析secret manager service中的密钥

  - 安装jq

    ```sh
    brew install jq
    ```

  - 解析密钥

    ```sh
    aws secretsmanager get-secret-value --secret-id simple_bank --query SecretString --out text | jq -r 'to_entries|map("\(.key)=\(.value)")|.[]' > app.env
    ```

  - 将解析命令复制到deploy.ci中

  - 提交代码自动push镜像

  - 登录docker

    ```sh
    aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 396605755172.dkr.ecr.ap-northeast-1.amazonaws.com
    ```

##### 第三十一章 K8s架构及如何在aws上创建EKS集群

- 创建集群
- 创建node group
- 伸缩node

##### 第三十二章 使用kubectl和k9s连接aws eks集群

- 更新集群配置信息到本地

  ```sh
  aws eks update-kubeconfig --name simple-bank --region ap-northeast-1
  ```

- 配置集群创建者的access_id和access_key

  ```sh
  cat ~/.aws/credentials
  ```

- 创建aws auth的configMap

  ```yaml
  apiVersion: v1 
  kind: ConfigMap 
  metadata: 
    name: aws-auth 
    namespace: kube-system 
  data: 
    mapUsers: | 
      - userarn: arn:aws:iam::396605755172:user/github-ci
        username: github-ci
        groups:
          - system:masters
  ```

- 基于集群创建者的认证信息部署confgiMap

  ```sh
  export AWS_PROFILE=default
  kubectl apply -f eks/aws-auth.yaml
  ```

- 配置非集群创建者的access_id和access_key

##### 第三十二章 基于AWS EKS部署应用

- 创建deployment
- 创建LoadBalancer类型的service

##### 第三十三章 使用AWS Route 53注册域名及配置DNS记录

- 购买域名
- 配置DNS的A记录，使其指向service的loadbalancer

##### 第三十四章 使用ingress路由网络请求

- 声明ingress.yaml

- 部署ingress

- 部署ingress controller

  ```sh
  kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.4.0/deploy/static/provider/aws/deploy.yaml
  ```

- 修改service的类型为ClusterIP并重新部署

- 修改DNS的A记录使其指向ingress loadbalancer

##### 第三十五章 基于cert-manager的TLS自动且免费认证

- 部署cert-manager组件

  ```sh
  kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.10.0/cert-manager.yaml
  ```

- 声明issuer.yaml

- 部署issuer

- 修改ingress.yaml

- 重新部署ingress

##### 第三十六章 通过github action自动部署应用到aws eks

- 参考配置文件



##### 常用第三方库

- [golang-migrate](https://github.com/golang-migrate/migrate)
- [Gorm](https://gorm.io/)
- [sqlx](https://github.com/jmoiron/sqlx) 
- [sqlc](https://sqlc.dev/)
- [testify](https://github.com/stretchr/testify)
- [Gin](https://github.com/gin-gonic/gin)
- [Beego](https://github.com/astaxie/beego)
- [Echo](https://github.com/labstack/echo)
- [Revel](https://github.com/revel/revel)
- [Martini](https://github.com/go-martini/martini)
- [Fiber](https://github.com/gofiber/fiber)
- [Buffalo](https://github.com/gobuffalo/buffalo)
- [Fast HTTP](https://github.com/valyala/fasthttp)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [HTTP Router](https://github.com/julienschmidt/httprouter)
- [Chi](https://github.com/go-chi/chi)
- [Viper](https://github.com/spf13/viper)
- [gomock](https://github.com/golang/mock)
- [google/uuid](https://github.com/google/uuid)
- [jwt-go](https://github.com/dgrijalva/jwt-go)
- [grpc](https://grpc.io)
- [proto](https://developers.google.com/protocol-buffers)
- 

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

##### 核心知识点

- 数据库设计
  - dbdigram
  - dbdoc
  - dbml
- 数据访问技术(DAO)
  - sqlc
- web框架
  - 路由
  - 参数校验
    1. 内置参数校验器
    2. 自定义参数校验器
  - 中间件
    1. 日志
    2. 异常恢复
    3. 认证授权
    4. 安全控制
- 事务处理技术
- 测试技术
  - mock
- 配置技术
- CI/CD流程