package main

import "github.com/leonardo-bonfim/go-design-paterns/patterns/adapter"

func adapterExample() {
	client := &adapter.Client{}
	mac := &adapter.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &adapter.Windows{}
	windowsMachineAdapter := adapter.WindowsAdapter{
		WindowsMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(&windowsMachineAdapter)
}

func main() {

	adapterExample()

}
