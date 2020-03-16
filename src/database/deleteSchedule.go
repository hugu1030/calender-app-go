package database

import (
	"../model"
)

func(d *db) DeleteSchedule(id) ([]model.Schedule,error){
	var schedules []model.Schedule
	d.Where("Id = ?",id).Delete(&schedules)
}