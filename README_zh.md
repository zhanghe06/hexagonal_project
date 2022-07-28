# 六边形架构（结合GORM）

[English Document](README.md)

## 项目依赖

解决`go mod`或`go get`超时
```
go env -w GOPROXY=https://goproxy.cn,direct
```

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/gin-gonic/gin
go get -u github.com/go-playground/validator/v10
go get -u github.com/nsqio/go-nsq
go get -u github.com/google/uuid
go get -u github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen@v1.6.0
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44.2
```

## 项目演示

```
golangci-lint run ./...
go generate ./...
go test -count=1 -cover ./...
```

## Gorm MySQL 字段对比

参考：[https://dev.mysql.com/doc/refman/8.0/en/integer-types.html](https://dev.mysql.com/doc/refman/8.0/en/integer-types.html)

| gorm | mysql | 字节长度 | 数字范围
| --- | --- | --- | ---
| uint8 | tinyint(4) unsigned | 1 | 0-255
| uint16 | smallint(6) unsigned | 2 | 0-65535
| - | mediumint(9) unsigned | 3 | 0-16777215
| uint32 | int(11) unsigned | 4 | 0-4294967295
| uint64 | bigint(21) unsigned | 8 | 0-18446744073709551615

说明：
1. mysql类型int(M) M 表示最大显示宽度，与所占多少存储空间并无任何关系
2. int 在32位系统中，和uint32相同；在64位系统中，和uint64相同

mysql 字符串类型

是否可变 | 字段 | 字节长度 | 字符数量
--- | --- | --- | ---
固定 | char(n) | 1 | 0-255
可变 | varchar(n) | 2 | 0-65535
可变 | tinytext | 1 | 0-255
可变 | text | 2 | 0-65535
可变 | mediumtext | 3 | 0-16777215
可变 | longtext | 4 | 0-4294967295

mysql 时间类型

字段 | 字节长度 | 含义
--- | --- | ---
date | 3 | 日期
time | 3 | 时间
datetime | 8 | 时间
timestamp | 4 | 时间
year | 1 | 年份

## GORM with Context

[GORM with Context](https://gorm.io/docs/context.html)

## 软删

[Without Hooks/Time Tracking](https://gorm.io/docs/update.html#Without-Hooks-x2F-Time-Tracking)

## 关于参数`model.Xxxx`和`entity.Xxxx`的使用场景

- service 使用`entity.Xxxx`，业务逻辑的参数可能混合不同表的数据，使用实体
- repository 使用`model.Xxxx`，数据操作直接对单表进行操作，使用model

## DTO、VO

- DTO - controller request
- VO - controller response

## 错误处理

- controller
  - 调用service
  - 错误类型：
    - panic，框架最外层捕获，无需处理
    - 参数校验错误，无需调用service，附上http status code（400），封装具体业务错误码，返回接口响应
    - 调用service报错，根据错误类型，判断http status code（4xx；500）之后，封装具体业务错误码，返回接口响应

- service
  - 调用repository
  - 错误类型：
    - panic，框架最外层捕获，无需处理
    - repository抛上来的错误
    - 业务本身产生的错误，直接抛给controller

- repository
  - 数据库操作，一般不含业务逻辑
  - 错误类型：
    - panic，框架最外层捕获，无需处理
    - 数据库操作错误，直接抛给调用方，无需处理

- http status code 和 封装具体业务错误码 仅仅在controller这一层处理（因为controller可能会调用不同业务模块的service，service无法确定调用方的具体业务模块）
- repository这一层的错误直接抛给调用方service

https://github.com/go-gorm/gorm/blob/master/errors.go

## validator

[go-playground validator](https://pkg.go.dev/github.com/go-playground/validator#hdr-Baked_In_Validators_and_Tags)

[]()

## TODO

- [X] 依赖反转（Dependency inversion principle，DIP）
- [ ] 控制反转（Inversion of Control，IoC）
