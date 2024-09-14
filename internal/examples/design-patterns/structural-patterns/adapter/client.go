package main

import (
	"fmt"

	"go-feather-lib/internal/examples/design-patterns/structural-patterns/adapter/machines"
)

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com machines.Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
