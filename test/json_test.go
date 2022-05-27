package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

// User is test for json
type AUser struct {
	ID   string
	Name string
}

func TestJson(t *testing.T) {
	u := AUser{ID: "user001", Name: "tom"}
	jsonU, _ := json.Marshal(u)
	fmt.Println(string(jsonU))

	var v AUser
	json.Unmarshal(jsonU, &v)
	fmt.Println(v)
}
