# Methods
+ A method is a function that is associated with a class/object.
+ All `Car` classes would have a `drive` method.
+ All `Car` classes also would have a `sayHello` method which just prints "Hello".
+ What is the difference between `sayHello` and `drive`?
+ `sayHello` is a method that does not change the state of the object.
+ `drive` is a method that changes the state of the object.
+ If you have a car that is already driving, you can't drive it again.
+ But you can say hello to it as many times as you want.

## Static Methods
+ `sayHello` is a static method.
+ Lets look at an example `sayHello` implementation in a `Car`.
```python
class Car:
    def __init__(self, color, model):
        self.color = color
        self.model = model
        self.is_driving = False

    def drive(self):
        case self.is_driving:
            True:
                print("Car is already driving")
            False:
                self.is_driving = True
                print("Driving")

    @staticmethod
    def sayHello():
        print("Hello")
```

+ To declare a static method which has nothing to do with the object, we use the `@staticmethod` decorator.
+ The static method does not take in a `self` parameter.
+ The static method here was `sayHello` and the instance method was `drive`.
