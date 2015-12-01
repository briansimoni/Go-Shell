package main


/* simple rule: If a type implenets all of the methods of that interface
then it is an implentation of that interface

There is no keyword 'implements'

since there is no method overloading, we can decide that the rule is this.
If a custom types method has a name that matches an interfaces name, then it must implement
that method with exactly the same number of types and arguments, and return the same type
And you cannot have a mismatch*/
import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct{

}

func (d Dog) Speak() string{
	return "Woof!"
}

func main() {
	poodle := Animal(Dog{}) // casting poodle to animal
	fmt.Println(poodle)
}