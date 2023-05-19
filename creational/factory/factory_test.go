package factory

import "testing"

func getRes(factory FoodMachineFactory, foodA, foodB string) string {
	synthesizer := factory.Create()
	synthesizer.AddFoodA(foodA)
	synthesizer.AddFoodB(foodB)
	return synthesizer.Cooking()
}

func TestCreateSubFactory(t *testing.T) {
	t.Run("create meatballs", func(t *testing.T) {
		factory := MeatballMachineFactory{}
		if res := getRes(factory, "meat", "berries"); res != "meat + berries = Fresh Meatballs" {
			t.Errorf("error with cooking meatballs: %s", res)
		}
	})

	t.Run("create fishsticks", func(t *testing.T) {
		factory := FishstickMachineFactory{}
		if res := getRes(factory, "fish", "sticks"); res != "fish + sticks = Fresh Fishsticks" {
			t.Errorf("error with cooking fishsticks: %s", res)
		}
	})
}
