package services

import (
	"context"

	"github.com/saki-engineering/graphql-sample/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Services interface {
	UserService
	// issueテーブルを扱うIssueServiceなど、他のサービスインターフェースができたらそれらを追加していく
}

type services struct {
	*userServices
	// issueテーブルを扱うissueServiceなど、他のサービス構造体ができたらフィールドを追加していく
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userServices: &userServices{exec: exec},
	}
}
