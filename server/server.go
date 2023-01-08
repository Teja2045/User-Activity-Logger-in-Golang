package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "task1/proto"

	"google.golang.org/grpc"
)

const grpcPort = ":8082"

type ActivityService struct {
	pb.UnimplementedActivityServiceServer
}

func main() {

	log.Println("Starting application")

	// start listening for grpc
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal(err)
	}

	// Instanciate new gRPC server
	server := grpc.NewServer()

	//register service
	pb.RegisterActivityServiceServer(server, &ActivityService{})

	log.Println("Starting grpc server on port " + grpcPort)

	// Start the gRPC server in goroutine
	server.Serve(listen)

}

func (a *ActivityService) RegisterUser(ctx context.Context, req *pb.User) (*pb.UserResponse, error) {
	username := req.Name
	phone := req.Phone
	email := req.Email

	//addUser, err := Dbx.Queryx("INSERT INTO user_table(Name, Phone, Email) VALUES(?,?,?)", username, phone, email)

	fmt.Println("uer has been added:", username, phone, email)

	resp := &pb.UserResponse{
		Response: "The user " + username + " has been added.",
	}

	return resp, nil
}

func (a *ActivityService) UpdateUserInfo(ctx context.Context, req *pb.UpdateUser) (*pb.UserResponse, error) {
	username := req.User.Name
	phone := req.User.Phone
	email := req.User.Email

	fmt.Println("uer has been updated:", username, phone, email)

	resp := &pb.UserResponse{
		Response: "The user " + username + " has been updated.",
	}

	return resp, nil
}

func (a *ActivityService) GetUser(ctx context.Context, req *pb.Name) (*pb.UserResponse, error) {
	username := req.Name

	fmt.Println("Getuser is working:", username)

	resp := &pb.UserResponse{
		Response: "The user " + username + " GetUser() is working.",
	}

	return resp, nil
}

func (a *ActivityService) AddActivity(ctx context.Context, req *pb.Activity) (*pb.UserResponse, error) {
	activityType := req.Type
	label := req.Label
	duration := req.Duration

	fmt.Println("activity has been added:", activityType, label, duration)

	resp := &pb.UserResponse{
		Response: "The activity " + activityType + " is added.",
	}

	return resp, nil
}

func (a *ActivityService) ActivityIsDone(ctx context.Context, req *pb.ActivityRequest) (*pb.Done, error) {
	username := req.Username
	activitytype := req.Type

	fmt.Println("ActivityIsDone() is working:", username, activitytype)

	resp := &pb.Done{
		Done: true,
	}

	return resp, nil
}

func (a *ActivityService) ActivityIsValid(ctx context.Context, req *pb.ActivityRequest) (*pb.Valid, error) {
	username := req.Username
	activitytype := req.Type

	fmt.Println("ActivityIsValid() is working:", username, activitytype)

	resp := &pb.Valid{
		Valid: true,
	}

	return resp, nil
}
