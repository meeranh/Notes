# What Is Matplotlib
+ In Python, when you want to graph some data, you need to use a graphing library such as Matplotlib.
+ This library allows you to visually represent arrays, lists, and other data structures in a graphical format.

+ **Note**: Matplotlib is extremely customizable when it comes to styling, therefore, I will not be covering all the styling options in this note. You may refer to their well-written documentation on all the styling options.

# Scatter Plot
+ This is a basic plot that shows a dot for each data point.
+ It is useful for showing the relationship between two variables.
+ The way to use a scatter plot is to call the `scatter()` function and pass in the x and y values.
```python
import numpy as np
import matplotlib.pyplot as plt

x = np.random.rand(100)
y = np.random.rand(100)

plt.scatter(x, y)

plt.title('Scatter Plot')
plt.xlabel('X-axis')
plt.ylabel('Y-axis')

plt.show()
```
![[Pasted image 20240420145656.png]]
+ You can also plot two different data points on the same graph with different colors.
+ To do this, you will have to call `scatter()` twice with different x and y values.
+ It helps if you had different colors for each data point.
```python
import numpy as np
import matplotlib.pyplot as plt

x1 = np.random.rand(100)
y1 = np.random.rand(100)

x2 = np.random.rand(100)
y2 = np.random.rand(100)

plt.scatter(x1, y1, color='red', label='Data 1')
plt.scatter(x2, y2, color='blue', label='Data 2')

plt.title('Colored Scatter Plot')
plt.xlabel('X-axis')
plt.ylabel('Y-axis')
plt.legend()

plt.show()
```
![[Pasted image 20240420150247.png]]
# Line Plot
+ Lets say you have a set of X and Y values.
+ You can plot these values on a graph using a line plot using the `plot()` function.
```python
import matplotlib.pyplot as plt

x = [1, 2, 3, 4, 5]
y = [1, 4, 9, 16, 8]

plt.plot(x, y)

plt.title('Line Plot')
plt.xlabel('X-axis')
plt.ylabel('Y-axis')

plt.show()
```
![[Pasted image 20240420145759.png]]
+ Again, you can do this with multiple data points in different colors:
```python
import matplotlib.pyplot as plt

x1 = [3, 4, 5, 6, 7]
y1 = [1, 1, 2, 4, 9]

x2 = [3, 6, 8, 9, 10]
y2 = [2, 5, 8, 7, 6]

plt.plot(x1, y1, color='red', label='Data 1')
plt.plot(x2, y2, color='blue', label='Data 2')

plt.title('Colored Line Plot')
plt.xlabel('X-axis')
plt.ylabel('Y-axis')
plt.legend()

plt.show()
```
![[Pasted image 20240420150404.png]]
# Bar Plots
+ Bar plots are used to show the relationship between a numerical variable and a categorical variable.
+ You can use the `bar()` function to plot a bar plot.
+ Lets say you want to create a bar chart of the number of votes for a programming language:
```python
import matplotlib.pyplot as plt

languages = ['Python', 'Rust', 'Go', 'Java', 'Kotlin']
votes = [100, 200, 90, 30, 70]

plt.bar(languages, votes)

plt.title('Bar Plot')
plt.xlabel('Programming Languages')
plt.ylabel('Votes')

plt.show()
```
![[Pasted image 20240420145836.png]]
# Histogram
+ Lets say you have a list of mean ages of people in a country.
+ You can plot a histogram to show the distribution of ages.
+ You can use the `hist()` function to plot a histogram.
```python
import numpy as np
import matplotlib.pyplot as plt

ages = np.random.normal(30, 10, 1000)

plt.hist(ages)

plt.title('Histogram')
plt.xlabel('Ages')
plt.ylabel('Frequency')

plt.show()
```
![[Pasted image 20240420145850.png]]
# Pie Chart
+ Pie charts are used to show the proportion of different categories in a dataset.
+ You can use the `pie()` function to plot a pie chart.
+ You will have to feed your data and the labels for each data point as `pie(data, labels=labels)`.
+ Lets say you have a list of the number of votes for different programming languages:
```python
import matplotlib.pyplot as plt

languages = ['Python', 'Rust', 'Go', 'Java', 'Kotlin']
votes = [100, 200, 90, 30, 70]

plt.pie(votes, labels=languages)
plt.title('Pie Chart')

plt.show()
```
![[Pasted image 20240420145910.png]]
# Box Plots
+ Lets say you want to create a box plot to show the distribution of heights of people in a country.
+ You can use the `boxplot()` function to plot a box plot.
```python
import numpy as np
import matplotlib.pyplot as plt

heights = np.random.normal(170, 10, 1000)

plt.boxplot(heights)

plt.title('Box Plot')
plt.ylabel('Heights')

plt.show()
```
![[Pasted image 20240420145958.png]]

