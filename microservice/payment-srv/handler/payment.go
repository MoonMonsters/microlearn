/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-20 09:37:40
 * @LastEditTime: 2019-08-20 09:38:33
 * @LastEditors: Please set LastEditors
 */
package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"
	"github.com/yuwe1/micolearn/microservice/payment-srv/model/payment"
	proto "github.com/yuwe1/micolearn/microservice/payment-srv/proto/payment"
)

var (
	paymentService payment.Service
)

type Service struct {
}

// Init 初始化handler
func Init() {
	paymentService, _ = payment.GetService()
}

// New 新增订单
func (e *Service) PayOrder(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	log.Log("[PayOrder] 收到支付请求")
	err = paymentService.PayOrder(req.OrderId)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Detail: err.Error(),
		}
		return
	}

	rsp.Success = true
	return
}
