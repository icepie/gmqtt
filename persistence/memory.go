package persistence

import (
	"github.com/DrmagicE/gmqtt/config"
	"github.com/DrmagicE/gmqtt/persistence/queue"
	mem_queue "github.com/DrmagicE/gmqtt/persistence/queue/mem"
	"github.com/DrmagicE/gmqtt/persistence/session"
	mem_session "github.com/DrmagicE/gmqtt/persistence/session/mem"
	"github.com/DrmagicE/gmqtt/persistence/subscription"
	mem_sub "github.com/DrmagicE/gmqtt/persistence/subscription/mem"
	"github.com/DrmagicE/gmqtt/persistence/unack"
	mem_unack "github.com/DrmagicE/gmqtt/persistence/unack/mem"
	"github.com/DrmagicE/gmqtt/server"
)

func init() {
	server.RegisterPersistenceFactory("memory", NewMemory)
}

func NewMemory(config config.Config, hooks server.Hooks) (server.Persistence, error) {
	return &memory{
		onMsgDropped: hooks.OnMsgDropped,
	}, nil
}

type memory struct {
	onMsgDropped server.OnMsgDropped
}

func (m *memory) NewUnackStore(config config.Config, clientID string) (unack.Store, error) {
	return mem_unack.New(mem_unack.Options{
		ClientID: clientID,
	}), nil
}

func (m *memory) NewSessionStore(config config.Config) (session.Store, error) {
	return mem_session.New(), nil
}

func (m *memory) Open() error {
	return nil
}
func (m *memory) NewQueueStore(config config.Config, clientID string) (queue.Store, error) {
	return mem_queue.New(mem_queue.Options{
		MaxQueuedMsg: config.MQTT.MaxQueuedMsg,
		ClientID:     clientID,
		DropHandler:  m.onMsgDropped,
	})
}

func (m *memory) NewSubscriptionStore(config config.Config) (subscription.Store, error) {
	return mem_sub.NewStore(), nil
}

func (m *memory) Close() error {
	return nil
}
