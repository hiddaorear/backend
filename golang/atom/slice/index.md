# slice

## `a := []string{}` 和 `a := make([]string{},0)` 哪种初始化方式更好？

二者都不够好。在考虑 `json.Marshal` 的情况下，使用 `var a []string` 最好。

```go
// 声明方式1
var a []string{}
// 等价下面的方式
var aH reflect.SliceHeader {
	Data: 0, // 没有分配内存空间
	Len: 0,
	Cap: 0,
}

// 声明方式2
b := []string{}
// 等价下面的方式
var bH reflect.SliceHeader {
	Data: zeroBase, // 分配了长度为0的内存空间，而不是没分配。如果切片对象为空，不管Len和Cap，也是zeroBase
	Len: 0,
	Cap: 0,
}
```

二者区别：

- 空切片：len和cap都是0，但值不等于nil，内存地址非空，是16进制的数字；
- nil切片：len和cap也是0，值是nil，内存地址是空指针；

`var a []string` 在go底层，如果 append，先做 nil 判断，然后执行 append。因此，空切片和nil切片，都可以 append。


> When declaring an empty slice, prefer

```go
var t []string
```

over

```go
t := []string{}
```

> The former declares a nil slice value, while the latter is non-nil but zero-length. They are functionally equivalent—their len and cap are both zero—but the nil slice is the preferred style.

> Note that there are limited circumstances where a non-nil but zero-length slice is preferred, such as when encoding JSON objects (a nil slice encodes to null, while []string{} encodes to the JSON array []).

> When designing interfaces, avoid making a distinction between a nil slice and a non-nil, zero-length slice, as this can lead to subtle programming errors.

> 来自[Declaring Empty Slices](https://go.dev/wiki/CodeReviewComments#declaring-empty-slices)

## 切片的底层实现带来的问题

- [Go编程模式：切片，接口，时间和性能](https://coolshell.cn/articles/21128.html)
