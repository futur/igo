package main

import (
"fmt"
"math"
"reflect"
"mypackages"
)

type myInt int
type UrIbt int

var c,js,objc bool

func main(){
	var prem myInt
	var dhilip UrIbt
	var integer int
	fmt.Println(integer, c,js,objc)
	prem = 10
	dhilip = UrIbt(prem)

	fmt.Println("prem", prem,"dhilip",  dhilip)

	fmt.Println("Hello world.")
	fmt.Println("Pi is ",math.Pi)
	fmt.Println("Sum is ", add(math.Pi, 98))
	x,y := swap("i live", "here")
	fmt.Println(reflect.TypeOf(x),y);
	fmt.Println(split(100));
	fmt.Printf(mypackages.Mypackage())

}

func add(x float64, y int) float64 {
	return x + float64(y)
}
func swap (a , b string) (string,string){
	return b,a
}
func split(sum int) (x,y int) {
	x = sum * 4/9;
	y = sum -x;
	return	
}
