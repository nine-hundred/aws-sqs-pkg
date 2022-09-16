package payloads

import "time"

type Payload struct {
	Title        string
	Author       string
	Body         string
	ReceivedTime time.Time
}
