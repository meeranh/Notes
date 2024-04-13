# Open Closed Principle
+ This principle states that a class should be open for extension but closed for modification.
+ Lets say that you are running an insurance company.
+ You want to take in a `InsuranceCustomerProfile` and identify how much total income you have generated from him.
+ The code to do this might look like the following:
```java
public class InsuranceCustomerProfile {
    public int totalSales() {
        // Return total sales
    }
}

public class InsuranceCustomerProfileAnalyzer {
    public int totalSales(InsuranceCustomerProfile profile) {
        return profile.totalSales();
    }
}
```

+ You can see that `InsuranceCustomerProfileAnalyzer` is dependent on `InsuranceCustomerProfile`.
+ The reason this is bad is because lets say that you now want to do car sales along with insurance sales.
+ You will have to modify the existing code to include the car sales.
+ Modifying existing code is bad because it can introduce bugs.
+ What you can do is when you create your software, you can create it in such a way that you can extend it without modifying the existing code.
```java
public class CustomerProfile {

    public int totalSales() {
        // Return total sales
    }
}

public class InsuranceCustomerProfile extends CustomerProfile {
    // Implement totalSales
}

public class CarCustomerProfile extends CustomerProfile {
    // Implement totalSales
}
```

+ You can see above that we are implementing a common `CustomerProfile` class, and then extending it to `InsuranceCustomerProfile` and `CarCustomerProfile`.
+ Therefore, now if you want to add even a `BikeCustomerProfile`, you can do so without modifying the existing code as
+ All you need to do is to create a `BikeCustomerProfile` class that is inherited from `CustomerProfile` and implement the `totalSales` method.
