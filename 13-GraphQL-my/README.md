# Chapter 13: GraphQL with Go using gqlgen

Welcome to Chapter 13, where we explore building GraphQL APIs in Go using `gqlgen`. `gqlgen` is a Go library for building GraphQL servers without the hassle, offering a great developer experience by generating codes that are powerful and easy to extend.

## Table of Contents
1. [Introduction to GraphQL](#introduction-to-graphql)
2. [Why gqlgen?](#why-gqlgen)
3. [Setup and Installation](#setup-and-installation)
   - [Prerequisites](#prerequisites)
   - [Installing gqlgen](#installing-gqlgen)
4. [Initializing a gqlgen Project](#initializing-a-gqlgen-project)
   - [Steps](#steps)
   - [Example Schema](#example-schema)
      - [`schema.graphqls` File](#schemagraphqls-file)
5. [Code Generation and Resolver Implementation](#code-generation-and-resolver-implementation)
   - [Generating Code](#generating-code)
   - [Implementing Resolvers](#implementing-resolvers)
6. [Testing the GraphQL API](#testing-the-graphql-api)
   - [Strategies](#strategies)
   - [Testing Tools](#testing-tools)
7. [Common Pitfalls and Solutions](#common-pitfalls-and-solutions)
   - [Issues](#issues)
8. [Best Practices](#best-practices)
   - [Tips](#tips)
9. [Additional Resources](#additional-resources)


## Introduction to GraphQL
Brief about what GraphQL is, its advantages, and why it is used.

## Why gqlgen?
An overview of why `gqlgen` is chosen as the GraphQL server library, its advantages, and a little on how it works.

## Setup and Installation
### Prerequisites
- Go installed (version x.x.x+)
- A familiar IDE or code editor
- Basic knowledge of Go programming

### Installing gqlgen
```shell
go get github.com/99designs/gqlgen
```

## Initializing a gqlgen Project

Starting a project with `gqlgen` involves generating initial configurations and stub codes that adhere to a GraphQL schema. Begin by initializing a new `gqlgen` project to create essential files that will structure your GraphQL API server.

### Steps
1. **Generate Initial Files:** Use the `gqlgen` CLI to create foundational files.
    ```shell
    go run github.com/99designs/gqlgen init
    ```
2. **Explore the Structure:** Understand the purpose of each generated file and directory.
    - `gqlgen.yml`: Configuration file for code generation.
    - `graph/`: Contains GraphQL schema and resolver implementation.

3. **Customization:** Adjust configurations and settings according to your project needs.

### Example Schema

In GraphQL, a schema defines the capabilities of the API by specifying the types of data that can be fetched and how they are related. For `gqlgen`, the schema is typically defined in a `schema.graphqls` file, which resides at the root of your GraphQL server directory.

#### `schema.graphqls` File

This file is fundamental in GraphQL API development as it dictates the operations (queries, mutations, and subscriptions) and shapes of the data your API will work with. Let's consider a simple example:

```graphql
type Category {
   id: ID!
   name: String!
   description: String
   courses: [Course!]!
}

type Course {
   id: ID!
   name: String!
   description: String
   category: Category!
   lessons: [Lesson!]!
}

type Lesson {
   id: ID!
   name: String!
   description: String
   course: Course!
   content: String!
}

input NewCategory {
   name: String!
   description: String
}

input NewCourse {
   name: String!
   description: String
   categoryId: ID!
}

input NewLesson {
   name: String!
   description: String
   courseId: ID!
   content: String!
}

type Query {
   categories: [Category!]!
   courses: [Course!]!
   lessons: [Lesson!]!
}

type Mutation {
   createCategory(input: NewCategory!): Category!
   createCourse(input: NewCourse!): Course!
   createLesson(input: NewLesson!): Lesson!
}
```

In the above snippet from a `schema.graphqls` file:
- A `Query` type defines a `user` query which expects an `ID` as input and returns a `User` type.
- The `User` type represents a user entity, defining relevant fields such as `id`, `name`, and `email`.

When you run `gqlgen`, it reads the schema from `schema.graphqls` and generates Go code for the defined types and operations. This schema acts as a contract between your API and the clients, ensuring that the data types and operations adhere to defined structures and rules.

In the development workflow:
1. Define or update the `schema.graphqls` with your data types and operations.
> **Note:** The automatically generated TODO has been replaced by the schema above.
2. Run the `gqlgen generate` command to produce Go code.
3. Implement resolvers for your types and operations to fetch and manipulate actual data.
> **Note:** Remember to check the `schema.resolvers.go` file for methods that are no longer used. Remove them to avoid confusion.

Remember, changes in your API's capabilities or data shapes should always be reflected and updated in the `schema.graphqls` file, followed by re-generating the code to keep things synchronized.

Code Generation and Resolver Implementation
-------------------------------------------

With a defined schema, `gqlgen` can generate Go code to facilitate the implementation of the GraphQL API.

### Generating Code

Use the `gqlgen` CLI to create Go code based on your GraphQL schema.

shellCopy code

`go run github.com/99designs/gqlgen generate`

### Implementing Resolvers

1.  Query and Mutation Resolvers: Implement logic for fetching and manipulating data as per your schema.
2.  Additional Resolvers: Provide resolver functions for any computed fields or relationships.

Testing the GraphQL API
-----------------------

Testing is paramount to verify that your API behaves as expected and to safeguard against regressions.

### Strategies

-   Unit Testing: Validate individual resolvers and utility functions.
-   Integration Testing: Test the API as a whole ensuring all parts work harmoniously.

### Testing Tools

-   Go's Testing Library: Use native testing capabilities.
-   Test Databases: Utilize a test DB to validate data-related operations without affecting production data.

Common Pitfalls and Solutions
-----------------------------

Here we discuss some typical challenges encountered during GraphQL API development and how to address them.

### Issues

-   N+1 Query Problem: Occurs when the API fetches data in a suboptimal manner.
-   Error Handling: Strategizing error formation and propagation.
-   Authorization: Securing resolvers and data.

Best Practices
--------------

Explore practices to improve your GraphQL API development with `gqlgen`.

### Tips

-   Optimize Resolvers: Ensure they fetch only necessary data.
-   Use DataLoader: To batch and cache requests.
-   Effective Logging: Implement logging for troubleshooting and auditing purposes.

Additional Resources
--------------------

-   [gqlgen Documentation](https://gqlgen.com/)
-   [GraphQL Official Website](https://graphql.org/)
-   [Apollo Federation](https://www.apollographql.com/docs/federation/)


