package main

import (
	"encoding/json"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"io/ioutil"
)

type Tutorial struct {
	Steps    []Step `json:"steps"`
	ID       string `json:"id"`
	Title    string `json:"title"`
	Language string `json:"language"`
}

type Step struct {
	Instruction string `json:"instruction"`
	Spec        string `json:"spec"`
	Source      string `json:"source"`
}

func renderTutorialIndex(ren render.Render) {
	ren.JSON(200, []map[string]string{{"id": "1", "title": "FizzBuzz"}})
}

func renderTutorial(ren render.Render) {
	var tutorial Tutorial
	file, err := ioutil.ReadFile("neutrino.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &tutorial)
	ren.JSON(200, tutorial)
}

func main() {
	m := martini.Classic()
	renderOptions := render.Options{
		Directory:  "public/templates",
		Extensions: []string{".tmpl", ".html"},
	}
	m.Use(render.Renderer(renderOptions))
	m.Get("/tutorials/:id", renderTutorial)
	m.Get("/tutorials", renderTutorialIndex)
	m.Run()
}
