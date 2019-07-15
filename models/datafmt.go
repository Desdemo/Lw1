package main

import (
	"Lw1/json"
)

type Person1 struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Birthday json.Time `json:"birthday"`
}

func main() {
	var a json.Time
	//a = json.Time(time.Now())
	//
	//p1 := Person1{
	//	Id:       1,
	//	Name:     "ddd",
	//	Birthday: json.Time(time.Now()),
	//}
	////println(json.Marshal(p1))
	println(a.String())
}
