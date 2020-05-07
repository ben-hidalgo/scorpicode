package hatmongo

import (
	"backend/internal/hats/hatdao"
	"context"

	"github.com/twitchtv/twirp"
	"go.mongodb.org/mongo-driver/mongo"
)

// ServerHooks is a Twirp middleware specific to the "hats" service for each dao in "hats"
func ServerHooks(mc *mongo.Client) *twirp.ServerHooks {

	return &twirp.ServerHooks{
		RequestReceived: func(ctx context.Context) (context.Context, error) {

			// each Dao will be added to the context separately
			ctx = context.WithValue(ctx, hatdao.Key, hatdao.New(mc))

			return ctx, nil
		},
	}
}
