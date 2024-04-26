# Single Responsibility Principle
+ This principle states that a class should only have responsibility & one reason to change.
+ For example, take the following code for example:
```java
public class User {
    private String name;
    private String email;
    private String password;
    private String DB_URL = "jdbc:mysql://localhost:3306/mydb";
    private string DB_USER = "mysql";
    private string EMAIL_HOST = "smtp.gmail.com";

    public void saveUser() {
        conn = new Connection(DB_URL, DB_USER);
        conn.connect();
        conn.save(this);
    }

    public void sendEmail() {
        client = new EmailClient(EMAIL_HOST);
        conn = client.connect();
        conn.send(this);
    }
}
```

+ The `User` class has two responsibilities and three reasons to change:

| Class | Responsibilities      | Reasons To Change                                         |
|-------|-----------------------|-----------------------------------------------------------|
| User  | Create a user         | A client might want to change the way the user is saved   |
|       | Save user to database | A client might want to switch from MySQL to PostgreSQL    |
|       | Send email to user    | A client might want to switch from Gmail to Yahoo         |

+ A function having either three responsibilities or reasons to change violates the Single Responsibility Principle. To fix this, we can refactor the code as follows:
```java
public class User {
    private String name;
    private String email;
    private String password;
}
```

```java
public class UserDB {
    private String DB_URL = "jdbc:mysql://localhost:3306/mydb";
    private string DB_USER = "mysql";

    public void saveUser(User user) {
        conn = new Connection(DB_URL, DB_USER);
        conn.connect();
        conn.save(user);
    }
}

```

```java
public class EmailService {
    private string EMAIL_HOST = "smtp.gmail.com";

    public void sendEmail(User user) {
        client = new EmailClient(EMAIL_HOST);
        conn = client.connect();
        conn.send(user);
    }
}
```

+ Now, we have three classes, with one responsibility and a reason for each. Lets list down the responsibilities and reasons to change:

| Class         | Responsibilities      | Reasons To Change                                         |
|---------------|-----------------------|-----------------------------------------------------------|
| User          | Create a user         | A client might want to change the way the user is saved   |
| UserDB        | Save user to database | A client might want to switch from MySQL to PostgreSQL    |
| EmailService  | Send email to user    | A client might want to switch from Gmail to Yahoo         |

# Cohesion
+ Cohesion is the degree to which the elements inside a class are related.
+ High cohesion means that the elements inside the class are very closely related.
+ Low cohesion means that the elements inside the class are not closely related.
+ For example, lets look at a low cohesion class:
```java
public class MiscellaneousTasks {
    private int[] numbers;

    public MiscellaneousTasks(int[] numbers) {
        this.numbers = numbers;
    }

    public void printCalendar(int month, int year) {
        // Simple example: this method prints a calendar for a given month and year
        System.out.println("Calendar for " + month + "/" + year);
        // Logic to print the calendar would go here
    }

    public int getDate() {
        // Simple example: this method returns the current date
        return 0;
    }

    public double calculateAverage() {
        // Calculate the average of numbers
        double sum = 0;
        for (int num : numbers) {
            sum += num;
        }
        return numbers.length > 0 ? sum / numbers.length : 0;
    }

    public void sendEmail(String recipient, String message) {
        // Simulates sending an email
        System.out.println("Sending email to " + recipient);
        System.out.println("Message: " + message);
    }
}
```

+ We have the following functions:
    1. `printCalendar`: Prints a calendar for a given month and year.
    2. `getDate`: Returns the current date.
    3. `calculateAverage`: Calculates the average of numbers.
    4. `sendEmail`: Simulates sending an email.

+ Look at these functions and see how they are not related to each other.
    + `printCalendar` and `sendEmail` are not related at all.
    + `getDate` and `calculateAverage` are also not related.
    + However, `getDate` and `printCalendar` are somewhat related.

+ For a class should have high cohesion, the functions inside the class should be closely related. For example, we can refactor the code as follows:
```java
public class Calendar {
    public void printCalendar(int month, int year) {
        // Simple example: this method prints a calendar for a given month and year
        System.out.println("Calendar for " + month + "/" + year);
        // Logic to print the calendar would go here
    }

    public int getDate() {
        // Simple example: this method returns the current date
        return 0;
    }
}
```

```java
public class Math {
    private int[] numbers;

    public Math(int[] numbers) {
        this.numbers = numbers;
    }

    public double calculateAverage() {
        // Calculate the average of numbers
        double sum = 0;
        for (int num : numbers) {
            sum += num;
        }
        return numbers.length > 0 ? sum / numbers.length : 0;
    }
}
```

```java
public class Email {
    public void sendEmail(String recipient, String message) {
        // Simulates sending an email
        System.out.println("Sending email to " + recipient);
        System.out.println("Message: " + message);
    }
}
```

+ Now, each function in the class are very closely related to each other, and therefore, this class has high cohesion.

# Coupling
+ Coupling is the degree to which one class's function depends on another function.
+ High coupling means that the function of a class are very dependent on each other.
+ This is undesirable when we want to make a change, as we will have to change multiple classes.
+ The following is an example of high coupling:
```java
public class User {
    private String name;
    private String email;
    private String password;

    public User(String name, String email, String password) {
        this.name = name;
        this.email = email;
        this.password = password;
    }

    public void saveUser() {
        conn = new Connection("jdbc:mysql://localhost:3306/mydb", "mysql");
        conn.connect();
        conn.save(this);
    }

    public void sendEmail() {
        client = new EmailClient("smtp.gmail.com");
        conn = client.connect();
        conn.send(this);
    }
}
```

+ We cannot change the database or email client without changing the `User` class. This is an example of high coupling.
+ So, we can refactor the code as follows:
```java
public class User {
    private String name;
    private String email;
    private String password;

    public User(String name, String email, String password) {
        this.name = name;
        this.email = email;
        this.password = password;
    }
```

```java
public class UserDB {
    public void saveUser(User user) {
        conn = new Connection("jdbc:mysql://localhost:3306/mydb", "mysql");
        conn.connect();
        conn.save(user);
    }
}
```
+ Now, if we wanted to switch from MongoDB to PostgreSQL, we only need to change the `UserDB` class, and not the `User` class. This is an example of low coupling.
