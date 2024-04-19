# Getters
+ Lets say that you have a protected or private member in a class.
+ There is no way you can access this from an instance of an object.
```python
class MyClass:
    def __init__(self):
        self.__privateMember = 100

some_class = MyClass()
print(some_class.__privateMember) # Prints an error
```

+ Therefore, if you want to access these members, you need to use a getter.
+ A getter is basically a function that returns a protected or private member.
+ Here is an example:
```python
class MyClass:
    def __init__(self):
        self.__privateMember = 100

    def getPrivateMember(self):
        return self.__privateMember

some_class = MyClass()
print(some_class.getPrivateMember()) # Prints 100
```

# Setters
+ A setter is basically the opposite of a getter.
+ It is a function that sets the value of a protected or private member.
+ Here is an example:
```python
class MyClass:
    def __init__(self):
        self.__privateMember = 100

    def setPrivateMember(self, value):
        self.__privateMember = value

    def getPrivateMember(self):
        return self.__privateMember

some_class = MyClass()
some_class.setPrivateMember(200)
print(some_class.getPrivateMember()) # Prints 200
```
