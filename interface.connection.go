package mom

import (
    "io"
    "time"
)

type IConn interface {
    Publish(topic TopicKey, payload Payload) error
    Request(topic TopicKey, payload Payload, timeout time.Duration) (IMessage, error)

    Subscribe(topic TopicKey, pool PoolKey, handler MessageHandler) (ISubscription, error)

    StreamWriter(topic TopicKey) (io.Writer, error)
    StreamReader(msg IMessage) (io.Reader, error)

    Close()
}