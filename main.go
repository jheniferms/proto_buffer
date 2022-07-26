package main

import (
	"fmt"
	"reflect"

	pb "github.com/jheniferms/proto_buffer/proto"
)

func main() {
	fmt.Println("\nDo Simples\n", doSimple(), "\n\nDo Complex")
	fmt.Println(doComplex(), "\n")

	doOneof(&pb.Result_Id{Id: 15})
	doOneof(&pb.Result_Message{Message: "jheni"})

	fmt.Println(doMap())

	writeToFile("teste.bin", doSimple())

	fmt.Println(writeToJson(doSimple()))

	jsonString := writeToJson(doSimple())
	fmt.Println(doFromJson(jsonString, reflect.TypeOf(pb.Simple{})))

	jsonString = writeToJson(doComplex())
	fmt.Println(doFromJson(jsonString, reflect.TypeOf(pb.Complex{})))
}

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          15,
		IsSimple:    true,
		Name:        "jheni",
		SampleLists: []int32{1, 2, 3, 4, 5},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{Id: 10, Name: "jheni"},
		MultipleDummies: []*pb.Dummy{
			{Id: 11, Name: "A"},
			{Id: 12, Name: "B"},
			{Id: 13, Name: "C"},
		},
	}
}

func doOneof(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Errorf("message has unexpected type: %v", x)
	}
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"1": {Id: 1},
			"2": {Id: 2},
			"3": {Id: 3},
		},
	}
}
