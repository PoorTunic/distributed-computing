package model

import "time"

// Data is the model in the server
type Data struct {
	ID string `json:"id"`
	// Msg contains an info row
	Data string `json:"data"`
	// Source info
	Source string `json:"source"`
	// LastUse time
	LastUse time.Time `json:"tmpstmp"`
}
