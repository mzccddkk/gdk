# slicex

Go slice functions

## Installation

```go
go get github.com/mzccddkk/gdk/slicex
```

## Functions

```go
func Column(input []map[interface{}]interface{}, columnKey interface{}, indexKey ...interface{}) interface{}

func In(needle interface{}, haystack []interface{}) bool

func Diff(slice1 []interface{}, sliceOthers ...[]interface{}) []interface{}

func Intersect(slice1 []interface{}, sliceOthers ...[]interface{}) []interface{}

func Unique(slice []interface{}) []interface{}

func CountValues(slice []interface{}) map[interface{}]uint

func Search(haystack []interface{}, needle interface{}) int
```
