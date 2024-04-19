# Modularity
+ OOP allows you to create code that can be segmented into modules.
+ These modules can be reused in other parts of the program.
+ The important thing is that these modules are independent of each other.
+ This is desirable because if you make a change in one module, it will not affect the other modules.
+ The following example show how you can modularize the code for a Dog and a Cat into two separate classes.
```python
class Dog:
    def __init__(self, name, age):
        self.name = name
        self.age = age

    def bark(self):
        print(f"{self.name} is barking")

class Cat:
    def __init__(self, name, age):
        self.name = name
        self.age = age

    def meow(self):
        print(f"{self.name} is meowing")
```

# Reusability
+ Lets say that you created a pet store that sells dogs and cats.
+ You now want to sell other pets like birds and fish.
+ If you created a common `Pet` class, you can now extend this class to create new classes for birds and fish.
+ This is what code reusability means, you don't have to code anything again.
```python
class Pet:
    def __init__(self, name, age):
        self.name = name
        self.age = age

class Dog(Pet):
    def bark(self):
        print(f"{self.name} is barking")

class Cat(Pet):
    def meow(self):
        print(f"{self.name} is meowing")

class Bird(Pet):
    def chirp(self):
        print(f"{self.name} is chirping")

class Fish(Pet):
    def swim(self):
        print(f"{self.name} is swimming")
```

# Extendability
+ OOP allows you to extend the functionality of a class.
+ Lets say you have a `Pet` class that has a `name` and `age` attribute.
+ You can extend this class to add more attributes.
+ Basically, you can improve your code without changing the existing code.
```python
class Pet:
    def __init__(self, name, age):
        self.name = name
        self.age = age
        self.health = 100
        self.happiness = 100
```
