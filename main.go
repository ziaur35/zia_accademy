package main

import(
"fmt" 
"reflect"
)
func main(){

//var students [3]string
//students[0]="mostain sir"
//students[1]="reza sir"
//students[2]="ziaur"

//x := students[0:2]
//fmt.Println(x)

//x := make([]string, 0)
//fmt.Println(x)

//fmt.Println(students)
//fmt.Println(len(students))
//fmt.Println(students[1])
//students := [...]string{"mostain sir", "reza sir", "ziaur"}
//fmt.Println(students, len(students))

var bread []string
bread = append(bread, "bread1", "bread2", "bread3")
//fmt.Println(bread)
//fmt.Println(bread, len(bread))
//fmt.Printf("%T\n", bread)
//fmt.Printf("%T", students)

a := reflect.TypeOf(bread).Kind().String()
fmt.Println(a)



}