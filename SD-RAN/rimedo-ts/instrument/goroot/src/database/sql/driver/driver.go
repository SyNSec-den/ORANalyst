// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/database/sql/driver/driver.go:5
// Package driver defines interfaces to be implemented by database
//line /usr/local/go/src/database/sql/driver/driver.go:5
// drivers as used by package sql.
//line /usr/local/go/src/database/sql/driver/driver.go:5
//
//line /usr/local/go/src/database/sql/driver/driver.go:5
// Most code should use package sql.
//line /usr/local/go/src/database/sql/driver/driver.go:5
//
//line /usr/local/go/src/database/sql/driver/driver.go:5
// The driver interface has evolved over time. Drivers should implement
//line /usr/local/go/src/database/sql/driver/driver.go:5
// Connector and DriverContext interfaces.
//line /usr/local/go/src/database/sql/driver/driver.go:5
// The Connector.Connect and Driver.Open methods should never return ErrBadConn.
//line /usr/local/go/src/database/sql/driver/driver.go:5
// ErrBadConn should only be returned from Validator, SessionResetter, or
//line /usr/local/go/src/database/sql/driver/driver.go:5
// a query method if the connection is already in an invalid (e.g. closed) state.
//line /usr/local/go/src/database/sql/driver/driver.go:5
//
//line /usr/local/go/src/database/sql/driver/driver.go:5
// All Conn implementations should implement the following interfaces:
//line /usr/local/go/src/database/sql/driver/driver.go:5
// Pinger, SessionResetter, and Validator.
//line /usr/local/go/src/database/sql/driver/driver.go:5
//
//line /usr/local/go/src/database/sql/driver/driver.go:5
// If named parameters or context are supported, the driver's Conn should implement:
//line /usr/local/go/src/database/sql/driver/driver.go:5
// ExecerContext, QueryerContext, ConnPrepareContext, and ConnBeginTx.
//line /usr/local/go/src/database/sql/driver/driver.go:5
//
//line /usr/local/go/src/database/sql/driver/driver.go:5
// To support custom data types, implement NamedValueChecker. NamedValueChecker
//line /usr/local/go/src/database/sql/driver/driver.go:5
// also allows queries to accept per-query options as a parameter by returning
//line /usr/local/go/src/database/sql/driver/driver.go:5
// ErrRemoveArgument from CheckNamedValue.
//line /usr/local/go/src/database/sql/driver/driver.go:5
//
//line /usr/local/go/src/database/sql/driver/driver.go:5
// If multiple result sets are supported, Rows should implement RowsNextResultSet.
//line /usr/local/go/src/database/sql/driver/driver.go:5
// If the driver knows how to describe the types present in the returned result
//line /usr/local/go/src/database/sql/driver/driver.go:5
// it should implement the following interfaces: RowsColumnTypeScanType,
//line /usr/local/go/src/database/sql/driver/driver.go:5
// RowsColumnTypeDatabaseTypeName, RowsColumnTypeLength, RowsColumnTypeNullable,
//line /usr/local/go/src/database/sql/driver/driver.go:5
// and RowsColumnTypePrecisionScale. A given row value may also return a Rows
//line /usr/local/go/src/database/sql/driver/driver.go:5
// type, which may represent a database cursor value.
//line /usr/local/go/src/database/sql/driver/driver.go:5
//
//line /usr/local/go/src/database/sql/driver/driver.go:5
// Before a connection is returned to the connection pool after use, IsValid is
//line /usr/local/go/src/database/sql/driver/driver.go:5
// called if implemented. Before a connection is reused for another query,
//line /usr/local/go/src/database/sql/driver/driver.go:5
// ResetSession is called if implemented. If a connection is never returned to the
//line /usr/local/go/src/database/sql/driver/driver.go:5
// connection pool but immediately reused, then ResetSession is called prior to
//line /usr/local/go/src/database/sql/driver/driver.go:5
// reuse but IsValid is not called.
//line /usr/local/go/src/database/sql/driver/driver.go:38
package driver

//line /usr/local/go/src/database/sql/driver/driver.go:38
import (
//line /usr/local/go/src/database/sql/driver/driver.go:38
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/database/sql/driver/driver.go:38
)
//line /usr/local/go/src/database/sql/driver/driver.go:38
import (
//line /usr/local/go/src/database/sql/driver/driver.go:38
	_atomic_ "sync/atomic"
//line /usr/local/go/src/database/sql/driver/driver.go:38
)

