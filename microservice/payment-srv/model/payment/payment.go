/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-20 08:47:27
 * @LastEditTime: 2019-08-20 09:10:05
 * @LastEditors: Please set LastEditors
 */
package payment

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	"github.com/yuwe1/micolearn/microservice/basic/common"
	invS "github.com/yuwe1/micolearn/microservice/inventory-srv/proto/inventory"
	ordS "github.com/yuwe1/micolearn/microservice/orders-srv/proto/orders"
	proto "github.com/yuwe1/micolearn/microservice/payment-srv/proto/payment"
)

var (
	s            *service
	m            sync.RWMutex
	invClient    invS.InventoryService
	ordSClient   ordS.OrdersService
	payPublisher micro.Publisher
)

type service struct{}

type Service interface {
	PayOrder(orrderId int64) error
}

func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("Getservice未被初始化")
	}
	return s, nil
}
func Init() {
	m.Lock()
	defer m.Unlock()
	if s != nil {
		return
	}
	invClient = invS.NewInventoryService("mu.micro.book.srv.inventory", client.DefaultClient)
	ordSClient = ordS.NewOrdersService("mu.micro.book.srv.orders", client.DefaultClient)
	payPublisher = micro.NewPublisher(common.TopicPaymentDone, client.DefaultClient)
	s = &service{}
}

// 发送支付事件
func (s *service) sendPayDoneEvt(orderId int64, state int32) {
	// 构建事件
	ev := &proto.PayEvent{
		Id:       uuid.New().String(),
		SentTime: time.Now().Unix(),
		OrderId:  orderId,
		State:    state,
	}

	log.Logf("[sendPayDoneEvt] 发送支付事件，%+v\n", ev)

	// 广播
	if err := payPublisher.Publish(context.Background(), ev); err != nil {
		log.Logf("[sendPayDoneEvt] 异常: %v", err)
	}
}
