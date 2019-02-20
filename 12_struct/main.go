package main

import(
  "fmt"
  "strconv"
)

type Pet struct {
    name string
    species string
    owner string
    gender string
    age int
    hungry bool
}

// value reciever
func (p Pet) greet() string  {
  if p.species == "Dog" && p.hungry == true {
    return "Woof! My name is " + p.name + " and I'm a good " + p.gender + "! I'm " + strconv.Itoa(p.age) + " years old. FEED ME!!!"
  } else if p.species == "Dog" && p.hungry == false {
    return "Woof! My name is " + p.name + " and I'm a good " + p.gender + "! I'm " + strconv.Itoa(p.age) + " years old."
  } else if p.species == "Cat" && p.hungry == true {
    return "Meow! My name is " + p.name + " give me canned food right now!"
  } else if p.species == "Cat" && p.hungry == false {
    return "Meow! My name is " + p.name + " give me canned food in a little bit!"
  } else {
    return "Hello, my name is " + p.name +"I don't know my species!"
  }
}

//pointer reciever
func (p *Pet) hasBirthday()  {
  p.age ++
}

func(p *Pet) feed() {
  if p.hungry == true{
    p.hungry = false
  }
}

func main() {
  pet1 := Pet{name:"Bub", species:"Cat", owner:"Hannah", gender:"Male", age:6, hungry: true}

  // pet2 := Pet{"George", "Dog", "Mark", "Male", 6, true}
  // pet3 := Pet{"Ada", "Dog", "Barb", "Female", 8, true}

  // pet1.hasBirthday()
  pet1.feed()
  fmt.Println(pet1.greet())
}
