package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	auth "github.com/go-micro-cn/tutorials/auth/proto/auth"
	us "github.com/go-micro-cn/tutorials/user-service/proto/user"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
)

var (
	serviceClient us.UserService
	authClient    auth.Service
)

// Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = us.NewUserService("mu.micro.book.service.user", client.DefaultClient)
	authClient = auth.NewService("mu.micro.book.service.auth", client.DefaultClient)
}

// Login 登录入口
func Login(w http.ResponseWriter, r *http.Request) {
	// 只接受POST请求
	if r.Method != "POST" {
		log.Warn("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()

	// 调用后台服务
	rsp, err := serviceClient.QueryUserByName(context.TODO(), &us.Request{
		UserName: r.Form.Get("userName"),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 返回结果
	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}

	if rsp.User.Pwd == r.Form.Get("pwd") {
		response["success"] = true

		// 干掉密码返回
		rsp.User.Pwd = ""
		response["data"] = rsp.User
		log.Info("[Login] 密码校验完成，生成token...")

		// 生成token
		rsp2, err := authClient.MakeAccessToken(context.TODO(), &auth.Request{
			UserId:   rsp.User.Id,
			UserName: rsp.User.Name,
		})
		if err != nil {
			log.Errorf("[Login] 创建token失败，err：%s", err)
			http.Error(w, err.Error(), 500)
			return
		}

		log.Infof("[Login] token %s", rsp2.Token)
		response["token"] = rsp2.Token

		// 同时将token写到cookies中
		w.Header().Add("set-cookie", "application/json; charset=utf-8")
		// 过期30分钟
		expire := time.Now().Add(30 * time.Minute)
		cookie := http.Cookie{Name: "remember-me-token", Value: rsp2.Token, Path: "/", Expires: expire, MaxAge: 90000}
		http.SetCookie(w, &cookie)

	} else {
		response["success"] = false
		response["error"] = &Error{
			Detail: "密码错误",
		}
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// Logout 退出登录
func Logout(w http.ResponseWriter, r *http.Request) {
	// 只接受POST请求
	if r.Method != "POST" {
		log.Warn("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	tokenCookie, err := r.Cookie("remember-me-token")
	if err != nil {
		log.Error("token获取失败")
		http.Error(w, "非法请求", 400)
		return
	}

	// 删除token
	_, err = authClient.DelUserAccessToken(context.TODO(), &auth.Request{
		Token: tokenCookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 清除cookie
	cookie := http.Cookie{Name: "remember-me-token", Value: "", Path: "/", Expires: time.Now().Add(0 * time.Second), MaxAge: 0}
	http.SetCookie(w, &cookie)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回结果
	response := map[string]interface{}{
		"ref":     time.Now().UnixNano(),
		"success": true,
	}

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
