# glualor

A Web Framework For Gopher Lua

## required

```shell
go get github.com/yuin/gopher-lua
go get github.com/zhu327/gluatemplate
go get github.com/zhu327/gluaweb
go get layeh.com/gopher-json
```

## run

```shell
go build main.go
./main
```

## lor

<http://lor.sumory.com/>

## description

这是一个在Gopher Lua中运行OpenResty lor框架的实例

```
├─app // 来至 github.com/lorlabs/lor-example
│  ├─config
│  ├─middleware
│  ├─model
│  ├─routes
│  ├─static
│  │  ├─css
│  │  └─js
│  └─views
├─lor // 来至 github.com/sumory/lor
│  └─lib
│      ├─middleware
│      ├─router
│      └─utils
└─pool // lua pool 库
```
