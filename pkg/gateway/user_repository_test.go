package gateway

import (
	"fmt"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"

	"github.com/pecolynx/golang-mysql-test/pkg/domain"
)

func TestAddUser(t *testing.T) {
	db, err := openMySQLForTest()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewUserRepository(db)
	fn := func() {
		id, err := repo.AddUser(domain.NewUserAddParameter("taro", 19))
		assert.NoError(t, err)
		fmt.Println(id)
	}
	withSetup(fn)
}
