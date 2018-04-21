# GraphQL - The Basics

## This holds the GraphQL Implementation of the goodreads Api

### About

- A Query Language For APIs, a possible replacement for REST.
- There requires a schema that has to be predefined for the data.
- A received query on a GraphQL Server is first checked to ensure it only refers to the types and fields defined (this makes it way more predictable) , then runs the provided functions to produce a result.

### Examples

The following query:

```js
{
  person {
    name
  }
}
```

Would produce the following json

```js

{
  "person" {
    "name":"Rob"
  }
}
```

For more information refer to [Official GraphQL Docs](http://graphql.org/learn/)