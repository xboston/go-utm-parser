# go-utm-parser
Fast UTM Parser

```go
package main

import (
	"log"

	utm "github.com/xboston/go-utm-parser"
)

func main() {

	url := "http://www.example.com/?utm_source=source&utm_campaign=campaign&utm_medium=medium&utm_content=content&utm_term=term&utm_extra1=extra1&utm_extra2=extra2"
	utm, _ := utm.ParseURL(&url)

	log.Println(utm)
}
```


```
$ go version
go version go1.7.3 darwin/amd64
$ go test -bench . -benchmem
BenchmarkParseURL-4                 	  300000	      4878 ns/op	     976 B/op	      12 allocs/op
BenchmarkParallelParseURL-4         	  500000	      2902 ns/op	     976 B/op	      12 allocs/op
BenchmarkParseQuery-4               	 3000000	       448 ns/op	       0 B/op	       0 allocs/op
BenchmarkParallelParseQuery-4       	10000000	       188 ns/op	       0 B/op	       0 allocs/op
BenchmarkParseQueryFull-4           	 2000000	       821 ns/op	     336 B/op	       2 allocs/op
BenchmarkParallelParseQueryFull-4   	 5000000	       351 ns/op	     336 B/op	       2 allocs/op
PASS
ok  	github.com/xboston/go-utm-parser	11.506s
```