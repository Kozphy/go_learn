# sqlx

## Source

[section_readme](https://pkg.go.dev/github.com/jmoiron/sqlx#section-readme)
[github_sqlx](http://jmoiron.github.io/sqlx/)

## intro

`sqlx` is a library which provides a set of extensions on go's standard database/sql library. The sqlx versions of `sql.DB, sql.TX, sql.Stmt`, et al. all **leave the underlying interfaces untouched**.

Major additional concepts are

- Marshal rows into structs (with embedded struct support), maps, and slices
- Named parameter support including prepared statements
- `Get` and `Select` to go quickly from query to struct/slice

`DB.Connx` returns an `sqlx.Conn`, which is an `sql.Conn`-alike consistent with sqlx's wrapping of other types.

`BindDriver` allows users to control the bindvars that sqlx will use for drivers, and add new drivers at runtime.

## Connecting to your Database

A **DB** instance is not a connection, but an abstraction representing a Database.

It maintains a `connection pool` internally, and will attempt to connect when a connection is first needed.

## The Connection Pool

`Statement preparation` and `query execution` require a connection, and the DB object will manage a pool of them so that **it can be safely used for concurrent querying**.

There are two ways to control the size of the connection pool as of Go 1.2:

- **DB.SetMaxIdleConns(n int)**
- **DB.SetMaxOpenConns(n int)**

**By default, the pool grows unbounded**, and connections will be created whenever there isn't a free connection available in the pool.

You can use **DB.SetMaxOpenConns** to set the **maximum size of the pool**.

Connections that are **not being used are marked idle** and then closed if they aren't required.

To avoid making and closing lots of connections, **set the maximum idle size with DB.SetMaxIdleConns** to a size that is sensible for your query loads.

It is easy to **get into trouble** by accidentally holding on to connections. **To prevent** this:

- Ensure you **Scan()** every Row object
- Ensure you either **Close()** or fully-iterate via **Next()** every Rows object
- Ensure every transaction returns its connection via **Commit()** or **Rollback()**

**If you neglect to do one of these things, the connections they use may be held until garbage collection**, and your db will end up creating far more connections at once in order to compensate for the ones its using.

Note that **Rows.Close()** `can be called multiple times safely`, so do not fear calling it where it might not be necessary.
