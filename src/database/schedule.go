package database

import (
	"../model"
)

// errorを返すかは要検討
func (d *db) FindAllSchedules() ([]model.Schedule, error) {
	var schedules []model.Schedule
	d.orm.Find(&schedules)

	return schedules, nil
}