package message

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getTierMessage() []byte {
	return []byte(`[{"message": "1231", "mes": { "id": 123 }, "flat_rate": 11.241, "array": [1, 12, 33]}]`)
}

func TestData_AccessProperty(t *testing.T) {
	msg := NewMessage(Snapshot, "flights", getTierMessage())
	assert.Equal(t, float64(123), msg.Data.AccessProperty("mes.id"))
	assert.Equal(t, "1231", msg.Data.AccessProperty("message"))
}

func TestData_SetProperty(t *testing.T) {
	msg := NewMessage(Snapshot, "flights", getTierMessage())
	msg.Data.SetProperty("ai_result_1111", "example_result_data")
	assert.Equal(t, "example_result_data", msg.Data.AccessProperty("ai_result_1111"))
}

func TestData_DopProperty(t *testing.T) {
	msg := NewMessage(Snapshot, "flights", getTierMessage())
	msg.Data.DropProperty("ai_result_1111")
	assert.Equal(t, nil, msg.Data.AccessProperty("ai_result_1111"))
}
