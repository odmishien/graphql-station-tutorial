package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/odmishien/graphql-station-tutorial/graph/generated"
	"github.com/odmishien/graphql-station-tutorial/graph/model"
)

func (r *queryResolver) StationByName(ctx context.Context, stationName *string) ([]*model.Station, error) {
	return r.StationByName(ctx, stationName)
}

func (r *queryResolver) StationByCd(ctx context.Context, stationCd *int) (*model.Station, error) {
	return r.StationByCd(ctx, stationCd)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
