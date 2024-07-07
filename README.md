
# Blog API Project in Golang

## Overview
This project is created to help me understand how to write functions, methods, and interfaces in Golang. It also aims to explore the development of hexagonal architecture in Golang.

## Data Structure
``` golang
type Posts struct {
    Title   string
    Content string
}
```

## Features
-  Get All: Retrieve all blog posts.
- Get By ID: Retrieve a blog post by its ID.
- Post: Create a new blog post.
- PUT: Update an existing blog post.
- Delete: Remove a blog post.

# Project Setup

## Installation

1. Clone the repository:
```bash
git clone https://github.com/PoowadolDev/blog-website-golang.git
```

2. Navigate to the project directory:
```bash
cd blog-website-golang/api-service
```

3. Install the dependencies:
```bash
go mod tidy
```

## Running the Project
To run the project, use the following command:
```bash
go run *
```
## Endpoints

- ### Get All Posts

```
  GET /GetAllPosts
```

- ### Get Post By ID:

```
  GET /GetPostById/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of blog post to fetch |

- ### Create New Post:
```
  POST /CreatePost
```

#### Request Body
```json
{
    "Title": "Post Title",
    "Content": "Post Content"
}
```

- ### Update Existing Post:

```
  PUT /UpdatePost/{id}
```
#### Request Body
```
{
    "Title": "Updated Title",
    "Content": "Updated Content"
}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of blog post to update |

- ### Delete Post:
```
DELETE /DeletePost/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of blog post to delete |

# Architecture

The project follows the hexagonal architecture pattern to separate concerns and make the code more maintainable and testable. The structure is as follows:

```bash
api-service
├───adapter
│   ├───database
│   └───handler
├───entity
├───ports
└───repository
```

### Description of Folders

- **adapter**: Contains the adapters for external systems.
    - **database**: Manages database interactions.
    - **handler**: Handles HTTP requests and responses.
- **entity**: Defines the core domain entities.
- **ports**: Contains the interfaces for driving and driven adapters.
- **repository**: Implements the data access layer.

### Key Concepts
- **Domain Layer (entity)**: Contains the core logic and domain entities.
- **Application Layer (ports)**: Manages use cases and application flow through interfaces.
- **Infrastructure Layer (adapter, repository)**: Handles communication with external systems (e.g., database, web server).
