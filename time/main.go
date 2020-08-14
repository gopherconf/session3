package main

import (
	"fmt"
	"time"
)

func main()  {
	getTime(time.Now())
	//Duration
	getTime(time.Now().Add(time.Minute*3).Add(time.Millisecond*800))

	fmt.Println("UnixTime:",time.Now().Unix())
	fmt.Println(time.Now().Format(time.RFC3339))
	fmt.Println(time.Now().Format(time.RFC850))

}


func getTime(t time.Time){
	fmt.Println(t.Minute())
}