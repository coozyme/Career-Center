package account

import (
	"CareerCenter/domain/entity/account"
	"CareerCenter/internal/repository/mapper"
	"CareerCenter/internal/repository/models"
	"CareerCenter/utils/exceptions"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (l AccountMysqlInteractor) GetByEmail(ctx context.Context, email string) (*account.AccountDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE email = ?`, models.GetTableNameAccount())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.AccountModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result := dbq.MustQ(ctx, l.DbConn, stmt, opts, email)
	if result == nil {
		return nil, exceptions.ErrorWrongEmailorPassword
	}

	account := mapper.ModelToEntity(result.(*models.AccountModel))

	return account, nil
}
