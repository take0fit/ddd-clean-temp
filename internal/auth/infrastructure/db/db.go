package db

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const TxKey string = "auth_tx"

type DB struct {
	*gorm.DB
}

var d *DB

func RDBConnect(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %v", err)
	}

	d = &DB{db}
	return nil
}

func GetDB() *DB {
	return d
}

func (d *DB) Transaction(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	tx := d.Begin()
	defer func(tx *gorm.DB) {
		if r := recover(); r != nil {
			tx.Rollback()
			err = errors.New(fmt.Sprintf("recovered from panic, err: %s", r))
		}
	}(tx)

	txCtx := context.WithValue(ctx, TxKey, tx)
	if err = fn(txCtx); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
