from abc import ABC, abstractmethod


class Food(ABC):
    """
    Declared the interface of the product.
    """

    def __init__(self, food_name: str):
        self.food_name = food_name

    @abstractmethod
    def restore_health(self):
        pass

    @abstractmethod
    def restore_hungry(self):
        pass

    @abstractmethod
    def restore_sanity(self):
        pass


class FoodMachine(ABC):
    """
    Food Machine, is an abstract class that declares the factory method and
    business logic methods
    """

    def __init__(self, material_a: str, material_b: str):
        self.material_a = material_a
        self.material_b = material_b

    @abstractmethod
    def cooking(self) -> Food:
        """
        Factory method, generally used to produce products.
        For example, this function in this class is used to
        cook something.
        """

    def get_food(self) -> str:
        """
        Core business logic, usual logical expansion of the product.
        For example, this function in this class is used to get the
        cooked food.
        """
        food = self.cooking()
        health = food.restore_health()
        hungry = food.restore_hungry()
        sanity = food.restore_sanity()
        return (
            f"{self.material_a} + {self.material_b} = {food.food_name}. "
            f"It restores health {health}, hungry {hungry}, sanity {sanity}"
        )


class MeatBallMachine(FoodMachine):
    """
    Concerte Creator. For the production of a product.
    """

    def __init__(self, material_a: str = "meat", material_b: str = "berries"):
        super().__init__(material_a, material_b)

    def cooking(self):
        return MeatBall()


class FishStickMachine(FoodMachine):
    def __init__(self, material_a: str = "fish", material_b: str = "stick"):
        super().__init__(material_a, material_b)

    def cooking(self):
        return FishStick()


class MeatBall(Food):
    """
    Concerte Product. Implemented the interface class.
    """

    def __init__(self, food_name: str = "MeatBalls"):
        super().__init__(food_name)

    def restore_health(self):
        return 3

    def restore_hungry(self):
        return 62.5

    def restore_sanity(self):
        return 5


class FishStick(Food):
    def __init__(self, food_name: str = "FishSticks"):
        super().__init__(food_name)

    def restore_health(self):
        return 40

    def restore_hungry(self):
        return 37.5

    def restore_sanity(self):
        return 5


def player_get_food(food_machine: FoodMachine) -> None:
    print(f"Player get food: {food_machine.get_food()}")


if __name__ == "__main__":
    print("case 1: get meatballs")
    player_get_food(MeatBallMachine())

    print("case 2: get fishsticks")
    player_get_food(FishStickMachine())