import (
	"context"
	"errors"
	"reflect"
)

// Value is a value that drivers must be able to handle.
//line /usr/local/go/src/database/sql/driver/driver.go:46
// It is either nil, a type handled by a database driver's NamedValueChecker
//line /usr/local/go/src/database/sql/driver/driver.go:46
// interface, or an instance of one of these types:
//line /usr/local/go/src/database/sql/driver/driver.go:46
//
//line /usr/local/go/src/database/sql/driver/driver.go:46
//	int64
//line /usr/local/go/src/database/sql/driver/driver.go:46
//	float64
//line /usr/local/go/src/database/sql/driver/driver.go:46
//	bool
//line /usr/local/go/src/database/sql/driver/driver.go:46
//	[]byte
//line /usr/local/go/src/database/sql/driver/driver.go:46
//	string
//line /usr/local/go/src/database/sql/driver/driver.go:46
//	time.Time
//line /usr/local/go/src/database/sql/driver/driver.go:46
//
//line /usr/local/go/src/database/sql/driver/driver.go:46
// If the driver supports cursors, a returned Value may also implement the Rows interface
//line /usr/local/go/src/database/sql/driver/driver.go:46
// in this package. This is used, for example, when a user selects a cursor
//line /usr/local/go/src/database/sql/driver/driver.go:46
// such as "select cursor(select * from my_table) from dual". If the Rows
//line /usr/local/go/src/database/sql/driver/driver.go:46
// from the select is closed, the cursor Rows will also be closed.
//line /usr/local/go/src/database/sql/driver/driver.go:61
type Value any

// NamedValue holds both the value name and value.
type NamedValue struct {
	// If the Name is not empty it should be used for the parameter identifier and
	// not the ordinal position.
	//
	// Name will not have a symbol prefix.
	Name	string

	// Ordinal position of the parameter starting from one and is always set.
	Ordinal	int

	// Value is the parameter value.
	Value	Value
}

// Driver is the interface that must be implemented by a database
//line /usr/local/go/src/database/sql/driver/driver.go:78
// driver.
//line /usr/local/go/src/database/sql/driver/driver.go:78
//
//line /usr/local/go/src/database/sql/driver/driver.go:78
// Database drivers may implement DriverContext for access
//line /usr/local/go/src/database/sql/driver/driver.go:78
// to contexts and to parse the name only once for a pool of connections,
//line /usr/local/go/src/database/sql/driver/driver.go:78
// instead of once per connection.
//line /usr/local/go/src/database/sql/driver/driver.go:84
type Driver interface {
	// Open returns a new connection to the database.
	// The name is a string in a driver-specific format.
	//
	// Open may return a cached connection (one previously
	// closed), but doing so is unnecessary; the sql package
	// maintains a pool of idle connections for efficient re-use.
	//
	// The returned connection is only used by one goroutine at a
	// time.
	Open(name string) (Conn, error)
}

// If a Driver implements DriverContext, then sql.DB will call
//line /usr/local/go/src/database/sql/driver/driver.go:97
// OpenConnector to obtain a Connector and then invoke
//line /usr/local/go/src/database/sql/driver/driver.go:97
// that Connector's Connect method to obtain each needed connection,
//line /usr/local/go/src/database/sql/driver/driver.go:97
// instead of invoking the Driver's Open method for each connection.
//line /usr/local/go/src/database/sql/driver/driver.go:97
// The two-step sequence allows drivers to parse the name just once
//line /usr/local/go/src/database/sql/driver/driver.go:97
// and also provides access to per-Conn contexts.
//line /usr/local/go/src/database/sql/driver/driver.go:103
type DriverContext interface {
	// OpenConnector must parse the name in the same format that Driver.Open
	// parses the name parameter.
	OpenConnector(name string) (Connector, error)
}

