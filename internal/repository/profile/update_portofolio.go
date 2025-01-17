package profile

import (
	"CareerCenter/internal/repository/models/profile"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (p ProfileMysqlInteractor) UpdatePortofolio(ctx context.Context, email string, path string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET portofolio ='%s' WHERE email = '%s' ",
		profile.GetTableNameProfile(), path, email)

	_, err := dbq.E(ctx, p.DbConn, query, nil)

	if err != nil {
		return err
	}

	return nil
}
