package db

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
	"os"
	"strconv"
)

const DuplicateEntryErrorNumber = 1062

var dbUser string
var dbPassword string
var dbHost string
var dbPort int
var dbName string
var dbInnodbLockWaitTimeout int

type TxKey string

const txKey TxKey = "lispo_tx"

type DB struct {
	*gorm.DB
}

var d *DB

func init() {
	dbUser = os.Getenv("DB_USER")
	if dbUser == "" {
		panic("DB_USER is unset")
	}
	dbPassword = os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		panic("DB_PASSWORD is unset")
	}
	dbHost = os.Getenv("DB_HOST")
	if dbHost == "" {
		panic("DB_HOST is unset")
	}
	dbPortStr := os.Getenv("DB_PORT")
	if dbPortStr == "" {
		panic("DB_PORT is unset")
	} else {
		var e error
		dbPort, e = strconv.Atoi(dbPortStr)
		if e != nil {
			panic(e)
		}
	}
	dbName = os.Getenv("DB_NAME")
	if os.Getenv("DB_NAME") == "" {
		panic("DB_NAME is unset")
	}

	if os.Getenv("INNODB_LOCK_WAIT_TIMEOUT_SECOND") != "" {
		var e error
		dbInnodbLockWaitTimeout, e = strconv.Atoi(os.Getenv("INNODB_LOCK_WAIT_TIMEOUT_SECOND"))
		if e != nil {
			panic(e)
		}
	}
}

func RDBConnect() error {
	escapedLoc := url.QueryEscape("Asia/Tokyo")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, escapedLoc)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("open db: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("get generic database object from gorm db: %v", err)
	}

	if dbInnodbLockWaitTimeout > 0 {
		q := fmt.Sprintf("SET innodb_lock_wait_timeout=%d", dbInnodbLockWaitTimeout)
		if _, err := sqlDB.Exec(q); err != nil {
			return fmt.Errorf("execute query: %v", err)
		}
	}

	maxIdleConnsStr := os.Getenv("DB_MAX_IDLE_CONNS")
	maxIdleConns, err := strconv.Atoi(maxIdleConnsStr)
	if err == nil && maxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(maxIdleConns)
	}

	maxOpenConnsStr := os.Getenv("DB_MAX_OPEN_CONNS")
	maxOpenConns, err := strconv.Atoi(maxOpenConnsStr)
	if err == nil && maxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(maxOpenConns)
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

	txCtx := context.WithValue(ctx, txKey, tx)
	if err = fn(txCtx); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
