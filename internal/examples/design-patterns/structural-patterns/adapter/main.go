package main

import (
	"go-feather-lib/internal/examples/design-patterns/structural-patterns/adapter/adapters"
	"go-feather-lib/internal/examples/design-patterns/structural-patterns/adapter/machines"
)

func main() {

	client := &Client{}
	mac := &machines.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &machines.Windows{}
	windowsMachineAdapter := &adapters.WindowsAdapter{
		WindowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
