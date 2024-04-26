# Classes
+ A class is a blueprint that mimics real world objects.
+ Think of it as like the code version of a real world object.
+ If you were to mimic a car using a class, you would have a class called `Car`.
+ A car would have properties like `color`, `make`, `model`, `year`, etc.
+ The following is a mimic of a car using a class:
```python
class Car:
    def __init__(self, color, make, model, year):
        self.color = color
        self.make = make
        self.model = model
        self.year = year
```
+ In the next section, we will discuss `__init__` and `self`.

# Objects
+ An object is an instance of a class.
+ If you wanted to create 10 different cars with different colors, makes, models, and years, you would create 10 different objects of the `Car` class.
+ The following is an example of creating 2 different cars:
```python
car1 = Car("Red", "Toyota", "Corolla", 2019)
car2 = Car("Blue", "Ford", "F-150", 2020)
```

+ Note that a `Car` is created using certain properties like `color`, `make`, `model`, and `year`.
+ This is the basic idea of the `__init__` method.

# The `__init__` Method (Constructor)
+ This method is commonly known as the constructor.
+ When you create an object using `Car("Your", "Properties", "Here")`, you are calling the `__init__` method (or the constructor).
+ So, if you want to create objects with different properties, you always need an `__init__` method.
+ In our `Car` example, lets look at the `__init__` method:
```python
class Car:
    def __init__(self, color, make, model, year):
        self.color = color
        self.make = make
        self.model = model
        self.year = year
```

+ What this does is whenever you create a `Car` object with a constructor, whatever you pass into that constructor will be assigned to `self`. We will describe `self` shortly.
+ So what this does is it assigns the `color` you pass in to the `self.color` property, the `make` you pass in to the `self.make` property, and so on.

# The `self` Keyword
+ Imagine you were designing a game where players can control characters.
+ You would have a class called `Character` and you would have properties like `health`, `attack`, `defense`, etc.
+ When you create a character, you would want to assign these properties to that character.
+ In programming, when you want to assign a property to an object, you use the `self` keyword.
+ The following example should explain it:
```python
class Character:
    def __init__(self, name, color, age):
        self.name = name
        self.color = color
        self.age = age

    def say_name(self):
        print(self.name)

mario = Character("Mario", "Red", 30)
luigi = Character("Luigi", "Green", 32)

mario.say_name() # Output: Mario
luigi.say_name() # Output: Luigi
```

+ What you should understand from this is that when you pass `Mario` and `Luigi` when creating the object, it gets passed to the `__init__` function as a `name`.
+ This `name` is passed into `self.name`
+ This `self.name` will be used throughout the entire class to refer to the `name` property.
+ Basically, `self` lets you hold onto the properties you pass in when creating an object.

# State
+ The properties of a class you assigned via `self` are called the state of the object.
+ For example, lets take the Super Mario example:
```python
mario = Character("Mario", "Red", 30)
```
+ This creates a `Character` object called `mario` with the state of `name = "Mario"`, `color = "Red"`, and `age = 30`.

```python
luigi = Character("Luigi", "Green", 32)
```
+ This creates a `Character` object called `luigi` with the state of `name = "Luigi"`, `color = "Green"`, and `age = 32`.

# Attributes
+ These are basically the properties of a class.
+ For example, lets identify the attributes of the `Character` class:
```python
class Character:
    health = 100

    def __init__(self, name, color, age):
        self.name = name
        self.color = color
        self.age = age
```

+ The attributes of the `Character` class are `name`, `color`, `age`, and `health`.
+ Anything assigned to `self` in the `__init__` method are attributes of the class.
+ Any variable assigned outside of the `__init__` method are also attributes of the class.
+ Any attribute can be accessed by using the dot operator.
```python
mario = Character("Mario", "Red", 30)
print(mario.name) # Output: Mario
print(mario.health) # Output: 100
```
