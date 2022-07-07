# Hexagonal Architecture with GORM

[中文文档](README_zh.md)

## Package

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

## Demo

```
golangci-lint run ./...
go generate ./...
go test -count=1 -cover ./...
```


## Gorm vs MySQL

Refer: [https://dev.mysql.com/doc/refman/8.0/en/integer-types.html](https://dev.mysql.com/doc/refman/8.0/en/integer-types.html)

| gorm | mysql | byte size | number range
| --- | --- | --- | ---
| uint8 | tinyint(4) unsigned | 1 | 0-255
| uint16 | smallint(6) unsigned | 2 | 0-65535
| - | mediumint(9) unsigned | 3 | 0-16777215
| uint32 | int(11) unsigned | 4 | 0-4294967295
| uint64 | bigint(21) unsigned | 8 | 0-18446744073709551615

Note: the mysql type int(M) M indicates the maximum display width, regardless of how much storage space is occupied

mysql string type:

fixed | field | byte size | number of characters
--- | --- | --- | ---
fixed | char(n) | 1 | 0-255
not fixed | varchar(n) | 2 | 0-65535
not fixed | tinytext | 1 | 0-255
not fixed | text | 2 | 0-65535
not fixed | mediumtext | 3 | 0-16777215
not fixed | longtext | 4 | 0-4294967295

mysql time type:

field | byte size | meaning
--- | --- | ---
date | 3 | date
time | 3 | time
datetime | 8 | time
timestamp | 4 | time
year | 1 | year

## usage scenarios of the parameters' model.xxxx 'and' entity.xxxx '

- service use`entity.Xxxx`，The parameters of the business logic may mix data from different tables, using entities
- repository use`model.Xxxx`，Data operations operate directly on a single table, using Model

## DTO、VO

- DTO - controller request
- VO - controller response

## TODO

- [X] Dependency inversion principle，DIP
- [ ] Inversion of Control，IoC
