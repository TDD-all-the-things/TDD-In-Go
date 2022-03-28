## API构思

2022.03.20与徐昊老师在TDD群里沟通发现API设计没有做好,主要原因还是没有理解好需求.

当前API如下:

```go
Parse(option *Option, args ...string)
```
因老师在[TDD演示（1）：任务分解法与整体工作流程](https://time.geekbang.org/column/article/494207)中用Java演示时出现Option类所以想当然就这么设计了.与老师在群里讨论+重新思考后发现,**理想**API如下:

```go
Parse(type Type, args ...string) (val Value, err error)
```
其中`type`表示任意数据的类型信息,`value`表示具有`type`类型数据的一个值,该值的填充过程取决于可变参数`args`中的信息,`err`表示在此过程中发生的任何错误.

在Go中是否可以做到呢?答案是,可以.看如下真实API:
```go
Parse(typ reflect.Type, args ...string) (val reflect.Value, err error)
```
但是,但是,但是使用你的API还需要依赖`reflect`包太不合理了.是否可以只使用核心语法呢?可以,看如下**折中**API:
```go
Parse(typ interface{}, args ...string) (val interface{}, err error)
```
将数据的类型信息装入`interface{}`中,将填充后的值也装入`interface{}`这样就不需要依赖`refect`包.此法可行但还是有易用性问题,看如下客户端调用伪代码:
```go
var option Option
val, err := Parse(&option, args)
if err != nil {
   //错误处理
}
realValue, ok := val.(&option)
```
上面这个代码有如下问题:
1. 将`&option`作为类型信息传入,将`option`作为类型信息传入不行吗?可以在`Parse`内部判断并报错来引导用户.
2. 返回值`val`是接口类型,需要用断言语法来获取真实数据——需要提供正确的类型信息,否则断言会失败,那么`realValue`到底是值还是指针呢?光靠写注释规定API使用方式是Go官方团队的特权,其他人这么做会喷.
3. 数据拷贝问题,如果传入`option`而其所占空间比较大怎么办?返回值`val`中装的也是一个占用空间比较大的值,那么这就造成了空间的浪费!

怎么办呢?借鉴Go标准库的API设计,比如官方`json`库:
```go
json.Unmarshal(data []byte, v any)
```
`json.Unmarshal`将`data`中的信息反序列化为`v`,`v`为`any`即`interface{}`类型.其内部做了检查——`v`中的值不能为空且必须是指针类型,这样设计API就可以解决上面那三个问题.

**最终**API如下:
```go
Parse(v interface{}, args...string) error
```
客户端使用方式如下:
```go
// 自定义类型option
var option Option
err := args.Parse(&option, "-l", "-p", "8080", "-d", "/usr/logs")
if err != nil {
   // 错误处理
}
// 直接操作option即可

// 自定义类型ListOption
var list ListOption
err := args.Parse(&list, "-g", "is", "list", "-d", "1", "2")
if err != nil {
   // 错误处理
}
// 直接操作list即可
```
`Option`和`ListOption`是自定义数据类型,`Parse`是如何知道要将哪个参数解析到数据值的哪个字段上的呢?答案是,结构体Tag.看如下自定义数据类型的定义:
```go
type Option struct {
   Logging   bool   `args:"l"`
   Port      int    `args:"p"`
   Directory string `args:"d"`
}

type ListOption struct {
   Group  []string `args:"g"`
   Digits []int    `args:"d"`
}
```
至此API构思才算真正完成!本想推倒重来的,现在就将其视作“具有测试的遗留代码”,然后一步一步重构到最终的API吧.

## 需求拆分

### 任务列表

1. 常规Option
2. 列表Option
3. 可扩展支持其他Option

### 测试列表

1. Single Option
    - ~~case 1~~
      - ~~`-l` bool~~
      - ~~`-p 8080` int~~
      - ~~`-d /usr/logs` string~~
    - case 2
      - `-g this is a list` []stirng
      - `-d 1 3 2 4` []int
2. Multiple Options
   - ~~`-l -p 8080 -d /usr/logs`~~
   - `-g this is a list -d 1 3 2 4`
3. Sad path
   - bool `-l t` `-l t f`
   - int `-p` `-p 8080 8081`
   - string `-d /usr/logs /usr/vars`
4. ~~Default Values~~
   - ~~bool `false`~~
   - ~~int `0`~~
   - ~~string `""`~~
5. Handle Error
   - v in args.Parse is not an pointer of struct?
   - v is CanSet?CanAddr?
   - v's field is unexported?
   - v's field CanSet?
   - ~~v's field don't have "args" tag~~
   - ~~v's field has "args" tag, but it's empty~~ 处理逻辑同缺失
   - func parserOption , if parser not found in PARSERS map?
   - ~~func Parse in OptionParser, parseValueFunc handle return error~~

### 调整测试

1. ~~BoolOptionParserTest~~
   - ~~sad path~~
     - `-l t`
     - `-l t f`
   - ~~default value~~
     - false
2. ~~SingleValuedOptionParserTest~~
   - ~~int~~
     - ~~sad path~~
       - int `-p` 
       - int `-p 8080 8081`
     - ~~default value~~
       - 0
   - ~~string~~
     - ~~sad path~~
       - string, `-d`
       - string, `-d /usr/logs /usr/vars`
     - ~~default value~~
       - ""