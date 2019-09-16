# uuid

The uuid package generates UUIDs based on [RFC 4122](https://tools.ietf.org/html/rfc4122 "RFC 4122"): Time-Based UUID

## Install

`go get -u github.com/m7shapan/uuid`

## How to use

```go
package main

import (
	"fmt"

	"github.com/m7shapan/uuid"
)

func main() {
	fmt.Println(uuid.NewUUID()) // 2bf08894-47a3978-bbdd8c85-70d16847
}
```
