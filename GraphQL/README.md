### GraphQL - The Basics

- A Query Language For APIs, a possible replacement for REST.
- There requires a schema that has to be predefined for the data.
- A received query on a GraphQL Server is first checked to ensure it only refers to the types and fields defined (this makes it way more predictable) , then runs the provided functions to produce a result.

### For example:
The following query:
```
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
    "name":"Shriram"
  }
}
```
For more information refer to http://graphql.org/learn/