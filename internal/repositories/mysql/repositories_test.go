package repositories

import (
	"database/sql/driver"
	"time"
)

/*
	By implementing the Match method on the AnyTime type, we're allowing flexibility in matching the time.Time
	values within your SQL mock expectations, ensuring that the test will accept any value of type time.Time
*/

// AnyTime used to mock time object
type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
