package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var url  = flag.String("url","http://gopherconf.ir/","Url")
func main(){
	flag.Parse()

	// Make Request
	req,err := http.NewRequest("GET",*url,nil)
	if err != nil{
		log.Fatal(err)
	}

	//Do request
	client := http.Client{}
	response,err := client.Do(req)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Status:",response.StatusCode)

	body,err := ioutil.ReadAll(response.Body)
	if err != nil{
		log.Fatal(err)
	}
	ioutil.WriteFile("./site.html",body,os.ModePerm)
}
