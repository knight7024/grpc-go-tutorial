package user

import (
	pbUser "github.com/knight7024/grpc-go-tutorial/proto/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var People = []*pbUser.Person{
	{
		Name:  "철수",
		Id:    1,
		Email: "chulsoo@test.com",
		Phones: []*pbUser.Person_PhoneNumber{
			{
				Number: "010-0000-0000",
				Type:   pbUser.Person_MOBILE,
			},
			{
				Number: "02-1111-1111",
				Type:   pbUser.Person_HOME,
			},
		},
	},
	{
		Name:        "영희",
		Id:          2,
		Email:       "younghee@test.com",
		LastUpdated: timestamppb.Now(),
	},
}
