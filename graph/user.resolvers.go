package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"microtips/graph/model"
	"microtips/user/pb"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	res, err := r.UserClient.Service.CreateUser(ctx, &pb.CreateUserRequest{
		UserInput: &pb.UserInput{
			Name:     input.Name,
			Password: input.Password,
		},
	})
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:       int(res.User.Id),
		Name:     res.User.Name,
		Password: res.User.Password,
	}, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	res, err := r.UserClient.Service.UpdateUser(ctx, &pb.UpdateUserRequest{
		Id:        int64(input.ID),
		UserInput: &pb.UserInput{Name: input.Name, Password: input.Password},
	})
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:       int(res.User.Id),
		Name:     res.User.Name,
		Password: res.User.Password,
	}, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input int) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, input int) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) User(ctx context.Context, input int) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) DeleteUser(ctx context.Context, input int) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
