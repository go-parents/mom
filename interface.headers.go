package mom

type IHeaders interface {
    Get(key string) string
}