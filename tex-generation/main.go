package main

import (
	"fmt"
	"github.com/naoina/toml"
	"io/ioutil"
	"os"
	"text/template"
)

type Personal struct {
	Firstname, Lastname, Mobile, Mail, Addr1, Addr2, Add1 string
}

type Experience struct {
	Date, Title, Dur, Exc, Org, Place, Sub, Add string
}

type Language struct {
	Title, Level string
}

type Interest struct {
	Title, Exc string
}

type Skill struct {
	Category            string
	Concept, Tech, Lang []string
}

type Resume struct {
	Name, Job, Exc, Exctex string
	Tags                   []string
	Personal               Personal
	Experiences            []Experience
	Volunteering           []Experience
	Educations             []Experience
	Languages              []Language
	Skills                 []Skill
	Interests              []Interest
}

func GetConfigFromFile(filename string, config interface{}) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Cannot read file %q", filename)
	}
	return GetConfig(b, config)
}

func GetConfig(b []byte, config interface{}) error {
	return toml.Unmarshal(b, config)
}

func main() {
	data := Resume{}
	err1 := GetConfigFromFile("../data/resume.toml", &data)
	if err1 != nil {
		panic(err1)
	}

	t, err2 := template.New("main.tex").Delims("#(", ")#").ParseFiles("main.tex")
	if err2 != nil {
		panic(err2)
	}

	// publish .tex resume
	file1, err3 := os.Create("../static/resume.tex")
	defer file1.Close()
	if err3 != nil {
		panic(err3)
	}

	err4 := t.Execute(file1, data)
	if err4 != nil {
		panic(err4)
	}

	//publish .toml resume
	file2, err5 := os.Create("../static/resume.toml")
	defer file2.Close()
	if err5 != nil {
		panic(err5)
	}

	tomlByte, err6 := toml.Marshal(data)
	if err6 != nil {
		panic(err6)
	}

	file2.Write(tomlByte)

}
