package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	//UserModel 是一个需要自定义的接口，在这里添加更多方法，并在 customUserModel 中实现添加的方法。
	UserModel interface {
		userModel
		FindByMobile(ctx context.Context, mobile string) (*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c, opts...),
	}
}

func (m *customUserModel) FindByMobile(ctx context.Context, mobile string) (*User, error) {
	var user User
	// UserModel中声明了这个方法  customUserModel中实现这个方法
	// fmt.Println("func (m *customUserModel) FindByMobile(ctx context.Context, mobile string) (*User, error) {")
	err := m.QueryRowNoCacheCtx(ctx, &user, fmt.Sprintf("select %s from %s where `mobile` = ? limit 1", userRows, m.table), mobile)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
