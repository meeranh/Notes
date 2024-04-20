# Why Use Numpy?
+ Numpy is a library in Python that allows you to work with arrays.
+ But why would you use Numpy arrays over Python lists? Aren't both the same? No.

1. Numpy is fast
2. Numpy is memory efficient

+ Lets take the following array: `[1, 2, 3, 4, 5]`
+ If this was declared using Numpy, the total size it would take on memory would be 20 bytes.
+ If this was declared using Python lists, the total size it would take on memory would be 104 bytes.
+ Because there is less memory usage, it is faster to read and write data.

+ The thing with Numpy arrays is that they are **homogeneous**.
+ This means that all the elements in the array must be of the same data type.
+ You cannot have a string and an integer in the same array.

# The Basics
### Declaring An Array In Numpy
+ To create a Numpy array, you can use the `np.array()` function.
```python
import numpy as np

arr = np.array([1, 2, 3, 4, 5])
print(arr) # [1 2 3 4 5]
```

+ If you wanted to create a two dimensional array, you can do so by passing in a list of lists.
```python
import numpy as np

arr = np.array([[1, 2, 3], [4, 5, 6]])
```

### Finding The Dimension Of An Array
+ In the previous example, we declared arrays that were both one and two dimensional.
+ We have a function called `ndim` that will return the number of dimensions of the array.
```python
import numpy as np

arr1 = np.array([1, 2, 3, 4, 5])
arr2 = np.array([
    [1, 2, 3],
    [4, 5, 6]
])

print(arr1.ndim) # 1
print(arr2.ndim) # 2
```

### Finding The Shape Of An Array
+ You might want to get the number of rows and columns of an array.
+ Lets take the following example:
```python
import numpy as np

arr = np.array([
    [1, 2, 3],
    [4, 5, 6]
])
arr.shape # (2, 3)
```

+ Basically, the `shape()` function will return a tuple with the number of rows and columns.

### Find The Data Type Of An Array
+ Because Numpy arrays are homogeneous, we can determine the data type of an array using the `dtype` property.
```python
import numpy as np

arr = np.array([1, 2, 3, 4, 5])
print(arr.dtype) # int32

arr = np.array([1.0, 2.0, 3.0, 4.0, 5.0])
print(arr.dtype) # float64

arr = np.array(['apple', 'banana', 'cherry'])
print(arr.dtype) # <U6
```

### Finding Array Count
+ You can find the number of elements in an array using the `size` property.
```python
import numpy as np

arr = np.array([1, 2, 3, 4, 5])

print(arr.size) # 5
```

### Storing Custom Data Types
+ The default data type for integer Numpy arrays is `int32`.
+ However, if your data is smaller, you can use `int8` or `int16` to save memory.
+ You can find the number of bytes that you are using per array with the `itemsize` property.
```python
import numpy as np

arr1 = np.array([1, 2, 3, 4, 5])
arr2 = np.array([1, 2, 3, 4, 5], dtype='int16')
arr3 = np.array([1, 2, 3, 4, 5], dtype='int8')

# Size of each element in bytes
print(arr1.itemsize) # 4
print(arr2.itemsize) # 2
print(arr3.itemsize) # 1

# Size of the entire array
print(arr1.size * arr1.itemsize) # 20
print(arr2.size * arr2.itemsize) # 10
print(arr3.size * arr3.itemsize) # 5
```

# Accessing & Changing Rows/Columns
### Accessing Elements In An Array
+ Lets say you have a 3 dimensional array.
+ You want to access the element in the second row, third column.
+ You can access it using the following syntax: `arr[rowIndex, colIndex]`
```python
import numpy as np

arr = np.array([
    [1, 2, 3],
    [4, 5, 6]
])

print(arr[1, 2]) # 6
print(arr[0, 0]) # 1
print(arr[1, -2]) # 5
```

+ Lets try this with a 3 dimensional array.
```python
import numpy as np

arr = np.array([
    [1, 2, 3, 4, 5],
    [6, 7, 8, 9, 10],
    [11, 12, 13, 14, 15]
])
```


