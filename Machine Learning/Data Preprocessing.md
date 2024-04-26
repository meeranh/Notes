# Encoding Categorical Data (One Hot Encoding)
+ Lets say that you have the following dataset:

| Name      | Role      |
|-----------|-----------|
| John      | Engineer  |
| Adam      | Doctor    |
| Meeran    | Engineer  |
| Kamal     | Doctor    |
| Kamala    | Cleaner   |

+ There are three categories here:
    1. Engineer
    2. Doctor
    3. Cleaner

+ There should be a way to convert these categories into numbers.
+ Unfortunately, we cannot just convert them into numbers like this:
    + Engineer = 1
    + Doctor = 2
    + Cleaner = 3

+ This is because models will assume that these numbers have some kind of order.
+ For example, cleaner has a higher number than engineer, so it might interpret that cleaner is better than engineer.
+ Therefore, we have to separate them into separate columns as follows:

| Name      | Engineer | Doctor | Cleaner |
|-----------|----------|--------|---------|
| John      | 1        | 0      | 0       |
| Adam      | 0        | 1      | 0       |
| Meeran    | 1        | 0      | 0       |
| Kamal     | 0        | 1      | 0       |
| Kamala    | 0        | 0      | 1       |

+ This kind of a transformation is called **One Hot Encoding**.
+ We can do this in Python using two methods:
    1. Using `pd.get_dummies()` method from pandas.
    2. Using `OneHotEncoder` class from `sklearn.preprocessing` module.

### Pandas Version
```python
import pandas as pd

df = pd.DataFrame({
    'Name': ['John', 'Adam', 'Meeran', 'Kamal', 'Kamala'],
    'Role': ['Engineer', 'Doctor', 'Engineer', 'Doctor', 'Cleaner']
})

df = pd.get_dummies(df, columns=['Role'])
```

### Sklearn Version
```python
from sklearn.preprocessing import OneHotEncoder

encoder = OneHotEncoder()

df = pd.DataFrame({
    'Name': ['John', 'Adam', 'Meeran', 'Kamal', 'Kamala'],
    'Role': ['Engineer', 'Doctor', 'Engineer', 'Doctor', 'Cleaner']
})

encoded = encoder.fit_transform(df[['Role']]).toarray()

df = pd.concat([df, pd.DataFrame(encoded, columns=encoder.categories_[0])], axis=1)
```

+ Both produce the same result, but the `pd.get_dummies()` method is more concise.

# Imputation
+ Imagine that you had a dataset, but some of the data in the `Age` column was missing.
+ Consider the following example:

| Name   | Age | Salary |
|--------|-----|--------|
| John   | 45  | 70,000 |
| Meeran | NaN | 52,000 |
| Adam   | 44  | 60,000 |

+ Imputation allows you to replace `NaN` with either the mean, median, or mode of the column.
+ Lets look at how we can do this via the `SimpleImputer` class from the `sklearn.impute` module.
```python
from sklearn.impute import SimpleImputer

imputer = SimpleImputer(strategy='mean')

df = pd.DataFrame({
    'Name': ['John', 'Meeran', 'Adam'],
    'Age': [45, np.nan, 44],
    'Salary': [70000, 52000, 60000]
})

df['Age'] = imputer.fit_transform(df[['Age']])
```

+ The above example (and the default strategy) is `mean`.
+ You can change this to either `median` or `mode` by changing the `strategy` parameter.
+ If you wanted to use `mean`, you do not have to explicitly specify it.
+ This is because `mean` is the default strategy.
+ So you initialize the `SimpleImputer` class without any parameters.

# Binary Choices
+ Most commonly, you might have a categorical column that has only two values.
+ The perfect example of this is `Gender` which can only be `Male` or `Female`.
+ Instead of hot encoding, you can just convert them into binary values.

```python
import pandas as pd

df = pd.DataFrame({
    'Name': ['John', 'Meeran', 'Devika'],
    'Gender': ['M', 'M', 'F']
})

genders = {
    'M': 1,
    'F': 0
}

df['Gender'] = [genders[g] for g in df['Gender']]
```

