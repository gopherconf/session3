package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main(){

	defer fmt.Println("World...")
	defer fmt.Println("World1...")
	defer fmt.Println("World2...")
	defer fmt.Println("World3...")
	defer fmt.Println("World4...")

	fmt.Println("Hello ...")

	s := struct {
		Name string `json:"name"`
		Family string
		Age int `json:"sen"`
	}{
		Name:   "GopherConf",
		Family: "Gopher",
		Age:    1,
	}

	out,err := json.Marshal(s)
	if err != nil{
		panic(err)
	}

	_ = ioutil.WriteFile("./sample.json",out,os.ModePerm)
}
