# Interface Segregation Principle
+ This principle states that a class should not have to implement an interface if it does not use the methods of that interface.
+ Lets assume that you have an office with multi-printers which can print, scan, and fax.
+ As an OOP developer, you might create an interface called `MultiFunctionDevice` that has methods for all three functionalities.
```java
pubilc interface MultiFunctionDevice {
    public void print() {
        // Print
    }

    public void getPrintSpoolDetails() {
        // Get print spool details
    }

    public void scan() {
        // Scan
    }

    public void scanPhoto() {
        // Scan photo
    }

    public void internetFax() {
        // Internet fax
    }
}
```
+ Later you realize, you have a printer that can only print and scan.
+ You now have to create a `Printer` class that implements the `MultiFunctionDevice` interface.
```java
public interface Printer extends MultiFunctionDevice {
    public void print() {
        // Print
    }

    public void getPrintSpoolDetails() {
        // Get print spool details
    }

    public void scan() {
        // Scan
    }

    public void scanPhoto() {
        // Scan photo
    }

    // These two are empty functions that are not used by the Printer class
    public void fax() {}
    public void internetFax() {}
    
}
```

+ We had to still implement two empty functions `fax` and `internetFax` in the `Printer` class.
+ This is bad practice empty functions should not be implemented.
+ We now discover that we have a device with only fax and internet fax capabilities.
+ We also create an interface for this:
```java
public interface FaxMachine extends MultiFunctionDevice {
    public void fax() {
        // Fax
    }

    public void internetFax() {
        // Internet fax
    }

    // These four are empty functions that are not used by the FaxMachine class
    public void print() {}
    public void getPrintSpoolDetails() {}
    public void scan() {}
    public void scanPhoto() {}

}
```
+ The `FaxMachine` class has too many empty functions and it is not a good practice to implement empty functions.
+ We can solve this by creating simple high cohesion interfaces for each functionality.
+ We can create a `IPrinter`, `IScanner`, and `IFaxMachine` interface and inherit them based off the functionality our devices have.
```java
public interface IPrinter {
    public void print();
    public void getPrintSpoolDetails();
}

public interface IScanner {
    public void scan();
    public void scanPhoto();
}

public interface IFaxMachine {
    public void fax();
    public void internetFax();
}
```

```java
public class Printer implements IPrinter, IScanner {
    public void print() {
        // Print
    }

    public void getPrintSpoolDetails() {
        // Get print spool details
    }

    public void scan() {
        // Scan
    }

    public void scanPhoto() {
        // Scan photo
    }
}
```

```java
public class FaxMachine implements IFaxMachine {
    public void fax() {
        // Fax
    }

    public void internetFax() {
        // Internet fax
    }
}
```

```java
public class MultiFunctionDevice implements IPrinter, IScanner, IFaxMachine {
    public void print() {
        // Print
    }

    public void getPrintSpoolDetails() {
        // Get print spool details
    }

    public void scan() {
        // Scan
    }

    public void scanPhoto() {
        // Scan photo
    }

    public void fax() {
        // Fax
    }

    public void internetFax() {
        // Internet fax
    }
}
```

+ Now we have refactored and cleaned our code so that each class implements only the methods that it needs.
+ This type of code does not violate the Interface Segregation Principle.

# Signs of Violation
1. **Fat interfaces**: Interfaces that have too many methods.
2. **Empty methods**: Methods that are not used by the class.
3. **Low cohesion**: Classes that implement interfaces that are not related.
