package serverapi

import (
	"lib/common/kafkautil"
	"time"

	"github.com/Shopify/sarama"
)

var GProducer *kafkautil.AsyncProducer = nil

func InitAsyncProducer(addrs []string) error {
	producerPtr, err := kafkautil.NewAsyncProducer(addrs, sarama.WaitForLocal, 1*time.Second)
	if err != nil {
		return err
	}

	GProducer = producerPtr

	return nil
}

const KLogUserOperTopic string = "logUserOper"

type UserOperation int32

const (
	UnKnown UserOperation = iota
	AnswerQuestion
)

type UserOperReq struct {
	UserID     string        `json:"uid"`
	RoomID     string        `json:"room_id"`
	RoomType   int64         `json:"room_type"`
	Operation  UserOperation `json:"oper"`
	Question   int64         `json:"question"`
	Answer     string        `json:"answer"`
	UserPick   string        `json:"user_pick"`
	AnswerTime time.Time     `json:"answer_time"`
}

func SendLogUserOperMsg(reqPtr *UserOperReq) error {
	key := "log"

	// TODO-ly: 并发?
	err := GProducer.SendJSON(KLogUserOperTopic, key, reqPtr, kafkautil.NeverTimeout)
	if err != nil {
		return err
	}

	return nil
}
