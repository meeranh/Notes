# Data Cleaning
+ If you get a CSV file full of data, it is very likely that you may have some inconsistent data values in your CSV file's columns.
+ For data science objectives, you need to have a very clean dataset.
+ The following is an example of a very clean dataset, lets call this dataset A:

| Name      | Phone Number  | Gender    |
|-----------|---------------|-----------|
| Meeran    | 076-6395-949  | Male      |
| Kamal     | 077-6788-541  | Male      |
| Kamala    | 076-2315-929  | Female    |

+ The following is **not** a clean dataset, lets call this dataset B:

| Name      | Phone Number  | Gender    |
|-----------|---------------|-----------|
| Meeran    | 0766395949    | Male      |
| Kamal     | +94776788541  | M         |
| Kamala    | 762315929     | Female    |

+ The process of cleaning a dataset is basically to convert dataset B into dataset A.
+ In this section, we will discuss some tools from Pandas that will come in handy when cleaning data.

# Removing Duplicates
+ Your dataset might have rows that have the exact same data.
+ Take the following example:

| Name      | Phone Number  | Gender    |
|-----------|---------------|-----------|
| Meeran    | 076-6395-949  | Male      |
| Meeran    | 076-6395-949  | Male      |
| Kamala    | 076-2315-929  | Female    |

+ The first and second columns are duplicates.
+ We can remove duplicates via the `drop_duplicates()` method.
```python
import pandas as pd

df = pd.read_csv('data.csv')

df = df.drop_duplicates()
```

+ Doing a `drop_duplicates()` operation will remove all duplicate rows.
+ Therefore, the above duplicate row example will now look like this:

| Name      | Phone Number  | Gender    |
|-----------|---------------|-----------|
| Meeran    | 076-6395-949  | Male      |
| Kamala    | 076-2315-929  | Female    |

# Dropping Unneeded Columns
+ We very often will run into columns that are not needed.
+ We have to remove these columns for effective data analysis.
+ Lets say we have the following table and want to get rid of `Unneeded_Column`

| Name      | Phone Number  | Gender    | Unneeded_Column |
|-----------|---------------|-----------|-----------------|
| Meeran    | 076-6395-949  | Male      | 1               |
| Kamal     | 077-6788-541  | Male      | 2               |
| Kamala    | 076-2315-929  | Female    | 3               |

+ If we want to get rid of the `Unneeded_Column`, we have to use the `drop()` method.
```python
import pandas as pd

df = pd.read_csv('data.csv')

df = df.drop(columns=['Unneeded_Column'])
```

+ What this will do is it will completely remove the `Unneeded_Column` from the dataset.
+ Therefore, once the operation is done, the dataset will look like this:

| Name      | Phone Number  | Gender    |
|-----------|---------------|-----------|
| Meeran    | 076-6395-949  | Male      |
| Kamal     | 077-6788-541  | Male      |
| Kamala    | 076-2315-929  | Female    |

# Strip
+ Lets say that we have some kind of unneeded characters at the start/end of a cell.
+ Lets take the following example:

| Name      | Phone Number  | Gender    |
|-----------|---------------|-----------|
| Meeran..  | 076-6395-949  | Male      |
| Kamal_    | 077-6788-541  | Male      |
| Ka+mala   | 076-2315-929  | Female    |

+ The `..`, `_` and `+` characters are unneeded.
+ We can remove this using the `str.strip()` method.
```python
import pandas as pd

df = pd.read_csv('data.csv')

df['Name'] = df['Name'].str.strip('.+_')
```

+ The above command will remove all instances of `.`, `_` and `+` from the `Name` column.
+ Therefore, the modified dataset will now look like this:

| Name      | Phone Number  | Gender    |
|-----------|---------------|-----------|
| Meeran    | 076-6395-949  | Male      |
| Kamal     | 077-6788-541  | Male      |
| Kamala    | 076-2315-929  | Female    |

# Replace
+ You have the ability to replace a string of text with another.
+ Say that you have a bunch of data containing the word 'Mister' and you want to replace it with 'Mr'.
+ Lets also say that you want to remove the starting 0s from the phone number.
+ An example table is as follows:

| Name          | Phone Number  | Gender    |
|---------------|---------------|-----------|
| Mister Meeran | 076-6395-949  | Male      |
| Mister Kamal  | 077-6788-541  | Male      |
| Missus Kamala | 076-2315-929  | Female    |


+ You can do this using the `str.replace()` method.
+ This function also has Regex support, which you will see in the example below.
```python
import pandas as pd

df = pd.read_csv('data.csv')

df['Name'] = df['Name'].str.replace('Mister', 'Mr')
df['Name'] = df['Name'].str.replace('Missus', 'Mrs')
df['Phone Number'] = df['Phone Number'].str.replace('^0', '')
```

+ Now, when we do this, the dataset will look like this:

| Name       | Phone Number | Gender    |
|------------|--------------|-----------|
| Mr Meeran  | 76-6395-949  | Male      |
| Mr Kamal   | 77-6788-541  | Male      |
| Mrs Kamala | 76-2315-929  | Female    |

# Apply
+ You might want to perform some kind of a function on each cell in a column.
+ The `apply()` method is used for this.
+ Basically, you can create a function that takes in 1 parameter.
+ You can then pass this function into `apply()`.
+ Whatever you return from your custom function will be replaced in the cell.
+ Lets say that you want to capitalize all the data in the name column:
```python
import pandas as pd

df = pd.read_csv('data.csv')

df['Name'] = df['Name'].apply(lambda x: x.capitalize())
```

+ This will convert the `Name` column to the following:

| Name       |
|------------|
| MR MEERAN  |
| MR KAMAL   |
| MRS KAMALA |
