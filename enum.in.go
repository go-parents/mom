package mom

type InEnumOption string

// Used in the Parameter model
const (
	InQuery  InEnumOption = "query"
	InHeader InEnumOption = "header"
	InPath   InEnumOption = "path"
	InCookie InEnumOption = "cookie"
)
