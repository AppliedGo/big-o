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
date = "2017-04-26"
publishdate = "2017-04-26"
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

Finally, the code inserts each item into a [binary search tree](https://appliedgo.net/bintree) that is [balanced](https://appliedgo.net/balancedtree/). Each insert takes about `log(n+1)` units of time (the height of a tree of `n` items), so all inserts take `n*log(n+1)` units of time. ("`log`" here means the logarithm to base 2.)

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

Can we simplify more? Let's take away the constant multiplier "`4`":

![4x^2 and x^2](parabolas.png)

The curve (cyan) looks less steep than the original one. However, when we move to larger scales, the difference shrinks substantially:

![4x^2 and x^2 scaled](parabolas_scaled.png)

After scaling both the x axis and the y axis by a factor of 100, the curves are much closer together, and what's more important, they share the same characteristics. Their shape is similar, and both curves dash of towards infinity.

So for estimating purposes, we can even leave out the constant factor. Our term has now shrunken to `n^2`.

Now we safely can say that our imaginary code needs *roughly* `n^2` units of time to process `n` items.

The simplified term `n^2` now stands for a whole range of terms, so let's give it a name.

## The Big-O notation





A Big-O function is a mathematical term that gives a *rough estimation* of how the speed of an algorithm changes with the size of the input data.

## Improving O(n)

### Change the algorithm


### Pre-process the input data

In some cases, the time complexity can be improved by re-arranging the input data.

For example, f you need to search through a list repeatedly, it will pay off to sort this list before searching it. Sure, sorting needs time, but you would have to do it only once to speed up those thousands of searches that occur afterwards.

A good sort algorithm takes O(n*log(n)) time but reduces all subsequent searches from O(n) to O(log(n)). Big win.

Or if you have a sorted list of data, inserting a new value takes O(n) time. If you insert new data very frequently, better turn this list into a balanced tree. Then inserting a new value only takes O(log(n+1)) time.


### Trade in time for space


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


**Happy coding!**

*/
