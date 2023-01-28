package service

import (
	"github.com/jinzhu/gorm"
	"todo_list/model"
	"todo_list/pkg/utils"
	"todo_list/serializer"
)

type UserService struct{
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"FanOne"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?" , service.UserName).First(&user).Count(&count)
	if count == 1{
		return serializer.Response{
			Status: 400,
			Msg: "已经有这个人了，无需再注册",
		}
	}
	user.UserName = service.UserName
	//密码加密
	if err := user.SetPassword(service.Password);err  != nil{
		return serializer.Response{
			Status: 400,
			Msg: err.Error(),
		}
	}
	//创建用户
	if err := model.DB.Create(&user).Error;err!=nil{
		return serializer.Response{
			Status: 500,
			Msg: "数据库操作错误",
		}
	}

	return serializer.Response{
		Status: 200,
		Msg: "用户注册成功了",
	}
}

func (service *UserService) Login() serializer.Response{
	var user model.User
	//先去找一下这个user
	if err := model.DB.Where("user_name=?",service.UserName).First(&user).Error; err != nil{
		if gorm.IsRecordNotFoundError(err){
			return serializer.Response{
				Status: 400,
				Msg: "用户不存在，请先登录",
			}
		}
		//如果不是用户不存在，其他原因
		return serializer.Response{
			Status: 500,
			Msg: "数据库错误",
		}
	}
	//密码验证
	if user.CheckPassword(service.Password) == false{
		return serializer.Response{
			Status: 400,
			Msg: "密码错误",
		}
	}
	//发一个token，为了验证身份（其他功能）所给前端存储的
	token,err := utils.GenerateToken(user.ID,service.UserName,0)
	if err != nil{
		return serializer.Response{
			Status: 500,
			Msg: "token签发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Data: serializer.TokenData{User: serializer.BuildUser(user),Token: token},
		Msg: "登陆成功",
	}
}