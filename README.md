# randy
Package randy provides a simple way to generate random data.

## Usage
```go
package main

import (
	"log"
	"randy"
)

func main() {
	name, err := randy.Name();
	if err != nil {
		log.Fatal(err.Error());
	}

	println(name);
}
```

### Output:
```bash
Craig
```