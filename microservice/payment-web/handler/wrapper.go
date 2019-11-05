/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-19 16:40:54
 * @LastEditTime: 2019-08-20 09:49:06
 * @LastEditors: Please set LastEditors
 */
package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/micro/go-micro/util/log"
	auth "github.com/yuwe1/micolearn/microservice/auth/proto/auth"
	"github.com/yuwe1/micolearn/microservice/basic/common"
	"github.com/yuwe1/micolearn/microservice/plugins/session"
)

// AuthWrapper 认证wrapper
func AuthWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ck, _ := r.Cookie(common.RememberMeCookieName)
		// token不存在，则状态异常，无权限
		if ck == nil {
			http.Error(w, "非法请求", 400)
			return
		}
		fmt.Print("ck:", ck.String())
		sess := session.GetSession(w, r)
		if sess.ID != "" {
			// 检测是否通过验证
			if sess.Values["valid"] != nil {
				h.ServeHTTP(w, r)
				return
			} else {
				userId := sess.Values["userId"].(int64)

				if userId != 0 {
					rsp, err := authClient.GetCachedAccessToken(context.TODO(), &auth.Request{
						UserId: uint64(userId),
					})
					if err != nil {
						log.Logf("[AuthWrapper]，err：%s", err)
						http.Error(w, "非法请求", 400)
						return
					}

					// token不一致
					if rsp.Token != ck.Value {
						log.Logf("[AuthWrapper]，token不一致")
						http.Error(w, "非法请求", 400)
						return
					}
				} else {
					log.Logf("[AuthWrapper]，session不合法，无用户id")
					http.Error(w, "非法请求", 400)
					return
				}
			}
		} else {
			http.Error(w, "非法请求", 400)
			return
		}

		h.ServeHTTP(w, r)
	})
}
