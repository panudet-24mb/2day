package client

import (
	"encoding/json"

	"github.com/lovoo/goka"
)

type Codec struct{}

// Encode encodes a event struct into an array.
func (c *Codec) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

// Decode decodes a event from byte encoded array.
func (c *Codec) Decode(data []byte) (interface{}, error) {
	event := new(Event)

	err := json.Unmarshal(data, event)
	return event, err
}

// Producer defines an interface whose events are produced on kafka.
type Producer interface {
	Emit(key string, event *Event) error
	Close() error
}

type kafkaProducer struct {
	emitter *goka.Emitter
}
type Event struct {
	UserID    string `json:"user_id"`
	Timestamp int64  `json:"timestamp"`
}

// NewProducer returns a new kafka producer.
func NewProducer(brokers []string, stream string) (Producer, error) {
	codec := new(Codec)
	emitter, err := goka.NewEmitter(brokers, goka.Stream(stream), codec)
	if err != nil {
		return nil, err
	}
	return &kafkaProducer{emitter}, nil
}

func (p *kafkaProducer) Emit(key string, event *Event) error {
	return p.emitter.EmitSync(key, event)
}

func (p *kafkaProducer) Close() error {
	p.emitter.Finish()
	return nil
}

// kafkaProducer, _ := client.NewProducer([]string{"kafka:29092"}, "IdentityProvider")
// event := &client.Event{
// 	UserID:    strconv.FormatInt(rand.Int63n(255), 10),
// 	Timestamp: time.Now().UnixNano(),
// }
// kafkaProducer.Emit("1", event)

// package client

// import (
// 	"log"

// 	"github.com/lovoo/goka"
// 	"github.com/lovoo/goka/codec"
// )

// type Message struct {
// 	UserID    string `json:"user_id"`
// 	Timestamp int64  `json:"timestamp"`
// }

// var (
// 	brokers             = []string{"kafka:29092"}
// 	topic   goka.Stream = "IdentityProvider"
// 	// group   goka.Group  = "example-group"
// )

// // Emit messages forever every second
// func InitKafka() goka.Emitter {
// 	emitter, err := goka.NewEmitter(brokers, topic, new(codec.String))
// 	if err != nil {
// 		log.Fatalf("error creating emitter: %v", err)
// 	}
// 	return *emitter
// 	// defer emitter.Finish()
// 	// for {
// 	// 	time.Sleep(1 * time.Second)
// 	// 	err = emitter.EmitSync("some-key", "some-value")
// 	// 	if err != nil {
// 	// 		log.Fatalf("error emitting message: %v", err)
// 	// 	}
// 	// }
// }
