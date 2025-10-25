package main 

import (
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"os"
	"chapter-C.30-golang-protobuf-implementation/model"
)

var user1 = &model.User{
	Id: "u001",
	Name: "Sylvana Windrunner",
	Password: "archer123",
	Gender: model.UserGender_FEMALE,
}

var userList = &model.UserList{
	List: []*model.User{
		user1,
	},
}

var garage1 = &model.Garage{
	Id:   "g001",
	Name: "Kalimdor",
	Coordinate: &model.GarageCoordinate{
		Latitude: 23.2212847,
		Longitude: 53.22033123,
	},
}

var garageList = &model.GarageList{
	List: []*model.Garage{
		garage1,
	},
}

var garageListByUser = &model.GarageListBuyer{
	List: map[string]*model.GarageList{
		user1.Id: garageList,
	},
}

func main() {
// =========== original
fmt.Printf("# ==== Original\n %#v \n", user1)
// =========== as string
fmt.Printf("# ==== As String\n %s \n", user1.String())

// =========== as json string
jsonb, err1 := protojson.Marshal(garageList)
if err1 != nil {
	fmt.Println("Failed to marshal to json:", err1.Error())
	os.Exit(0)
}
fmt.Printf("# ==== As JSON String\n %s \n", string(jsonb))

protoObject := &model.GarageList{}
err2 := protojson.Unmarshal(jsonb, protoObject)
if err2 != nil {
	fmt.Println("Failed to unmarshal from json:", err2.Error())
	os.Exit(0)
}
fmt.Printf("# ==== From JSON String to Proto Object\n %#v \n", protoObject.String())

}
