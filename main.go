package main

import"fmt"

type book struct{
 title string
 author string
 isbn string
 price float32
 pages int
}
func main(){
var mormon book
mormon.title= "book of mormon"
mormon.author= "sacred text"
mormon.isbn= "9781537696775"
mormon.price=35.05
mormon.pages=238



fmt.Println(mormon)

}