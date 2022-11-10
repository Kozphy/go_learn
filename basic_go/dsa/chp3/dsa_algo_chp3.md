# Linear Data Structures

## We will cover the following linear data structures in this chapter

- Lists
- Sets
- Tuples
- Stacks

## List

A list is a `collection of ordered elements` that are used to store list of items. Unlike array lists, these `can expand and shrink dynamically`.

## LinkList

`LinkedList` is **a sequence of nodes that have properties and a reference to the next node** in the sequence. It is a linear data structure that is used to store data.

They are `not stored contiguously in memory`, which makes them different arrays.

## Doubly linked list

[double_linked_list](https://pkg.go.dev/container/list)

## Sets

A `Set` is a linear data structure that has a collection of values that are **not repeated**. A set can store unique values **without any particular order**.

In the real world, sets can be used to collect all tags for blog posts and conversation participants in a chat.

## Tuples

Tuples are **finite ordered sequences of objects**. They **can contain a mixture of other data types and are used to group related data into a data structure**.

In a relational database, a tuple is a row of a table. Tuples have a fixed size compared to lists, and are also faster. A finite set of tuples in the relational database is referred to as a relation instance.

A **tuple** can be assigned in a single statement, which is **useful for swapping values**.

**Lists** usually contain values of the same data type, while **tuples contain different data**.

## Queues

A queue **consists of elements to be processed in a particular order or based on priority**.

A **priority-based queue of orders** is shown in the following code, structured as a **heap**.

Operations such as enqueue, dequeue, and peek can be performed on queue. A queue is a **linear data structure and a sequential collection**.

Queues are commonly used for **storing tasks that need to be done**, or **incoming HTTP requests that need to be processed by a server**.

In real life, **handling interruptions in real-time systems, call handling, and CPU task scheduling** are good examples for using queues.

### Add method

The `Add` method on the `Queue` class takes the `order` parameter and adds it to `Queue` **based on the priority**.

Based on this, the location of the `order` parameter **is found by comparing it with the priority parameter**:
