package simple_factory

type Animal interface {
	Speak() string
}

func NewAnimal(animalType string) Animal {
	switch animalType {
	case "Dog":
		return &Dog{}
	case "Cat":
		return &Cat{}
	default:
		panic ("Invalid type")
	}
}

type Cat struct {
}

func (c *Cat) Speak() string {
	return "Meow!"
}

type Dog struct {
}

func (d *Dog) Speak() string {
	return "Woof!"
}
