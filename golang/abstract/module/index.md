# Go 入门

# module的组织

## 综述

拓展
维护

## 用法

### 项目初始化

【注意，不是路径，而是项目顶层文件夹名称】

- `go mod init xxx_dir`
- `go mod tidy`


# 目录的组织

## 综述

基本的 Go 项目考虑 pkg、cmd、interal，3 个目录即可。

## Go 项目目录需要规避的做法

- 不建议使用 src 目录。原因是Go项目会放在`$GOPATH/src`中，使用 src 会导致路径中有两个src；
- 不建议使用复数，统一使用单数；

## 平铺式目录

当项目是一个架构或库的时候，可以考虑平铺式目录。在项目根目录存放代码，目录看起来层级少。优点是引用的路径短。如[glog](https://github.com/golang/glog)的目录。

``` shell

LICENSE                 glog_file_linux.go      glog_test.go
README.md               glog_file_nonwindows.go glog_vmodule_test.go
glog.go                 glog_file_other.go      go.mod
glog_bench_test.go      glog_file_posix.go      go.sum
glog_context_test.go    glog_file_windows.go    internal
glog_file.go            glog_flags.go

```


## 结构化目录

结构化目录大致有 2 种维度，一种是通用的架构分层（如：MVC），另一种是按照业务功能分层。按照架构分层，便于拓展架构，架构变动以后，业务可能需要修改。而按照业务功能分层，便于添加业务，新业务可能导致架构需要修改。Go 倾向于按照功能职责分层。

设有web服务，有业务A、B、C，每一个业务有：model，trpc，repository，service，validate。考察一下目录的组织结构。

### 按功能职责维度设计

![layout_by_business](./go_primer_img/layout_by_business.png)

当我们看业务逻辑的时候，按照功能职责的结构，代码聚合在一个目录，一目了然。当我们看A业务的逻辑的适合，从3个目录中找其中一个。记人查看代码只需进入1次目录。修改业务代码，也只需在一个目录下操作。

![layout_by_business_2](./go_primer_img/layout_by_business_2.png)

这是业务代码各自独立的场景。但实际业务代码之间，会有依赖。业务B 的service依赖A、B、C的model，这样就形成了平级之间的依赖。在 Go 项目中，不推荐这样的依赖。从依赖角度来看，A和C的model，实际上已经不仅仅这两个业务在使用，通常会把二者提取出来。这样业务多了，model作为上层目录，存放个业务的model会更适合。model放在独立上层目录以后，人查看代码，就不只进入一次目录。进入业务目录，进入model目录，至少2次。当从model目录看代码的时候，就容易有一个疑问，这些model文件是在哪些地方用了，需要搜索或依赖IDE的跳转到业务调用处。

当业务逐渐发展的时候，类似的目录聚合可能多起来，这样就容易形成，按照架构层次的维度去设计目录。

### 按架构层次维度设计

![layout_by_architecture](./go_primer_img/layout_by_architecture.png)

按照架构维度，目录有5个：model、repository、service、validate、trpc，每一个目录下有3个业务各自的独立文件。查看某一个业务代码的时候，就需要在5个目录中跳来跳去。



# 资料

- [目录设计-Go 面向包的设计和架构分层](https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md)
