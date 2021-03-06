package main

import (
	"errors"
	"log"

	"github.com/bygui86/go-protobuf/domain"
	"github.com/golang/protobuf/proto"
)

func main() {
	// Test to be marshalled
	marshalTest := &domain.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Reps:  []int64{1, 2, 3},
		Optionalgroup: &domain.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
		Union: &domain.Test_Name{"fred"},
	}

	// Marshalling
	marshalledData, marshalErr := proto.Marshal(marshalTest)
	if marshalErr != nil {
		log.Fatalln("marshaling error: ", marshalErr)
		panic(marshalErr)
	}

	// Test to contain unmarshalled data
	unmarshalTest := &domain.Test{}

	// Unmarshalling
	unmarshalErr := proto.Unmarshal(marshalledData, unmarshalTest)
	if unmarshalErr != nil {
		log.Fatalln("unmarshaling error: ", unmarshalErr.Error)
		panic(unmarshalErr)
	}

	// Now test and newTest contain the same data.
	if marshalTest.GetLabel() != unmarshalTest.GetLabel() {
		log.Fatalln("data mismatch %q != %q", marshalTest.GetLabel(), unmarshalTest.GetLabel())
		panic(errors.New("data mismatch"))
	}

	// Use a type switch to determine which oneof was set.
	switch oneOfType := unmarshalTest.Union.(type) {
	case *domain.Test_Number: // marshalTest.Union.Number contains the number.
		log.Println("oneof union contains a number:", oneOfType.Number)
	case *domain.Test_Name: // marshalTest.Union.Name contains the string.
		log.Println("oneof union contains a name:", oneOfType.Name)
	default:
		log.Fatalln("oneof union type not recognized")
		panic(errors.New("oneof union type not recognized"))
	}
}
