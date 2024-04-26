# Linear Regression
+ Lets say that we want to predict the amount of marks a student might get for every hour of studying per day.
+ Lets say that we already have a set of data that:
    1. Has the number of hours a student studied per day.
    2. Has the marks that the student got for studying that many hours.

| Hours Studied | Marks |
|---------------|-------|
| 1             | 15    |
| 2             | 29    |
| 3             | 42    |
| 4             | 40    |
| 5             | 60    |
| 6             | 65    |
| 7             | 75    |
| 8             | 70    |
| 9             | 90    |
| 10            | 95    |

+ We need to first plot these onto a graph.
![[Pasted image 20240422194742.png]]

+ Then, we need to draw a line that goes through all the data points.
+ This line is known as the regression line:
![[Pasted image 20240422194853.png]]

+ As humans, we cannot perfectly find this line using our eyes.
+ We have to use a formula to find the **slope** and the **height** of this line.

# Crafting The Perfect Regression Line
+ The equation to draw a straight line is:
**$$y = mx + c$$**
    - **$y$**: The marks a student might get.
    - **$x$**: The number of hours a student studied.
    - **$m$**: The slope of the line.
    - **$c$**: The height of the line.

+ However, the equation to find the slope of the line (**$m$**), the equation is:
**$$m = \frac{n(\sum xy) - (\sum x)(\sum y)}{n(\sum x^2) - (\sum x)^2}$$**
    - **$n$**: The number of data points.
    - **$\sum xy$**: The sum of the product of the number of hours studied and the marks.
    - **$\sum x$**: The sum of the number of hours studied.
    - **$\sum y$**: The sum of the marks.
    - **$\sum x^2$**: The sum of the square of the number of hours studied.

+ The formula might look intimidating, but it is really simple.
+ To approach this problem, lets bring back the table of data we referred previously:

| Hours Studied | Marks |
|---------------|-------|
| 1             | 15    |
| 2             | 29    |
| 3             | 42    |
| 4             | 40    |
| 5             | 60    |
| 6             | 65    |
| 7             | 75    |
| 8             | 70    |
| 9             | 90    |
| 10            | 95    |

+ Lets name the *Hours Studied* column as **$x$** and the *Marks* column as **$y$**.

| **$x$** | **$y$** |
|---------|---------|
| 1       | 15      |
| 2       | 29      |
| 3       | 42      |
| 4       | 40      |
| 5       | 60      |
| 6       | 65      |
| 7       | 75      |
| 8       | 70      |
| 9       | 90      |
| 10      | 95      |

+ Now, lets find the total sum of four things:
    1. The sum of all the **$x$** values in the table
    2. The sum of all the **$y$** values in the table
    3. Multiply the **$x$** and **$y$** values and sum them up.
    4. The sum of the square of the **$x$** values.
    5. The total sums of all the above values.

| **$x$** | **$y$** | **$x*y$** | **$x^2$** |
|---------|---------|-----------|-----------|
| 1       | 15      | 15        | 1         |
| 2       | 29      | 58        | 4         |
| 3       | 42      | 126       | 9         |
| 4       | 40      | 160       | 16        |
| 5       | 60      | 300       | 25        |
| 6       | 65      | 390       | 36        |
| 7       | 75      | 525       | 49        |
| 8       | 70      | 560       | 64        |
| 9       | 90      | 810       | 81        |
| 10      | 95      | 950       | 100       |
| **55**  | **581** | **3894**  | **385**   |

+ We actually calculated the values we needed to find the slope (**$m$**) of the line.
    + **$n$**: 10
    + **$\sum xy$**: 3894
    + **$\sum x$**: 55
    + **$\sum y$**: 581
    + **$\sum x^2$**: 385

+ Now, lets plug these values into the formula to find the slope of the line:

**$$m = \frac{n(\sum xy) - (\sum x)(\sum y)}{n(\sum x^2) - (\sum x)^2} = \frac{10(3894) - (55)(581)}{10(385) - (55)^2} = 8.47$$**

+ We now found the slope of the line: **$m = 8.47$**.
+ We now have to find the height of the line (**$c$**).
+ Since the equation is **$y = mx + c$**, we can rearrange the equation to find **$c$**:
    + **$c = y - mx$**

+ We have both **$m$** and **$c$** with a pair of **$x$** and **$y$** values.
+ Therefore, we can just plug in the values to find **$c$**.
    + **$x = 1$**
    + **$y = 15$**
    + **$m = 8.47$**
    + **$c = y - mx = 15 - (8.47)(1) = 6.53$**

+ Finally, we now have the equation of the line:
    + **$$y = 8.47x + 6.53$$**

+ If we plot this line with our data points, it will look like the following:
![[Pasted image 20240422201804.png]]

+ You can find how linear regression actually works [here](https://www.youtube.com/watch?v=YwZYSTQs-Hk).

# Graphing A Linear Regression Line Using Python
+ We will now implement the exact same example in Python.
+ We will be using `LinearRegression` from `sklearn.linear_model` to find the slope and height of the line.

```python
import numpy as np
from sklearn.linear_model import LinearRegression
import matplotlib.pyplot as plt

# Study time vs scores
x = np.array([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]).reshape(-1, 1)
y = np.array([15, 29, 42, 40, 60, 65, 75, 70, 90, 95])

# Create a Linear Regression model
model = LinearRegression().fit(x, y)

# Get the slope and height of the line
slope = model.coef_[0]
height = model.intercept_

# Plot the data points
plt.scatter(x, y)

# Plot the regression line
plt.plot(x, slope*x + height, color='red')

plt.show()
```

+ This will give us the following linear regression line:
![[Pasted image 20240422221443.png]]

# Predicting The Marks Of A Student Using Python
+ The previous section proved that the `LinearRegression` model can find the slope and height of the line.
+ We will now give in a number of hours a student studied and predict the marks the student might get.

```python
import numpy as np
import pandas as pd
from sklearn.linear_model import LinearRegression
from sklearn.model_selection import train_test_split
from sklearn.metrics import r2_score

# Convert the dictionary to a DataFrame
df = pd.DataFrame({
    'Hours': [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
    'Marks': [15, 29, 42, 40, 60, 65, 75, 70, 90, 95]
})

# Split the dataset into train/test, ensuring X is 2D
X = df['Hours'].values.reshape(-1, 1)
y = df['Marks']
x_train, x_test, y_train, y_test = train_test_split(X, y, test_size=0.2)

# Create a Linear Regression model
model = LinearRegression()
model.fit(x_train, y_train)

# Calculate the accuracy score
accuracy = model.score(x_test, y_test)

# Predict the marks on the test set
y_pred = model.predict(x_test)

print("Accuracy:", accuracy)
print("Predictions:", y_pred)
```

+ In linear regression models, the `X` values must be a 2D array even if there is only one feature.
+ This is why we are using `reshape(-1, 1)` to convert the `X` values to a 2D array.
