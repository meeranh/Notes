# Abstraction
+ Abstraction is the concept of hiding the complex implementation details and showing only the necessary features of the object.
+ Imagine you are creating a library that is used by millions of developers.
+ You don't want to expose the internal implementation details of your library to the developers.
+ This is because you may make some breaking changes in the future, and it might break all of the codebases of these millions of developers.
+ You only want to show them a set of functions that they can just pass their inputs and get the output.
+ These types of simple functions are called **abstractions**.
+ The following code example is how you implement a class without abstractions:
```python
class Car
	def __init__(self, make, model, year, color, price, owner):
        self.make = make
        self.model = model
        self.year = year
        self.color = color
        self.price = price
        self.owner = owner

	def start_engine(self):
        print("Engine started")

	def send_power_to_wheels(self):
        print("Power sent to wheels")

	def enable_stabilizers(self):
        print("Stabilizers enabled")
```

+ The above class gives too much control over the car to the user. This is not a good practice.
+ Instead you need to create one public method that the user can use to start the car:
```python
class Car:
	def __init__(self, make, model, year, color, price, owner):
		self.make = make
		self.model = model
		self.year = year
		self.color = color
		self.price = price
		self.owner = owner

	def start_car(self):
		self.start_engine()
		self.send_power_to_wheels()
		self.enable_stabilizers()
```

+ This way, the user only needs to call the `start_car()` method to start the car.
