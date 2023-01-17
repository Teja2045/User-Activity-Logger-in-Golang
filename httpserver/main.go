package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

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

type Userj struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
type Actj struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Duration int    `json:"duration"`
}

type Actreqj struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Actresj struct {
	Response string `json:"response"`
}

type Getj struct {
	Name string `json:"name"`
}

type Act struct {
	Type     string
	Time     time.Time
	Duration int32
	Label    string
}

func main() {
	//Init Router
	r := mux.NewRouter()

	// arrange our route
	r.HandleFunc("/adduser", RegisterUser).Methods("POST")
	r.HandleFunc("/getuser", GetUser).Methods("GET")
	r.HandleFunc("/update", UpdateUserInfo).Methods("PUT")
	r.HandleFunc("/addactivity", AddActivity).Methods("POST")
	r.HandleFunc("/activityisdone", ActivityIsDone).Methods("GET")
	r.HandleFunc("/activityisvalid", ActivityIsValid).Methods("GET")

	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))

}

func RegisterUser(w http.ResponseWriter, req *http.Request) {
	var p Userj
	// use json.NewDecoder to read the JSON payload from the request body
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	username := p.Name
	phone := p.Phone
	email := p.Email
	activities := []Act{}
	user := User{

		Name:       username,
		Email:      email,
		Phone:      phone,
		Activities: activities,
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

	json.NewEncoder(w).Encode(result)

}

// update user service
// input : user
// functionality : update user in database
// output : string
func UpdateUserInfo(w http.ResponseWriter, req *http.Request) {
	var p Userj
	// use json.NewDecoder to read the JSON payload from the request body
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	username := p.Name
	phone := p.Phone
	email := p.Email

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

	update = bson.M{"$set": bson.M{"phone": phone}}

	result1, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	fmt.Println("user has been updated:", result, result1)

}

// get user service
// input : name
// functionality : gets user details
// output : user
// url = http://localhost:8080/name=saiteja?name=John
func GetUser(w http.ResponseWriter, req *http.Request) {
	var p Getj
	// use json.NewDecoder to read the JSON payload from the request body
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	username := p.Name

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

	filter := bson.M{"name": username}
	user := &User{}
	err = coll.FindOne(context.TODO(), filter).Decode(user)
	if err != nil {
		log.Fatal("find...", err)
	}
	fmt.Println(user)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)

}

//____________________________________________________________________________

// add activity service
// input : acivity
// functionality : add to database
// output : string
func AddActivity(w http.ResponseWriter, req *http.Request) {
	var actj Actj
	// use json.NewDecoder to read the JSON payload from the request body
	err := json.NewDecoder(req.Body).Decode(&actj)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	activityType := actj.Type
	// here I am using 'label' as key for user.. so label is user name
	name := actj.Name
	label := actj.Type + "ing"
	duration := int32(actj.Duration)
	time1 := time.Now()

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

	act := Act{
		Type:     activityType,
		Duration: duration,
		Time:     time1,
		Label:    label,
	}

	filter := bson.D{{Key: "name", Value: name}}

	var finduser User
	err = coll.FindOne(context.TODO(), filter).Decode(&finduser)
	if err != nil {
		log.Fatal(err)
	}
	activities := finduser.Activities
	activities = append(activities, act)
	//fmt.Println(result)
	Key := "name"
	username := name
	filter1 := bson.M{Key: username}
	update := bson.M{"$set": bson.M{"activities": activities}}

	res, err := coll.UpdateOne(context.TODO(), filter1, update)
	if err != nil {
		panic(err)
	}

	fmt.Println("activity has been added:", res)

}

// activty isDone serive
// input : username, activity
// functionality : checks if that user had done that activity
// output : boolean
func ActivityIsDone(w http.ResponseWriter, req *http.Request) {
	var actreqj Actreqj
	// use json.NewDecoder to read the JSON payload from the request body
	err := json.NewDecoder(req.Body).Decode(&actreqj)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	username := actreqj.Name

	activitytype := actreqj.Type
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

	var finduser User
	err = coll.FindOne(context.TODO(), filter).Decode(&finduser)
	if err != nil {
		log.Fatal(err)
	}
	activities := finduser.Activities
	flag := false
	fmt.Println("ActivityIsDone() is working:", username, activitytype)
	for i := 0; i < len(activities); i++ {
		act := activities[i]
		if act.Type == activitytype {
			flag = true
			break
		}
	}
	if flag {
		rsp := Actresj{Response: "Yes"}

		// Marshal the struct into a JSON string
		//resp, _ := json.Marshal(rsp)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(rsp)
	} else {
		rsp := Actresj{Response: "No"}

		// Marshal the struct into a JSON string
		//resp, _ := json.Marshal(rsp)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(rsp)
	}
}

// activty isValid service
// input : username, activity
// functionality : checks if the activity done by user is valid
// output : boolean
func ActivityIsValid(w http.ResponseWriter, req *http.Request) {
	var actreqj Actreqj
	// use json.NewDecoder to read the JSON payload from the request body
	err := json.NewDecoder(req.Body).Decode(&actreqj)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	username := actreqj.Name

	activitytype := actreqj.Type

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

	var finduser User
	err = coll.FindOne(context.TODO(), filter).Decode(&finduser)
	if err != nil {
		log.Fatal(err)
	}
	flag := false
	activities := finduser.Activities
	fmt.Println("ActivityIsValid() is working:", activities[0].Type)
	for i := 0; i < len(activities); i++ {
		act := activities[i]
		if act.Type == activitytype {
			if act.Duration >= 6 {
				flag = true
			}
		}
	}
	if flag {
		rsp := Actresj{Response: "Yes"}

		// Marshal the struct into a JSON string
		//resp, _ := json.Marshal(rsp)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(rsp)
	} else {
		rsp := Actresj{Response: "No"}

		// Marshal the struct into a JSON string
		//resp, _ := json.Marshal(rsp)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(rsp)
	}
}
