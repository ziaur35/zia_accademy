package main

import"fmt"


func main(){

//var students [3]string

//students[0]="mostain sir"
//students[1]="reza sir"
//students[2]="ziaur"


//fmt.Println(students)
//fmt.Println(len(students))
//fmt.Println(students[1])
students := [...]string{"mostain sir", "reza sir", "ziaur"}
fmt.Println(students, len(students))


//var fruits []string

//fruits = append(fruits, "apple", "banana", "mango")

//fmt.Println(fruits, len(fruits))

//fmt.Printf("%T\n", fruits)

//fmt.Printf("%T", students)

//a := reflect.Typeof(fruits).Kind().string()
//fmt.Println(a)

}