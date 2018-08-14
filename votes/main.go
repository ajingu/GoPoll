package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	Age  int           `bson:"age"`
}

func main() {
	session, _ := mgo.Dial("localhost")
	defer session.Close()
	db := session.DB("test")

	alice := &Person{
		ID:   bson.NewObjectId(),
		Name: "Alice",
		Age:  20,
	}

	col := db.C("people")
	if err := col.Insert(alice); err != nil {
		log.Fatalln(err)
	}

	p := new(Person)
	query := db.C("people").Find(bson.M{})
	query.One(&p)

	fmt.Printf("%+v\n", p)
}
