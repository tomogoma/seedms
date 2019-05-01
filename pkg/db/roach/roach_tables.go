package roach

import (
	"time"
)

const (
	// Database definition version
	Version = 0

	// Table names
	TblConfigurations = "configurations"
	TblAPIKeys        = "apiKeys"

	// DB Table Columns
	ColID         = "ID"
	ColCreateDate = "createDate"
	ColUpdateDate = "updateDate"
	ColUserID     = "userID"
	ColKey        = "key"
	ColValue      = "value"

	// CREATE TABLE DESCRIPTIONS
	TblDescConfigurations = `
	CREATE TABLE IF NOT EXISTS ` + TblConfigurations + ` (
		` + ColKey + ` VARCHAR(56) PRIMARY KEY NOT NULL CHECK (` + ColKey + ` != ''),
		` + ColValue + ` BYTEA NOT NULL CHECK (` + ColValue + ` != ''),
		` + ColCreateDate + ` TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		` + ColUpdateDate + ` TIMESTAMPTZ NOT NULL
	);
	`
	TblDescAPIKeys = `
	CREATE TABLE IF NOT EXISTS ` + TblAPIKeys + ` (
		` + ColID + ` SERIAL PRIMARY KEY NOT NULL CHECK (` + ColID + `>0),
		` + ColUserID + ` INTEGER NOT NULL,
		` + ColKey + ` VARCHAR(256) NOT NULL CHECK ( LENGTH(` + ColKey + `) >= 56 ),
		` + ColCreateDate + ` TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		` + ColUpdateDate + ` TIMESTAMPTZ NOT NULL
	);
	`
)

// AllTableDescs lists all CREATE TABLE DESCRIPTIONS in order of dependency
// (tables with foreign key references listed after parent table descriptions).
var AllTableDescs = []string{
	TblDescConfigurations,
	TblDescAPIKeys,
}

// AllTableNames lists all table names in order of dependency
// (tables with foreign key references listed after parent table descriptions).
var AllTableNames = []string{
	TblConfigurations,
	TblAPIKeys,
}

type ApiKey struct {
	ID        uint      `json:"id"gorm:"primary_key"`
	UserId    int       `json:"user_id"gorm:"not null"`
	Key       string    `json:"key"gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Configuration struct {
	Key       string    `json:"key"gorm:"primary_key,not null"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
