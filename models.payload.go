package mom

import "context"

type Payload struct {
    Headers map[string]string
    Parameters map[string][]string
    Body interface{}
    Context context.Context
}