module github.com/zixas/go_learn

go 1.19

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.5.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/pkg/errors v0.9.1
	golang.org/x/tour v0.1.0
)

require github.com/lib/pq v1.10.7 // indirect

replace example.com/hello => ../hello
