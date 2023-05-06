class Animal:
    def __init__(self, name: str) -> None:
        self.name = name

    def speak(self):
        pass


class Dog(Animal):
    def speak(self):
        return "Woof!"


class Cat(Animal):
    def speak(self):
        return "Meow!"


class AnimalFactory:
    def create_animal(self, name: str) -> Animal:
        if name == "dog":
            return Dog(name)
        elif name == "cat":
            return Cat(name)
        else:
            return None


def main():
    factory = AnimalFactory()
    dog = factory.create_animal("dog")
    print(dog.speak())

    cat = factory.create_animal("cat")
    print(cat.speak())


if __name__ == "__main__":
    main()
