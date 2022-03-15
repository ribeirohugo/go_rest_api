//go:generate mockgen -package postgres -source=postgres.go -destination postgres_mock.go

package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const emailTest = "email@domain"

func TestDatabase_GetUserByEmail(t *testing.T) {
	t.Run("returns error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sqlMock := NewMockSQL(ctrl)

		db := Database{sql: sqlMock}

		sqlRows := &sql.Rows{}

		sqlMock.EXPECT().
			QueryContext(gomock.Any(), gomock.Any(), emailTest).
			Return(sqlRows, fmt.Errorf("error")).
			Times(1)

		_, err := db.GetUserByEmail(context.Background(), emailTest)
		assert.Error(t, err)
	})
}
