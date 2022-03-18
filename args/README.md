# Args

## 需求描述

### 英文原文

Command-Line Argument Parser

Most of us have had to parse command-line arguments from time to time. If we don’t have a convenient utility, then we simply walk the array of strings that is passed into the `main` function. There are several good utilities available from various sources,but none of them do exactly what I want. So, of course, I decided to write my own. I call it: `Args`.

`Args` is very simple to use. You simply construct the Args class with the input arguments and a format string, and then query the Args instance for the values of the arguments. Consider the following simple example:

```java
public static void main(String[] args) { 
    try {
        Args arg = new Args("l,p#,d*", args);
        boolean logging = arg.getBoolean('l');
        int port = arg.getInt('p');
        String directory = arg.getString('d');
        executeApplication(logging, port, directory);
    } catch (ArgsException e) {
        System.out.printf("Argument error: %s\n", e.errorMessage());
    } 
}
```

We just create an instance of the Args class with two parameters. The first parameter is the format, or schema, string: "l,p#,d*." It defines three command-line arguments. The first, –l, is a boolean argument. The second, -p, is an integer argument. The third, -d, is a string argument. The second parameter to the Args constructor is simply the array of command-line argument passed into main.

If the constructor returns without throwing an ArgsException, then the incoming command-line was parsed, and the Args instance is ready to be queried. Methods like getBoolean, getInteger, and getString allow us to access the values of the arguments by their names.
If there is a problem, either in the format string or in the command-line arguments themselves, an ArgsException will be thrown. A convenient description of what went wrong can be retrieved from the errorMessage method of the exception.

### 中文翻译

我们中的大多数人都不得不时不时地解析一下命令行参数。如果我们没有一个方便的工具，那么我们就简单地处理一下传入main函数的字符串数组。有很多开源工具可以完成这个任务，但它们可能并不能完全满足我们的要求。所以我们再写一个吧.

传递给程序的参数由标志和值组成。标志应该是一个字符，前面有一个减号。每个标志都应该有零个或多个与之相关的值。
例如:`-l -p 8080 -d /usr/logs`其中“l”（日志）没有相关的值，它是一个布尔标志，如果存在则为 true，不存在则为 false。“p”（端口）有一个整数值“d”（目录）有一个字符串值。

标志后面如果存在多个值，则该标志表示一个列表：`-g this is a list -d 1 2 -3 5`　"g"表示一个字符串列表[“this”, “is”, “a”, “list”]，“d"标志表示一个整数列表[1, 2, -3, 5]。

如果参数中没有指定某个标志，那么解析器应该指定一个默认值。例如，false 代表布尔值，0 代表数字，”"代表字符串，[]代表列表。
如果给出的参数与模式不匹配，重要的是给出一个好的错误信息，准确地解释什么是错误的。　
确保你的代码是可扩展的，即如何增加新的数值类型是直接和明显的。