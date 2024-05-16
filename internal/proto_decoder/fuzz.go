package proto_decoder

import (
	"github.com/sirupsen/logrus"
	"github.com/ysmood/grpc-tools/internal"
)

func Fuzz(data []byte) int {
	dec := NewDecoder(logrus.New())

	_, err := dec.Decode("", &internal.Message{
		RawMessage: data,
	})
	if err != nil {
		return 0
	}

	return 1
}
