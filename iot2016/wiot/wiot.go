package wiot

import (
	"fmt"
	OW "misterjulz/iot2016/iot2016/openwhisk"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type WatsonIoT struct {
	Client MQTT.Client
	Topic  string
}

func getMqttOpts(url string, clientID string, username string, password string) *MQTT.ClientOptions {

}

func New(url string, topic string, clientID string, username string, password string) *WatsonIoT {

}

func (wiot *WatsonIoT) Connect() {

}

func (wiot *WatsonIoT) Subscribe(ow *OW.OpenWhisk) {

}

func (wiot *WatsonIoT) Unsubscribe() {

}

func (wiot *WatsonIoT) Disconnect() {

}

func OWMessageHandler(fn func(MQTT.Client, MQTT.Message, *OW.OpenWhisk), ow *OW.OpenWhisk) MQTT.MessageHandler {
	return func(client MQTT.Client, msg MQTT.Message) {
		fn(client, msg, ow)
	}
}

func sayHello(client MQTT.Client, msg MQTT.Message, ow *OW.OpenWhisk) {
	fmt.Printf("the messsage is %s", msg.Payload())
	fmt.Println("trigger ow action")
	ow.TriggerAction("hello", `{"arg":"watson"}`)

}
