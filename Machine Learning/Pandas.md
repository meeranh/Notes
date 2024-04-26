# Why Use Pandas?
+ Pandas is a Python library which is very useful for importing and manipulating datasets.
+ Imagine you want to do Excel-like operations on a dataset, Pandas is the way to go.
+ It allows you to load very large datasets, which is not possible on Excel.

# Loading Data
### Loading CSV Files
+ In Pandas, you can import a dataset using the `read_csv()` function:
```python
import pandas as pd

df = pd.read_csv('data.csv')

print(df.head(3)) # Prints the first 3 rows of the dataset
print(df.tail(3)) # Prints the last 3 rows of the dataset
```

+ Once you load a dataset into a variable using Pandas, it is called a DataFrame.
+ You can think of a DataFrame as a table with rows and columns (just like an Excel sheet).

### Loading Excel Files
+ You can load an Excel file using the `read_excel()` function:
```python
import pandas as pd

df = pd.read_excel('data.xlsx')
```

### Loading CSVs With Custom Delimiters
+ The default delimiter for `read_csv()` is a comma (`,`).
+ However, you may have a dataset with a different delimiter (e.g. a semicolon `;`).
+ We will look at `Tab` separated files as an example:
```python
import pandas as pd

df = pd.read_csv('data.txt', delimiter='\t')
```
+ The key takeaway here is that you can specify a custom delimiter.

### Loading Parts Of Datasets To Save Memory
+ Lets say that you have a very large dataset on a PC with like 1GB or RAM.
+ You will not be able to load the entire dataset into memory.
+ You can load parts of the dataset, do some operations, and then load the next part.
+ Or you could even load the first 1000 rows of the dataset:
```python
import pandas as pd

df = pd.read_csv('data.csv', nrows=1000) # Loads the first 1000 rows
```

+ If you wanted to load 10 rows at a time, do some operations, and then load the next 10 rows:
```python
import pandas as pd

for df in pd.read_csv('data.csv', chunksize=10):
    print(df) # Prints out 10 rows at a time
```

# Reading Data
+ Lets say that you have now imported a CSV file.
+ This imported file is now a DataFrame.
+ You can do multiple operations onto this DataFrame.

### Print Column Headers
+ You can print the column headers of the DataFrame using the `columns` attribute:
```python
import pandas as pd

df = pd.read_csv('data.csv')

print(df.columns) # Prints out the dataset column headers
```

### Print An Entire Column
+ You now want to print out all the data that falls under a column:
```python
import pandas as pd

df = pd.read_csv('data.csv')

print(df['column_name']) # Prints out the entire column
```

+ This will also work if you need to print out multiple columns.
+ You will have to pass an array of arrays
```python
import pandas as pd

df = pd.read_csv('data.csv')

print(df[['column_one', 'column_two']]) # Prints out multiple columns
```

### Print A Specific Row
+ Lets say that you now need to print out the first row of the dataset:
```python
import pandas as pd

df = pd.read_csv('data.csv')

print(df.iloc[0]) # Prints out the first row of the dataset
```
+ What `iloc` stands for is `index location`.
+ The first index in Python stars with 0, and hence the first row is at index 0.

+ You can also get multiple rows by passing an array of indexes:
```python
import pandas as pd

df = pd.read_csv('data.csv')

print(df.iloc[[0:3]]) # Prints out the first 3 rows of the dataset
print(df.iloc[[0, 3, 5]]) # Prints out the 1st, 4th and 6th rows of the dataset
```

### Iterate Through Every Row In The Dataset
+ You can iterate through every row in the dataset using the `iterrows()` function:
```python
import pandas as pd

df = pd.read_csv('data.csv')

for index, row in df.iterrows():
    print(index, row) # Prints out all rows with their indexes
```

+ Lets say that you want to print out a 'Name' column for each row:
```python
import pandas as pd

df = pd.read_csv('data.csv')

for index, row in df.iterrows():
    print(index, row['Name']) # Prints out the 'Name' column for each row
```

### Describe Your Data In The Dataset
+ You can get a summary of your dataset using the `describe()` function:
```python
import pandas as pd

df = pd.read_csv('data.csv')

print(df.describe()) # Prints out a summary of the dataset
```

+ This summary contains the following details:
    1. **Count**: The number of rows in the dataset.
    2. **Mean**: The average value of each column.
    3. **Std**: The standard deviation of each column.
    4. **Min**: The minimum value of each column.
    5. **25%**: The 25th percentile of each column.
    6. **50%**: The 50th percentile of each column.
    7. **75%**: The 75th percentile of each column.
    8. **Max**: The maximum value of each column.

