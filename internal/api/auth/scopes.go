package auth

type Scope string

const (
	AuthScopeApp  Scope = "app"
	AuthScopeUser Scope = "user"
)

func (s Scope) String() string {
	return string(s)
}
