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
)

type User struct {
	Name       string `bson:"name,omitempty"`
	Email      string `bson:"email,omitempty"`
	Phone      string `bson:"phone,omitempty"`
	Activities []Act  `bson:"activities,omitempty"`
}

type Act struct {
	Type     string
	Time     time.Time
	Duration int32
	Label    string
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
	ts := time.Now()

	act1 := Act{

		Type:     "Eat",
		Label:    "Eating",
		Time:     ts,
		Duration: 6,
	}
	act2 := Act{
		Type:     "Sleep",
		Label:    "Sleeping",
		Time:     ts,
		Duration: 78,
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
	fmt.Println(finduser.Name)
	fmt.Println(len(finduser.Activities))
	// fmt.Println(finduser.Email)
	// fmt.Println(finduser.Phone)
	fmt.Println(a[0])
	fmt.Println(a[1])

	act3 := Act{
		Type:     "Play",
		Label:    "Playing",
		Time:     ts,
		Duration: 62,
	}
	activities = append(activities, act3)
	Key := "name"
	username := "saiteja"
	filter1 := bson.M{Key: username}
	update := bson.M{"$set": bson.M{"activities": activities}}

	res, err := coll.UpdateOne(context.TODO(), filter1, update)
	if err != nil {
		panic(err)
	}
	fmt.Print(res)

}
