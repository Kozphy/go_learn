# Arrays, slices (and strings): The mechanics of 'append'

> last updated: 2022/10/20

## [Source](https://go.dev/blog/slices)

## intro

One of the most `common features of procedural programming languages` is the concept of an array. Arrays seem like simple things but there are many questions that must be answered when adding them to a language, such as:

- fixed-size or variable-size?
- is the size part of the type?
- what do multidimensional arrays look like?
- does the empty array have meaning?

The answers to these questions affect whether arrays are just a feature of the language or a core part of its design.

In the early development of Go, it took about a year to decide the answers to these questions before the design felt right. The key step was the introduction of slices, which built on fixed-size arrays to give a flexible, extensible data structure. To this day, however, programmers new to Go often stumble over the way slices work, perhaps because experience from other languages has colored their thinking.

`In this post` we’ll attempt to clear up the confusion. We’ll do so by building up the pieces `to explain how the append built-in function works, and why it works the way it does`.
