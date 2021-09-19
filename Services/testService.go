package Services

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin_mongodb/config"
	"gin_mongodb/util"
	"log"
	"testing"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Test struct {
	Id      int    `json:"id"`
	Testcol string `json:"testcol"`
}

var URL = "127.0.0.1:27017"

func init() {
	env, err := config.Get()
	if err != nil {
		log.Fatal("env read error : ", err)
	}
	URL = env.MONGODB_URL
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

// C create/insert
func TestInsert() {
	insert(p, col)
	pb := findByName(p.Name, col)

	fmt.Println("Insert Result")
	fmt.Println(pb)
}

func Log(live_id string, insert_data string, msg_type string) (data interface{}, err error) {

	fmt.Println(live_id)
	fmt.Println(msg_type)
	if err != nil {
		return err, errors.New("文件下载失败")
	}
	session, err := mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB(col.DB).C(live_id)
	if msg_type == "WebcastChatMessage" {
		fmt.Println("等于")
		insert_obj := util.WebcastMemberMessage{}
		err = json.Unmarshal([]byte(insert_data), &insert_obj)
		if err != nil {
			fmt.Println(err)
			return "插入失败。解码失败", nil
		}
		err = c.Insert(insert_obj)
		if err != nil {
			log.Fatal(err)
			return "插入失败。mongoDb插入失败", nil

		}
		return "插入成功", nil
	} else if msg_type == "WebcastMemberMessage" {
		insert_obj := util.WebcastMemberMessage{}
		err = json.Unmarshal([]byte(insert_data), &insert_obj)
		if err != nil {
			fmt.Println(err)
			return "插入失败。解码失败", nil
		}
		err = c.Insert(insert_obj)
		if err != nil {
			log.Fatal(err)
			return "插入失败。mongoDb插入失败", nil

		}
		return "插入成功", nil
	} else if msg_type == "WebcastSocialMessage" {
		insert_obj := util.WebcastSocialMessage{}
		err = json.Unmarshal([]byte(insert_data), &insert_obj)
		if err != nil {
			fmt.Println(err)
			return "插入失败。解码失败", nil
		}
		err = c.Insert(insert_obj)
		if err != nil {
			log.Fatal(err)
			return "插入失败。mongoDb插入失败", nil

		}
		return "插入成功", nil
	} else if msg_type == "WebcastLikeMessage" {
		insert_obj := util.WebcastLikeMessage{}
		err = json.Unmarshal([]byte(insert_data), &insert_obj)
		if err != nil {
			fmt.Println(err)
			return "插入失败。解码失败", nil
		}
		err = c.Insert(insert_obj)
		if err != nil {
			log.Fatal(err)
			return "插入失败。mongoDb插入失败", nil

		}
		return "插入成功", nil
	} else if msg_type == "WebcastRoomUserSeqMessage" {
		insert_obj := util.WebcastRoomUserSeqMessage{}
		err = json.Unmarshal([]byte(insert_data), &insert_obj)
		if err != nil {
			fmt.Println(err)
			return "插入失败。解码失败", nil
		}
		err = c.Insert(insert_obj)
		if err != nil {
			log.Fatal(err)
			return "插入失败。mongoDb插入失败", nil

		}
		return "插入成功", nil
	} else if msg_type == "WebcastGiftMessage" {
		insert_obj := util.WebcastGiftMessage{}
		err = json.Unmarshal([]byte(insert_data), &insert_obj)
		if err != nil {
			fmt.Println(err)
			return "插入失败。解码失败", nil
		}
		err = c.Insert(insert_obj)
		if err != nil {
			log.Fatal(err)
			return "插入失败。mongoDb插入失败", nil

		}
		return "插入成功", nil
	} else if msg_type == "WebcastRoomIntroMessage" {
		insert_obj := util.WebcastRoomIntroMessage{}
		err = json.Unmarshal([]byte(insert_data), &insert_obj)
		if err != nil {
			fmt.Println(err)
			return "插入失败。解码失败", nil
		}
		err = c.Insert(insert_obj)
		if err != nil {
			log.Fatal(err)
			return "插入失败。mongoDb插入失败", nil

		}
		return "插入成功", nil
	} else {
		return "插入失败，没有找到这个类型", nil
	}
}

func detail_insert() {

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
