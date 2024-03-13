package routes

import (
	"github.com/emoral435/swole-goal/api/token"
	db "github.com/emoral435/swole-goal/db/sqlc"
	util "github.com/emoral435/swole-goal/utils"
)

// ServerStore represents the information that we need to share accross the application.
// We need to manage out tokens, and we also need to manage our configuration for the server.
type ServerStore struct {
	TokenMaker token.Maker
	Config     util.Config
	Store      *db.Store
}

// CreateServerStore creates a new ServerStore instance
func CreateServerStore(t token.Maker, c util.Config, s *db.Store) *ServerStore {
	return &ServerStore{
		t,
		c,
		s,
	}
}
