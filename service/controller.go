package service

import (
	"github.com/ant0ine/go-json-rest/rest"
	"strconv"
	"log"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/xpp666/go.orm/model"
)

// 查询用户
func QueryUserById(w rest.ResponseWriter, r *rest.Request) {
	returnJson := make(map[string]interface{})

	//解析获取数据
	id := r.PathParam("id")
	userId, _ := strconv.Atoi(id)

	userInfo, err := FindUserByID(userId)
	if err != nil {
		log.Println(err)
		//操作失败返回结果
		returnJson["code"] = 1
		returnJson["msg"] = err.Error()
		panic(w.WriteJson(returnJson))
		return
	}

	//操作成功返回结果
	returnJson["code"] = 0
	returnJson["msg"] = "query userInfo success!"
	returnJson["user"] = userInfo
	w.WriteJson(returnJson)
}

// 注册用户
func Register(w rest.ResponseWriter, r *rest.Request)  {

	returnJson := make(map[string]interface{})
	//解析获取数据
	//id := r.Body
	data,err :=ioutil.ReadAll(r.Body)
	var user model.User
	json.Unmarshal(data,&user)
	fmt.Println(user)
	b, err := AddUser(user)
	if err != nil {
		log.Println(err)
		//操作失败返回结果
		returnJson["code"] = 1
		returnJson["msg"] = err.Error()
		panic(w.WriteJson(returnJson))
		return
	}

	//操作成功返回结果
	returnJson["code"] = 0
	returnJson["msg"] = "add user success!"
	returnJson["result"] = b
	w.WriteJson(returnJson)
}

// 删除用户
func DeleteUser(w rest.ResponseWriter, r *rest.Request) {
	returnJson := make(map[string]interface{})

	//解析获取数据
	id := r.PathParam("id")
	userId, _ := strconv.Atoi(id)

	//通过gorm操作数据库
	b, err := DelByID(userId)
	if err != nil {
		log.Println(err)
		//操作失败返回结果
		returnJson["code"] = 1
		returnJson["msg"] = err.Error()
		panic(w.WriteJson(returnJson))
		return
	}

	//操作成功返回结果
	returnJson["code"] = 0
	returnJson["msg"] = "delete user success!"
	returnJson["result"] = b
	w.WriteJson(returnJson)
}

func Update(w rest.ResponseWriter, r *rest.Request)  {
	returnJson := make(map[string]interface{})
	//解析获取数据
	//id := r.Body
	data,err :=ioutil.ReadAll(r.Body)
	var user model.User
	json.Unmarshal(data,&user)
	fmt.Println(user)
	b, err := UpdateUserByID(&user)
	if err != nil {
		log.Println(err)
		//操作失败返回结果
		returnJson["code"] = 1
		returnJson["msg"] = err.Error()
		panic(w.WriteJson(returnJson))
		return
	}

	//操作成功返回结果
	returnJson["code"] = 0
	returnJson["msg"] = "update user success!"
	returnJson["result"] = b
	w.WriteJson(returnJson)
}

