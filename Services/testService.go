package Services

import (
	"fmt"
	"testing"
    "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Test struct {
    Id int `json:"id"` 
    Testcol string `json:"testcol"`
}

func (this *Test) Insert() (id int, err error) {
    insert
    return 1,nil
}
const (
	URL = "127.0.0.1:27017"
)

func insert(p Person, col Collection) {
	session, err := mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB(col.DB).C(col.Name)
	err = c.Insert(p)
	if err != nil {
		log.Fatal(err)
	}
}

func findByName(name string, col Collection) Person {
	session, err := mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(col.DB).C(col.Name)
	result := Person{}
	err = collection.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}


func update(p Person, col Collection) {
	session, err := mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	collection := session.DB(col.DB).C(col.Name)
	err = collection.Update(bson.M{"name": p.Name}, bson.M{"$set": bson.M{"phone": p.Phone}})
	if err != nil {
		log.Fatal(err)
	}
}

func deleteData(p Person, col Collection) {
	session, err := mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	collection := session.DB(col.DB).C(col.Name)
	err = collection.Remove(bson.M{"name": p.Name})
}

type Person struct {
	Name  string
	Phone string
}
type Collection struct {
	DB   string
	Name string
}

var col = Collection{"testDB", "contacts"}
var p = Person{"Jack_Green", "9987650"}

// C create/insert
func TestInsert(t *testing.T) {
	insert(p, col)
	pb := findByName(p.Name, col)
	if pb.Name != p.Name || pb.Phone != p.Phone {
		t.Error("insert failed")
	}
	fmt.Println("Insert Result")
	fmt.Println(pb)
}

// R read/find
func TestFindByName(t *testing.T) {
	p := findByName(p.Name, col)
	if p.Name == "" || p.Phone == "" {
		t.Error("find by name test failed")
	}
	fmt.Println("Find Result")
	fmt.Println(p)

}

// U update
func TestUpdate(t *testing.T) {
	p := Person{"Jack_Green", "121212122"}
	update(p, col)
	ub := findByName(p.Name, col)
	fmt.Println("Update Result")
	fmt.Println(ub)
}

// D delete
func TestDeleteData(t *testing.T) {
	deleteData(p, col)
}
