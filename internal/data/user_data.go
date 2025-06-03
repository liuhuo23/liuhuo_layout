package data

import (
	"fmt"
	"liuhuo23/liuos/internal/model"

	"gorm.io/gorm"
)

type UserData struct {
	data *Data
	db   *gorm.DB
}

func NewUserData(data *Data) *UserData {
	db := data.Db
	return &UserData{data: data, db: db}
}

func (u *UserData) CreateUser(user *model.User) error {
	return u.db.Create(user).Error
}

func (u *UserData) GetUserInfo(userID uint) (*model.User, error) {
	var user model.User
	if u.db == nil {
		fmt.Println("db is nil")
		return nil, nil
	}
	if err := u.db.First(&user, 1).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserData) GetUserRoles(userID uint) ([]model.Role, error) {
	var roles []model.Role
	if err := u.db.Model(&model.UserRole{UserID: userID}).
		Joins("JOIN roles ON user_roles.role_id = roles.id").
		Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (u *UserData) GetRolePermissions(roleID uint) ([]model.Permission, error) {
	var permissions []model.Permission
	if err := u.db.Model(&model.RolePermission{RoleID: roleID}).
		Joins("JOIN permissions ON role_permissions.permission_id = permissions.id").
		Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (u *UserData) CheckPermission(userID uint, resource string, action string) (bool, error) {
	var count int64
	if err := u.db.Model(&model.RolePermission{}).
		Joins("JOIN user_roles ON role_permissions.role_id = user_roles.role_id").
		Joins("JOIN permissions ON role_permissions.permission_id = permissions.id").
		Where("user_roles.user_id = ? AND permissions.resource = ? AND permissions.action = ?", userID, resource, action).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
