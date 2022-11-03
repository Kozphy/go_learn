# graph

> last updated: 2022/11/03

## What is graph

**A graph models a set of connections.**

Each graph is made up of `nodes` and `edges`.
![graph_node_edge](./asset/grokking/ch6/graph_node_edge.drawio.svg)

A node can be directly connected to many other nodes. Those nodes are called its `neighbors`.

### direct graph vs undirect graph

An **directed graph have arrows** pointing to them, but not arrows from them to someone else which relationship is only one way.

An **undirected graph doesn’t have any arrows**, and both nodes are each other’s neighbors.
![direct_vs_undirect](./asset/grokking/ch6/direct_undirect.drawio.svg)

## bfs (breadth-first search)

The algorithm to solve a **shortest-path** problem is called **breadth-first** search.

It can help answer **two types of questions**:

1. Question type 1: Is there a path from node A to node B?
2. Question type 2: What is the shortest path from node A to node B?

![bfs-bridge](./asset/grokking/ch6/bfs_bridge.drawio.svg)

You can use `queue` to complete bfs.
