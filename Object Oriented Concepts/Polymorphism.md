# Polymorphism
+ Polymorphism states that objects of a class might behave differently.
+ For example, if you have a twin sister.
+ Your twin might look the exact same as you, but she might have a different personality.
+ Polymorphism lets you simulate this behavior in programming.

### Method Overriding
+ Lets say that you have the following animals:
    1. Dog
    2. Cat
    3. Cow

+ What we have in common here is that all animals eat, but they eat differently.
+ So, in polymorphism, we can create an `Animal` class with an `eat()` method, and then create a `Dog`, `Cat`, and `Cow` class that extends the `Animal` class and overrides the `eat()` method with a custom implementation of `eat()`.

```python
class Animal:
    def eat(self):
        print("This animal eats food")
```

```python
class Dog(Animal):
    def eat(self):
        print("This dog eats bones")
```

```python
class Cat(Animal):
    def eat(self):
        print("This cat eats fish")
```

```python
class Cow(Animal):
    def eat(self):
        print("This cow eats grass")
```

+ However, lets say that for some reason you wanted to reset the `eat()` method to the original implementation.
+ Lets say you wanted to reset the `Cow`'s `eat()` method to the original implementation.
+ You can do that with the help of `super()`.
```python
class Cow(Animal):
    def eat(self):
        print("This cow eats grass")

    def resetEat(self):
        super().eat()

eat_instance = Cow()
eat_instance.eat() # Prints "This cow eats grass"
eat_instance.resetEat() # Prints "This animal eats food"
```
+ Basically, you can access the parent class's methods using `super()`.

### Method Overloading
+ Method overloading is not supported in Python (kind of).
+ We will start with a real example of method overloading in Java.
+ Lets say that we have a `Calculator` class with an `add()` method.
+ The `add()` method can take two integers, two floats, or two strings.
```java
public class Calculator {
    public int add(int a, int b) {
        return a + b;
    }

    public float add(float a, float b) {
        return a + b;
    }

    public String add(String a, String b) {
        return a + b;
    }
}
```

+ However, you now have the need to add three integers.
+ You can create a new `add` method that takes three integers.
```java
public class Calculator {
    public int add(int a, int b) {
        return a + b;
    }

    public int add(int a, int b, int c) {
        return a + b + c;
    }

    public float add(float a, float b) {
        return a + b;
    }

    public String add(String a, String b) {
        return a + b;
    }
}
```
+ The process of creating multiple methods with the same name but different parameters is called method overloading.
+ Now lets see a similar example in Python.
+ Lets say that we have a `Square` class with an `area()` method.
+ You now have created two square objects.
```python
class Square:
    def __init__(self, side):
        self.side = side

    def area(self):
        return self.side * self.side

square1 = Square(5)
square2 = Square(10)
```

+ You now want to add the areas of the two squares.
+ But you cannot do `square1 + square2`.
+ This is because `Square` objects do not support the `+` operator.
+ We can make the `Square` class support the `+` operator:
```python
class Square:
    def __init__(self, side):
        self.side = side

    def area(self):
        return self.side * self.side

    def __add__(self, other):
        return self.area() + other.area()
```

+ By doing this, you basically allowed the `Square` class to support the `+` operator.
+ These are know as magic/special methods in Python.
