package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
)

var (
	//go:embed dict.tmpl
	tmplStr string
	tmpl    *template.Template
)

func init() {
	// 解析模板
	tmpl = template.New("dict")
	_, err := tmpl.Parse(tmplStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// LoadDict 从指定文件中读取字典
func LoadDict(filename string) (map[string]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	text := string(data)

	// 将文本按行分割
	lines := regexp.MustCompile(`\r?\n`).Split(text, -1)
	spacesRE := regexp.MustCompile(`\s`)

	dict := make(map[string]string)
	for _, line := range lines {
		// 将一行按空格分割，格式符合要求则加入字典
		items := spacesRE.Split(strings.TrimSpace(line), -1)
		if len(items) >= 2 {
			dict[items[0]] = items[1]
		}
	}

	return dict, nil
}

// DictFile Go 语言 map 类型格式字典文件生成模板参数
type DictFile struct {
	PackageName  string
	VariableName string
	Dictionary   map[string]string
}

func NewDictFile(pkgName, varName string) *DictFile {
	return &DictFile{
		PackageName:  pkgName,
		VariableName: varName,
		Dictionary:   make(map[string]string),
	}
}

// UpdateDict 更新字典
func (d *DictFile) UpdateDict(filenames ...string) error {
	for _, filename := range filenames {
		dict, err := LoadDict(filename)
		if err != nil {
			return err
		}
		for k, v := range dict {
			d.Dictionary[k] = v
		}
	}
	return nil
}

// Generate 执行模板，生成 .go 形式的字典文件
func (d *DictFile) Generate() (string, error) {
	buffer := new(bytes.Buffer)
	err := tmpl.Execute(buffer, d)
	return buffer.String(), err
}
