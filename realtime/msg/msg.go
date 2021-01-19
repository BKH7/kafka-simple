package msg

import (
	"encoding/json"

	"github.com/BKH7/kafka-simple/realtime/conn"
)

// Msgbody ...
type Msgbody struct {
	ID     int    `json:"Msg_id"`
	Sender string `json:"Sender"`
	Msg    string `json:"Msg"`
}

// Producer ...
func Producer(data *Msgbody) error {
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}

	conn.Producer("mytopic", j)
	return nil
}
