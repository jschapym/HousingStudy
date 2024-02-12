# Concurrency Assignment

MSDS 431

# Overview

Linear Regression was run on the Boston House Study dataset using Golang. This is used to predict several values and variables. One program is run with concurrency and one without.

**mainwcon.go** is the program that implements concurrency to run 100 iterations of regression analysis.

**mainwocon.go** is the program that does not use concurrency, instead performing analyses sequentially.

**boston.csv** contains the data from the Boston Housing Study.

The results were placed into .txt files labeled **mainwcon.txt** and **mainwocon.txt**.

Based on the comparison, we can see that, when performing regressions without concurrency, it takes approximately 0.15s. With concurrency, it takes approximately 0.16s. This may be because goroutines must each be created to perform each iteration of the linear regression analysis. 