### Changing Elements In An Array
+ Like in the previous example, you could locate an element using the `arr[rowIndex, colIndex]` syntax.
+ Then you could assign a new value to that element.
```python
import numpy as np

arr = np.array([
    [1, 2, 3],
    [4, 5, 6]
])

arr[1, 2] = 9
print(arr) # [[1  2  3] [4  5  9]]
```
### Get An Entire Row/Column
+ If you wanted to print the whole column or the whole row, you can use the following syntax:
    + **Row**: `arr[rowIndex, :]`
    + **Column**: `arr[:, colIndex]`

```python
import numpy as np

arr = np.array([
    [1, 2, 3],
    [4, 5, 6]
])

print(arr[0, :]) # [1 2 3]
print(arr[:, 1]) # [2 5]
```

### Slicing Rows/Columns
+ If you wanted to modify the 3rd column of every row to a certain number, you can use the following syntax:
    + `arr[:, colIndex] = newValue`

```python
import numpy as np

arr = np.array([
    [1, 2, 3],
    [4, 5, 6]
])

arr[:, 2] = 9

print(arr) # [[1 2 9] [4 5 9]]
```

### Stepping Through Rows/Columns
+ Lets say you wanted to print every second element in the first row.
+ You can step through every second element in the first row using the following syntax:
    + `arr[rowIndex, start:stop:step]`

```python
import numpy as np

arr = np.array([
    [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
    [11, 12, 13, 14, 15, 16, 17, 18, 19, 20],
    [21, 22, 23, 24, 25, 26, 27, 28, 29, 30]
])

print(arr[0, 1:10:2]) # [2 4 6 8 10]
print(arr[0, 1:10:3]) # [2 5 8]
```

# Initialization Of Matrices
### Zero Matrix
+ If you wanted to create a matrix with all zeros, you can use the `np.zeros()` function.
+ Syntax: `np.zeros((rows, cols))`
```python
import numpy as np

arr = np.zeros((2, 3))
print(arr) # [[0. 0. 0.] [0. 0. 0.]]

arr = np.zeros((2, 3), dtype='int8')
print(arr) # [[0 0 0] [0 0 0]]
```

### Ones Matrix
+ If you wanted to create a matrix with all ones, you can use the `np.ones()` function.
+ Syntax: `np.ones((rows, cols))`
```python
import numpy as np

arr = np.ones((2, 3))
print(arr) # [[1. 1. 1.] [1. 1. 1.]]

arr = np.ones((2, 3), dtype='int8')
print(arr) # [[1 1 1] [1 1 1]]
```

### Custom Value Matrix
+ If you wanted to create a matrix with a custom value, you can use the `np.full()` function.
+ Lets say we wanted an array with all 9s:
+ Syntax: `np.full((rows, cols), value)`
```python
import numpy as np

arr = np.full((2, 3), 9)
print(arr) # [[9 9 9] [9 9 9]]
```

### Random Decimal Numbers From 0 to 1
+ If you wanted to create a matrix with random decimal numbers from 0 to 1, you can use the `np.random.rand()` function.
+ Syntax: `np.random.rand(rows, cols)`
```python
import numpy as np

arr = np.random.rand(2, 3)

print(arr) # [[0.42829726 0.96451452 0.5762536 ] [0.36603201 0.57044302 0.52083468]]
```

### Random Integer Numbers
+ If you wanted to create a matrix with random integer numbers, you can use the `np.random.randint()` function.
+ Syntax: `np.random.randint(start, stop, size=(rows, cols))`
+ If you wanted an array with 2 rows and 3 columns with random numbers from 0 to 9:
```python
import numpy as np

arr = np.random.randint(0, 10, size=(2, 3))

print(arr) # [[5 7 8] [0 1 3]]
```

### A Random Normal Distribution
+ If you wanted to simulate a Gaussian distribution, you can use the `np.random.normal()` function.
+ Syntax: `np.random.normal(mean, std, size=your_size)`
```python
import numpy as np

arr = np.random.normal(0, 1, size=3)

print(arr) # [-0.07368877 -0.07368877 -0.07368877]
```

