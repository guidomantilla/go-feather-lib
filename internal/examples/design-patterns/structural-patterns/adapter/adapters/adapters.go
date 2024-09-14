package adapters

import (
	"fmt"

	"go-feather-lib/internal/examples/design-patterns/structural-patterns/adapter/machines"
)

var _ machines.Computer = (*WindowsAdapter)(nil)

type WindowsAdapter struct {
	WindowMachine *machines.Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.WindowMachine.InsertIntoUSBPort()
}