+ What that will give out is the following dataset:

| Name   | Gender |
|--------|--------|
| John   | 1      |
| Meeran | 1      |
| Devika | 0      |
   
# Feature Scaling

+ Features are the columns of the dataset.
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

+ There are two common ways to scale features:
    1. **Normalization**
    2. **Standardization**

## Min-Max Normalization
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
+ So basically, this process is known as `Min-Max Normalization`.
+ The reason it is knows as `Min-Max` is because we are using the minimum and maximum values of the feature to normalize the data.

+ In Python, the way to apply normalization is as follows:
```python
from sklearn.preprocessing import MinMaxScaler

scaler = MinMaxScaler()

df[['Age', 'Salary']] = scaler.fit_transform(df[['Age', 'Salary']])
```

+ What the above code does is it scales the `Age` and `Salary` columns of the dataset `df` using the `MinMaxScaler` class from the `sklearn.preprocessing` module.

## Standardization (Z-Score Normalization)
+ In normalization, all the data was scaled to a number between 0 and 1.
+ In standardization, the mean of the data is 0 and the standard deviation is 1.
+ So, the data is evenly distributed around 0.
+ Standardization is useful when building linear/logistic regression models where the distance between data points is important.
+ Also, standardization is useful when there are too many outliers in your dataset.
+ The equation for standardization is:
**$$X_{std} = \frac{X - \mu}{\sigma}$$**
    + **$X_{std}$** = The standardized value.
    + **$X$** = The value of the feature/column.
    + **$\mu$** = The mean of the feature/column.
    + **$\sigma$** = The standard deviation of the feature/column.

+ If we were to apply this to our dataset, we would have to do the following for each row:

| Name   | Age (Standardized) | Salary (Standardized) |
|--------|--------------------|-----------------------|
| John   | 0.926              | 1.268                 |
| Adam   | 0.463              | -0.091                |
| Meeran | -1.389             | -1.177                |

+ If we wanted to scale our data based on standardization, we would have to do the following in Python:
```python
from sklearn.preprocessing import StandardScaler

scaler = StandardScaler()

df[['Age', 'Salary']] = scaler.fit_transform(df[['Age', 'Salary']])
```

# Pre-Processing Pipelines
+ In the above sections, we have discussed different methods of pre-processing data.
+ The issue is, lets say we have multiple datasets that come on a regular basis.
+ We will have to perform pre-processing on each and every data.
+ So we need to run functions one by one and pre-process all datasets.
+ Pipelines are a feature from `sklearn` that allow you to run multiple functions at once.
+ It can be imported as `from sklearn.pipeline import Pipeline`.
+ `sklearn.base` has two classes that are used to create pipelines:
    1. `BaseEstimator`
    2. `TransformerMixin`

+ We have to inherit from these classes to create our own custom classes.
+ The base class has the following two methods:
    1. `fit()`
    2. `transform()`

+ We must implement both these methods or else we will get an error.
+ Lets see an example implementation:

