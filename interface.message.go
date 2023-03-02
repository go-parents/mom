package mom

import "context"

type IMessage interface {
    Headers() IHeaders
    Parameters() IParameters
    Ack(response interface{})
    Raw() []byte
    Unmarshal(v interface{})
    Context() context.Context
}