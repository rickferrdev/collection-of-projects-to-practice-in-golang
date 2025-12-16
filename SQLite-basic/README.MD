# SQLite Basic

A Go project demonstrating basic SQLite database operations with a clean, modular structure.

## Project Structure

```
SQLite-basic/
├── main.go              # Entry point
├── go.mod              # Go module definition
├── connection/         # Database connection module
│   └── connection.go
└── pets/               # Pets data model and operations
    └── pets.go
```

## Overview

This project provides a foundational example of:
- Establishing SQLite database connections
- Managing pet records
- Basic CRUD operations with SQL

## Getting Started

### Prerequisites
- Go 1.16+
- SQLite3

### Installation

```bash
git clone <repository-url>
cd SQLite-basic
go mod download
```

### Running

```bash
go run main.go
```

## Module Description

- **connection/**: Handles database connection setup and initialization
- **pets/**: Defines pet data structure and database operations

## Technologies

- Go
- SQLite3
