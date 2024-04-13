package services

import (
	"context"

	"github.com/saki-engineering/graphql-sample/graph/db"
	"github.com/saki-engineering/graphql-sample/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userServices struct {
	exec boil.ContextExecutor
}

func converUser(user *db.User) *model.User {
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}
}

func (u *userServices) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	// 1. SQLBoilerで生成されたORMコードを呼び出す
	user, err := db.Users(
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name),
		db.UserWhere.Name.EQ(name),
	).One(ctx, u.exec)
	// 2. エラー処理
	if err != nil {
		return nil, err
	}
	// 3. 戻り値の*model.User型を作る
	return converUser(user), nil
}
