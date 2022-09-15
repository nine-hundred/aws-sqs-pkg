package consumers

import "testing"

func Test_Consume(t *testing.T) {
	t.Run("consume test", func(t *testing.T) {
		c := Consumer{}
		c.Consume()
	})
}
