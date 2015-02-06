# godoc
godoc命令说明

遵守几个规则：
　注释需要紧贴对应的包声明和函数之前，不能有空行。
　如果注释内要有空行，应该使用空白注释行。
　开发者可以直接使用//BUG(author):的方式记录该代码片段中的遗留问题。

实际操作例子：
最常用的：
　godoc --http=":8080"
　这样就打开了跟golang.org一样的网站，如果不能访问golang.org的时候就可以这样访问，而且建议一般平时就这样访问速度更快。
　并且pkg里面会有所有安装过的包，如：原来  go get *****　或者go install 到$GOROOT/src/pkg中也是可以查看到的。

终端下使用：
　godoc fmt Println 
　就是打印fmt.Println这个函数的使用方法，如果习惯在终端命令行下可以如此使用

查看外部的包:
  godoc -http=":8080" -path="."
  因为godoc默认是去$GOROOT/src/pkg读取的，如果有project不在这个下面的，我们可以使用path
