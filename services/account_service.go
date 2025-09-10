package services

import m "mock-api/model"

type AccountService struct {
	db m.Database
}

func (a *AccountService) SetDatabase(db m.Database) {
	a.db = db
}
