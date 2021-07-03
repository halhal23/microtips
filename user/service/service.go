package service

import (
	"context"
	"fmt"
	"microtips/user/pb"
	"microtips/user/pkg/jwt"
	"microtips/user/repository"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	ReadUserById(ctx context.Context, req *pb.ReadUserByIdRequest) (*pb.ReadUserByIdResponse, error)
	ReadUserByName(ctx context.Context, req *pb.ReadUserByNameRequest) (*pb.ReadUserByNameResponse, error)
	UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
	ListUser(req *pb.ListUserRequest, stream pb.UserService_ListUserServer) error
	SignUp(ctx context.Context, input *pb.SignUpRequest) (*pb.SignUpResponse, error)
	SignIn(ctx context.Context, input *pb.SignInRequest) (*pb.SignInResponse, error)
}

type service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{r}
}

func (s *service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	input := req.GetUserInput()

	hassedPassword, err := HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	id, err := s.repository.InsertUser(ctx, &pb.UserInput{
		Name:     input.Name,
		Password: hassedPassword,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		User: &pb.User{
			Id:       id,
			Name:     input.Name,
			Password: hassedPassword,
		},
	}, nil
}

func (s *service) ReadUserById(ctx context.Context, req *pb.ReadUserByIdRequest) (*pb.ReadUserByIdResponse, error) {
	id := req.GetId()
	user, err := s.repository.SelectUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &pb.ReadUserByIdResponse{
		User: &pb.User{
			Id:       id,
			Name:     user.Name,
			Password: user.Password,
		},
	}, nil
}

func (s *service) ReadUserByName(ctx context.Context, req *pb.ReadUserByNameRequest) (*pb.ReadUserByNameResponse, error) {
	name := req.GetName()
	user, err := s.repository.SelectUserByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return &pb.ReadUserByNameResponse{
		User: &pb.User{
			Id:       user.Id,
			Name:     user.Name,
			Password: user.Password,
		},
	}, nil
}

func (s *service) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	id := req.GetId()
	input := req.GetUserInput()
	err := s.repository.UpdateUser(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{
		User: &pb.User{
			Id:       id,
			Name:     input.Name,
			Password: input.Password,
		},
	}, nil
}

func (s *service) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	id := req.GetId()
	err := s.repository.DeleteUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{Id: id}, nil
}

func (s *service) ListUser(req *pb.ListUserRequest, stream pb.UserService_ListUserServer) error {
	rows, err := s.repository.SelectAllUsers()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var user pb.User
		err := rows.Scan(&user.Id, &user.Name, &user.Password)
		if err != nil {
			return err
		}

		stream.Send(&pb.ListUserResponse{User: &user})
	}
	return nil
}

func (s *service) SignUp(ctx context.Context, input *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	hashedPassword, err := HashPassword(input.UserInput.Password)
	if err != nil {
		return nil, err
	}
	_, err = s.repository.InsertUser(ctx, &pb.UserInput{Name: input.UserInput.Name, Password: hashedPassword})
	if err != nil {
		return nil, err
	}
	token, err := jwt.GenerateToken(input.UserInput.Name)
	if err != nil {
		return nil, err
	}
	fmt.Printf("これが発行されたJWTです: %v\n", token)
	return &pb.SignUpResponse{
		Token: token,
	}, nil
}

func (s *service) SignIn(ctx context.Context, input *pb.SignInRequest) (*pb.SignInResponse, error) {
	user, err := s.repository.SelectUserByName(ctx, input.UserInput.Name)
	if err != nil {
		return nil, err
	}
	ok := CheckPasswordHash(input.UserInput.Password, user.Password)
	if !ok {
		return nil, fmt.Errorf("パスワードが合っていませんでした。")
	}
	token, err := jwt.GenerateToken(input.UserInput.Name)
	if err != nil {
		return nil, err
	}
	fmt.Printf("これが発行されたJWTです: %v\n", token)
	return &pb.SignInResponse{
		Token: token,
	}, nil
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
