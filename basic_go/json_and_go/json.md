# json and go

[source](https://go.dev/blog/json)

## Encoding (Marshal)

1. Only data structures that can be represented as valid JSON will be encoded:
`JSON` objects only `support strings as keys`; to encode a Go map type it must be of the form `map[string]T` (where T is any Go type supported by the json package).

2. Channel, complex, and function types `cannot` be encoded.

3. Cyclic data structures are `not supported`; they will cause Marshal to go into an infinite loop.

4. Pointers `will be encoded` as the values they point to (or ‘null’ if the pointer is nil).

## Decoding (Unmarshal)

1. How does Unmarshal `identify the fields` in which to store the decoded data?
For a given JSON key "Foo", Unmarshal will look through the destination struct’s fields to find (in order of preference):
    - An exported field with a tag of "Foo".
    - An exported field named "Foo", or
    - An exported field named "FOO" or "FoO" or some other case-insensitive match of "Foo".
