package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const outputfile = "output.json"
const urlFormat = "https://xkcd.com/%d/info.0.json"

type Index struct {
	Infos    map[int]*Info
	FilePath string
}

type Info struct {
	Title      string
	Transcript string
	ImgURL     string `json:"img"`
}

func newInfoFromURL(comicURL string) (*Info, error) {
	resp, err := http.Get(comicURL)
	if resp.StatusCode != http.StatusOK || err != nil {
		return nil, err
	}
	resp.Body.Close()

	var info Info
	if err = json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}
	return &info, nil
}

func newIndex(filePath string) *Index {
	var infos map[int]*Info

	out, err := ioutil.ReadFile(filePath)
	if err != nil {
		infos = make(map[int]*Info)
	} else {
		json.Unmarshal(out, &infos)
	}

	return &Index{
		Infos:    infos,
		FilePath: filePath,
	}
}
