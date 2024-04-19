# Single Inheritance
+ Inheritance is a mechanism in which one class acquires the property of another class.
+ Lets say that you have a list of birds:
    1. Eagle
    2. Sparrow
    3. Pigeon

+ All of these birds have a `wing_length`, `beak_length`, and `color`.
+ But a `Sparrow` has a `nest` property, and an `Eagle` has a `talons` property, and a `Pigeon` has a `coo` property.
+ We can implement a common `Bird` class that has the `wing_length`, `beak_length`, and `color` properties
+ Then we can create a `Sparrow`, `Eagle`, and `Pigeon` class that extends the `Bird` class and adds the `nest`, `talons`, and `coo` properties respectively.

```python
class Bird:
    wing_length: int
    beak_length: int
    color: str

class Sparrrow(Bird):
	nest: str

class Eagle(Bird):
	talons: str

class Pigeon(Bird):
	coo: str
```

# Base Class vs Derived Class
+ In the above example, we used the `Bird` class to create the `Sparrow`, `Eagle`, and `Pigeon` classes.
+ Therefore, the `Bird` class is known as the  **Base Class** and the `Sparrow`, `Eagle`, and `Pigeon` classes are known as the **Derived Classes**.
+ This is because the `Sparrow`, `Eagle`, and `Pigeon` classes derive their properties from the `Bird` class.

# Multiple Inheritance
+ In the above example, we were using just one base class `Bird` to create the `Sparrow`, `Eagle`, and `Pigeon` classes.
+ But what if we want to create a `Penguin` class that has the properties of both a `Bird` and a `Fish`?
+ We can use multiple inheritance to achieve this.
+ You first have to create the `Bird` and `Fish` classes:
```python
class Bird:
    wing_length: int
    color: str

class Fish:
    fin_length: int
```

+ Now you can create the `Penguin` class that inherits from both the `Bird` and `Fish` classes:
```python
class Penguin(Bird, Fish):
    def __init__(self, wing_length, color, fin_length):
        self.wing_length = wing_length
        self.color = color
        self.fin_length = fin_length
```

+ This will behave just like a normal class that inherits from a single class.
```python
penguin = Penguin(10, "Black", 5)
print(penguin.wing_length) # 10
print(penguin.color) # Black
```

# Multilevel Inheritance
+ In multilevel inheritance, a class is derived from a class that is also derived from another class.
+ Lets say we have a `Animal` class that has a `name` property.
+ We have a `Dog` class that inherits from the `Animal` class and has a `bark` property.
+ We have a `Husky` class that inherits from the `Dog` class and has a `fur` property.
+ The `Husky` class has the `name`, `bark`, and `fur` properties.
+ Basically, `Husky` will have all the properties of higher classes.
```python
class Animal:
    name: str

class Dog(Animal):
    bark: str

class Husky(Dog):
    fur: str
```

+ The theory is that `Husky` will inherit both `Animal` and `Dog` properties.
```python
husky = Husky()
husky.name = "Husky"
husky.bark = "Woof"
husky.fur = "Thick"
```
