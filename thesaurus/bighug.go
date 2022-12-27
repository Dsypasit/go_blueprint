package thesaurus

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

type BigHuge struct {
	APIKey string
}

type synonyms struct {
	Noun *words `json:"none"`
	Verb *words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}

func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string
	response, err := http.Get("https://words.bighugelabs.com/api/2/" + b.APIKey + "/" + term + "/json")
	if err != nil {
		return syns, errors.New("bighuge: failed when looking for synonyms for " + term + ": " + err.Error())
	}
	var data synonyms
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}
	if data.Noun != nil {
		syns = append(syns, data.Noun.Syn...)
	}
	if data.Verb != nil {
		syns = append(syns, data.Verb.Syn...)
	}
	return syns, nil
}

func main() {
	godotenv.Load()
	fmt.Println("vim-go")
}
