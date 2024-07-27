package cli

import (
	"fmt"
	"scout-store-admin-cli/handler"
)

type Cli struct {
	Handler *handler.Handler
}

func New(handler *handler.Handler) *Cli {
	return &Cli{
		Handler: handler,
	}
}

func (c *Cli) Init() {
	fmt.Println("Selamat Datang di Toko Aksesoris Pramuka!")
}
