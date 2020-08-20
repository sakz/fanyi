package config

type Config struct {
	Iciba         string `json:"iciba"`
	Youdao        string `json:"youdao"`
	Dictionaryapi string `json:"dictionaryapi"`
}

var SourceCfg Config

func init() {
	SourceCfg = Config{
		Iciba:         "http://dict-co.iciba.com/api/dictionary.php?key=D191EBD014295E913574E1EAF8E06666&w=${word}",
		Youdao:        "http://fanyi.youdao.com/openapi.do?keyfrom=node-fanyi&key=110811608&type=data&doctype=json&version=1.1&q=${word}",
		Dictionaryapi: "http://www.dictionaryapi.com/api/v1/references/collegiate/xml/${word}?key=82c5d495-ccf0-4e72-9051-5089e85c2975",
	}
}

//var filepath = "config/source.json"
//var SourceCfg Config
//
//func init() {
//	file, err := os.OpenFile(filepath, os.O_RDONLY, 0)
//	defer file.Close()
//	if err != nil {
//		log.Fatal(err)
//	}
//	data, err := ioutil.ReadAll(file)
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = json.Unmarshal(data, &SourceCfg)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
