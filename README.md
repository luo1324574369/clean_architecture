
![图片](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

[Blog链接](https://juejin.cn/post/7139141106737872909/)

# 依赖原则
1. 依赖方向往内
2. 业务代码独立于框架 独立于基础设施
4. 外圈可以依赖任意一层内圈

# 各层介绍

1. 层级之间通过`interface`沟通,如有必要可以定义`assembler`转换格式,但注意`entity`层应不能依赖外层

## entity
1. 既可以是有方法的对象,也可以是一组函数的集合
2. 可以供所有层级使用

## logic(use cases)
1. 负责逻辑具体实现,编排entities的数据变更以及使用他们的属性完成目的,负责回答的层
2. 建议各业务分拆子目录实现

## service (controllers) 和 repository (gateways)
1. port与adapter层,用于沟通业务和外部设施的桥梁,例如数据库和web
2. 不应该有业务逻辑

```
| -- main.go            # 项目启动入口, 调用service各业务的new方法
| -- entity
    | -- log              # 定义常用的通用方法,例如: 日志,隔离核心代码与外部设施的依赖
    | -- gift             # 每个业务定义一个文件夹,存在各个实体 
| -- logic
    | -- gift             # 每个业务定义一个文件夹
        | -- port.go        # 定义接口,供service调用
        | -- adapter.go     # port的具体实现
| -- repository 
    | -- gift
        | -- adapter.go     # 实现gateway的port
        | -- assembler.go   # entity与外部设施(mysql,redis,http)沟通的数据结构转换器,非必须
        | -- mysql          # mysql/redis/http 视各自业务而定
        | -- port.go        # 接口文件
| -- service            # 服务入口,建议各自业务定义一个文件
    | -- gift.go          # 调用logic的interface, 发起提问的作用
```


# QA
1. entity中的log包的作用
> 用于隔离核心代码与外部依赖

2. repository中必须定义po吗
> 否,只是按照业务需要,可直接依赖entity

3. service 调用 repository 需要定义接口吗
> 没有强制规定,由于都属于外层,是不稳定的,可以不定义