# Go Reflection Sample

## Simple

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	look(3)
}

func look(v interface{}) {
	fmt.Println(reflect.TypeOf(v))
  fmt.Println(reflect.ValueOf(v))
  fmt.Println(reflect.ValueOf(v).Kind())
}
```
