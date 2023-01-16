package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	pb "task1/proto"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//Note: short description tells the inputs for the commands, Long description, about the command functionality
func main() {
	grpcServer := "localhost:8082"

	conn, err := grpc.Dial(grpcServer, grpc.WithInsecure())
	if err != nil {
		fmt.Println("error connecting grpc: ", err)
	}
	defer conn.Close()

	client := pb.NewActivityServiceClient(conn)

	rootCmd := &cobra.Command{
		Use:   "client",
		Short: "grpc-client",
		Long:  `grpc-client`,
	}

	registerCmd := &cobra.Command{
		Use:   "register",
		Short: "| NAME | PHONE | EMAIL |",
		Long:  "takes 3 inputs which are user details and registers a user",
		Run: func(cmd *cobra.Command, args []string) {
			details := &pb.User{
				Name:  args[0],
				Phone: "" + args[1],
				Email: args[2],
			}
			addUser, err := client.RegisterUser(context.Background(), details)
			if err != nil {

				fmt.Println("Error in registering user using grpc client", err.Error())

			}
			fmt.Println(addUser)
		},
	}

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "| NAME | PHONE | EMAIL |",
		Long:  "takes 3 inputs, user name and details and updates the database",
		Run: func(cmd *cobra.Command, args []string) {
			details := &pb.UpdateUser{
				User: &pb.User{
					Name:  args[0],
					Phone: args[1],
					Email: args[2],
				},
			}
			updateUser, err := client.UpdateUserInfo(context.Background(), details)

			if err != nil {

				fmt.Println("Error in updating user using grpc client", err.Error())

			}
			fmt.Println(updateUser)
		},
	}

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "| NAME |",
		Long:  "takes 1 input which is the user name and retrieves his/her details",
		Run: func(cmd *cobra.Command, args []string) {
			name := &pb.Name{
				Name: args[0],
			}
			userDetails, err := client.GetUser(context.Background(), name)

			if err != nil {

				fmt.Println("Error in calling users using grpc client", err.Error())

			}
			fmt.Println("user details are:", userDetails)
		},
	}

	add_activityCmd := &cobra.Command{
		Use:   "add_activity",
		Short: "| USERNAME | ACTIVITY TYPE | ACTIVITY DURATION |",
		Long:  "takes 3 inputs, username and activity details",
		Run: func(cmd *cobra.Command, args []string) {
			ts := timestamppb.Now()
			i, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				panic(err)
			}

			userActivity := &pb.Activity{

				Label:    args[0],
				Type:     args[1],
				Time:     ts,
				Duration: int32(i),
			}

			addActivity, err := client.AddActivity(context.Background(), userActivity)
			if err != nil {

				fmt.Println("Error in adding activity", err.Error())

			}
			fmt.Println(addActivity)
		},
	}

	activity_isDoneCmd := &cobra.Command{
		Use:   "activity_isDone",
		Short: "| USERNAME | ACTIVITY TYPE |",
		Long:  "takes 2 inputs, username, activity type and checks if the activity is done by the user",
		Run: func(cmd *cobra.Command, args []string) {
			activityDoneRequest := &pb.ActivityRequest{

				Username: args[0],
				Type:     args[1],
			}

			checkIsDone, err := client.ActivityIsDone(context.Background(), activityDoneRequest)
			if err != nil {

				fmt.Println("Error in isDone()", err.Error())

			}
			if checkIsDone.Done {
				fmt.Println("Activity status(isDone)", "Yes")
			} else {
				fmt.Println("Activity status(isDone)", "No")
			}

		},
	}

	activity_isValid := &cobra.Command{
		Use:   "activity_isValid",
		Short: "| USERNAME | ACTIVITY TYPE |",
		Long:  "takes 2 inputs, username, activity type and checks if the activity by the user is valid or not",
		Run: func(cmd *cobra.Command, args []string) {
			ActivityRequest := &pb.ActivityRequest{

				Username: args[0],
				Type:     args[1],
			}

			checkIsValid, err := client.ActivityIsValid(context.Background(), ActivityRequest)
			if err != nil {

				fmt.Println("Error in isValid()", err.Error())

			}
			if checkIsValid.Valid {
				fmt.Println("Activity is valid(isValid)", "Yes")
			} else {
				fmt.Println("Activity is valid(isValid)", "No")
			}

		},
	}

	rootCmd.AddCommand(registerCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(add_activityCmd)
	rootCmd.AddCommand(activity_isDoneCmd)
	rootCmd.AddCommand(activity_isValid)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

//_________________________________________________________________________
