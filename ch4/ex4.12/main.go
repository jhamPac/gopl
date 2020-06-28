package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

func (i *Index) build(fromID, toID int) {
	newInfos := make(map[int]*Info)
	for j := fromID; j < toID; j++ {
		if _, exists := i.Infos[j]; exists {
			continue
		}
		newInfo, err := newInfoFromURL(fmt.Sprintf(urlFormat, j))
		if err != nil {
			continue
		}
		newInfos[j] = newInfo
	}
	i.addInfos(newInfos)
}

func (i *Index) addInfos(infos map[int]*Info) {
	for id := range infos {
		i.Infos[id] = infos[id]
	}
}

func (i *Index) search(query string) []*Info {
	var foundInfos []*Info
	for _, info := range i.Infos {
		if strings.Contains(info.Title, query) || strings.Contains(info.Transcript, query) {
			foundInfos = append(foundInfos, info)
		}
	}
	return foundInfos
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
