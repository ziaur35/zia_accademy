package main

import"fmt"

func main(){


mormon := struct{
 title string
 author string
 isbn string
 price float32
 pages int
}{
title: "book of mormon",
author: "sacred text",
isbn: "9781537696775",
price:35.05,
pages:238,
}
fmt.Println(mormon)

}