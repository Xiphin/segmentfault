# SegmentFault光棍节(双11)程序员闯关秀通关工具
------
**编译环境：**
> golang && (windows|linux)

**使用方法(2种)：**

*获取代码:*
> go get github.com/Xiphin/segmentfault

 1. 编译成可执行程序
> * 进入 github.com/Xiphin/segmentfault 目录后 go build
> * 运行 segmentfault -sf 闯关链接，形如：
```
> segmentfault -sf https://1111.segmentfault.com
> segmentfault -sf https://1111.segmentfault.com?k=1573402aa6086d9ce42cfd5991027022
> segmentfault -sf https://1111.segmentfault.com/?k=1573402aa6086d9ce42cfd5991027022
```
  
 2. 直接 go run code.go
> * 进入 github.com/Xiphin/segmentfault 目录后 go run code.go -sf 闯关链接，形如：
```
> go run code.go -sf https://1111.segmentfault.com
> go run code.go -sf https://1111.segmentfault.com?k=1573402aa6086d9ce42cfd5991027022
> go run code.go -sf https://1111.segmentfault.com/?k=1573402aa6086d9ce42cfd5991027022
```
**运行结果类似如下：**
```
SegmentFault 1111 URL: https://1111.segmentfault.com

[=>]你从第 1 关开始的 =>
[=>]通往第 2 关的密码： 1d801c00d1f1b1c2b2ac850401468adf
[=>]通往第 3 关的密码： d3e11317c7d1cb58b27c043480d74b0f
[=>]通往第 4 关的密码： a87ff679a2f3e71d9181a67b7542122c
[=>]通往第 5 关的密码： e4da3b7fbbce2345d7772b0674a318d5
[=>]通往第 6 关的密码： bdbf46a337ac08e6b4677c2826519542
[=>]通往第 7 关的密码： 1573402aa6086d9ce42cfd5991027022
[=>]通往第 8 关的密码： d1a571d7ba84dcf2d146fb5b666efcb9
[=>]通往第 9 关的密码： fb7d71730d6af16078286a31b0b61357

[**]注意：第 9 关不能直接通过密码访问，需要制下面链接才能访问通关:

https://1111.segmentfault.com/?k=d1a571d7ba84dcf2d146fb5b666efcb9$post_k=fb7d71730d6af16078286a31b0b61357

[=>]通往第 10 关的文件已生成，请自行解压通关！

[**]你也可以复制下面链接直接通行第 10 关：

https://1111.segmentfault.com/?k=e4a4a96a69a1b2b530b3bec6734cdf52
```


