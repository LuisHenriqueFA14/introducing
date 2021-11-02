package api

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/LuisHenriqueFA14/introducing/internal/info"
)

/**
* toString - Converts a struct to a string
*/
func toString(jsonStruct interface{}) string {
	str, err := json.Marshal(jsonStruct)

	if err != nil {
		return err.Error()
	}

	return string(str)
}

/**
* Full - Route to return all information
*/
func Full(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, toString(info.GetAllInfo()))
}

/**
* About - Route to return information about me
*/
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, toString(info.GetAllInfo().AboutMe))
}

/**
* Personal - Route to return personal information about me
*/
func Personal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, toString(info.GetAllInfo().Personal))
}

/**
* Technical - Route to return technical information about me
*/
func Technical(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, toString(info.GetAllInfo().TechnicalInfo))
}

/**
* Student - Route to return student information about me
*/
func Student(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, toString(info.GetAllInfo().StudentInfo))
}
