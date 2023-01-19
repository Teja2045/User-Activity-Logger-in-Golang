package mongo_test

import (
	"context"
	"fmt"
	"testing"

	pb "task1/proto"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestRegisterUser(t *testing.T) {

	grpcServer := "localhost:8082"

	conn, err := grpc.Dial(grpcServer, grpc.WithInsecure())
	if err != nil {
		fmt.Println("error connecting grpc: ", err)
	}
	defer conn.Close()

	client := pb.NewActivityServiceClient(conn)

	addUserDetails := &pb.User{

		Name:  "tester",
		Phone: "1456123789",
		Email: "tester@yahoo.com",
	}

	addUser, err := client.RegisterUser(context.Background(), addUserDetails)
	if err != nil {

		fmt.Println("Error in updating  users using grpc client", err.Error())

	}
	assert.NoError(t, err)
	assert.NotNil(t, addUser)
}

func TestUpdateUserInfo(t *testing.T) {

	grpcServer := "localhost:8082"

	conn, err := grpc.Dial(grpcServer, grpc.WithInsecure())
	if err != nil {
		fmt.Println("error connecting grpc: ", err)
	}
	defer conn.Close()

	client := pb.NewActivityServiceClient(conn)

	details := &pb.UpdateUser{
		User: &pb.User{
			Name:  "tester",
			Phone: "09876543",
			Email: "tester@gmail.com",
		},
	}

	updateUser, err := client.UpdateUserInfo(context.Background(), details)

	if err != nil {

		fmt.Println("Error in updating  users using grpc client", err.Error())

	}
	assert.NoError(t, err)
	assert.NotNil(t, updateUser)
}

func TestGetUser(t *testing.T) {

	grpcServer := "localhost:8082"

	conn, err := grpc.Dial(grpcServer, grpc.WithInsecure())
	if err != nil {
		fmt.Println("error connecting grpc: ", err)
	}
	defer conn.Close()

	client := pb.NewActivityServiceClient(conn)

	name := &pb.Name{
		Name: "tester",
	}
	userDetails, err := client.GetUser(context.Background(), name)

	if err != nil {

		fmt.Println("Error in calling users using grpc client", err.Error())

	}
	assert.NoError(t, err)
	assert.NotNil(t, userDetails)

	inputuser := &pb.User{
		Name:  "tester",
		Phone: "09876543",
		Email: "tester@gmail.com",
	}

	outputuser := &pb.User{
		Name:  userDetails.Name,
		Phone: userDetails.Phone,
		Email: userDetails.Email,
	}
	assert.Equal(t, inputuser.Email, outputuser.Email)
	assert.Equal(t, inputuser.Phone, outputuser.Phone)
	// name = &pb.Name{
	// 	Name: "tester1",
	// }
	// userDetails, err = client.GetUser(context.Background(), name)

	// if err != nil {

	// 	fmt.Println("Error in calling users using grpc client", err.Error())

	// }
	// assert.NoError(t, err)
	// assert.NotNil(t, userDetails)

}

func TestAddActivity(t *testing.T) {

	grpcServer := "localhost:8082"

	conn, err := grpc.Dial(grpcServer, grpc.WithInsecure())
	if err != nil {
		fmt.Println("error connecting grpc: ", err)
	}
	defer conn.Close()

	client := pb.NewActivityServiceClient(conn)

	ts := timestamppb.Now()

	userActivity := &pb.Activity{

		Type:     "eat",
		Label:    "tester",
		Time:     ts,
		Duration: 20,
	}

	addActivity, err := client.AddActivity(context.Background(), userActivity)
	if err != nil {

		fmt.Println("Error in adding activity", err.Error())

	}
	assert.NoError(t, err)
	assert.NotNil(t, addActivity)
}

func TestActivityIsDone(t *testing.T) {

	grpcServer := "localhost:8082"

	conn, err := grpc.Dial(grpcServer, grpc.WithInsecure())
	if err != nil {
		fmt.Println("error connecting grpc: ", err)
	}
	defer conn.Close()

	client := pb.NewActivityServiceClient(conn)

	activityDoneRequest := &pb.ActivityRequest{

		Username: "tester",
		Type:     "eat",
	}

	checkIsDone, err := client.ActivityIsDone(context.Background(), activityDoneRequest)
	if err != nil {

		fmt.Println("Error in isDone()", err.Error())

	}
	assert.NoError(t, err)
	assert.NotNil(t, checkIsDone)
	assert.Equal(t, checkIsDone.Done, true)
}

func TestActivityIsValid(t *testing.T) {

	grpcServer := "localhost:8082"

	conn, err := grpc.Dial(grpcServer, grpc.WithInsecure())
	if err != nil {
		fmt.Println("error connecting grpc: ", err)
	}
	defer conn.Close()

	client := pb.NewActivityServiceClient(conn)

	ActivityRequest := &pb.ActivityRequest{

		Username: "saiteja",
		Type:     "eat",
	}

	checkIsValid, err := client.ActivityIsValid(context.Background(), ActivityRequest)
	if err != nil {

		fmt.Println("Error in isValid()", err.Error())

	}
	assert.NoError(t, err)
	assert.NotNil(t, checkIsValid)
	assert.Equal(t, checkIsValid.Valid, true)
}
