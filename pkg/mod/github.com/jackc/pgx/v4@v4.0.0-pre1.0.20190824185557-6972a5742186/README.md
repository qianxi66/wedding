[![](https://godoc.org/github.com/jackc/pgx?status.svg)](https://godoc.org/github.com/jackc/pgx)
[![Build Status](https://travis-ci.org/jackc/pgx.svg)](https://travis-ci.org/jackc/pgx)

# pgx - PostgreSQL Driver and Toolkit

pgx is a pure Go driver and toolkit for PostgreSQL. It is usable through database/sql but also offers a native interface similar to database/sql that offers better performance and more features.

## Example Usage

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string
	var weight int64
	err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, weight)
}
```

## Prerelease v4

This is the `v4` branch. `v4` is still under development but is ready for testing. See [v4 milestone](https://github.com/jackc/pgx/milestone/2) for currently known outstanding issues.

## Features

pgx supports many additional features beyond what is available through database/sql.

* Support for approximately 60 different PostgreSQL types
* Batch queries
* Single-round trip query mode
* Full TLS connection control
* Binary format support for custom types (can be much faster)
* Copy protocol support for faster bulk data loads
* Extendable logging support including built-in support for log15 and logrus
* Connection pool with after connect hook to do arbitrary connection setup
* Listen / notify
* PostgreSQL array to Go slice mapping for integers, floats, and strings
* Hstore support
* JSON and JSONB support
* Maps inet and cidr PostgreSQL types to net.IPNet and net.IP
* Large object support
* NULL mapping to Null* struct or pointer to pointer.
* Supports database/sql.Scanner and database/sql/driver.Valuer interfaces for custom types
* Notice response handling (this is different than listen / notify)

## Related Libraries

pgx is part of a family of PostgreSQL libraries. Many of these can be used independently. Many can also be accessed from pgx for lower-level control.

## github.com/jackc/pgconn

pgconn is a lower-level PostgreSQL database driver that operates at nearly the same level is the C library libpq.

## github.com/jackc/pgx/v4/pgxpool

pgxpool is a connection pool for pgx. pgx is entirely decoupled from its default pool implementation. This means pgx can be used with a different pool without any pool at all.

## github.com/jackc/pgx/v4/stdlib

database/sql compatibility layer for pgx. pgx can be used as a normal database/sql driver, but at any time the native interface may be acquired for more performance or PostgreSQL specific functionality.

## github.com/jackc/pgx/pgtype

Approximately 60 PostgreSQL types are supported including uuid, hstore, json, bytea, numeric, interval, inet, and arrays. These types support database/sql interfaces and are usable even outside of pgx. They are fully tested in pgx and pq. They also support a higher performance interface when used with the pgx driver.

## github.com/jackc/pgproto3

pgproto3 provides standalone encoding and decoding of the PostgreSQL v3 wire protocol. This is useful for implementing very low level PostgreSQL tooling.

## github.com/jackc/pglogrepl

pglogrepl provides function to act as a client for PostgreSQL logical replication.

## github.com/jackc/pgx/v4/pgmock

pgmock offers the ability to create a server that mocks the PostgreSQL wire protocol. This is used internally to test pgx by purposely inducing unusual errors. pgproto3 and pgmock together provide most of the foundational tooling required to implement a PostgreSQL proxy or MitM (such as for a custom connection pooler).

## github.com/jackc/tern

tern is a stand-alone SQL migration system.

## Alternatives

* [pq](http://godoc.org/github.com/lib/pq)
* [go-pg](https://github.com/go-pg/pg)

For normal queries with small result sets all drivers perform similarly, but pgx can have a significant advantage with large result sets or when lower level features are used.  But for most application use cases the performance difference will be irrelevant. See [go_db_bench](https://github.com/jackc/go_db_bench) to run tests for yourself.

The primary difference between the drivers is features and API style.

pq is exclusively used with database/sql. go-pg does not use database/sql at all. pgx supports database/sql as well as a faster and more featureful native interface.

go-pg has an ORM and schema migration support baked in. pq and pgx do not.

When possible, pgx decouples functionality into separate packages to make it easy to reuse outside of pgx and replace functionality inside. For example, the [pgtype](https://github.com/jackc/pgtype) package that provides support for PostgreSQL types is usable with pq and database/sql and the bundled connection pool is entirely replaceable.

## Documentation

pgx includes extensive documentation in the godoc format. It is viewable online at [godoc.org](https://godoc.org/github.com/jackc/pgx).

## Testing

pgx tests need a PostgreSQL database. It will connect to the database specified in the `PGX_TEST_DATABASE` environment
variable. The `PGX_TEST_DATABASE` environment variable can be a URL or DSN. In addition, the standard `PG*` environment
variables will be respected. Consider using [direnv](https://github.com/direnv/direnv) to simplify environment variable
handling.

### Example Test Environment

Connect to your PostgreSQL server and run:

```
create database pgx_test;
```

Connect to the newly created database and run:

```
create domain uint64 as numeric(20,0);
```

Now you can run the tests:

```
PGX_TEST_DATABASE="host=/var/run/postgresql database=pgx_test" go test ./...
```

## Version Policy

pgx follows semantic versioning for the documented public API on stable releases. This is the prerelease of `v4`. Branch `v3` is the latest stable release. `master` can contain new features or behavior that will change or be removed before being merged to the stable `v3` branch (in practice, this occurs very rarely).
