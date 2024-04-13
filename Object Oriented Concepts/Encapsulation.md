# Encapsulation
+ In procedural programming, we can create a variable and assign a value to it.
+ We can also create a function and call it.
+ But in object-oriented programming, we can create a class and define variables and functions in it.
+ This is called encapsulation.

+ This is an example of a program without encapsulation:
```javascript
let name = "John";
let age = 30;
let currentYear = 2021;

function getYearOfBirth(year, age) {
    return year - age;
}
```

+ This is an example of a program with encapsulation:
```javascript
class Person {
    constructor(name, age) {
        this.name = name;
        this.age = age;
    }

    getYearOfBirth(year) {
        return year - this.age;
    }
}
```

+ The benefit of encapsulation is that we can group related variables and functions together in a class.
+ We also can specify methods with zero parameters.
