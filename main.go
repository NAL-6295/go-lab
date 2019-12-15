package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/NAL-6295")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	data := new(userJson)
	if err := json.Unmarshal(b, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	repourl := data.ReposURL

	reporesp, err := http.Get(repourl)
	if err != nil {
		return
	}
	defer reporesp.Body.Close()
	repoBytes, err := ioutil.ReadAll(reporesp.Body)
	if err != nil {
		return
	}

	repos := new([]user_repos)

	if err := json.Unmarshal(repoBytes, repos); err != nil {
		return
	}

	for _, repo := range *repos {
		fmt.Println(repo.Name)
		fmt.Println(repo.URL)
		fmt.Println(repo.Description)

	}

}
