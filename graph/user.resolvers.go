package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"io"
	"log"
	"microtips/graph/model"
	"microtips/user/middleware/auth"
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
	res, err := r.UserClient.Service.DeleteUser(ctx, &pb.DeleteUserRequest{Id: int64(input)})
	if err != nil {
		return 0, err
	}
	return int(res.Id), nil
}

func (r *mutationResolver) SignUp(ctx context.Context, input model.CreateUserInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SignIn(ctx context.Context, input model.CreateUserInput) (string, error) {
	res, err := r.UserClient.Service.SignIn(ctx, &pb.SignInRequest{
		UserInput: &pb.UserInput{Name: input.Name, Password: input.Password},
	})
	if err != nil {
		return "", nil
	}
	return res.Token, nil
}

func (r *queryResolver) User(ctx context.Context, input int) (*model.User, error) {
	res, err := r.UserClient.Service.ReadUserById(ctx, &pb.ReadUserByIdRequest{Id: int64(input)})
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:       int(res.User.Id),
		Name:     res.User.Name,
		Password: res.User.Password,
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	user := auth.ForContext(ctx)
	log.Printf("コンテキストから見つけたuser in graphql: %v\n", user)
	res, err := r.UserClient.Service.ListUser(ctx, &pb.ListUserRequest{})
	if err != nil {
		return nil, err
	}
	var users []*model.User
	for {
		r, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		users = append(users, &model.User{
			ID:       int(r.User.Id),
			Name:     r.User.Name,
			Password: r.User.Password,
		})
	}
	return users, nil
}
