package hatmongo

import (
	"backend/internal/hats/hatdao"
	"backend/internal/hats/orderdao"
	"context"

	"github.com/twitchtv/twirp"
	"go.mongodb.org/mongo-driver/mongo"
)

// ServerHooks is a Twirp middleware specific to the "hats" service for each dao in "hats"
func ServerHooks(mc *mongo.Client) *twirp.ServerHooks {

	return &twirp.ServerHooks{
		RequestReceived: func(ctx context.Context) (context.Context, error) {

			// each DAO is added to the context separately
			ctx = context.WithValue(ctx, hatdao.Key, hatdao.New(mc))
			ctx = context.WithValue(ctx, orderdao.Key, orderdao.New(mc))

			return ctx, nil
		},
	}
}
