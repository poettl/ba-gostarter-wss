package data

import (
	"context"

	"allaboutapps.dev/aw/go-starter/internal/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const (
	PlainTestUserPassword  = "password"
	HashedTestUserPassword = "$argon2id$v=19$m=65536,t=1,p=4$RFO8ulg2c2zloG0029pAUQ$2Po6NUIhVCMm9vivVDuzo7k5KVWfZzJJfeXzC+n+row" //nolint:gosec
)

// Live Service fixtures to be applied by manually running the CLI "app db seed"
// Note that these fixtures are not available while testing
// see the separate internal/test/fixtures.go file

type Upsertable interface {
	Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error
}

type FixtureMap struct {
	User1 *models.User
}

func Fixtures() FixtureMap {
	f := FixtureMap{}

	f.User1 = &models.User{
		ID:       "f6ede5d8-e22a-4ca5-aa12-67821865a3e5",
		IsActive: true,
		Username: null.StringFrom("user1@email.com"),
		Password: null.StringFrom(HashedTestUserPassword),
		Scopes:   []string{"app"},
	}
	return f

}

func Upserts() []Upsertable {
	return []Upsertable{
		Fixtures().User1,
	}
}
