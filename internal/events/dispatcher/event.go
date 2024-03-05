package dispatcher

import "github.com/electivetechnology/utility-library-go/pubsub"

type Event struct {
	Name    string
	Payload pubsub.Message
}
