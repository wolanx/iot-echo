# c use go

- Go 与 C 的桥梁：cgo 入门，剖析与实践 包含数据结构 https://zhuanlan.zhihu.com/p/349197066
- Go语言高级编程(Advanced Go Programming) https://chai2010.cn/advanced-go-programming-book/ch2-cgo/ch2-01-hello-cgo.html

```sh
# 生成 .so .h
go build -o libddd.so -buildmode=c-shared main.go

# c 连接 lib
gcc test.c -L ./ -lddd -o test

# 默认使用 /usr/lib，临时env
LD_LIBRARY_PATH=$LD_LIBRARY_PATH:. ./test
```

# go use c

//#cgo CFLAGS: -I./
//#cgo LDFLAGS: -L./ -lddd
//#include "libddd.h"
CFLAGS .h 路径
LDFLAGS .so 路径

```sh
LD_LIBRARY_PATH=$LD_LIBRARY_PATH:. go run use_so.go
```
