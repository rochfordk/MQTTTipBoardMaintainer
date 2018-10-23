package main

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/rochfordk/tipboardsubscriber"
)

func main() {

	fmt.Println("Is this thing on?......")
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	fmt.Println("Setting up MQTT options......")
	opts := MQTT.NewClientOptions().AddBroker("tcp://192.168.1.6:1883")
	//http.ListenAndServe(":8080", nil)
	//opts.SetDefaultPublishHandler(f)

	fmt.Println("Creating MQTT client......")
	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	fmt.Println("Creating Dashboard......")
	dash := tipboardsubscriber.TipboardDash{DashHost: "127.0.0.1", DashPort: 7272, DashAPIKey: "13446d45c9544d4da0b981bd946de743"}

	fmt.Println("Creating SimpleSubscriber......")
	sub1 := &tipboardsubscriber.SimpleSubscriber{Dashboard: dash, TileKey: "clients_total", Title: "Total", Description: "(Total number of registered clients)", Value: 9}

	//fmt.Println("Dash", dash.DashHost)
	//fmt.Println("Dash", sub.Dashboard.DashHost)


	//sub1.Subscribe("192.168.1.6", 9001, "test_topic/msg")
	fmt.Println("Subscribing to test topic......")
	go sub1.Subscribe(c, "test_topic/msg")

	//curl http://localhost:7272/api/v0.1/13446d45c9544d4da0b981bd946de743/push -X POST -d "tile=just_value" -d "key=clients_connected" -d 'data={"title": "Connected:", "description": "(Currently Connected Clients)", "just-value": "23"}'

	/*sub2 := &tipboardsubscriber.SimpleSubscriber{Dashboard: dash, TileKey: "clients_connected", Title: "Connected", Description: "(Currently Connected Clients)", Value: 19}
	sub2.Subscribe("192.168.1.6", 9001, "test_topic/msg")*/
}
