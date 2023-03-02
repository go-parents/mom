package mom

// PoolKey represents an identifier for a specific pool(s) that should receive messages for a given topic
// MOM will round-robin the messages between services within the same pool.
// MOM will duplicate the message between services in different pools.
type PoolKey string