```python
from sklearn.base import BaseEstimator, TransformerMixin
import pandas as pd
import numpy as np
from sklearn.impute import SimpleImputer
from sklearn.preprocessing import MinMaxScaler
from sklearn.pipeline import Pipeline

df = pd.DataFrame({
    'Name': ['John', 'Meeran', 'Adam'],
    'Age': [45, np.nan, 44],
    'Salary': [70000, 52000, 60000],
    'Gender': ['M', 'M', 'F'],
    'Unneeded_Column': [1, 2, 3]
})

class UnneededColumnDropper(BaseEstimator, TransformerMixin):
    def fit(self, X, y=None):
        return self

    def transform(self, X):
        X = X.drop(columns=['Unneeded_Column'])
        return X

class GenderEncoder(BaseEstimator, TransformerMixin):
    def fit(self, X, y=None):
        return self

    def transform(self, X):
        # Numeric Gender
        genders = {
            'M': 1,
            'F': 0
        }

        X['Gender'] = [genders[g] for g in X['Gender']]

        return X

class AgeImputer(BaseEstimator, TransformerMixin):
    def fit(self, X, y=None):
        return self

    def transform(self, X):
        imputer = SimpleImputer(strategy='mean')
        X['Age'] = imputer.fit_transform(X[['Age']])
        return X

class FeatureScaler(BaseEstimator, TransformerMixin):
    def fit(self, X, y=None):
        return self

    def transform(self, X):
        scaler = MinMaxScaler()
        X[['Age', 'Salary']] = scaler.fit_transform(X[['Age', 'Salary']])
        return X

pipeline = Pipeline([
    ('column_dropper', UnneededColumnDropper()),
    ('gender_encoder', GenderEncoder()),
    ('age_imputer', AgeImputer()),
    ('feature_scaler', FeatureScaler())
])

pipeline.fit_transform(df)
```

+ What the above defined pipeline does is it runs the following command:
```python
df = UnneededColumnDropper().fit_transform(df)
df = GenderEncoder().fit_transform(df)
df = AgeImputer().fit_transform(df)
df = FeatureScaler().fit_transform(df)
```

+ It does basically the above commands in one line: `pipeline.fit_transform(df)`.

# Splitting the Dataset Into Training and Testing Sets
+ If you need to calculate an accuracy score for a model, you need to split the dataset into training and testing sets.
+ You can do this via the `train_test_split()` method from the `sklearn.model_selection` module.

### Train-Test Split
+ Imagine you wanted to build a model that predicts if a person is a thief or not.
+ Lets say we have the following dataset:

| Name      | Age | Occupation  | Thief |
|-----------|-----|-------------|-------|
| John      | 45  | Engineer    | No    |
| Mahinda   | 55  | Politician  | Yes   |
| Meeran    | 40  | Doctor      | No    |
| Kamal     | 35  | Engineer    | No    |
| Ranil     | 60  | Politician  | Yes   |
| Kamala    | 45  | Doctor      | No    |
| Madush    | 25  | Thief       | Yes   |
| Mahesh    | 30  | Engineer    | No    |
| Leela     | 35  | Cleaner     | No    |
| Kusum     | 40  | Doctor      | No    |

+ In the above dataset, the column that tells us if a person is a thief or not is the `Thief` column.
+ This column is known as the **Label** and the other columns are known as **Features**.
+ We need to actually split the *Label* and the *Features* into two separate sets.
+ If we do that, we will have the following two sets:
    1. **Features Set**:
    
        | Index | Name      | Age | Occupation  |
        |-------|-----------|-----|-------------|
        | 0     | John      | 45  | Engineer    |
        | 1     | Mahinda   | 55  | Politician  |
        | 2     | Meeran    | 40  | Doctor      |
        | 3     | Kamal     | 35  | Engineer    |
        | 4     | Ranil     | 60  | Politician  |
        | 5     | Kamala    | 45  | Doctor      |
        | 6     | Madush    | 25  | Thief       |
        | 7     | Mahesh    | 30  | Engineer    |
        | 8     | Leela     | 35  | Cleaner     |
        | 9     | Kusum     | 40  | Doctor      |
    
    2. **Label Set**:
    
        | Index | Thief |
        |-------|-------|
        | 0     | No    |
        | 1     | Yes   |
        | 2     | No    |
        | 3     | No    |
        | 4     | Yes   |
        | 5     | No    |
        | 6     | Yes   |
        | 7     | No    |
        | 8     | No    |
        | 9     | No    |

+ Perfect, now we need to split these into training and testing sets.
+ We specify a percentage for splitting, lets assume that this value is 80% training and 20% testing.
+ That means, we now will have four sets:
    1. **Training Features Set**: 80% of the features.
    2. **Training Label Set**: 80% of the labels.
    3. **Testing Features Set**: 20% of the features.
    4. **Testing Label Set**: 20% of the labels.

