# dsa algo with go

> last updated: 2022/10/23

## Classification of data structures

If the situation requires various datatypes within a data structure, we can choose `heterogeneous` data structures. Linked, ordered, and unordered lists are grouped as `heterogeneous` data structures.

`Linear` data structures are lists, sets, tuples, queues, stacks and heaps.

Trees, tables, and containers are categorized as `nonlinear` data structures.

Two-dimensional and multidimensional arrays are grouped as `homogeneous` data structures.

`Dynamic data structures` are dictionaries, tree sets, and sequences.

![classification](./classification.drawio.svg)

## Lists

A list is a sequence of elements. Each element can be connected to another with a link in a `forward or backward direction`.

Data items within a list need `not be contiguous in memory` or on disk.

## Tuples

A tuple is a `finite sorted list of elements`. It is a data structure that groups data. Tuples are typically `immutable` sequential collections. The element has related fields of `different datatypes`. The `only way to modify` a tuple is to `change the fields`.

`Operators` such as `+ and *` can be applied to tuples. A database record is referred to as a tuple.

## Heaps

A heap is a data structure that is based on the heap property. The `heap` data structure `is used in selection, graph, and k-way merge algorithms`.

`Operations` such as `finding, merging, insertion, key changes, and deleting` are performed on heaps. Heaps are part of the `container/heap` package `in Go`.

According to the heap order (maximum heap) property, the value stored at each node is greater than or equal to its children.

If the `order` is `descending`, it is referred to as a `maximum heap`; `otherwise`, it's a `minimum heap`.

## Devide and conquer

A divide and conquer algorithm `breaks a complex problem into smaller problems` and solves these smaller problems. The smaller problem will be further broken down till it is a known problem. The approach is to recursively solve the sub-problems and `merge the solutions of the sub-problems`.

Recursion, quick sort, binary search, fast Fourier transform, and merge sort `are good examples of divide and conquer algorithms`

## Backtracking algorithms

A backtracking algorithm solves a problem by `constructing the solution incrementally`.

Multiple options are evaluated, and the algorithm chooses to go to the next component of the solution through recursion.

Backtracking can be a chronological type or can traverse the paths, depending on the problem that you are solving.

Backtracking is an algorithm that `finds candidate solutions and rejects a candidate on the basis of its feasibility and validity`.

It is useful in scenarios such as `finding a value in an unordered table`.

Constraint satisfaction problems such as parsing, rules engine, knapsack problems, and combinatorial optimization `are solved using backtracking`.

## Summary

This chapter covered the definition of `abstract datatypes`, classifying data structures into linear, nonlinear, homogeneous, heterogeneous, and dynamic types. Abstract datatypes such as container, list, set, map, graph, stack, and queue were presented in this chapter. The chapter covered the performance analysis of data structures and structural design patterns.
