# Karni


## Overview

**Karni** is a Go library designed to provide an Object-Document Mapping (ODM) platform for MongoDB. It simplifies the interaction with MongoDB by allowing developers to work with Go structs instead of raw BSON documents. This library aims to streamline database operations, making it easier to manage and manipulate data within a Go application.

## Features

- **Seamless MongoDB Connection**: Easily connect to MongoDB using a URL.
- **Intuitive API**: Work with Go structs to perform CRUD operations.
- **Error Handling**: Comprehensive error management to ensure robust applications.
- **Scalable Architecture**: Designed to handle large datasets efficiently.

## Getting Started

### Prerequisites

- Go 1.16 or later
- MongoDB 4.0 or later

## Installation

To install Karni, use `go get`:

```bash
go get github.com/swagisays/karni/karni

```
### Usage

Here's a quick example to get you started:

```go
package main

import (
    "fmt"
    "log"
    "github.com/swagisays/karni/karni"
)

func main() {
    err := karni.Connect("mongodb://localhost:27017", "yourDatabaseName")

    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }
    
    fmt.Println("Connected to MongoDB successfully!")
    // Add your ODM operations here
}
```

## Documentation

Under development



## License

This project is licensed under the MIT License - see the `LICENSE` file for details.

## Contact

For any inquiries, please contact [swagisays@icloud.com](mailto:swagisays@icloud.com).

---

*Happy Coding with Karni!* 🎉