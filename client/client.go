package main

import (
	"context"
	"fmt"
	pb "task1/proto"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	grpcServer := "localhost:8082"

	conn, err := grpc.Dial(grpcServer, grpc.WithInsecure())
	if err != nil {
		fmt.Println("error connecting grpc: ", err)
	}
	defer conn.Close()

	client := pb.NewActivityServiceClient(conn)

	addUserDetails := &pb.User{

		Name:  "SaiTeja",
		Phone: "1456123789",
		Email: "teja@yahoo.com",
	}

	addUser, err := client.RegisterUser(context.Background(), addUserDetails)
	if err != nil {

		fmt.Println("Error in updating  users using grpc client", err.Error())

	}
	fmt.Println(addUser)

	name := &pb.Name{
		Name: "Saiteja",
	}
	userDetails, err := client.GetUser(context.Background(), name)

	if err != nil {

		fmt.Println("Error in calling users using grpc client", err.Error())

	}
	fmt.Println("user details are:", userDetails)

	details := &pb.UpdateUser{
		User: &pb.User{
			Name:  "SaiTeja",
			Phone: "8639218758",
			Email: "teja@gmail.com",
		},
	}

	updateUser, err := client.UpdateUserInfo(context.Background(), details)

	if err != nil {

		fmt.Println("Error in updating  users using grpc client", err.Error())

	}
	fmt.Println(updateUser)

	//startTime, _ := ptypes.TimestampProto(time.Now())
	ts := timestamppb.Now()

	userActivity := &pb.Activity{

		Type:     "Eat",
		Label:    "Eating",
		Time:     ts,
		Duration: 6,
	}

	addActivity, err := client.AddActivity(context.Background(), userActivity)
	if err != nil {

		fmt.Println("Error in updating  users using grpc client", err.Error())

	}
	fmt.Println(addActivity)

	activityDone := &pb.ActivityRequest{

		Username: "SaiTeja",
		Type:     "Eat",
	}

	checkIsDone, err := client.ActivityIsDone(context.Background(), activityDone)
	if err != nil {

		fmt.Println("Error in updating  users using grpc client", err.Error())

	}
	fmt.Println("Activity status(isDone)", checkIsDone)

	activityValid := &pb.ActivityRequest{

		Username: "SaiTeja",
		Type:     "Eat",
	}

	checkIsValid, err := client.ActivityIsValid(context.Background(), activityValid)
	if err != nil {

		fmt.Println("Error in updating  users using grpc client", err.Error())

	}
	fmt.Println("Activity status(isDone)", checkIsValid)
}
