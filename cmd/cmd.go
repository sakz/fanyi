package cmd

import (
	"fanyi/config"
	"fanyi/print"
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Execute() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s word\n\n", os.Args[0])
		flag.PrintDefaults()
		eg := `Examples:
  $ fanyi word
  $ fanyi world peace
  $ fanyi 中文`
		fmt.Println(eg)
	}
	flag.Parse()
	var queryString string
	if len(os.Args[1:]) == 0 {
		text, err := clipboard.ReadAll()
		if err != nil || text == "" {
			//读取剪切板失败或者没内容
			flag.Usage()
			return
		}
		fmt.Printf(" \n 默认读取剪贴板: %s\n", text)
		queryString = text
	} else {
		queryString = strings.Join(flag.Args(), " ")
	}
	queryString = url.QueryEscape(queryString)
	fmt.Println()
	ch := make(chan string)
	go youdao(queryString, ch)
	go iciba(queryString, ch)
	for i := 0; i < 2; i++ {
		<-ch
	}
}

func youdao(queryString string, ch chan<- string) {
	cfg := config.SourceCfg
	youdaoUrl := strings.Replace(cfg.Youdao, "${word}", queryString, 1)
	resp, err := http.Get(youdaoUrl)
	if err != nil {
		log.Println("有道翻译接口问题")
		ch <- "youdao failed"
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	print.Youdao(data)
	ch <- "youdao done"
}

func iciba(queryString string, ch chan<- string) {
	cfg := config.SourceCfg
	icibaUrl := strings.Replace(cfg.Iciba, "${word}", queryString, 1)
	resp, err := http.Get(icibaUrl)
	if err != nil {
		log.Println("iciba翻译接口问题")
		ch <- "iciba failed"
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	print.Iciba(data)
	ch <- "iciba done"
}