### Read Your Dataset In Sorted Order
+ You can print out the content of your dataset in ascending/descending order.
+ This can be done using the `sort_values()` function:
```python
import pandas as pd

df = pd.read_csv('data.csv')

print(df.sort_values('column_name')) # Ascending order
print(df.sort_values('column_name', ascending=False)) # Descending order

# You can also sort multiple columns
print(df.sort_values(['column_one', 'column_two']))
```

+ After specifying two or more columns, you may want one of them to be ascending, and one of them to be descending.
+ Lets say that you want to print out a list of names sorted by their age (oldest first).
+ You also want to then sort the names in alphabetical order:
```python
import pandas as pd

df = pd.read_csv('data.csv')

print(df.sort_values(['Age', 'Name'], ascending=[0, 1]))
```

# Manipulating Data
### Adding A New Total Column
+ Lets say that you have a list of costs to produce a certain product.

| ID | Name | Tax | Vat | Unit Cost |
|----|------|-----|-----|-----------|
| 1  | A    | 10  | 5   | 100       |
| 2  | B    | 20  | 10  | 200       |
| 3  | C    | 30  | 15  | 300       |

+ It might be very helpful to have a total cost column.
+ We can do this very easily in Pandas:
```python
import pandas as pd

df = pd.read_csv('data.csv')

df['Total Cost'] = df['Unit Cost'] + df['Tax'] + df['Vat']
```

+ What will happen is now you will generate a new table that looks like the following:

| ID | Name | Tax | Vat | Unit Cost | Total Cost |
|----|------|-----|-----|-----------|------------|
| 1  | A    | 10  | 5   | 100       | 115        |
| 2  | B    | 20  | 10  | 200       | 230        |
| 3  | C    | 30  | 15  | 300       | 345        |

+ You can now see that a new column has been added to the DataFrame.

+ You could have done it in a different method, by using `iloc`:
```python
import pandas as pd

df = pd.read_csv('data.csv')

df['Total Cost'] = df.iloc[:, 2:5].sum(axis=1)
```
+ Basically, you are adding the 3rd to 5th columns and storing the result in a new column.
+ `axis=1` means that you are adding the columns horizontally.
+ If you want to add the columns vertically, you can specify `axis=0`.

### Deleting A Column
+ You might want to now delete that `Total Cost` column that you previously created.
+ The `drop()` function will return a new DataFrame without the specified column:
+ Lets say that you specify `df.drop(columns=['Total Cost'])`:
+ It will return a new DataFrame without the `Total Cost` column, but it will not delete it.
+ You will have to reassign the DataFrame to the new DataFrame:
```python
import pandas as pd

df = pd.read_csv('data.csv')

df = df.drop(columns=['Total Cost'])
```

### Exporting A DataFrame To A CSV
+ Once you are done modifying your DataFrame, you might want to save it to a CSV file.
+ You can do this using the `to_csv()` function:
```python
import pandas as pd

df = pd.read_csv('data.csv')

# Do some operations on the DataFrame

df.to_csv('new_data.csv')
```

+ The problem with just using `df.to_csv(filename)` is that it will save the DataFrame with the index.
+ So you might get a CSV file such as the following:

|   | ID | Name | Tax | Vat | Unit Cost |
|---|----|------|-----|-----|-----------|
| 0 | 1  | A    | 10  | 5   | 100       |
| 1 | 2  | B    | 20  | 10  | 200       |
| 2 | 3  | C    | 30  | 15  | 300       |

+ As you can see, there is an index column. The way to prevent this is by including `index=False`:
```python
import pandas as pd

df = pd.read_csv('data.csv')

# Do some operations on the DataFrame

df.to_csv('new_data.csv', index=False)
```

### Filtering Data In The Dataset
+ You might want to filter for all rows which contain "Meeran" in the "Name" column.
+ You can do this using the `loc` function:
```python
import pandas as pd

df = pd.read_csv('data.csv')

df.loc[df['Name'] == 'Meeran']
```

+ You can also filter for certain logic in columns.
+ Lets say that you want to filter for all rows where the "Age" column is greater than 25:
```python
import pandas as pd

df = pd.read_csv('data.csv')

df.loc[df['Age'] > 25]
```
+ This will return a list of rows where the "Age" column is greater than 25.
+ You can then use this list to perform further operations.

+ Also, you can specify multiple conditions when filtering.
+ Lets say you want everyone older than 25 **and** whose "Name" is "Meeran":
```python
import pandas as pd

df = pd.read_csv('data.csv')

df.loc[(df['Age'] > 25) & (df['Name'] == 'Meeran')]
```

+ If you wanted everyone older than 25 **or** whose "Name" is "Meeran":
```python
import pandas as pd

df = pd.read_csv('data.csv')

df.loc[(df['Age'] > 25) | (df['Name'] == 'Meeran')]
```

#### Resetting The Index
+ Very often, when you filter data and create a new DataFrame, you will be left with the old index.
+ For example lets imagine that we have this table:

