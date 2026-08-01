[![Go Reference](https://pkg.go.dev/badge/github.com/marcelo-sjr/go_datastructures.svg)](https://pkg.go.dev/github.com/marcelo-sjr/go_datastructures)
[![Go Version](https://img.shields.io/github/go-mod/go-version/marcelo-sjr/go_datastructures)](go.mod)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

# Data Structures in Go

A simple DS implementation written in Go for **study purposes**.   

This repository is part of my journey learning data structures and algorithms using Go. The goal is to understand how ds work internally by implementing them from scratch, without relying on Go's standard library implementations.

> **Note:** This project is intended for educational purposes and is **not** designed to be a production-ready library.

## Features

Currently implemented:

- ✅ Singly Linked List
- ✅ Insert elements
- ✅ Remove elements
- ✅ Search elements
- ✅ Traverse the list
- ✅ Support for Go Generics

## Roadmap

Planned improvements include:


- [ ] Doubly Linked List
- [ ] Circular Linked List
- [ ] Iterators
- ✅ More utility methods
- [ ] Benchmarks
- [ ] Unit tests
- [ ] Documentation with complexity analysis
- [ ] Stack
- [ ] Queue
- [ ] Tree
- [ ] Iterators

## Project Structure

```text
.
├── linkedlist.go
├── node.go
├── go.mod
├── LICENSE
└── README.md
```

## Complexity

| Operation | Time |
|----------|------|
| Search | O(n)   | 
| Insert (head)   | O(1) | 
| Insert (tail)   | O(1) | 
| Insert (after)  | O(n) | 
| Delete | O(n)   | O(1)*| *head deletion
| Find   | O(n)   |
| Values | O(n)   |   



## Future

The current implementation supports **Go Generics**, making the linked list reusable with any type. Soon new methods will be created to allow search and deletion by value.

## License

This project is licensed under the MIT License.
