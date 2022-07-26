package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func writeToJson(ppb proto.Message) string {
	fmt.Println("Iniciando escrita do json")

	option := protojson.MarshalOptions{
		Multiline: true,
	}
	out, err := option.Marshal(ppb)

	if err != nil {
		log.Fatal("Can't convert to json", err)
		return ""
	}

	return string(out)
}

func doFromJson(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJson(jsonString, message)

	return message
}

func fromJson(in string, ppb proto.Message) {

	fmt.Println("Obtendo valor do json")
	option := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	if err := option.Unmarshal([]byte(in), ppb); err != nil {
		log.Fatalln("Couldn't unmarshall from JSON", err)
	}

}

func writeToFile(name string, ppb proto.Message) {
	out, err := proto.Marshal(ppb)

	if err != nil {
		log.Fatal("Error", err)
	}

	if err = ioutil.WriteFile(name, out, 0644); err != nil {
		log.Fatal("Can't write to file", err)
		return
	}

	fmt.Println("Data has been written!")
}
