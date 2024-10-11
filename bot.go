//go:build bot
// +build bot

package main

import (
	"fmt"

	"github.com/piheta/sept/backend/sn"
)

func main() {
	// Bot logic
	fmt.Println("Bot is running...")
	sn.SuperNode(8080)
}
