package test

import "strconv"

//BDD TDD DDD software aproches

func Hello(name string) string{
	if name == ""{
		return "Hello World"
	}
	_,err:=strconv.Atoi(name)
	if err==nil{
		return "Hello Adam"
	}
	return "Hello world"+name
}
func main(){

}