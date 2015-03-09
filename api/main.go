// Package api defines data types representing core collectd data types.
package api

import (
	"fmt"
	"time"
)

// Value represents either a Gauge or a Derive. It is Go's equivalent to the C
// union value_t.
type Value interface {
	String() string
}

// Gauge represents a gauge metric value, such as a temperature. It is Go's
// equivalent to the C type gauge_t.
type Gauge float64

func (g Gauge) String() string {
	return fmt.Sprintf("%g", g)
}

// Derive represents a derive or counter metric value, such as bytes sent over
// the network. It is Go's equivalent to the C type derive_t.
type Derive int64

func (d Derive) String() string {
	return fmt.Sprintf("%d", d)
}

// Identifier identifies one metric.
type Identifier struct {
	Host                   string
	Plugin, PluginInstance string
	Type, TypeInstance     string
}

// ValueList represents one (set of) data point(s) of one metric. It is Go's
// equivalent of the C type value_list_t.
type ValueList struct {
	Identifier
	Time     time.Time
	Interval time.Duration
	Values   []Value
}

// String returns a string representation of the Identifier.
func (id Identifier) String() string {
	str := id.Host + "/" + id.Plugin
	if id.PluginInstance != "" {
		str += "-" + id.PluginInstance
	}
	str += "/" + id.Type
	if id.TypeInstance != "" {
		str += "-" + id.TypeInstance
	}
	return str
}