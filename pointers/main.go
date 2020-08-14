package main

import "fmt"

func main(){
	i := 700
	getPointer(&i)
	fmt.Println(i)

	getNil(nil)
	getNil(&i)
}


func getPointer(i *int){
	fmt.Println(i)
	fmt.Println(*i)

	*i = 500
}

func getNil(i *int)  {
	if i != nil{
		fmt.Println(*i)
	}else{
		fmt.Println("i is nil")
	}

}