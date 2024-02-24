package grpc

import (
	"context"
	"errors"
	"strconv"

	"github.com/allan7yin/grpc-go/internal/models"
	interfaces "github.com/allan7yin/grpc-go/pkg/v1"
	pb "github.com/allan7yin/grpc-go/proto"
	"google.golang.org/grpc"
)

type UserServStruct struct {
	useCase interfaces.UseCaseInterface
	pb.UnimplementedUserServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase interfaces.UseCaseInterface) {
	userGrpc := &UserServStruct{useCase: usecase}
	pb.RegisterUserServiceServer(grpcServer, userGrpc)
}

// Create
//
// This function creates a user whose data is supplied
// through the CreateUserRequest message of proto
func (srv *UserServStruct) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserProfileResponse, error) {
	data := srv.transformUserRPC(req)
	if data.Email == "" || data.Name == "" {
		return &pb.UserProfileResponse{}, errors.New("please provide all fields")
	}

	user, err := srv.useCase.Create(data)
	if err != nil {
		return &pb.UserProfileResponse{}, err
	}
	return srv.transformUserModel(user), nil
}

// Get
//
// This function returns the user instance of which ID
// is supplied through the SingleUserRequest message field of proto
func (srv *UserServStruct) Read(ctx context.Context, req *pb.SingleUserRequest) (*pb.UserProfileResponse, error) {
	id := req.GetId()
	if id == "" {
		return &pb.UserProfileResponse{}, errors.New("id cannot be blank")
	}
	user, err := srv.useCase.Get(id)
	if err != nil {
		return &pb.UserProfileResponse{}, err
	}
	return srv.transformUserModel(user), nil
}

// Delete
//
// This function deletes the user, returning a success message if successful
// is supplied through the SingleUserRequest message field of proto
func (srv *UserServStruct) Delete(ctx context.Context, req *pb.SingleUserRequest) (*pb.SuccessResponse, error) {
	id := req.GetId()
	if id == "" {
		return &pb.SuccessResponse{}, errors.New("id cannot be blank")
	}
	if err := srv.useCase.Delete(id); err != nil {
		return &pb.SuccessResponse{Response: "0"}, err
	}
	return &pb.SuccessResponse{Response: "1"}, nil
}

func (srv *UserServStruct) transformUserRPC(req *pb.CreateUserRequest) models.User {
	return models.User{Name: req.GetName(), Email: req.GetEmail()}
}

func (srv *UserServStruct) transformUserModel(user models.User) *pb.UserProfileResponse {
	return &pb.UserProfileResponse{Id: strconv.Itoa(int(user.ID)), Name: user.Name, Email: user.Email}
}
