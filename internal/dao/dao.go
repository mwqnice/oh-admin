package dao

import (
	"context"
	"gorm.io/gorm"
)

type Dao struct {
	engine *gorm.DB
	ctx    context.Context
}

func New(engine *gorm.DB, ctx context.Context) *Dao {
	return &Dao{engine: engine, ctx: ctx}
}

func (d *Dao) SetTxEngine(engine *gorm.DB) {
	d.engine = engine
}

func (d *Dao) Begin() {
	d.engine = d.engine.Begin()
}

func (d *Dao) Commit() {
	d.engine.Commit()
}

func (d *Dao) Rollback() {
	d.engine.Rollback()
}

func (d *Dao) TxError() error {
	return d.engine.Error
}
