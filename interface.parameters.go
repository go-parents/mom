package mom

type IParameters interface {
    Get(InEnumOption, string) []string
    First(InEnumOption, string) string
}