| Index | Name | Age |
|-------|------|-----|
| 0     | A    | 20  |
| 1     | B    | 30  |
| 2     | C    | 40  |
| 3     | D    | 50  |

+ We write a simple filter operation that returns only the rows where the "Age" column is greater than 30:
```python
import pandas as pd

df = pd.read_csv('data.csv')

new_df = df.loc[df['Age'] > 30]
```

+ What will be printed is the following

| Index | Name | Age |
|-------|------|-----|
| 2     | C    | 40  |
| 3     | D    | 50  |

+ You might not want your indexes to start at a non-zero value.
+ So you have a couple of options to reset your index back to 0.

###### Option 1: Using `reset_index()`
+ You can use the `reset_index()` function to reset the index:
```python
import pandas as pd

df = pd.read_csv('data.csv')

df = df.loc[df['Age'] > 30]
new_df = df.reset_index()
```

+ This will return a new DataFrame with the index reset to 0.
+ However, you will have an additional column called `index` which contains the old index values.
+ If you don't want this column, you can use `drop=True`:
```python
import pandas as pd

df = pd.read_csv('data.csv')

df = df.loc[df['Age'] > 30]
new_df = df.reset_index(drop=True)
```

+ This will return a new DataFrame with the index reset to 0 and without the `index` column.
+ But the issue is that this creates a new DataFrame and assigns it to a new variable.
+ If we have a dataset with 100GB of data, this will be very inefficient.
+ Therefore, you can use the `inplace=True` parameter.
+ This will modify the existing DataFrame without creating a new one:
```python
import pandas as pd

df = pd.read_csv('data.csv')

df = df.loc[df['Age'] > 30]
df.reset_index(drop=True, inplace=True)
```

#### Filtering via Regex
+ If you wanted to filter for all rows where the "Name" column contains "Meeran" and starts with "M":
```python
import pandas as pd
import re

df = pd.read_csv('data.csv')

df.loc[df['Name'].str.contains('^M'), regex=True, flags=re.IGNORECASE] # Starts with M
df.loc[df['Name'].str.contains('Meeran')] # Contains Meeran
```

+ You can use any Regex flag you want.

### Conditional Changes
+ Lets say you have a dataset with a column called "Fruits" with a list of fruits.
+ You don't want to call them "Apples" anymore, you want to call them "Red Apples".
+ You can do this via a conditional change.
+ The syntax for a conditional change is as follows:
`df.loc[condition, 'column_name'] = 'New Value'`

+ Lets look at this in action:
```python
import pandas as pd

df = pd.read_csv('data.csv')

df.loc[df['Fruits'] == 'Apples', 'Fruits'] = 'Red Apples'
```

+ What this does is it goes through all the rows in the "Fruits" column.
+ If the value is "Apples", it will first fetch the 'Fruits' column and then change the value to "Red Apples".

+ You can do these kinds of operations on any kind of conditions (including multiple conditions).
+ For example, lets say that you wanted to change all 'Status' of a person to 'Rich' whose 'Salary' is greater than 1000. You also want to change their 'Poor' status to False:
```python
import pandas as pd

df = pd.read_csv('data.csv')

df.loc[df['Salary'] > 1000, ['Status', 'Poor']] = ['Rich', False]
```

# Aggregation Using GroupBy
+ Using the `groupby()` function, you can group data together based off of a column.
+ For example, you can calculate the average, sum, count, etc. of a column based on the group.
+ Lets say that we have the following dataset:

| Employee | Department | Age | Salary |
|----------|------------|-----|--------|
| A        | HR         | 30  | 1000   |
| B        | HR         | 40  | 2000   |
| C        | IT         | 25  | 1500   |
| D        | IT         | 35  | 2500   |
| E        | IT         | 45  | 3000   |
| F        | Marketing  | 20  | 500    |
| G        | Marketing  | 30  | 1000   |

+ We now want to find the average salary of a person in each department.
+ We can do this using the `groupby()` function:
```python
import pandas as pd

df = pd.read_csv('data.csv')

df.groupby('Department').mean()
```

+ What this will do is it will group the data by the 'Department' column.
+ It will then calculate the mean of the 'Age' and 'Salary' columns and give the following output:

| Department | Age | Salary |
|------------|-----|--------|
| HR         | 35  | 1500   |
| IT         | 35  | 2333.33|
| Marketing  | 25  | 750    |

+ The `groupby()` function allows three main operations:
    1. **count()**: Counts the number of non-null values in each group.
    2. **mean()**: Calculates the mean of the values in each group.
    3. **sum()**: Calculates the sum of the values in each group.

```python
import pandas as pd

df = pd.read_csv('data.csv')

df.groupby('Department').count()
df.groupby('Department').mean()
df.groupby('Department').sum()
```
