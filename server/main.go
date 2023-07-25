package main

import (
	_ "database/sql"
	"fmt"
	_ "log"
	_ "net/http"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello, World!")
}