// A Connector represents a driver in a fixed configuration
//line /usr/local/go/src/database/sql/driver/driver.go:109
// and can create any number of equivalent Conns for use
//line /usr/local/go/src/database/sql/driver/driver.go:109
// by multiple goroutines.
//line /usr/local/go/src/database/sql/driver/driver.go:109
//
//line /usr/local/go/src/database/sql/driver/driver.go:109
// A Connector can be passed to sql.OpenDB, to allow drivers
//line /usr/local/go/src/database/sql/driver/driver.go:109
// to implement their own sql.DB constructors, or returned by
//line /usr/local/go/src/database/sql/driver/driver.go:109
// DriverContext's OpenConnector method, to allow drivers
//line /usr/local/go/src/database/sql/driver/driver.go:109
// access to context and to avoid repeated parsing of driver
//line /usr/local/go/src/database/sql/driver/driver.go:109
// configuration.
//line /usr/local/go/src/database/sql/driver/driver.go:109
//
//line /usr/local/go/src/database/sql/driver/driver.go:109
// If a Connector implements io.Closer, the sql package's DB.Close
//line /usr/local/go/src/database/sql/driver/driver.go:109
// method will call Close and return error (if any).
//line /usr/local/go/src/database/sql/driver/driver.go:121
type Connector interface {
	// Connect returns a connection to the database.
	// Connect may return a cached connection (one previously
	// closed), but doing so is unnecessary; the sql package
	// maintains a pool of idle connections for efficient re-use.
	//
	// The provided context.Context is for dialing purposes only
	// (see net.DialContext) and should not be stored or used for
	// other purposes. A default timeout should still be used
	// when dialing as a connection pool may call Connect
	// asynchronously to any query.
	//
	// The returned connection is only used by one goroutine at a
	// time.
	Connect(context.Context) (Conn, error)

	// Driver returns the underlying Driver of the Connector,
	// mainly to maintain compatibility with the Driver method
	// on sql.DB.
	Driver() Driver
}

// ErrSkip may be returned by some optional interfaces' methods to
//line /usr/local/go/src/database/sql/driver/driver.go:143
// indicate at runtime that the fast path is unavailable and the sql
//line /usr/local/go/src/database/sql/driver/driver.go:143
// package should continue as if the optional interface was not
//line /usr/local/go/src/database/sql/driver/driver.go:143
// implemented. ErrSkip is only supported where explicitly
//line /usr/local/go/src/database/sql/driver/driver.go:143
// documented.
//line /usr/local/go/src/database/sql/driver/driver.go:148
var ErrSkip = errors.New("driver: skip fast-path; continue as if unimplemented")

// ErrBadConn should be returned by a driver to signal to the sql
//line /usr/local/go/src/database/sql/driver/driver.go:150
// package that a driver.Conn is in a bad state (such as the server
//line /usr/local/go/src/database/sql/driver/driver.go:150
// having earlier closed the connection) and the sql package should
//line /usr/local/go/src/database/sql/driver/driver.go:150
// retry on a new connection.
//line /usr/local/go/src/database/sql/driver/driver.go:150
//
//line /usr/local/go/src/database/sql/driver/driver.go:150
// To prevent duplicate operations, ErrBadConn should NOT be returned
//line /usr/local/go/src/database/sql/driver/driver.go:150
// if there's a possibility that the database server might have
//line /usr/local/go/src/database/sql/driver/driver.go:150
// performed the operation. Even if the server sends back an error,
//line /usr/local/go/src/database/sql/driver/driver.go:150
// you shouldn't return ErrBadConn.
//line /usr/local/go/src/database/sql/driver/driver.go:150
//
//line /usr/local/go/src/database/sql/driver/driver.go:150
// Errors will be checked using errors.Is. An error may
//line /usr/local/go/src/database/sql/driver/driver.go:150
// wrap ErrBadConn or implement the Is(error) bool method.
//line /usr/local/go/src/database/sql/driver/driver.go:162
var ErrBadConn = errors.New("driver: bad connection")

// Pinger is an optional interface that may be implemented by a Conn.
//line /usr/local/go/src/database/sql/driver/driver.go:164
//
//line /usr/local/go/src/database/sql/driver/driver.go:164
// If a Conn does not implement Pinger, the sql package's DB.Ping and
//line /usr/local/go/src/database/sql/driver/driver.go:164
// DB.PingContext will check if there is at least one Conn available.
//line /usr/local/go/src/database/sql/driver/driver.go:164
//
//line /usr/local/go/src/database/sql/driver/driver.go:164
// If Conn.Ping returns ErrBadConn, DB.Ping and DB.PingContext will remove
//line /usr/local/go/src/database/sql/driver/driver.go:164
// the Conn from pool.
//line /usr/local/go/src/database/sql/driver/driver.go:171
type Pinger interface {
	Ping(ctx context.Context) error
}

