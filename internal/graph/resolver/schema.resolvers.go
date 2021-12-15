package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kaitoimai01/gqlgen_auth0_sample/internal/graph/generated"
	"github.com/kaitoimai01/gqlgen_auth0_sample/internal/graph/model"
	"github.com/kaitoimai01/gqlgen_auth0_sample/internal/user"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, input model.UserInfo) (*model.Auth0UserID, error) {
  // user/user.go を参照
	var u = &user.User{
		Email: input.Email,
		FamilyName: input.FamilyName,
		GivenName: input.GivenName,
	}

	res, err := user.RegisterUser(ctx, &user.RegisterUserInput{
		User: u,
		Auth0ManagementClient: r.Auth0Client,
		MailClient:            r.MailClient,
	})
	if err != nil {
		return nil, fmt.Errorf("could not register user: %w", err)
	}

	return &model.Auth0UserID{
		Auth0UserID: res,
	}, nil
}

func (r *queryResolver) MemberID(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
