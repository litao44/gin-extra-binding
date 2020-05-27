package ebinding

import (
	"bytes"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"

	"github.com/litao44/gin-extra-binding/testpb"
)

func Test_decodePB(t *testing.T) {
	as := assert.New(t)

	p1 := testpb.Person{
		Name:  "hello",
		Id:    100,
		Email: "world@163.com",
	}

	body, err := proto.Marshal(&p1)
	as.Nil(err)

	var p2 testpb.Person
	err = decodePB(bytes.NewReader(body), &p2)
	as.Nil(err)

	as.Equal(p1.Name, p2.Name)
	as.Equal(p1.Id, p2.Id)
	as.Equal(p1.Email, p2.Email)
}