// Execer is an optional interface that may be implemented by a Conn.
//line /usr/local/go/src/database/sql/driver/driver.go:175
//
//line /usr/local/go/src/database/sql/driver/driver.go:175
// If a Conn implements neither ExecerContext nor Execer,
//line /usr/local/go/src/database/sql/driver/driver.go:175
// the sql package's DB.Exec will first prepare a query, execute the statement,
//line /usr/local/go/src/database/sql/driver/driver.go:175
// and then close the statement.
//line /usr/local/go/src/database/sql/driver/driver.go:175
//
//line /usr/local/go/src/database/sql/driver/driver.go:175
// Exec may return ErrSkip.
//line /usr/local/go/src/database/sql/driver/driver.go:175
//
//line /usr/local/go/src/database/sql/driver/driver.go:175
// Deprecated: Drivers should implement ExecerContext instead.
//line /usr/local/go/src/database/sql/driver/driver.go:184
type Execer interface {
	Exec(query string, args []Value) (Result, error)
}

// ExecerContext is an optional interface that may be implemented by a Conn.
//line /usr/local/go/src/database/sql/driver/driver.go:188
//
//line /usr/local/go/src/database/sql/driver/driver.go:188
// If a Conn does not implement ExecerContext, the sql package's DB.Exec
//line /usr/local/go/src/database/sql/driver/driver.go:188
// will fall back to Execer; if the Conn does not implement Execer either,
//line /usr/local/go/src/database/sql/driver/driver.go:188
// DB.Exec will first prepare a query, execute the statement, and then
//line /usr/local/go/src/database/sql/driver/driver.go:188
// close the statement.
//line /usr/local/go/src/database/sql/driver/driver.go:188
//
//line /usr/local/go/src/database/sql/driver/driver.go:188
// ExecContext may return ErrSkip.
//line /usr/local/go/src/database/sql/driver/driver.go:188
//
//line /usr/local/go/src/database/sql/driver/driver.go:188
// ExecContext must honor the context timeout and return when the context is canceled.
//line /usr/local/go/src/database/sql/driver/driver.go:198
type ExecerContext interface {
	ExecContext(ctx context.Context, query string, args []NamedValue) (Result, error)
}

// Queryer is an optional interface that may be implemented by a Conn.
//line /usr/local/go/src/database/sql/driver/driver.go:202
//
//line /usr/local/go/src/database/sql/driver/driver.go:202
// If a Conn implements neither QueryerContext nor Queryer,
//line /usr/local/go/src/database/sql/driver/driver.go:202
// the sql package's DB.Query will first prepare a query, execute the statement,
//line /usr/local/go/src/database/sql/driver/driver.go:202
// and then close the statement.
//line /usr/local/go/src/database/sql/driver/driver.go:202
//
//line /usr/local/go/src/database/sql/driver/driver.go:202
// Query may return ErrSkip.
//line /usr/local/go/src/database/sql/driver/driver.go:202
//
//line /usr/local/go/src/database/sql/driver/driver.go:202
// Deprecated: Drivers should implement QueryerContext instead.
//line /usr/local/go/src/database/sql/driver/driver.go:211
type Queryer interface {
	Query(query string, args []Value) (Rows, error)
}

// QueryerContext is an optional interface that may be implemented by a Conn.
//line /usr/local/go/src/database/sql/driver/driver.go:215
//
//line /usr/local/go/src/database/sql/driver/driver.go:215
// If a Conn does not implement QueryerContext, the sql package's DB.Query
//line /usr/local/go/src/database/sql/driver/driver.go:215
// will fall back to Queryer; if the Conn does not implement Queryer either,
//line /usr/local/go/src/database/sql/driver/driver.go:215
// DB.Query will first prepare a query, execute the statement, and then
//line /usr/local/go/src/database/sql/driver/driver.go:215
// close the statement.
//line /usr/local/go/src/database/sql/driver/driver.go:215
//
//line /usr/local/go/src/database/sql/driver/driver.go:215
// QueryContext may return ErrSkip.
//line /usr/local/go/src/database/sql/driver/driver.go:215
//
//line /usr/local/go/src/database/sql/driver/driver.go:215
// QueryContext must honor the context timeout and return when the context is canceled.
//line /usr/local/go/src/database/sql/driver/driver.go:225
type QueryerContext interface {
	QueryContext(ctx context.Context, query string, args []NamedValue) (Rows, error)
}

