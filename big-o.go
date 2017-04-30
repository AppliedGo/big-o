/*
<!--
Copyright (c) 2017 Christoph Berger. Some rights reserved.

Use of the text in this file is governed by a Creative Commons Attribution Non-Commercial
Share-Alike License that can be found in the LICENSE.txt file.

Use of the code in this file is governed by a BSD 3-clause license that can be found
in the LICENSE.txt file.

The source code contained in this file may import third-party source code
whose licenses are provided in the respective license files.
-->

<!--
NOTE: The comments in this file are NOT godoc compliant. This is not an oversight.

Comments and code in this file are used for describing and explaining a particular topic to the reader. While this file is a syntactically valid Go source file, its main purpose is to get converted into a blog article. The comments were created for learning and not for code documentation.
-->

+++
title = "Big Oh!"
description = "When optimizing your Go code, be aware of Big-O."
author = "Christoph Berger"
email = "chris@appliedgo.net"
date = "2017-04-30"
draft = "true"
domains = ["Algorithms And Data Structures"]
tags = ["Performance", "Big-O", "Optimization"]
categories = ["Tutorial"]
+++

You worked hard to save a few CPU cycles in the central loop, but your code is still slow? Time to think about the time complexity of your algorithm.

<!--more-->

Imagine you just finished writing some code to solve a particular problem. The input is an array of values, and your code calculates some result from this array. You benchmark the code and find that it is reasonably fast for your test data.

Later, you get your hands on some real big data set from your production systems. A hundred times larger than your test data. "Ok," you tell yourself, "if I run the benchmark with this input, my code should take a hundred times as long as with my test data."

To your surprise, the code takes a thousand times longer! What happened?


## The problem of input size

Whenever you design a function that works through a set of input data, ask yourself,

**"How does my algorithm scale with the input size?"**

In other words, how does execution time change when the size of the input is doubled, tripled, quadrupled, and so forth?

To answer this, you have to look how your code processes the input data.


## From input size to execution time

"Input" can mean anything: A file to process, a single number in an array, a character in a string, a string in a list of strings, a record in the result set of a database query, and so forth. The exact type of input is not relevant here. The number of input items counts.

But for now let's assume we have an array of integers to process. And let's assume some imaginary code that does a couple of things with this array:

First, a loop runs over the array and multiplies each number by two. For `n` items, the loop needs `n` units of time.

```
loop
    process array
end
```

Then, two nested loops run over the array. The outer one does nothing but executing the inner one, so we can neglect the time this loop uses. The inner loop runs `n*n` times, and needs four times the time than the single loop to process each item. So for our `n` input items, the double loop runs for `4*n*n` time units.

```
loop
    loop
       heavy processing
    end
end
```

Finally, the code inserts each item into a [binary search tree](https://appliedgo.net/bintree) that is [balanced](https://appliedgo.net/balancedtree/). Each insert takes about `log(n+1)` units of time (the height of a tree of `n` items is `log(n+1)`), so all inserts take `n*log(n+1)` units of time. ("`log`" here means the logarithm to base 2.)

```
loop
    insert array into balanced tree
end
```


In summary, we have an execution time of `4*n^2 + n + n*log(n)` time units.

Now this term is a bit unwieldy, isn't it?

Luckily, we do not need to take the complete, unmodified term to answer our question. After all, we just want to get a rough estimation and not a precise timing.


## Simplifying the term

Plotted as a graph where the x axis is our input size `n` and the y axis are the time units needed to process `n` items, our term looks like this:

![4*n^2 + n + n*log(n+1)](4nn+n+nlogn+1.png)

We can already see that with a raising number of items (x axis), processing time raises quickly - one item takes 5 time units, two items 20 units, three items 40 units, and so forth.

Can we leave some parts of the term away without changing this graph too much?

Obviously, we need to identify the part of the term that has the biggest influence on the shape of the graph. Let's divide the term into its components:

![separated](separated.png)

The orange graph corresponds to the term `4*n^2`, the blue one to `n`, and the green one to `n*log(n+1)`. It looks like wie can reduce our term to `4*n^2` without changing the characteristic shape of the graph.

Can we simplify more? Let's take away the constant factor "`4`":

![4x^2 and x^2](parabolas.png)

The new curve (cyan) looks less steep than the original one. However, their shapes are very similar, and so is their behavior when the input size is growing. The curves first grow slowly, then faster, and finally they shoot towards infinity, as we can see in this image where both axes have been scaled by a factor of 100:

![4x^2 and x^2 scaled](parabolas_scaled.png)

So for estimating purposes, we can even leave out the constant factor. Our term has now shrunken to a simple `n^2`, and the shape still resembles the original curve.

Now we safely can say that our imaginary code needs *roughly* `n^2` units of time to process `n` items. Much easier than the bulky term we started with!

**Takeaway: From all the terms that describe the time a function needs for processing an input of specific size, we only need that part of the term that has the largest effect.**


## The Big-O notation

The simplified term `n^2` stands for a whole class of terms, so let's give it a name.

- - -
A function that needs roughly `n^2` units of time to process `n` input items has a **time complexity of `O(n^2)`**.
- - -

In general, we can say that

**A Big-O function is a mathematical term that gives a *rough estimation* of how the speed of an algorithm changes with the size of the input data.**

Now we have a means to categorize time complexity into classes. Here are the most common classes:

### O(1) - constant time

**Characteristic:** The code always needs the same time, no matter how large the input set is.

**Curve:**

![O(1)](constant.png)

**Example:** Accessing a single element in an array of `n` items is an `O(1)` operation.


### O(log(n)) - logarithmic time

**Characteristic:** The time the code needs to process `n` elements raises much slower than `n` for large values of `n`.

**Curve:**

![O(log(n))](logarithmic.png)

**Examples:** Inserting or retrieving elements into or from a balanced binary tree. Performing a binary search on a sorted array.


### O(n) - linear time

**Characteristic:** Time needed for processing `n` element grows linearly with `n`.

**Curve:**

![O(n)](linear.png)

**Examples:** Iterating over an array, reading a file from start to end,


### O(n*log(n)) - linearithmic time

Whoa, who invents these names?! "Linearithmic" is an artificial blend of "linear" and "logarithmic", in the spirit of ["Brangelina"](https://en.wiktionary.org/wiki/Brangelina). (Bonus question: Which came first?)

**Characteristic:** Somewhere between `O(n)` and `O(n^2)`. Looks slightly like a flat `O(n^2)` curve for small values of `n`, and turns into an almost linear curve for large values of `n`.

**Curve:**

![O(n*log(n))](linearithmic.png)

**Examples:** Heapsort, merge sort, binary tree sort


### O(n^m) - polynomial time

For a given `m > 0`.

Our `O(n^2)` problem belongs this class. This specific case is also called "quadratic time".

(Nitpicking: `O(n)` would also belong to the polynomial class, since `n` is the same as `n^1` and therefore a polynomial term, but due to its special "linear" nature it defines a separate class.)

**Characteristic:** Processing time raises fast. For `O(n^2)`, when input size doubles, execution time quadruples.

Here is where the danger zone comes into sight. With increasing `n`, polynomial algorithms get slow fast. Still, [Cobham's thesis](https://en.wikipedia.org/wiki/Cobham%27s_thesis) considers polynomial problems as *"easy, fast, and practical".* Which means there are worse time complexities on the horizon.

**Curve:**

![O(n^m)](polynomial.png)

**Examples:** Bubble sort (O(n^2)). Any nested loops: Two nested loops - `O(n^2)`. Three nested loops - `O(n^3)` and so on.


### O(2^n) - exponential time

**Characteristic:** **Here be dragons.** Even for moderate values of `n`, processing time goes through the roof.

**Curve:**

![O(2^n)](exponential.png)

**Example:** The famous Traveling Salesman Problem: Find the shortest route between `n` towns.


The different characteristics of these time complexity classes become apparent when we put all those functions into one figure:

(Note: "`n`" has been replaced by "`x`" in the following two figures, as the graph tool did not let me specify `n` instead of `x` in the formulas.)

![Complexity graphs](complexitygraphs.png)

Here we can see immediately how the computation time diverges between the different complexity classes.

The graphs of `O(n^2)` and `O(2^n)` look as if they move very close together; however, after scaling the y axis to [0-200], the difference becomes obvious.

![Complexity graphs with x axis scaled to 0-200](complexitygraphs_scaled.png)

## Improving O(n)

Looking at some of the curves, one question arises: Can we improve a given code, mabye even shift it into a faster Big-O class?


### Change the algorithm

Let me give you a very practical answer here:

If the algorithm you strive to improve is one of the well-known, well-researched algorithms, the chances to find an equivalent algorithm in a better complexity class is roughly zero. If you do find one, you are the hero of the day.

If you are looking at some piece of lousy home-grown code, hacked together quickly while recovering from a severe hangover, then your chances aren't that bad. Read on...


### Pre-process the input data

In some cases, the time complexity can be improved by re-arranging the input data.

For example, f you need to search through a list repeatedly (and much more often than adding something to this list), it will pay off to sort this list before searching it. Sure, sorting needs time, but you would have to do it only once to speed up those, say, thousands of searches that occur afterwards.

A good sort algorithm takes O(n*log(n)) time but reduces each of the subsequent searches from O(n) to O(log(n)). Or all `n` of them from `O(n^2)` to `O(n*log(n)`. Big win!


### Change your data structure

Sometimes it pays off inspecting the way your input data is stored.

Example: Inserting a new value into a sorted linear list takes `O(n*log(n))` time: `O(log(n)` time to find the right cell, and `O(n)` time to shift all subsequent cells by one cell. (Note we are talking about average times here, in individual cases, these operations may take less time. E.g., if the new value is inserted at the end, there is nothing to shift.)

If you insert new data very frequently, better turn this list into a balanced tree. Then inserting a new value only takes `O(log(n))` time. (Remember, `log(n)` is about the height of a balanced binary tree.)


### Trade in time for space

So far we discussed only time complexity, but the same problem (and the same complexity classes) exists for memory consumption as well. Why not trading in one for the other?

## Can concurrency change the complexity class?

##

## The code
*/

// ## Imports and globals
package main

func main() {}

/*
## How to get and run the code

Step 1: `go get` the code. Note the `-d` flag that prevents auto-installing
the binary into `$GOPATH/bin`.

    go get -d github.com/appliedgo/TODO:

Step 2: `cd` to the source code directory.

    cd $GOPATH/src/github.com/appliedgo/TODO:

Step 3. Run the binary.

    go run TODO:.go


## Odds and ends
## Some remarks
## Tips
## Links

[Wikipedia: Time complexity](https://en.wikipedia.org/wiki/Time_complexity)

[Wikipedia: Big O notation](https://en.wikipedia.org/wiki/Big_O_notation)

**Happy coding!**

*/
