package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"text/template"
)

var wg sync.WaitGroup

type UserInfoAgg struct {
	ID   string
	Info map[string]Attributes
}

type Attributes struct {
	Slug       string   `json:"slug"`
	AliasNames []string `json:"alias_names"`
	Animagus   string   `json:"animagus"`
	Name       string   `json:"name"`
	Gender     string   `json:"gender"`
}

type Character struct {
	ID   string     `json:"id"`
	Type string     `json:"type"`
	Info Attributes `json:"attributes"`
}

type response struct {
	Data []Character `json:"data"`
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("https://api.potterdb.com/v1/characters")
	// defer resp.Body.Close()
	usersInfo := make(map[string]Attributes)

	bytes, _ := io.ReadAll(resp.Body)

	var hResponse response
	if err := json.Unmarshal(bytes, &hResponse); err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, character := range hResponse.Data {
		usersInfo[character.ID] = Attributes{
			Slug:       character.Info.Slug,
			Name:       character.Info.Name,
			Gender:     character.Info.Gender,
			AliasNames: character.Info.AliasNames,
			Animagus:   character.Info.Animagus,
		}
	}

	p := UserInfoAgg{Info: usersInfo}
	// fmt.Println(p)
	t, _ := template.ParseFiles("basic_template.html")
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi there")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/users", usersHandler)
	http.ListenAndServe(":9090", nil)
}
