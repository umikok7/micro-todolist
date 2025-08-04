package core

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"user/model"
	"user/service"
)

// BuildUser 数据库模型转换为返回给服务层的模型
func BuildUser(item model.User) *service.UserModel {
	userModel := service.UserModel{
		ID:        uint32(item.ID),
		UserName:  item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &userModel
}

func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest, resp *service.UserDetailResponse) error {
	var user model.User
	resp.Code = 200
	if err := model.DB.Where("user_name=?", req.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			resp.Code = 400
			return nil
		}
		resp.Code = 500
		return nil
	}
	// 检验密码
	if user.CheckPassword(req.Password) == false {
		resp.Code = 400
		return nil
	}
	// 进行序列化
	resp.UserDetail = BuildUser(user)
	return nil
}

func (*UserService) UserRegister(ctx context.Context, req *service.UserRequest, resp *service.UserDetailResponse) error {
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码输入不一致")
		return err
	}
	var count int64 = 0
	if err := model.DB.Model(&model.User{}).Where("user_name=?", req.UserName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		err := errors.New("用户名已存在")
		return err
	}
	user := model.User{
		UserName: req.UserName,
	}
	if err := user.SetPassword(req.Password); err != nil {
		return err
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return err
	}
	resp.UserDetail = BuildUser(user)
	return nil
}
