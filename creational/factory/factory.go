package factory

import "fmt"

type FoodMachineFunction interface {
	AddFoodA(string)
	AddFoodB(string)
	Cooking() string
}

type FoodMachineFactory interface {
	Create() FoodMachineFunction
}

type FoodMachine struct {
	FoodA string
	FoodB string
}

func (f *FoodMachine) AddFoodA(food string) {
	f.FoodA = food
}

func (f *FoodMachine) AddFoodB(food string) {
	f.FoodB = food
}

type MeatballMachine struct {
	*FoodMachine
}

func (mm MeatballMachine) Cooking() string {
	if mm.FoodA == "meat" || mm.FoodB == "meat" {
		return fmt.Sprintf("%s + %s = Fresh Meatballs", mm.FoodA, mm.FoodB)
	}

	return "Wet Goop"
}

type MeatballMachineFactory struct{}

func (MeatballMachineFactory) Create() FoodMachineFunction {
	return &MeatballMachine{
		FoodMachine: &FoodMachine{},
	}
}

// FishstickSynthesizer cooking a fishstick
type FishstickMachine struct {
	*FoodMachine
}

func (fm FishstickMachine) Cooking() string {
	if fm.FoodA == "fish" || fm.FoodB == "fish" {
		return fmt.Sprintf("%s + %s = Fresh Fishsticks", fm.FoodA, fm.FoodB)
	}

	return "Wet Goop"
}

type FishstickMachineFactory struct{}

func (FishstickMachineFactory) Create() FoodMachineFunction {
	return &FishstickMachine{
		FoodMachine: &FoodMachine{},
	}
}
