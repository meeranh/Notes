# Inheritance
+ Inheritance is a mechanism in which one class acquires the property of another class.
+ Lets say that you have a list of birds:
    1. Eagle
    2. Sparrow
    3. Pigeon

+ All of these birds have a `wingLength`, `beakLength`, and `color`.
+ But a `Sparrow` has a `nest` property, and an `Eagle` has a `talons` property, and a `Pigeon` has a `coo` property.
+ We can implement a common `Bird` class that has the `wingLength`, `beakLength`, and `color` properties
+ Then we can create a `Sparrow`, `Eagle`, and `Pigeon` class that extends the `Bird` class and adds the `nest`, `talons`, and `coo` properties respectively.

```java
class Bird {
    int wingLength;
    int beakLength;
    String color;
}

class Sparrow extends Bird {
    String nest;
}

class Eagle extends Bird {
    String talons;
}

class Pigeon extends Bird {
    String coo;
}
```
