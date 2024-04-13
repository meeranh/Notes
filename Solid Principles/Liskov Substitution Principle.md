# Liskov Substitution Principle
+ This principle states that objects of a superclass should be replaceable with objects of its subclasses without affecting the functionality of the program.
+ For example, lets say that we have a `Car` class.
+ We want to create two instances of `Car`, one for a `normalCar` and one for a `racingCar`:
```java
public class Car {
    public void getCabinWidth() {
        // Return cabin width
    }
}

List<Car> cars = new ArrayList<>();
normalCar = new Car();
racingCar = new Car();

System.out.println(normalCar.getCabinWidth()); // This will print out fine
System.out.println(racingCar.getCabinWidth()); // This will give an error
```

+ The reason why a racing car will give an error is because a racing car does not have a cabin. It has a cockpit.
+ Therefore, when developing software, you have to create it in a way where these kinds of errors do not occur.
+ To fix this what we can do is:
    1. Create a new `Vehicle` class.
    2. Create a `Car` class that extends `Vehicle`.
    3. Create a `RacingCar` class that extends `Car`.

```java
public class Vehicle {
    public double getVehicleWidth() {
        // Return vehicle width
    }
}
```

```java
public class Car extends Vehicle {
    
    @Override
    pubilc double getVehicleWidth() {
        return this.getCabinWidth();
    }

    public double getCabinWidth() {
        // Return cabin width
    }

}
```

```java
public class RacingCar extends Car {
    
    @Override
    public double getVehicleWidth() {
        return this.getCockpitWidth();
    }

    public double getCockpitWidth() {
        // Return cockpit width
    }

}
```

+ Now, we can create a `RacingCar` object and call the `getVehicleWidth` method without any errors.
+ Therefore, the Listov Substitution Principle states that objects of a superclass should be replaceable with objects of its subclasses without affecting the functionality of the program.
