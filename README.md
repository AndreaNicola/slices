# slices

```golang
package main

import (
	"fmt"
	"github.com/AndreaNicola/slices"
)

var mapper = func(a string) interface{} {
	return a + "aaaa"
}

var filter = func(a string) bool {
	return "B" == a
}

func main() {

	s := slices.New[string]([]string{"A", "B", "C", "D", "E", "F", "G", "H"})
	filtered := s.Filter(filter)
	mapped := filtered.Map(mapper)
	fmt.Println(s, filtered, mapped)

}

```

Output

```
[B C D E F G H A] [B] [Baaaa]
```