package service

import (
	"time"

	"github.com/xpp666/go.orm/dbinit"
	"github.com/xpp666/go.orm/model"
)

// AddUser 增加新用户
func AddUser(user model.User) (bool, error) {
	//user.ID = string()
	t1, _ := time.ParseDuration("8h")
	user.CreateTime = time.Now().Add(t1)
	user.UpdateTime = time.Now().Add(t1)
	if err := dbinit.Db.Create(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}

// DelByID 通过ID删除用户
func DelByID(ID int) (bool, []model.User, error) {
	if err := dbinit.Db.Where("ID=?", ID).Delete(&model.User{}).Error; err != nil {
		return false, []model.User{}, err
	}
	userList, _ := FindAllUser()
	return true, userList, nil
}

// DelByName 通过用户名删除用户
func DelByName(Name string) (bool, error) {
	if err := dbinit.Db.Where("Name=?", Name).Delete(&model.User{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

//FindAllUser 查询所有用户
func FindAllUser() ([]model.User, error) {
	users := []model.User{}
	if err := dbinit.Db.Model(&model.User{}).Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

// FindUserByID 通过ID查找用户
func FindUserByID(ID int) (model.User, error) {
	// return db.Model(&User{}).Where(&User{ID: ID}).First(&User)
	//defer dbinit.Db.Close()
	var u model.User
	if err := dbinit.Db.Where("ID = ?", ID).Find(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

// FindUserByName 通过用户名查找用户
func FindUserByName(Name string) (model.User, error) {
	// return db.Model(&User{}).Where(&User{ID: ID}).First(&User)
	var u model.User
	if err := dbinit.Db.Where("Name = ?", Name).Find(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

// 通过ID修改用户信息
func UpdateUserByID(user *model.User) (bool, []model.User, error) {
	t1, _ := time.ParseDuration("8h")
	user.UpdateTime = time.Now().Add(t1)
	if err := dbinit.Db.Model(&user).Where("ID = ?", user.ID).Update(&user).Error; err != nil {
		return false, []model.User{}, err
	}
	userList, _ := FindAllUser()
	return false, userList, nil
}

// 通过用户名修改用户信息
func UpdateUserByName(user model.User) (bool, []model.User, error) {
	t1, _ := time.ParseDuration("8h")
	user.UpdateTime = time.Now().Add(t1)
	if err := dbinit.Db.Model(&user).Where("Name = ?", user.Name).Update(user).Error; err != nil {
		return false, []model.User{}, err
	}
	userList, _ := FindAllUser()
	return true, userList, nil
}
