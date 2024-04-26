# Dependency Inversion Principle
+ This principle states that high-level modules should not depend on low-level modules. Both should depend on abstractions.
+ For example, lets assume that we have a `Store` class that depends on calling Stripe API to process payments.
+ The code might look like the following:
```java
public class Store {
    private string stripeAPI;
    private string authToken;

    public void purchaseProduct() {
        makeHttpCall(stripeAPI, authToken);
    }
}
```

+ The above code is bad because the `Store` class is very dependent on the Stripe API.
+ What if Stripe changes their API structure? You will have to change the entire `Store` class logic.
+ This is a perfect example of tight coupling as the `Store` class is tightly coupled with the Stripe API.
+ What we can do is create a class between the `Store` class and the Stripe API.
+ These kinds of classes are known as **Abstractions**.
+ Lets create an abstraction class known as `PaymentProcessor`:
```java
public interface PaymentProcessor {
    private void callStripeAPI() {
        // Calls the Stripe API
    }

    public void processPayment(serviceName) {
       switch(serviceName) {
           case "Stripe":
               callStripeAPI();
               break;
           default:
               // Do nothing
       } 
    }
}
```

+ Now we can modify the `Store` class to depend on the `PaymentProcessor` abstraction:
```java
public class Store {
    private PaymentProcessor paymentProcessor;

    public void purchaseProduct() {
        paymentProcessor.processPayment();
    }
}
```

+ Lets now say that we wanted to add a new payment processor like PayPal.
+ We can easily modify the `PaymentProcessor` class to include a new case for PayPal:
```java
public interface PaymentProcessor {
    private void callStripeAPI() {
        // Calls the Stripe API
    }

    private void callPayPalAPI() {
        // Calls the PayPal API
    }

    public void processPayment(serviceName) {
       switch(serviceName) {
           case "Stripe":
               callStripeAPI();
               break;
           case "PayPal":
               callPayPalAPI();
               break;
           default:
               // Do nothing
       } 
    }
}
```

+ So with this change, what is happening is that our `Store` class stays the exact same since it depends on an abstraction of `PaymentProcessor`.
+ Therefore, the dependency inversion principle states that high-level modules should not depend on low-level modules. Both should depend on abstractions.
