# gorilla mux

> last updated: 2022/10/18
> version: v1.8.0

- [gorilla mux](#gorilla-mux)
  - [Source](#source)
  - [intro](#intro)
  - [Example](#example)
  - [Matching Routes](#matching-routes)
    - [match path prefixes](#match-path-prefixes)
    - [subrouting](#subrouting)
      - [example](#example-1)

## Source

[gorilla/mux](https://pkg.go.dev/github.com/gorilla/mux#readme-middleware)
[what is http request multiplexer?](https://stackoverflow.com/questions/40478027/what-is-an-http-request-multiplexer)

## intro

Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler.

The name mux stands for "HTTP request multiplexer". Like the standard `http.ServeMux`, `mux.Router` matches incoming requests against a list of registered routes and calls a handler for the route that matches the URL or other conditions.

## Example

registering a couple of URL paths and handlers:

```go
func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
    r.HandleFunc("/products", ProductsHandler)
    r.HandleFunc("/articles", ArticlesHandler)
    http.Handle("/", r)
}
```

Here we register three routes mapping URL paths to handlers. This is equivalent to how `http.HandleFunc()` works: if an incoming request URL matches one of the paths, the corresponding handler is called passing `(http.ResponseWriter, *http.Request)` = `type Handler` as parameters.

`Paths can have variables`. They `are defined` using the format `{name}` or `{name:pattern}`. If a `regular expression pattern is not defined`, the matched variable will be `anything` until the next slash. For example:

```go
r := mux.NewRouter()
r.HandleFunc("/products/{key}", ProductHandler)
r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
```

The `names` are used to `create` a `map of route variables` which can `be retrieved calling mux.Vars()`:

```go
func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
```

And this is all you need to know about the basic usage. More advanced options are explained below.

## Matching Routes

`Routes` can also be restricted to a `domain` or `subdomain`. Just define a host pattern to be matched. They can also have variables:

```go
r := mux.NewRouter()
// Only matches if domain is "www.example.com".
r.Host("www.example.com")
// Matches a dynamic subdomain.
r.Host("{subdomain:[a-z]+}.example.com")
```

### match path prefixes

There are several other matchers that can be added. To `match path prefixes`:

```text
r.PathPrefix("/products/")
```

...or `HTTP methods`:

```go
r.Methods("GET", "POST")
```

...or `URL schemes`:

```go
r.Schemes("https")
```

...or `header values`:

```go
r.Headers("X-Requested-With", "XMLHttpRequest")
```

...or `query values`:

```go
r.Queries("key", "value")
```

...or to use a `custom matcher function`:

```go
r.MatcherFunc(func(r *http.Request, rm *RouteMatch) bool {
    return r.ProtoMajor == 0
})
```

...and finally, it is possible to `combine several matchers in a single route`:

```go
r.HandleFunc("/products", ProductsHandler).
  Host("www.example.com").
  Methods("GET").
  Schemes("http")
```

`Routes` are tested `in the order` they were added to the router. `If two routes match, the first one wins`:

```go
r := mux.NewRouter()
r.HandleFunc("/specific", specificHandler)
r.PathPrefix("/").Handler(catchAllHandler)
```

### subrouting

`Setting the same matching conditions` again and again can be `boring`, so we have a way to `group several routes that share the same requirements`. We call it `"subrouting"`.

#### example

For example, let's say we have several URLs that should only match when the host is www.example.com. Create a route for that host and get a "subrouter" from it:

```go
r := mux.NewRouter()
s := r.Host("www.example.com").Subrouter()
```

Then `register routes in the subrouter`:

```go
s.HandleFunc("/products/", ProductsHandler)
s.HandleFunc("/products/{key}", ProductHandler)
s.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
```

The `three URL paths` we registered above will only `be tested` if the `domain is www.example.com`, because the subrouter is tested first. This is not only convenient, but also optimizes request matching. You can create subrouters combining any attribute matchers accepted by a route.

`Subrouters` can be used to `create domain or path "namespaces"`: you define subrouters in a central place and then parts of the app can register its paths relatively to a given subrouter.

There's one more thing about subroutes. When a `subrouter has a path prefix`, the `inner routes use it as base` for their paths:

```go
r := mux.NewRouter()
s := r.PathPrefix("/products").Subrouter()
// "/products/"
s.HandleFunc("/", ProductsHandler)
// "/products/{key}/"
s.HandleFunc("/{key}/", ProductHandler)
// "/products/{key}/details"
s.HandleFunc("/{key}/details", ProductDetailsHandler)
```