// Conn is a connection to a database. It is not used concurrently
//line /usr/local/go/src/database/sql/driver/driver.go:229
// by multiple goroutines.
//line /usr/local/go/src/database/sql/driver/driver.go:229
//
//line /usr/local/go/src/database/sql/driver/driver.go:229
// Conn is assumed to be stateful.
//line /usr/local/go/src/database/sql/driver/driver.go:233
type Conn interface {
	// Prepare returns a prepared statement, bound to this connection.
	Prepare(query string) (Stmt, error)

	// Close invalidates and potentially stops any current
	// prepared statements and transactions, marking this
	// connection as no longer in use.
	//
	// Because the sql package maintains a free pool of
	// connections and only calls Close when there's a surplus of
	// idle connections, it shouldn't be necessary for drivers to
	// do their own connection caching.
	//
	// Drivers must ensure all network calls made by Close
	// do not block indefinitely (e.g. apply a timeout).
	Close() error

	// Begin starts and returns a new transaction.
	//
	// Deprecated: Drivers should implement ConnBeginTx instead (or additionally).
	Begin() (Tx, error)
}

// ConnPrepareContext enhances the Conn interface with context.
type ConnPrepareContext interface {
	// PrepareContext returns a prepared statement, bound to this connection.
	// context is for the preparation of the statement,
	// it must not store the context within the statement itself.
	PrepareContext(ctx context.Context, query string) (Stmt, error)
}

// IsolationLevel is the transaction isolation level stored in TxOptions.
//line /usr/local/go/src/database/sql/driver/driver.go:264
//
//line /usr/local/go/src/database/sql/driver/driver.go:264
// This type should be considered identical to sql.IsolationLevel along
//line /usr/local/go/src/database/sql/driver/driver.go:264
// with any values defined on it.
//line /usr/local/go/src/database/sql/driver/driver.go:268
type IsolationLevel int

// TxOptions holds the transaction options.
//line /usr/local/go/src/database/sql/driver/driver.go:270
//
//line /usr/local/go/src/database/sql/driver/driver.go:270
// This type should be considered identical to sql.TxOptions.
//line /usr/local/go/src/database/sql/driver/driver.go:273
type TxOptions struct {
	Isolation	IsolationLevel
	ReadOnly	bool
}

// ConnBeginTx enhances the Conn interface with context and TxOptions.
type ConnBeginTx interface {
	// BeginTx starts and returns a new transaction.
	// If the context is canceled by the user the sql package will
	// call Tx.Rollback before discarding and closing the connection.
	//
	// This must check opts.Isolation to determine if there is a set
	// isolation level. If the driver does not support a non-default
	// level and one is set or if there is a non-default isolation level
	// that is not supported, an error must be returned.
	//
	// This must also check opts.ReadOnly to determine if the read-only
	// value is true to either set the read-only transaction property if supported
	// or return an error if it is not supported.
	BeginTx(ctx context.Context, opts TxOptions) (Tx, error)
}

// SessionResetter may be implemented by Conn to allow drivers to reset the
//line /usr/local/go/src/database/sql/driver/driver.go:295
// session state associated with the connection and to signal a bad connection.
//line /usr/local/go/src/database/sql/driver/driver.go:297
type SessionResetter interface {
	// ResetSession is called prior to executing a query on the connection
	// if the connection has been used before. If the driver returns ErrBadConn
	// the connection is discarded.
	ResetSession(ctx context.Context) error
}

// Validator may be implemented by Conn to allow drivers to
//line /usr/local/go/src/database/sql/driver/driver.go:304
// signal if a connection is valid or if it should be discarded.
//line /usr/local/go/src/database/sql/driver/driver.go:304
//
//line /usr/local/go/src/database/sql/driver/driver.go:304
// If implemented, drivers may return the underlying error from queries,
//line /usr/local/go/src/database/sql/driver/driver.go:304
// even if the connection should be discarded by the connection pool.
//line /usr/local/go/src/database/sql/driver/driver.go:309
type Validator interface {
	// IsValid is called prior to placing the connection into the
	// connection pool. The connection will be discarded if false is returned.
	IsValid() bool
}

// Result is the result of a query execution.
type Result interface {
	// LastInsertId returns the database's auto-generated ID
	// after, for example, an INSERT into a table with primary
	// key.
	LastInsertId() (int64, error)

	// RowsAffected returns the number of rows affected by the
	// query.
	RowsAffected() (int64, error)
}