+ Therefore, if we were to split the above dataset, we would have the following sets:
    1. **Training Features Set**:
    
        | Index | Name      | Age | Occupation  |
        |-------|-----------|-----|-------------|
        | 0     | John      | 45  | Engineer    |
        | 1     | Mahinda   | 55  | Politician  |
        | 2     | Meeran    | 40  | Doctor      |
        | 3     | Kamal     | 35  | Engineer    |
        | 4     | Ranil     | 60  | Politician  |
        | 5     | Kamala    | 45  | Doctor      |
        | 6     | Madush    | 25  | Thief       |
        | 7     | Mahesh    | 30  | Engineer    |
    
    2. **Training Label Set**:
    
        | Index | Thief |
        |-------|-------|
        | 0     | No    |
        | 1     | Yes   |
        | 2     | No    |
        | 3     | No    |
        | 4     | Yes   |
        | 5     | No    |
        | 6     | Yes   |
        | 7     | No    |
    
    3. **Testing Features Set**:
    
        | Index | Name      | Age | Occupation  |
        |-------|-----------|-----|-------------|
        | 8     | Leela     | 35  | Cleaner     |
        | 9     | Kusum     | 40  | Doctor      |
    
    4. **Testing Label Set**:
    
        | Index | Thief |
        |-------|-------|
        | 8     | No    |
        | 9     | No    |

+ This is basically the process of splitting the dataset into training and testing sets.
+ Note that the splitting is done randomly, so the indices might not be the same as the above example.
+ Now, lets implement this in Python:

```python
from sklearn.model_selection import train_test_split

df = pd.DataFrame({
    'Name': ['John', 'Mahinda', 'Meeran', 'Kamal', 'Ranil',
             'Kamala', 'Madush', 'Mahesh', 'Leela', 'Kusum'],

    'Occupation': ['Engineer', 'Politician', 'Doctor', 'Engineer', 'Politician',
                   'Doctor', 'Thief', 'Engineer', 'Cleaner', 'Doctor'],

    'Age': [45, 55, 40, 35, 60,
            45, 25, 30, 35, 40],

    'Thief': [0, 1, 0, 0, 1,
              0, 1, 0, 0, 0]
})

features = df[['Name', 'Age', 'Occupation']]
labels = df['Thief']

X_train, X_test, Y_train, Y_test = train_test_split(features, labels, test_size=0.2)
```

+ The above code will then split out the training and testing sets similar to the diagrams we had seen earlier.

### Stratified Shuffle Split
+ The issue with the `train_test_split()` method is that it assigns the data to the train and test variables randomly.
+ There is a chance that this will result in a skewed dataset.
+ With the stratified shuffle split, that chance is reduced (almost zero).
+ This will be useful to create a very evenly distributed dataset.
+ The Python code for this is as follows:

```python
from sklearn.model_selection import StratifiedShuffleSplit

df = pd.DataFrame({
    'Name': ['John', 'Mahinda', 'Meeran', 'Kamal', 'Ranil',
             'Kamala', 'Madush', 'Mahesh', 'Leela', 'Kusum'],

    'Occupation': ['Engineer', 'Politician', 'Doctor', 'Engineer', 'Politician',
                   'Doctor', 'Thief', 'Engineer', 'Cleaner', 'Doctor'],

    'Age': [45, 55, 40, 35, 60,
            45, 25, 30, 35, 40],

    'Thief': [0, 1, 0, 0, 1,
              0, 1, 0, 0, 0]
})

features = df[['Name', 'Age', 'Occupation']]
labels = df['Thief']

splitter = StratifiedShuffleSplit(test_size=0.2)

for train_index, test_index in splitter.split(features, labels):
    X_train, X_test = features.loc[train_index], features.loc[test_index]
    Y_train, Y_test = labels.loc[train_index], labels.loc[test_index]
```
+ The `splitter.split()` function returns a `generator` object.
+ Therefore, a `for` loop should be used to iterate over the generator object.
