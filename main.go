package main

import (
	"os"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
	"fmt"
)

var BaseConfig map[string]string
const CONFIGNAME = "config.dat"

func main() {
	initConfig(CONFIGNAME)

	//os.Mkdir(BaseConfig["ROOT_PATH"], 777)

	prefixList := strings.Split(BaseConfig["ALLOW_PREFIX"], ",")


	for _, prefix := range prefixList {
		p := fmt.Sprintf("/%s/", prefix)
		path := fmt.Sprintf("%s/%s", BaseConfig["ROOT_PATH"], prefix)
		http.Handle(p, http.StripPrefix(p, http.FileServer(http.Dir(path))))
	}

	err := http.ListenAndServe(":8899", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func initConfig(file string) map[string]string {
	fp, _ := os.OpenFile(file, os.O_RDONLY, 0644)
	b, _ := ioutil.ReadAll(fp)

	s := string(b)

	BaseConfig = make(map[string]string)

	list := strings.Split(s, "\n")
	for _, l := range list {
		if strings.TrimSpace(l) == "" {
			continue
		}

		item := strings.Split(l, "=")
		if len(item) != 2 {
			continue
		}

		BaseConfig[trim(item[0])] = trim(item[1])
	}

	return BaseConfig
}

func trim(s string) string {
	return strings.Trim(strings.TrimSpace(s), `"`)
}
