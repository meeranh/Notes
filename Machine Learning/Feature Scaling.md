# Features
+ Feature are the columns of the dataset.
+ For example, lets say we had the following as our dataset:

| Name   | Age | Salary |
|--------|-----|--------|
| John   | 45  | 70,000 |
| Adam   | 44  | 60,000 |
| Meeran | 40  | 52,000 |

+ The features of this dataset are `Age` and `Salary`.
+ Lets say that you are given the following problem:
    + *Which two people are most similar to each other?*

+ To do this, you first need someway to compare `Age` and `Salary`.
+ The problem here is that `Age` and `Salary` are two different units.
+ Comparing `Age` and `Salary` is like comparing apples and oranges, they are not the same.
+ Therefore, we first need to convert the data into a common format. This is know as `Feature Scaling`.

# Feature Scaling
+ There are two common ways to scale features:
    1. **Normalization**
    2. **Standardization**

## Normalization
+ What normalization basically does is it converts all the data values to a number between 0 and 1.
+ We first have to identify some values:
    + **$X_{min}$** = The lowest value in the feature/column.
    + **$X_{max}$** = The highest value in the feature/column.
    + **$X$** = The value of the feature/column.

+ The formula used to normalize the data is:
**$$X_{norm} = \frac{X - X_{min}}{X_{max} - X_{min}}$$**

+ If we were to apply this to our dataset, we would have to do the following for each row:

| Name   | Age (Normalized) | Salary (Normalized) |
|--------|------------------|---------------------|
| John   | 1                | 1                   |
| Adam   | 0.8              | 0.444               |
| Meeran | 0                | 0                   |

+ Now the data is in a common format and we can compare the two columns.
+ So basically, this process is known as `Normalization`.
