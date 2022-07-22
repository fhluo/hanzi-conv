package main

import (
	"github.com/fhluo/zhconv/cli"
	"github.com/fhluo/zhconv/conv"
	"github.com/fhluo/zhconv/dict/st"
)

func main() {
	STConv := conv.New()
	STConv.UpdateDict(st.Characters, st.Phrases)
	cli.Execute(STConv, "s2t", "简体中文 -> 繁体中文")
}
