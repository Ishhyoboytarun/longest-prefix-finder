
# Longest Prefix Finder using Trie and Goroutines

Finding the longest prefix string for a given input string from a set of prefixes is a common task in text processing. This project demonstrates how to efficiently accomplish this using a Trie data structure and goroutines in Go.

## Table of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
- [Usage](#usage)
    - [File Format](#file-format)
    - [Optimal Goroutines](#optimal-goroutines)
- [Example](#example)
- [Performance Optimization](#performance-optimization)
- [Contributing](#contributing)

## Introduction

The Longest Prefix Finder is designed to efficiently find the longest prefix from a set of sample prefixes for a given input string. It uses a Trie data structure to organize and search for prefixes and employs goroutines to optimize performance.

## Getting Started

### Prerequisites

To run this project, you need to have [Go](https://golang.org/) installed on your system.

### Installation

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/Ishhyoboytarun/longest-prefix-finder.git
   ```

2. Navigate to the project directory:

   ```bash
   cd longest-prefix-finder
   ```

3. Build the executable:

   ```bash
   go build
   ```
4. Run the exe: 

    ```bash
    ./longest-prefix-finder
    ```

## Usage

To use this Longest Prefix Finder, follow these steps:

1. Create a file named `sample_prefixes.txt` that contains a list of sample prefixes, one per line.

2. Run the `longest-prefix-finder` executable with your input string:

### File Format

The `sample_prefixes.txt` file should contain one prefix per line. This file serves as the dictionary of prefixes for the input string.

### Optimal Goroutines

You can specify the optimal number of goroutines to be used for processing the prefixes. The program will ensure that only the specified number of goroutines is used for concurrent processing. This helps balance concurrency and resource utilization.

## Example

Suppose you have the following `sample_prefixes.txt` file:

```
pre
prefix
prefixes
long
longest
example
```

Running the program with an input string:

```
"longestexample"
```

The output will be:

```
"longest"
```

## Performance Optimization

The program uses goroutines to process prefixes concurrently. You can specify the number of goroutines to use for processing with the `-goroutines` flag. The program will ensure that only the specified number of goroutines is used to strike a balance between concurrency and resource utilization.

## Contributing

Contributions are welcome! If you have ideas for improvements, please open an issue or submit a pull request.

---
