package message

import (
	"encoding/json"
	"github.com/ivancorrales/knoa"
	"github.com/thedevsaddam/gojsonq/v2"
)

type Data struct {
	jsonq  *gojsonq.JSONQ
	packet knoa.Knoa[[]any]
}

func NewData(bytes []byte) Data {
	data := Data{}
	var proxyObj []interface{}
	json.Unmarshal(bytes, &proxyObj)
	data.packet = knoa.FromArray(proxyObj)

	return data
}

// AccessProperty returns property value from the raw data bytes
// You can access nested properties by separate nesting levels with '.' like
// 'driver.driver.profile.name'
func (d Data) AccessProperty(property string) interface{} {
	return gojsonq.New().FromString(d.packet.JSON()).Find("[0]." + property)
}

// SetProperty appends top level property to the response
// It requires copying the entire raw object
func (d Data) SetProperty(property string, data string) {
	d.packet.Set("[0]."+property, data)
	d.jsonq = gojsonq.New().FromInterface(d.packet.Out())
}

func (d Data) DropProperty(property string) {
	d.packet.Unset("[0]." + property)
	d.jsonq = gojsonq.New().FromInterface(d.packet.Out())
}
