package actor

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type DummyMessage struct{}
type BlackHoleActor struct{}

var testTimeout = 1 * time.Second

func (state *BlackHoleActor) Receive(context Context) {}

func NewBlackHoleActor() Actor {
	return &BlackHoleActor{}
}

func TestSpawnProducesActorRef(t *testing.T) {
	actor := Spawn(FromProducer(NewBlackHoleActor))
	defer actor.Stop()
	assert.NotNil(t, actor)
}

type EchoRequest struct{}

type EchoResponse struct{}

type EchoActor struct{}

func NewEchoActor() Actor {
	return &EchoActor{}
}

func (*EchoActor) Receive(context Context) {
	switch context.Message().(type) {
	case EchoRequest:
		context.Respond(EchoResponse{})
	}
}

func TestActorCanReplyToMessage(t *testing.T) {
	actor := Spawn(FromProducer(NewEchoActor))
	defer actor.Stop()
	err := actor.RequestFuture(EchoRequest{}, testTimeout).Wait()
	if err != nil {
		assert.Fail(t, "timed out")
		return
	}
}
