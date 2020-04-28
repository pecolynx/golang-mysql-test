package gateway

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/pecolynx/golang-mysql-test/pkg/application"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) application.Repository {
	return &repository{
		db: db,
	}
}

type db struct {
	db *gorm.DB
}

type tx struct {
	db *gorm.DB
}

func (r *repository) NewDB() application.DB {
	return &db{
		db: r.db,
	}
}

func (r *repository) NewTX() application.TX {
	return &tx{
		db: r.db,
	}
}

func (d *db) DB() *gorm.DB {
	return d.db
}

func (d *tx) begin() (*gorm.DB, error) {
	// START TRANSACTION
	tx := d.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tx, nil
}

func (d *tx) rollback(tx *gorm.DB) error {
	// ROLLBACK
	result := tx.Rollback()
	return result.Error
}

func (d *tx) commit(tx *gorm.DB) error {
	// COMMIT
	result := tx.Commit()
	return result.Error
}

func (d *tx) Transaction(fn func(db *gorm.DB) error) error {
	tx, err := d.begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			log.Info("recover")
			d.rollback(tx)
			panic(p)
		} else if err != nil {
			log.Info("rollback")
			d.rollback(tx)
		} else {
			log.Debug("commit")
			err = d.commit(tx)
		}
	}()
	err = fn(tx)
	return err

}
