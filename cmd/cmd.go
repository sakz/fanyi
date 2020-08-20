package cmd

import (
	"fanyi/config"
	"fanyi/print"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Execute() {
	args := os.Args[1:]
	queryString := strings.Join(args, " ")
	queryString = url.QueryEscape(queryString)
	youdao(queryString)
}

func youdao(queryString string) {
	cfg := config.SourceCfg
	youdaoUrl := strings.Replace(cfg.Youdao, "${word}", queryString, 1)
	resp, err := http.Get(youdaoUrl)
	if err != nil {
		log.Println("有道翻译接口问题")
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	print.Youdao(data)
}
