package print

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/gookit/color"
	"log"
	"regexp"
	"strings"
)

func Youdao(data []byte) {
	magenta := color.FgMagenta.Render
	gray := color.FgGray.Render
	green := color.FgGreen.Render
	cyan := color.FgCyan.Render

	json, err := simplejson.NewJson(data)
	if err != nil {
		log.Fatal(err)
	}
	query, _ := json.Get("query").String()
	phonetic, err := json.Get("basic").Get("phonetic").String()
	var phoneticStr string
	if err != nil {
		phoneticStr = ""
	} else {
		phoneticStr = fmt.Sprintf("[ %s ]", magenta(phonetic))
	}
	fmt.Printf("%s %s %s\n\n", query, phoneticStr, gray("~  fanyi.youdao.com"))
	explains, _ := json.Get("basic").Get("explains").Array()
	for _, value := range explains {
		fmt.Printf("%s%s\n", gray("- "), green(value))
	}
	fmt.Println()
	web, _ := json.Get("web").Array()
	for i, value := range web {
		val := value.(map[string]interface{})
		line1 := fmt.Sprintf("%d. %s", i+1, highlight(val["key"].(string), query))
		fmt.Println(line1)
		valuelen := len(val["value"].([]interface{}))
		valArr := make([]string, valuelen, valuelen)
		for i, value := range val["value"].([]interface{}) {
			valArr[i] = value.(string)
		}
		valueStr := strings.Join(valArr, ", ")
		fmt.Printf("   %s\n", cyan(valueStr))
	}
}

func highlight(str string, query string) string {
	yellow := color.FgYellow.Render
	r := regexp.MustCompile("(?i)" + query)
	f := func(s string) string {
		return yellow(s)
	}
	res := r.ReplaceAllStringFunc(str, f)
	return res
}
