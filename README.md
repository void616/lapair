Random Latin words pair (adjective + noun)


```go
package main

import (
	"log"
	"math/rand"

	"github.com/void616/lapair"
)

func main() {
	x := uint16(rand.Int())
	y := uint16(rand.Int())
	noun, adj := lapair.Pair(x, y)
	log.Println(noun, adj)
}
```

Output:
```
provincia fortis
```
