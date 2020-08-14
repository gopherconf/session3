package main

import "fmt"

func main(){
	firstFunc("GopherConf 3", func(s string){
		fmt.Println(s)
	})
}

func firstFunc(str string,prnt func(string)){
	 prnt("Hello " + str)
}