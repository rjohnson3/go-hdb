package sqltrace_test

import (
	"github.com/rjohnson3/go-hdb/driver/sqltrace"
)

func Example() {
	sqltrace.SetOn(true)  // set SQL trace output active
	sqltrace.SetOn(false) // set SQL trace output inactive
}
