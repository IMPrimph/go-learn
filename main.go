package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

// sample data
// [    {
// 	"id": "20448207-20d2-48cb-85b4-d61ee8f50f97",
// 	"type": "character",
// 	"attributes": {
// 	  "slug": "1992-gryffindor-vs-slytherin-quidditch-match-spectators",
// 	  "alias_names": [],
// 	  "animagus": null,
// 	  "blood_status": null,
// 	  "boggart": null,
// 	  "born": null,
// 	  "died": null,
// 	  "eye_color": null,
// 	  "family_members": [],
// 	  "gender": null,
// 	  "hair_color": null,
// 	  "height": null,
// 	  "house": null,
// 	  "image": null,
// 	  "jobs": [],
// 	  "marital_status": null,
// 	  "name": "1992 Gryffindor vs Slytherin Quidditch match spectators",
// 	  "nationality": null,
// 	  "patronus": null,
// 	  "romances": [],
// 	  "skin_color": null,
// 	  "species": null,
// 	  "titles": [],
// 	  "wands": [],
// 	  "weight": null,
// 	  "wiki": "https://harrypotter.fandom.com/wiki/1992_Gryffindor_vs_Slytherin_Quidditch_match_spectators"
// 	},
// 	"links": {
// 	  "self": "/v1/characters/20448207-20d2-48cb-85b4-d61ee8f50f97"
// 	}
//   },]

type response struct {
	Data []Character `json:"data"`
}

func main() {
	resp, _ := http.Get("https://api.potterdb.com/v1/characters")
	defer resp.Body.Close()
	usersInfo := make(map[string]string)

	bytes, _ := io.ReadAll(resp.Body)

	var hResponse response
	if err := json.Unmarshal(bytes, &hResponse); err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, character := range hResponse.Data {
		usersInfo[character.Info.Name] = character.Info.Gender
	}

	fmt.Println(usersInfo)
}
