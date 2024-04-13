# Abstraction
+ Abstraction is the concept of hiding the complex implementation details and showing only the necessary features of the object.
+ Imagine you are creating a library that is used by millions of developers.
+ You don't want to expose the internal implementation details of your library to the developers.
+ This is because you may make some breaking changes in the future, and it might break all of the codebases of these millions of developers.
+ You only want to show them a set of functions that they can just pass their inputs and get the output.
+ These types of simple functions are called **abstractions**.
+ The following code example is how you implement a class without abstractions:
```java
class Car {
    String make;
    String model;
    int year;
    String color;
    int price;
    String owner;
    
    public void startEngine() {
        System.out.println("Engine started");
    }

    public void sendPowerToWheels() {
        System.out.println("Power sent to wheels");
    }
    
    public void enableStabizers() {
        System.out.println("Stabilizers enabled");
    }
}
```

+ The above class gives too much control over the car to the user. This is not a good practice.
+ Instead you need to create one public method that the user can use to start the car:
```java
class Car {
    String make;
    String model;
    int year;
    String color;
    int price;
    String owner;
    
    public void startCar() {
        startEngine();
        sendPowerToWheels();
        enableStabizers();
    }

    private void startEngine() {
        System.out.println("Engine started");
    }

    private void sendPowerToWheels() {
        System.out.println("Power sent to wheels");
    }
    
    private void enableStabizers() {
        System.out.println("Stabilizers enabled");
    }
}
```

+ This way, the user only needs to call the `startCar()` method to start the car.