// Stmt is a prepared statement. It is bound to a Conn and not
//line /usr/local/go/src/database/sql/driver/driver.go:327
// used by multiple goroutines concurrently.
//line /usr/local/go/src/database/sql/driver/driver.go:329
type Stmt interface {
	// Close closes the statement.
	//
	// As of Go 1.1, a Stmt will not be closed if it's in use
	// by any queries.
	//
	// Drivers must ensure all network calls made by Close
	// do not block indefinitely (e.g. apply a timeout).
	Close() error

	// NumInput returns the number of placeholder parameters.
	//
	// If NumInput returns >= 0, the sql package will sanity check
	// argument counts from callers and return errors to the caller
	// before the statement's Exec or Query methods are called.
	//
	// NumInput may also return -1, if the driver doesn't know
	// its number of placeholders. In that case, the sql package
	// will not sanity check Exec or Query argument counts.
	NumInput() int

	// Exec executes a query that doesn't return rows, such
	// as an INSERT or UPDATE.
	//
	// Deprecated: Drivers should implement StmtExecContext instead (or additionally).
	Exec(args []Value) (Result, error)

	// Query executes a query that may return rows, such as a
	// SELECT.
	//
	// Deprecated: Drivers should implement StmtQueryContext instead (or additionally).
	Query(args []Value) (Rows, error)
}

// StmtExecContext enhances the Stmt interface by providing Exec with context.
type StmtExecContext interface {
	// ExecContext executes a query that doesn't return rows, such
	// as an INSERT or UPDATE.
	//
	// ExecContext must honor the context timeout and return when it is canceled.
	ExecContext(ctx context.Context, args []NamedValue) (Result, error)
}

// StmtQueryContext enhances the Stmt interface by providing Query with context.
type StmtQueryContext interface {
	// QueryContext executes a query that may return rows, such as a
	// SELECT.
	//
	// QueryContext must honor the context timeout and return when it is canceled.
	QueryContext(ctx context.Context, args []NamedValue) (Rows, error)
}

// ErrRemoveArgument may be returned from NamedValueChecker to instruct the
//line /usr/local/go/src/database/sql/driver/driver.go:381
// sql package to not pass the argument to the driver query interface.
//line /usr/local/go/src/database/sql/driver/driver.go:381
// Return when accepting query specific options or structures that aren't
//line /usr/local/go/src/database/sql/driver/driver.go:381
// SQL query arguments.
//line /usr/local/go/src/database/sql/driver/driver.go:385
var ErrRemoveArgument = errors.New("driver: remove argument from query")

// NamedValueChecker may be optionally implemented by Conn or Stmt. It provides
//line /usr/local/go/src/database/sql/driver/driver.go:387
// the driver more control to handle Go and database types beyond the default
//line /usr/local/go/src/database/sql/driver/driver.go:387
// Values types allowed.
//line /usr/local/go/src/database/sql/driver/driver.go:387
//
//line /usr/local/go/src/database/sql/driver/driver.go:387
// The sql package checks for value checkers in the following order,
//line /usr/local/go/src/database/sql/driver/driver.go:387
// stopping at the first found match: Stmt.NamedValueChecker, Conn.NamedValueChecker,
//line /usr/local/go/src/database/sql/driver/driver.go:387
// Stmt.ColumnConverter, DefaultParameterConverter.
//line /usr/local/go/src/database/sql/driver/driver.go:387
//
//line /usr/local/go/src/database/sql/driver/driver.go:387
// If CheckNamedValue returns ErrRemoveArgument, the NamedValue will not be included in
//line /usr/local/go/src/database/sql/driver/driver.go:387
// the final query arguments. This may be used to pass special options to
//line /usr/local/go/src/database/sql/driver/driver.go:387
// the query itself.
//line /usr/local/go/src/database/sql/driver/driver.go:387
//
//line /usr/local/go/src/database/sql/driver/driver.go:387
// If ErrSkip is returned the column converter error checking
//line /usr/local/go/src/database/sql/driver/driver.go:387
// path is used for the argument. Drivers may wish to return ErrSkip after
//line /usr/local/go/src/database/sql/driver/driver.go:387
// they have exhausted their own special cases.
//line /usr/local/go/src/database/sql/driver/driver.go:402
type NamedValueChecker interface {
	// CheckNamedValue is called before passing arguments to the driver
	// and is called in place of any ColumnConverter. CheckNamedValue must do type
	// validation and conversion as appropriate for the driver.
	CheckNamedValue(*NamedValue) error
}

