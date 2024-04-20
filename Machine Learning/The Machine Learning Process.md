# The Machine Learning Process
+ When building a machine learning model, there are some steps you need to follow:

## 1. Data Pre-Processing
+ This is the process of preparing the dataset that you have to train your model.
+ This step comes with a few tasks:
    1. **Importing the data**: You need to import the dataset that you will use to train your model.
    2. **Clean the data**: You have to remove any missing values, duplicates, or irrelevant data.
    3. **Split train & test sets**: Create two segments of the data, one to train the model, and the other to test it.

## 2. Modelling
+ This is the process of training your model on the dataset.
+ This also comes with a few tasks:
    1. **Choose a model**: You need to choose a model that you will use to train your data.
    2. **Train the model**: You need to train the model on the dataset.
    3. **Make predictions**: Use the model to make predictions on the test set.

## 3. Evaluation
+ Finally, once you are done with training the model, you need to evaluate it.
    1. **Calculate performance metrics**: This process involves calculating metrics like accuracy, precision, recall, etc.
    2. **Make a verdict**: Based on the performance metrics, you need to make a decision on whether the model is good enough or not.

# Training & Testing Sets
+ In data pre-processing, we mentioned that you need to split the data into two segments: training and testing sets.
+ For example, lets say that you taught a toddler how to do addition using the following examples:
    1. 2 + 2 = 4
    2. 3 + 3 = 6
    3. 4 + 4 = 8
    4. 5 + 5 = 10

+ The toddler now know exactly the answers to the above examples.
+ Basically, the toddler knows the answers to the questions.
+ But how do we know if the toddler has properly learnt how to do addition?
+ We have to present the toddler with a new set of questions:
    1. 6 + 6 = ?
    2. 7 + 7 = ?
    3. 8 + 8 = ?
    4. 9 + 9 = ?

+ Now, we can properly verify if the toddler has learnt how to do addition.
+ What you just saw is a simple example of how training and testing sets work in machine learning.
+ The set of questions the toddler was trained on is known as the **training set**.
+ The set of questions the toddler was tested on is known as the **testing set**.

# How Training & Testing Sets Work In Machine Learning?
+ If you have a dataset with 1000 rows, you can split it into two segments:
    1. **Training set**: This would consist of 800 rows.
    2. **Testing set**: This would consist of 200 rows.

+ Once you are done training your model with the training set, you can test it if it can make accurate predictions on the testing set.
