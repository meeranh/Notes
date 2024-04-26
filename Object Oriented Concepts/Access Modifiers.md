# Access Modifiers/Specifiers
+ Lets say that you do not want certain members of a class to be accessible from outside the class.
+ In this case, you can use access modifiers to restrict access to certain members of the class.
+ We have three kinds of access modifiers:

| Name      | Description                                       | Syntax       |
|-----------|---------------------------------------------------|--------------|
| Public    | Accessible from anywhere                          | memberName   |
| Protected | Accessible within the class and its subclasses    | _memberName  |
| Private   | Accessible only within the class                  | __memberName |

### Public Access Modifier
+ The public access modifier is just the member name itself, nothing special:
```python
class MyClass:
    def __init__(self):
        self.publicMember = 100

some_class = MyClass()
print(some_class.publicMember) # Prints 100
```

### Protected Access Modifier
+ The protected access modifier is a single underscore (_) before the member name.
+ This just tells the developers that it should not be accessed from outside the class:
```python
class MyClass:
    def __init__(self):
        self._protectedMember = 10

some_class = MyClass()
print(some_class._protectedMember) # Prints 10
```
+ There is actually no difference in function between a public and protected member.
+ It is just a convention to tell developers to not directly access the member.
+ These kinds of members are accessed using getters, which we will discuss later.

### Private Access Modifier
+ These kinds of members use two underscores (__) before the member name.
+ You basically cannot access these members from outside the class:
```python
class MyClass:
    def __init__(self):
        self.__privateMember = 0

some_class = MyClass()
print(some_class.__privateMember) # Prints an error
```
