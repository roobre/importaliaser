package importaliaser

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Store interface {
	Config() Config
	Alias(string) (Alias, bool)
}

type Alias struct {
	Protocol string
	URI      string
}

type Config struct {
	RootURL             string
	Speculative         bool
	SpeculativeFormat   string
	SpeculativeProtocol string
}

type JSONStorer struct {
	JSONConfig Config `json:"config"`

	Aliases map[string]Alias
}

func NewJSONStorer(path string) Store {
	js := &JSONStorer{}

	jscontents, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jscontents, js)
	if err != nil {
		log.Fatal(err)
	}

	return js
}

func (js *JSONStorer) Alias(name string) (Alias, bool) {
	a, ok := js.Aliases[name]
	return a, ok
}

func (js *JSONStorer) Config() Config {
	return js.JSONConfig
}