// ColumnConverter may be optionally implemented by Stmt if the
//line /usr/local/go/src/database/sql/driver/driver.go:409
// statement is aware of its own columns' types and can convert from
//line /usr/local/go/src/database/sql/driver/driver.go:409
// any type to a driver Value.
//line /usr/local/go/src/database/sql/driver/driver.go:409
//
//line /usr/local/go/src/database/sql/driver/driver.go:409
// Deprecated: Drivers should implement NamedValueChecker.
//line /usr/local/go/src/database/sql/driver/driver.go:414
type ColumnConverter interface {
	// ColumnConverter returns a ValueConverter for the provided
	// column index. If the type of a specific column isn't known
	// or shouldn't be handled specially, DefaultValueConverter
	// can be returned.
	ColumnConverter(idx int) ValueConverter
}

// Rows is an iterator over an executed query's results.
type Rows interface {
	// Columns returns the names of the columns. The number of
	// columns of the result is inferred from the length of the
	// slice. If a particular column name isn't known, an empty
	// string should be returned for that entry.
	Columns() []string

	// Close closes the rows iterator.
	Close() error

	// Next is called to populate the next row of data into
	// the provided slice. The provided slice will be the same
	// size as the Columns() are wide.
	//
	// Next should return io.EOF when there are no more rows.
	//
	// The dest should not be written to outside of Next. Care
	// should be taken when closing Rows not to modify
	// a buffer held in dest.
	Next(dest []Value) error
}

// RowsNextResultSet extends the Rows interface by providing a way to signal
//line /usr/local/go/src/database/sql/driver/driver.go:445
// the driver to advance to the next result set.
//line /usr/local/go/src/database/sql/driver/driver.go:447
type RowsNextResultSet interface {
	Rows

	// HasNextResultSet is called at the end of the current result set and
	// reports whether there is another result set after the current one.
	HasNextResultSet() bool

	// NextResultSet advances the driver to the next result set even
	// if there are remaining rows in the current result set.
	//
	// NextResultSet should return io.EOF when there are no more result sets.
	NextResultSet() error
}

// RowsColumnTypeScanType may be implemented by Rows. It should return
//line /usr/local/go/src/database/sql/driver/driver.go:461
// the value type that can be used to scan types into. For example, the database
//line /usr/local/go/src/database/sql/driver/driver.go:461
// column type "bigint" this should return "reflect.TypeOf(int64(0))".
//line /usr/local/go/src/database/sql/driver/driver.go:464
type RowsColumnTypeScanType interface {
	Rows
	ColumnTypeScanType(index int) reflect.Type
}

// RowsColumnTypeDatabaseTypeName may be implemented by Rows. It should return the
//line /usr/local/go/src/database/sql/driver/driver.go:469
// database system type name without the length. Type names should be uppercase.
//line /usr/local/go/src/database/sql/driver/driver.go:469
// Examples of returned types: "VARCHAR", "NVARCHAR", "VARCHAR2", "CHAR", "TEXT",
//line /usr/local/go/src/database/sql/driver/driver.go:469
// "DECIMAL", "SMALLINT", "INT", "BIGINT", "BOOL", "[]BIGINT", "JSONB", "XML",
//line /usr/local/go/src/database/sql/driver/driver.go:469
// "TIMESTAMP".
//line /usr/local/go/src/database/sql/driver/driver.go:474
type RowsColumnTypeDatabaseTypeName interface {
	Rows
	ColumnTypeDatabaseTypeName(index int) string
}

// RowsColumnTypeLength may be implemented by Rows. It should return the length
//line /usr/local/go/src/database/sql/driver/driver.go:479
// of the column type if the column is a variable length type. If the column is
//line /usr/local/go/src/database/sql/driver/driver.go:479
// not a variable length type ok should return false.
//line /usr/local/go/src/database/sql/driver/driver.go:479
// If length is not limited other than system limits, it should return math.MaxInt64.
//line /usr/local/go/src/database/sql/driver/driver.go:479
// The following are examples of returned values for various types:
//line /usr/local/go/src/database/sql/driver/driver.go:479
//
//line /usr/local/go/src/database/sql/driver/driver.go:479
//	TEXT          (math.MaxInt64, true)
//line /usr/local/go/src/database/sql/driver/driver.go:479
//	varchar(10)   (10, true)
//line /usr/local/go/src/database/sql/driver/driver.go:479
//	nvarchar(10)  (10, true)
//line /usr/local/go/src/database/sql/driver/driver.go:479
//	decimal       (0, false)
//line /usr/local/go/src/database/sql/driver/driver.go:479
//	int           (0, false)
//line /usr/local/go/src/database/sql/driver/driver.go:479
//	bytea(30)     (30, true)
//line /usr/local/go/src/database/sql/driver/driver.go:491
type RowsColumnTypeLength interface {
	Rows
	ColumnTypeLength(index int) (length int64, ok bool)
}

