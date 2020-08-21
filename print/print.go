package print

import (
	"encoding/xml"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/gookit/color"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type IcibaResp struct {
	//xmlName     xml.Name `xml:"dict"`
	Key         string   `xml:"key"`
	Ps          []string `xml:"ps"`
	Pron        []string `xml:"pron"`
	Pos         []string `xml:"pos"`
	Acceptation []string `xml:"acceptation"`
	Sent        []Sent   `xml:"sent"`
}

type Sent struct {
	Orig  string `xml:"orig"`
	Trans string `xml:"trans"`
}

var magenta = color.FgMagenta.Render
var gray = color.FgGray.Render
var green = color.FgGreen.Render
var cyan = color.FgCyan.Render

func Youdao(data []byte) {
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
	fmt.Printf(" %s %s %s\n\n", query, phoneticStr, gray("~  fanyi.youdao.com"))
	explains, _ := json.Get("basic").Get("explains").Array()
	for _, value := range explains {
		fmt.Printf(" %s %s\n", gray("-"), green(value))
	}
	fmt.Println()
	web, _ := json.Get("web").Array()
	for i, value := range web {
		val := value.(map[string]interface{})
		fmt.Printf(" %s %s\n", gray(strconv.Itoa(i+1)+"."), highlight(val["key"].(string), query))
		valuelen := len(val["value"].([]interface{}))
		valArr := make([]string, valuelen)
		for i, value := range val["value"].([]interface{}) {
			valArr[i] = value.(string)
		}
		valueStr := strings.Join(valArr, ", ")
		fmt.Printf("    %s\n", cyan(valueStr))
	}
	fmt.Println()
	fmt.Println(gray("   --------"))
	fmt.Println()
}

func Iciba(data []byte) {
	v := IcibaResp{}
	err := xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	var phoneticStr string
	for i, value := range v.Ps {
		if i == 0 {
			phoneticStr += "英" + "[ " + value + "] "
		} else {
			phoneticStr += "美" + "[ " + value + "] "
		}
	}
	fmt.Printf(" %s %s %s\n\n", v.Key, magenta(phoneticStr), gray("~  iciba.com"))
	if !isChinese(v.Key) {
		for i := 0; i < len(v.Pos); i++ {
			fmt.Printf(" %s %s %s", gray("-"), green(v.Pos[i]), green(v.Acceptation[i]))
		}
	}
	fmt.Println()
	for i := 0; i < len(v.Sent); i++ {
		fmt.Printf(" %s %s\n", gray(strconv.Itoa(i+1)+"."), highlight(del(v.Sent[i].Orig), v.Key))
		fmt.Printf("    %s\n", cyan(del(v.Sent[i].Trans)))
	}
	fmt.Println()
	fmt.Println(gray("   --------"))
	fmt.Println()
}

// 高亮句子中的单词
func highlight(str string, query string) string {
	yellow := color.FgYellow.Render
	//r := regexp.MustCompile("(?i)" + query)
	//f := func(s string) string {
	//	return yellow(s)
	//}
	//res := r.ReplaceAllStringFunc(str, f)

	// 句子中单词用黄色，其他用灰色
	r := regexp.MustCompile("(?i)" + "(.*)" + "(" + query + ")" + "(.*)")
	res1 := r.ReplaceAllString(str, "$1$2"+gray("$3"))
	r2 := regexp.MustCompile("(?i)" + "(.*?)" + "(" + query + ")")
	res2 := r2.ReplaceAllString(res1, gray("$1")+yellow("$2"))
	return res2
}

// 删除string中的换行符
func del(str string) string {
	r := regexp.MustCompile("\n")
	res := r.ReplaceAllString(str, "")
	return res
}

// 是否包含中文
func isChinese(str string) bool {
	count := 0
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	return count > 0
}
