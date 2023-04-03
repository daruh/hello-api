package main

import (
	"gopkg.in/yaml.v2"
	"hello-api/handlers/rest"
	"io/ioutil"
	"log"
	"net/http"
)

const CONFIG_FILE = "/etc/config/configmap-microservice-demo.yaml"

type Config struct {
	Message string `yaml:"message"`
}

func check(err error) {
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func loadConfig(configFile string) *Config {
	conf := &Config{}
	configData, err := ioutil.ReadFile(configFile)
	check(err)

	err = yaml.Unmarshal(configData, conf)
	check(err)
	log.Println(conf)
	return conf
}

func main() {

	addr := ":8083"

	mux := http.NewServeMux()

	mux.HandleFunc("/translate/hello", rest.TranslateHandler)

	log.Printf("listening on %s\n", addr)
	log.Println("config map demo")
	loadConfig(CONFIG_FILE)

	log.Fatal(http.ListenAndServe(addr, mux))
}

type Resp struct { // <6>
	Language    string `json:"language"`
	Translation string `json:"translation"`
}
