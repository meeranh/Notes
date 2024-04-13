# Polymorphism
+ Lets say that you have the following animals:
    1. Dog
    2. Cat
    3. Cow

+ What we have in common here is that all animals eat, but they eat differently.
+ So, in polymorphism, we can create an `Animal` class with an `eat()` method, and then create a `Dog`, `Cat`, and `Cow` class that extends the `Animal` class and overrides the `eat()` method with a custom implementation of `eat()`.

```java
class Animal {
    public void eat() {
        System.out.println("This animal eats food");
    }
}
```

```java
class Dog extends Animal {
    @Override
    public void eat() {
        System.out.println("This dog eats bones");
    }
}
```

```java
class Cat extends Animal {
    @Override
    public void eat() {
        System.out.println("This cat eats fish");
    }
}
```

```java
class Cow extends Animal {
    @Override
    public void eat() {
        System.out.println("This cow eats grass");
    }
}
```
