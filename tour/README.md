# 命令行程序编写

该章节主要编写了三个命令行程序，意在熟悉外部包`cobro`的使用方法，另外提供了命令行程序项目布局的思路



#### 项目布局

根目录所含文件及文件夹

- cmd
- pkg
- internal
- main.go

`pkg` 文件夹下包含一些通用的公共组建，例如配置管理，数据库连接，日志写入等。这些组件个人编写，可以复用在任何项目中，在该demo中pkg文件夹下没有文件



`internal` 文件夹下包含不对外开放的包的源码，供cmd中的程序调用



`cmd`下包含命令行的定义，其中`root.go`文件提供一个根命令，其他命令只需要在该`init()`方法中注册即可。

```go
package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(sqlCmd)
  
  // register sub-command
  // Use AddCommand 
}
```



这里给一个字命令定义,以供参考, 声明一个Command的命令，Run参数是调用该command之后的函数操作。

在 `init()` 方法里面用将变量与Flag绑定。

`StringVarP()`参数5个参数， 第一个参数为绑定变量的指针，第二个为命令名称，第三个为命令名称的简称，第四个为默认参数，第五个为命令注释。

```go
var str string
var mode int8

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "transfer word format",
	Long:  "support many word formats trasfer",
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUppperCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderScore(str)
		default:
			log.Fatalf("no support for this mode, please check help doc")
		}
		log.Printf("Output: %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "please input string")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "please input mode of transferation")
}
```











