package main

import (
	conn "hisoka/src"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()

	conn.StartClient()
}
