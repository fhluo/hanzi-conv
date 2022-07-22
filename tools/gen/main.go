package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	outputFilename string // 输出文件名
	packageName    string // 包名
	variableName   string // 变量名
)

func init() {
	rootCmd.Flags().StringVarP(&outputFilename, "out", "o", "dict.go", "输出文件名")
	rootCmd.Flags().StringVarP(&packageName, "pkg", "p", "dict", "包名")
	rootCmd.Flags().StringVarP(&variableName, "var", "v", "dict", "变量名")
}

// check 检查是否发生错误，如果发生错误，打印错误信息并停止程序的执行
func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "gen",
	Short: "从文本文件中生成以 .go 格式的字典文件",
	Run: func(cmd *cobra.Command, args []string) {
		dictFile := NewDictFile(packageName, variableName)
		check(dictFile.UpdateDict(args...))

		text, err := dictFile.Generate()
		check(err)

		check(os.WriteFile(outputFilename, []byte(text), 0666))
	},
}

func main() {
	check(rootCmd.Execute())
}
