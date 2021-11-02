package info

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

/**
* Structs to be used in the json file
*/
type Json struct {
	AboutMe AboutMe `json:"about_me"`
	Personal Personal `json:"personal"`
	TechnicalInfo TechnicalInfo `json:"technical_info"`
	StudentInfo StudentInfo `json:"student_info"`
}

type AboutMe struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Specialty string `json:"specialty"`
	Nacionality string `json:"nacionality"`
}

type Personal struct {
	BirthDay string `json:"birth_day"`
	FavoriteFood string `json:"favorite_food"`
	FavoriteSubject string `json:"favorite_subject"`
}

type TechnicalInfo struct {
	MostUsedLanguages []string `json:"most_used_languages"`
	MostUsedFrameworks []struct {
		Framework string `json:"framework"`
		Lang string `json:"lang"`
	} `json:"most_used_frameworks"`
}

type StudentInfo struct {
	Olympics []struct {
		Initials string `json:"initials"`
		Name string `json:"name"`
	} `json:"olympics"`
}

/**
* GetAllInfo - Get all the info from the json file
*/
func GetAllInfo() Json {
	jsonFile, err := os.Open("internal/info/info.json")
	
	if err != nil {
		panic(err.Error())
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var full Json

	err = json.Unmarshal(byteValue, &full)

	if err != nil {
		panic(err.Error())
	}

	return full
}
