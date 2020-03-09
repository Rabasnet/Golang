package main

import ("fmt")
func main(){
a:= person{
	name:"rabi basnet",
	
}
fmt.Println(a) //original value is printed
changeMe(&a) //function called with the name address of the original name
fmt.Println(a) //printing the new name of the original name
}
type person struct{
name string

}
 func changeMe(p *person){
p.name="good man"
 }