# Subplotting
+ Lets say that you wanted to plot multiple graphs on the same figure.
+ You can use the `subplot()` function to create multiple plots on the same figure.
+ Lets say you need four plots on the same figure:
+ You first need to allocate space for the plots using `subplot(2, 2)`.
+ What this does is it creates a 2x2 grid of plots.
+ So you will have 4 plots in total:
    1. The first plot will be at the top left *(1, 1)*.
    2. The second plot will be at the top right *(1, 2)*.
    3. The third plot will be at the bottom left *(2, 1)*.
    4. The fourth plot will be at the bottom right *(2, 2)*.
```python
import numpy as np
import matplotlib.pyplot as plt

x = np.random.rand(100)
y = np.random.rand(100)

languages = ['Python', 'Rust', 'Go', 'Java', 'Kotlin']
votes = [100, 200, 90, 30, 70]

ages = np.random.normal(30, 10, 1000)

fig, axs = plt.subplots(2, 2)

axs[0, 0].scatter(x, y)
axs[0, 0].set_title('Scatter Plot')

axs[0, 1].plot(x, y)
axs[0, 1].set_title('Line Plot')

axs[1, 0].bar(languages, votes)
axs[1, 0].set_title('Bar Plot')

axs[1, 1].hist(ages)
axs[1, 1].set_title('Histogram')

plt.show()
```
![[Pasted image 20240420151131.png]]

# 3D Plotting
+ Lets say that you wanted to plot an x, y, and z coordinate in 3D.
+ You will have to create a 3D axis using `ax = plt.axes(projection='3d')`.
+ You can then do whatever you want with the 3D axis.
```python
import numpy as np
import matplotlib.pyplot as plt

ax = plt.axes(projection='3d')

x = np.random.rand(100)
y = np.random.rand(100)
z = np.random.rand(100)

ax.scatter(x, y, z)

ax.set_title('3D Scatter Plot')
ax.set_xlabel('X-axis')
ax.set_ylabel('Y-axis')

plt.show()
```
![[Pasted image 20240420151646.png]]

+ You may use any kinds of plots in 3D such as line plots, bar plots, etc.
+ We will do a line plot in 3D:
```python
import numpy as np
import matplotlib.pyplot as plt

ax = plt.axes(projection='3d')

x = np.random.rand(100)
y = np.random.rand(100)
z = np.random.rand(100)

ax.plot(x, y, z)

ax.set_title('3D Line Plot')
ax.set_xlabel('X-axis')
ax.set_ylabel('Y-axis')

plt.show()
```
![[Pasted image 20240420151852.png]]

# Surface Plots
+ Lets say that you want to plot the surface of a 3D function.
+ A 3D function will have X, Y, and Z values.
+ You can use the `plot_surface()` function to plot the surface of a 3D function.
```python
import numpy as np
import matplotlib.pyplot as plt

def func_3d(x, y):
    return np.sin(np.sqrt(x**2 + y**2))

x = np.linspace(-6, 6, 30)
y = np.linspace(-6, 6, 30)

X, Y = np.meshgrid(x, y)
Z = func_3d(X, Y)

ax = plt.axes(projection='3d')
ax.plot_surface(X, Y, Z)

ax.set_title('3D Surface Plot')
ax.set_xlabel('X-axis')

plt.show()
```
![[Pasted image 20240420152928.png]]