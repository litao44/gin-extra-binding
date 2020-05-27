package ebinding

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin/binding"
	"github.com/golang/protobuf/proto"
)

var _ binding.Binding = &pbBinding{}

type pbBinding struct{}

func (p pbBinding) Name() string {
	return "pb"
}

func (p pbBinding) Bind(req *http.Request, obj interface{}) error {
	if req == nil || req.Body == nil {
		return fmt.Errorf("invalid request")
	}
	return decodePB(req.Body, obj)
}

func decodePB(r io.Reader, obj interface{}) error {
	message, ok := obj.(proto.Message)
	if !ok {
		return fmt.Errorf("invalid object")
	}

	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	err = proto.Unmarshal(body, message)
	if err != nil {
		return err
	}

	return validate(obj)
}
