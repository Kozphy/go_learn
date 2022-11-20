# sqlx

## [Source](https://pkg.go.dev/github.com/jmoiron/sqlx#section-readme)

## intro

`sqlx` is a library which provides a set of extensions on go's standard database/sql library. The sqlx versions of `sql.DB, sql.TX, sql.Stmt`, et al. all **leave the underlying interfaces untouched**.

Major additional concepts are

- Marshal rows into structs (with embedded struct support), maps, and slices
- Named parameter support including prepared statements
- `Get` and `Select` to go quickly from query to struct/slice

`DB.Connx` returns an `sqlx.Conn`, which is an `sql.Conn`-alike consistent with sqlx's wrapping of other types.

`BindDriver` allows users to control the bindvars that sqlx will use for drivers, and add new drivers at runtime.
