# Go 入门

## module的组织

### 综述：管理依赖的复杂度

拓展
维护

### 用法

#### 项目初始化

【注意，不是路径，而是项目顶层文件夹名称】

- `go mod init xxx_dir`
- `go mod tidy




## 目录的组织

小项目考虑 pkg、cmd、interal，3 个目录即可。

### 综述：管理目录的复杂度

### Go 项目目录需要规避的做法

- 不建议使用 src 目录。原因是Go项目会放在`$GOPATH/src`中，使用 src 会导致路径中有两个src；
- 不建议使用复数，统一使用单数；

### 平铺式目录和结构化目录

当项目是一个架构或库的时候，可以考虑平铺式目录。在项目根目录存放代码，目录看起来层级少。优点是引用的路径短。如[glog](https://github.com/golang/glog)的目录。

``` shell

LICENSE                 glog_file_linux.go      glog_test.go
README.md               glog_file_nonwindows.go glog_vmodule_test.go
glog.go                 glog_file_other.go      go.mod
glog_bench_test.go      glog_file_posix.go      go.sum
glog_context_test.go    glog_file_windows.go    internal
glog_file.go            glog_flags.go

```





### 以功能拆分

### 以类型拆分

## 代码的组织

### 综述：管理代码的复杂度

群论

### 以类型拆分

#### 面向对象


#### 设计模式

### error 的问题

### goroutin
