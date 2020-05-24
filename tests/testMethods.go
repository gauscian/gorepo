/*
 * Created on Sat May 16 2020
 *
 * Primary Author - Achin Gupta
 */

package main


import (
    "fmt"
)

// Person is a general person
type Person struct {
	first string
	second string
}

// SecretAgent : Person with license to kill
type SecretAgent struct {
	Person
	ltk bool
}

// method connected with secret agent. Returns nothing.
func (s SecretAgent) speak () {
	fmt.Println("Name is ", s.second, ",", s.first, s.second)
	fmt.Println("And I have license to kill. Correct ?", s.ltk)
}


func main() {
	
    sa1 := SecretAgent {
		Person: Person {
			first: "James",
			second: "Bond",
		},
		ltk: true,
	}
	sa1.speak()

	sa2 := SecretAgent {
		Person: Person {
			first: "Miss",
			second: "Monneypenny",
		},
		ltk: false,
	}

	sa2.speak()
}