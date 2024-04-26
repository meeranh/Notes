# Abstract Base Class
+ Lets say that you are planning what your classes would look like.
+ You have a plan of what your method names would be, but you don't know what the implementation would be.
+ You can specify this plan using an abstract base class.
```python
from abc import ABC, abstractmethod

class Shape(ABC):
    @abstractmethod
    def area(self):
        pass

    @abstractmethod
    def perimeter(self):
        pass
```

+ Basically, you now have a plan for what your classes would look like.
+ Also, when you create your actual classes, it has to follow the exact structure of the abstract base class.
+ Lets see an example:
```python
from abc import ABC, abstractmethod

class Shape(ABC):
    @abstractmethod
    def area(self):
        pass

    @abstractmethod
    def perimeter(self):
        pass

class Square(Shape):
    def __init__(self, side):
        self.side = side

    def area(self):
        return self.side * self.side

    def perimeter(self):
        return 4 * self.side


s = Square(5) # This is fine
```
+ The above example is fine because the `Square` class follows the structure of the `Shape` class.
+ However, lets say that you forgot to implement the `perimeter` method in the `Square` class.
```python
from abc import ABC, abstractmethod

class Shape(ABC):
    @abstractmethod
    def area(self):
        pass

    @abstractmethod
    def perimeter(self):
        pass

class Square(Shape):
    def __init__(self, side):
        self.side = side

    def area(self):
        return self.side * self.side


s = Square(5) # This will give an error
```

+ Basically, when you define an abstract base class, you need to have all the methods implemented in the actual class.
+ Lets discuss the `ABC` and `abstractmethod` decorators.

## `ABC` Class
+ This is the way you tell Python that this is an abstract base class.
+ You have to import `ABC` from the `abc` module.
+ Then you have to create your class as a child of the `ABC` class.
```python
from abc import ABC

class Shape(ABC):
    pass
```

+ Then, you need to use the `abstractmethod` decorator to specify that this is an abstract method.
+ You have to import `abstractmethod` from the `abc` module.
+ This `abstractmethod` decorator should be used to decorate the method that you want to make abstract.
```python
from abc import ABC, abstractmethod

class Shape(ABC):
    @abstractmethod
    def area(self):
        pass
```
