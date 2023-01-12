package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	Name       string `bson:"name,omitempty"`
	Email      string `bson:"email,omitempty"`
	Phone      string `bson:"phone,omitempty"`
	Activities []Act  `bson:"activities,omitempty"`
}

type Act struct {
	typ      string
	time     timestamppb.Timestamp
	duration int
	label    string
}

func main() {
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
	ts := timestamppb.Now()

	act1 := Act{

		typ:      "Eat",
		label:    "Eating",
		time:     *ts,
		duration: 6,
	}
	act2 := Act{
		typ:      "Sleep",
		label:    "Sleeping",
		time:     *ts,
		duration: 78,
	}

	activities := []Act{
		act1, act2,
	}
	user := User{Name: "saiteja", Email: "teja123gmail.com", Phone: "8622212341", Activities: activities}
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)
	filter := bson.D{{Key: "name", Value: "saiteja"}}

	var finduser User
	err = coll.FindOne(context.TODO(), filter).Decode(&finduser)
	if err != nil {
		log.Fatal(err)
	}

	a := finduser.Activities
	fmt.Println(finduser.Activities)
	// fmt.Println(finduser.Email)
	// fmt.Println(finduser.Phone)
	fmt.Println(a[0].typ)
	fmt.Println(a[1].typ)

}