// RowsColumnTypeNullable may be implemented by Rows. The nullable value should
//line /usr/local/go/src/database/sql/driver/driver.go:496
// be true if it is known the column may be null, or false if the column is known
//line /usr/local/go/src/database/sql/driver/driver.go:496
// to be not nullable.
//line /usr/local/go/src/database/sql/driver/driver.go:496
// If the column nullability is unknown, ok should be false.
//line /usr/local/go/src/database/sql/driver/driver.go:500
type RowsColumnTypeNullable interface {
	Rows
	ColumnTypeNullable(index int) (nullable, ok bool)
}

// RowsColumnTypePrecisionScale may be implemented by Rows. It should return
//line /usr/local/go/src/database/sql/driver/driver.go:505
// the precision and scale for decimal types. If not applicable, ok should be false.
//line /usr/local/go/src/database/sql/driver/driver.go:505
// The following are examples of returned values for various types:
//line /usr/local/go/src/database/sql/driver/driver.go:505
//
//line /usr/local/go/src/database/sql/driver/driver.go:505
//	decimal(38, 4)    (38, 4, true)
//line /usr/local/go/src/database/sql/driver/driver.go:505
//	int               (0, 0, false)
//line /usr/local/go/src/database/sql/driver/driver.go:505
//	decimal           (math.MaxInt64, math.MaxInt64, true)
//line /usr/local/go/src/database/sql/driver/driver.go:512
type RowsColumnTypePrecisionScale interface {
	Rows
	ColumnTypePrecisionScale(index int) (precision, scale int64, ok bool)
}

// Tx is a transaction.
type Tx interface {
	Commit() error
	Rollback() error
}

// RowsAffected implements Result for an INSERT or UPDATE operation
//line /usr/local/go/src/database/sql/driver/driver.go:523
// which mutates a number of rows.
//line /usr/local/go/src/database/sql/driver/driver.go:525
type RowsAffected int64

var _ Result = RowsAffected(0)

func (RowsAffected) LastInsertId() (int64, error) {
//line /usr/local/go/src/database/sql/driver/driver.go:529
	_go_fuzz_dep_.CoverTab[179201]++
								return 0, errors.New("LastInsertId is not supported by this driver")
//line /usr/local/go/src/database/sql/driver/driver.go:530
	// _ = "end of CoverTab[179201]"
}

func (v RowsAffected) RowsAffected() (int64, error) {
//line /usr/local/go/src/database/sql/driver/driver.go:533
	_go_fuzz_dep_.CoverTab[179202]++
								return int64(v), nil
//line /usr/local/go/src/database/sql/driver/driver.go:534
	// _ = "end of CoverTab[179202]"
}

// ResultNoRows is a pre-defined Result for drivers to return when a DDL
//line /usr/local/go/src/database/sql/driver/driver.go:537
// command (such as a CREATE TABLE) succeeds. It returns an error for both
//line /usr/local/go/src/database/sql/driver/driver.go:537
// LastInsertId and RowsAffected.
//line /usr/local/go/src/database/sql/driver/driver.go:540
var ResultNoRows noRows

type noRows struct{}

var _ Result = noRows{}

func (noRows) LastInsertId() (int64, error) {
//line /usr/local/go/src/database/sql/driver/driver.go:546
	_go_fuzz_dep_.CoverTab[179203]++
								return 0, errors.New("no LastInsertId available after DDL statement")
//line /usr/local/go/src/database/sql/driver/driver.go:547
	// _ = "end of CoverTab[179203]"
}

func (noRows) RowsAffected() (int64, error) {
//line /usr/local/go/src/database/sql/driver/driver.go:550
	_go_fuzz_dep_.CoverTab[179204]++
								return 0, errors.New("no RowsAffected available after DDL statement")
//line /usr/local/go/src/database/sql/driver/driver.go:551
	// _ = "end of CoverTab[179204]"
}

//line /usr/local/go/src/database/sql/driver/driver.go:552
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/database/sql/driver/driver.go:552
var _ = _go_fuzz_dep_.CoverTab
