package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ AccountModel = (*customAccountModel)(nil)

type (
	// PolicyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPolicyModel.
	AccountModel interface {
		accountModel
	}

	customAccountModel struct {
		*defaultAccountModel
	}
)

// NewPolicyModel returns a model for the mongo.
func NewAccountModel(url, db, collection string) AccountModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customAccountModel{
		defaultAccountModel: newDefaultAccountModel(conn),
	}
}
