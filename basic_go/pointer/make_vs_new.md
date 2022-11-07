# make and new

## Source

[Why would I make() or new()?](https://stackoverflow.com/questions/9320862/why-would-i-make-or-new)

- `new(T)` - Allocates memory, and sets it to the **zero value** for type **`T..`** `..`that is 0 for int, **""** for string and **nil** for referenced types (slice, map, chan)
- `make(T)` - Allocates memory for referenced data types (slice, map, chan), plus **initializes** their **underlying data structures**
