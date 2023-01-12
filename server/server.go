package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "task1/proto"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserDB struct {
	Name  string `bson:"name,omitempty"`
	Email string `bson:"email,omitempty"`
	Phone string `bson:"phone,omitempty"`
}

type Record struct {
	Name     string `bson:"name,omitempty"`
	Type     string `bson:"type,omitempty"`
	Duration int
	Time     timestamppb.Timestamp
}

const grpcPort = ":8082"

type ActivityService struct {
	pb.UnimplementedActivityServiceServer
}

func main() {

	log.Println("Starting application...")

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

// simplest way to handle errors
func HandleError(err error) {
	if err != nil {
		log.Fatal("error: ", err)
	}
}

// add user service
// input : user
// functionality : add to database
// output : string
func (a *ActivityService) RegisterUser(ctx context.Context, req *pb.User) (*pb.UserResponse, error) {
	username := req.Name
	phone := req.Phone
	email := req.Email

	user := UserDB{

		Name:  username,
		Email: email,
		Phone: phone,
	}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://saiteja:saiteja@cluster0.ugdvlxb.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	dbctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(dbctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(dbctx)

	err = client.Ping(dbctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	coll := client.Database("task1").Collection("user")

	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	resp := &pb.UserResponse{
		Response: "The user " + username + " has been added.",
	}

	return resp, nil
}

// update user service
// input : user
// functionality : update user in database
// output : string
func (a *ActivityService) UpdateUserInfo(ctx context.Context, req *pb.UpdateUser) (*pb.UserResponse, error) {
	username := req.User.Name
	//phone := req.User.Phone
	email := req.User.Email

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://saiteja:saiteja@cluster0.ugdvlxb.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	dbctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	err = client.Connect(dbctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(dbctx)

	err = client.Ping(dbctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	coll := client.Database("task1").Collection("user")
	Key := "name"
	filter := bson.M{Key: username}
	update := bson.M{"$set": bson.M{"email": email}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	fmt.Println("user has been updated:", result)

	resp := &pb.UserResponse{
		Response: "The user " + username + " has been updated.",
	}

	return resp, nil
}

// get user service
// input : name
// functionality : gets user details
// output : user
func (a *ActivityService) GetUser(ctx context.Context, req *pb.Name) (*pb.UserResponse, error) {
	username := req.Name

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://saiteja:saiteja@cluster0.ugdvlxb.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	dbctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err = client.Connect(dbctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(dbctx)

	err = client.Ping(dbctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	coll := client.Database("task1").Collection("user")

	filter := bson.D{{Key: "name", Value: username}}
	user := &UserDB{}
	err = coll.FindOne(context.TODO(), filter).Decode(user)
	if err != nil {
		log.Println("problem here")
		log.Fatal("find...", err)
	}
	fmt.Println(user)

	resp := &pb.UserResponse{
		Response: "The user " + username + " details were given.",
	}

	return resp, nil
}

//____________________________________________________________________________

// add activity service
// input : acivity
// functionality : add to database
// output : string
func (a *ActivityService) AddActivity(ctx context.Context, req *pb.Activity) (*pb.UserResponse, error) {
	activityType := req.Type
	// here I am using 'label' as key for user.. so label is user name
	label := req.Label
	duration := req.Duration

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://saiteja:saiteja@cluster0.ugdvlxb.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	dbctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err = client.Connect(dbctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(dbctx)

	err = client.Ping(dbctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	coll := client.Database("task1").Collection("record")

	record := Record{
		Name:     label,
		Type:     activityType,
		Duration: int(duration),
		Time:     *req.Time,
	}

	result, err := coll.InsertOne(context.TODO(), record)
	if err != nil {
		panic(err)
	}
	//fmt.Println(result)

	fmt.Println("activity has been added:", result)

	resp := &pb.UserResponse{
		Response: "The activity " + activityType + " is added.",
	}

	return resp, nil
}

// activty isDone serive
// input : username, activity
// functionality : checks if that user had done that activity
// output : boolean
func (a *ActivityService) ActivityIsDone(ctx context.Context, req *pb.ActivityRequest) (*pb.Done, error) {
	username := req.Username
	activitytype := req.Type
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://saiteja:saiteja@cluster0.ugdvlxb.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	dbctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err = client.Connect(dbctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(dbctx)

	err = client.Ping(dbctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	coll := client.Database("task1").Collection("record")

	filter := bson.D{{Key: "name", Value: username}} //, {Key: "type", Value: activitytype}}
	resp := &pb.Done{
		Done: false,
	}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {

		log.Fatal("find...", err)
	}
	fmt.Println("ActivityIsDone() is working:", username, activitytype)
	var records []Record
	if err = cursor.All(ctx, &records); err != nil {
		if err == mongo.ErrNoDocuments {
			return resp, nil
		}
	}
	fmt.Println(len(records))

	resp = &pb.Done{
		Done: true,
	}
	return resp, nil
}

// activty isValid serive
// input : username, activity
// functionality : checks if the activity done by user is valid
// output : boolean
func (a *ActivityService) ActivityIsValid(ctx context.Context, req *pb.ActivityRequest) (*pb.Valid, error) {
	username := req.Username
	activitytype := req.Type

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://saiteja:saiteja@cluster0.ugdvlxb.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	dbctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err = client.Connect(dbctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(dbctx)

	err = client.Ping(dbctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	coll := client.Database("task1").Collection("record")

	filter := bson.M{"name": username, "type": activitytype}
	record := &Record{}
	err = coll.FindOne(context.TODO(), filter).Decode(record)
	if err != nil {
		log.Println("problem here")
		log.Fatal("find...", err)
	}
	fmt.Println(record.Duration)

	resp := &pb.Valid{
		Valid: false,
	}
	if record.Duration > 6 {
		resp = &pb.Valid{
			Valid: true,
		}
	}

	return resp, nil
}
