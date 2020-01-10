package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const (
	DB_NAME       = "fangkm_test"
	MONGODB_HOST  = "192.168.96.42:27017"
	DB_COLLECTION = "student"
)

type Student struct {
	Name   string `bson: "name"`
	Age    int    `bson: "age"`
	Sid    string `bson: "sid"`
	Status int    `bson: "status"`
}

type Per struct {
	Students []Student
}

func main() {
	session, err := mgo.Dial(MONGODB_HOST)
	if err != nil {
		log.Fatalf("Counld not connect to host:%s %v", MONGODB_HOST, err)
	}
	defer session.Close()

	//AddData(session)
	//Search(session)
	//SearchALl(session)
	UpdateAge(session)
}

func AddData(session *mgo.Session) {
	newSession := session.Clone()
	defer newSession.Close()
	connection := newSession.DB(DB_NAME).C(DB_COLLECTION)

	data := &Student{
		//Name:   "fangkm",
		//Age:    18,
		//Sid:    "s20180907",
		//Status: 1,
		Name:   "seeta1",
		Age:    18,
		Sid:    "s20180908",
		Status: 1,
	}

	err := connection.Insert(data)
	if err != nil {
		log.Println(err)
	}
}

func Search(session *mgo.Session) {
	newSession := session.Clone()
	defer newSession.Close()
	connection := newSession.DB(DB_NAME).C(DB_COLLECTION)

	user := Student{}
	err := connection.Find(bson.M{"sid": "s20180907"}).One(&user)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%+v\n", user)
}

func SearchALl(session *mgo.Session) {
	newSession := session.Clone()
	defer newSession.Close()
	connection := newSession.DB(DB_NAME).C(DB_COLLECTION)

	var stu Student
	var stus Per

	iter := connection.Find(bson.M{"status": 1}).Sort("_id").Iter()

	for iter.Next(&stu) {
		stus.Students = append(stus.Students, stu)
	}

	if err := iter.Close(); err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%+v", stus)
}

func UpdateAge(session *mgo.Session) {
	newSession := session.Clone()
	defer newSession.Close()
	connection := newSession.DB(DB_NAME).C(DB_COLLECTION)

	err := connection.Update(bson.M{"status": 1}, bson.M{"$set": bson.M{"age": 20}})
	if err !=nil{
		log.Printf("%+v", err)
		return

	}
}
