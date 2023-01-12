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

	//_________________________________________________________________________
	//adding user details

	addUserDetails := &pb.User{

		Name:  "saiteja",
		Phone: "1456123789",
		Email: "teja@yahoo.com",
	}

	addUser, err := client.RegisterUser(context.Background(), addUserDetails)
	if err != nil {

		fmt.Println("Error in updating  users using grpc client", err.Error())

	}
	fmt.Println(addUser)

	//_________________________________________________________________________
	//getting user

	name := &pb.Name{
		Name: "saiteja",
	}
	userDetails, err := client.GetUser(context.Background(), name)

	if err != nil {

		fmt.Println("Error in calling users using grpc client", err.Error())

	}
	fmt.Println("user details are:", userDetails)

	//____________________________________________________________________
	//update user

	details := &pb.UpdateUser{
		User: &pb.User{
			Name:  "saiteja",
			Phone: "8639218758",
			Email: "teja@gmail.com",
		},
	}

	updateUser, err := client.UpdateUserInfo(context.Background(), details)

	if err != nil {

		fmt.Println("Error in updating  users using grpc client", err.Error())

	}
	fmt.Println(updateUser)

	//_____________________________________________________________________________________
	// add activity
	//startTime, _ := ptypes.TimestampProto(time.Now())

	ts := timestamppb.Now()

	userActivity := &pb.Activity{

		Type:     "Eat",
		Label:    "saiteja",
		Time:     ts,
		Duration: 6,
	}

	addActivity, err := client.AddActivity(context.Background(), userActivity)
	if err != nil {

		fmt.Println("Error in adding activity", err.Error())

	}
	fmt.Println(addActivity)

	//______________________________________________________________________________
	// activity is done
	activityDoneRequest := &pb.ActivityRequest{

		Username: "saiteja",
		Type:     "Eat",
	}

	checkIsDone, err := client.ActivityIsDone(context.Background(), activityDoneRequest)
	if err != nil {

		fmt.Println("Error in isDone()", err.Error())

	}
	fmt.Println("Activity status(isDone)", checkIsDone)

	//______________________________________________________________________________
	// activity is valid

	ActivityRequest := &pb.ActivityRequest{

		Username: "saiteja",
		Type:     "Eat",
	}

	checkIsValid, err := client.ActivityIsValid(context.Background(), ActivityRequest)
	if err != nil {

		fmt.Println("Error in isValid()", err.Error())

	}
	fmt.Println("Activity status(isDone)", checkIsValid)
}
