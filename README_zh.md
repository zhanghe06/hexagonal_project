# 六边形架构（结合GORM）

[English Document](README.md)

## 项目依赖

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

## Gorm MySQL 字段对比

参考：[https://dev.mysql.com/doc/refman/8.0/en/integer-types.html](https://dev.mysql.com/doc/refman/8.0/en/integer-types.html)

gorm | mysql | 字节长度 | 数字范围
--- | --- | --- | ---
uint8 | tinyint(4) unsigned | 1 | 0-255
uint16 | smallint(6) unsigned | 2 | 0-65535
- | mediumint(9) unsigned | 3 | 0-16777215
uint32 | int(11) unsigned | 4 | 0-4294967295
uint64 | bigint(21) unsigned | 8 | 0-18446744073709551615

说明：mysql类型int(M) M 表示最大显示宽度，与所占多少存储空间并无任何关系

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

## 关于参数`model.Xxxx`和`entity.Xxxx`的使用场景

- service 使用`entity.Xxxx`，业务逻辑的参数可能混合不同表的数据，使用实体
- repository 使用`model.Xxxx`，数据操作直接对单表进行操作，使用model
