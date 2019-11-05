/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-20 09:33:11
 * @LastEditTime: 2019-08-20 09:33:52
 * @LastEditors: Please set LastEditors
 */
package subscriber

import (
	"context"

	"github.com/yuwe1/micolearn/microservice/orders-srv/model/orders"
	payS "github.com/yuwe1/micolearn/microservice/payment-srv/proto/payment"
	"github.com/micro/go-micro/util/log"
)

var (
	ordersService orders.Service
)

// Init 初始化handler
func Init() {
	ordersService, _ = orders.GetService()
}

// PayOrder 订单支付消息
func PayOrder(ctx context.Context, event *payS.PayEvent) (err error) {
	log.Logf("[PayOrder] 收到支付订单通知，%d，%d", event.OrderId, event.State)

	err = ordersService.UpdateOrderState(event.OrderId, int(event.State))
	if err != nil {
		log.Logf("[PayOrder] 收到支付订单通知，更新状态异常，%s", err)
		return
	}
	return
}