### A Sequence Of Linear Numbers
+ Imagine you wanted to create an array that had 1, 5, 10, 15, 20.
+ Basically adding 4 to the previous number from 1 to 20
+ You can use the `np.linspace()` function.
+ Syntax: `np.linspace(start, stop, number_of_elements)`
```python
import numpy as np

arr = np.linspace(1, 20, 5).astype('int8')

print(arr) # [ 1 5 10 15 20]
```

### Repeated Array
+ If you wanted to create a matrix with a repeated array, you can use the `np.repeat()` function.
+ Syntax: `np.repeat(array, times)`

```python
import numpy as np

arr = np.array([1, 2, 3])
arr = np.repeat(arr, 3)

print(arr) # [1 1 1 2 2 2 3 3 3]
```

# Copying Arrays
+ Lets say you wanted to make a copy of an array.
+ You might do the common mistake as follows:
```python
import numpy as np

arr1 = np.array([1, 2, 3])
arr2 = arr1

arr2[0] = 100

print(arr1) # [100 2 3]
print(arr2) # [100 2 3]
```

+ As you can see, changing `arr2` also changed `arr1`.
+ This is because `arr2` is just a reference to `arr1`.
+ You did not make a copy of `arr1`, you just pointed the memory address of `arr2` to `arr1`.
+ Luckily, Numpy has a function called `copy()` that will make a copy of the array.
```python
import numpy as np

arr1 = np.array([1, 2, 3])
arr2 = arr1.copy()

arr2[0] = 100

print(arr1) # [1 2 3]
print(arr2) # [100 2 3]
```

# Mathematics In Numpy
+ It is very easy to do mathematical operations in Numpy.
+ The following code snippet shows how flexible Numpy is with mathematical operations.
```python
import numpy as np

arr = np.array([1, 2, 3, 4, 5])

print(arr + 2) # [3 4 5 6 7]
print(arr - 2) # [-1  0  1  2  3]
print(arr * 2) # [2 4 6 8 10]
print(arr / 2) # [0.5 1.  1.5 2.  2.5]
print(arr ** 2) # [ 1  4  9 16 25]
print(arr += 2) # [3 4 5 6 7]
print(np.sin(arr)) # [ 0.84147098  0.90929743  0.14112001 -0.7568025  -0.95892427]
```

# Statistics In Numpy
+ Numpy has built-in functions for statistics.
+ Lets say you have an array of numbers.
+ You want to find the min, max, mean, and sum of the array.
```python
import numpy as np

arr = np.array([
    [1, 2, 3],
    [4, 5, 6]
])

print(np.min(arr)) # 1
print(np.max(arr)) # 6
print(np.mean(arr)) # 3.5
print(np.sum(arr)) # 21
```

# Horizontal & Vertical Stacking
+ Lets say you have two arrays.
+ You want to stack them horizontally or vertically.
+ You can use the `np.hstack()` and `np.vstack()` functions.

### Horizontal Stacking
+ If you wanted to stack two arrays horizontally, you can use the `np.hstack()` function.
```python
import numpy as np

arr1 = np.array([1, 2, 3])
arr2 = np.array([4, 5, 6])

arr = np.hstack((arr1, arr2))
print(arr) # [1 2 3 4 5 6]
```

### Vertical Stacking
+ If you wanted to stack two arrays vertically, you can use the `np.vstack()` function.
```python
import numpy as np

arr1 = np.array([1, 2, 3])
arr2 = np.array([4, 5, 6])

arr = np.vstack((arr1, arr2))
print(arr) # [[1 2 3] [4 5 6]]
```

# Reading From Files
+ Lets say you have a file with an array of numbers.
+ Imagine this is your file: `data.txt`
```
1 2 3 4 5
6 7 8 9 10
```

+ You can read this file using the `np.genfromtxt()` function.
```python
import numpy as np

arr = np.genfromtxt('data.txt', delimiter=' ')
print(arr) # [[ 1.  2.  3.  4.  5.] [ 6.  7.  8.  9. 10.]]
```

# Boolean Masking
+ Lets say you have an array of numbers.
+ If you want to return an array of boolean values based on a condition, you can use the following syntax:
```python
import numpy as np

arr = np.array([1, 2, 3, 4, 5])

print(arr > 3) # [False False False True True]
```
