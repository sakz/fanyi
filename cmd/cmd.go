package cmd

import (
	"fanyi/config"
	"fanyi/print"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func Execute() {
	args := os.Args[1:]
	queryString := strings.Join(args, " ")
	youdao(queryString)
}

func youdao(queryString string) {
	cfg := config.SourceCfg
	youdaoUrl := strings.Replace(cfg.Youdao, "${word}", queryString, 1)
	resp, err := http.Get(youdaoUrl)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(data))
	print.Youdao(data)
}
