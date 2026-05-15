package service

import (
	"errors"

	"project/itStep/internal/models"
	"project/itStep/internal/repository"
)

func CreateSchedule(schedule *models.Schedule) error {

	if len(schedule.Title) < 3 {
		return errors.New("title too short")
	}

	return repository.CreateSchedule(schedule)
}

func GetCommunitySchedule(
	communityID uint,
) ([]models.Schedule, error) {

	return repository.GetCommunitySchedule(
		communityID,
	)
}
func UpdateSchedule(
	scheduleID uint,
	userID uint,
	role string,
	title string,
	description string,
	date string,
	time string,
) error {

	schedule, err := repository.GetScheduleByID(scheduleID)
	if err != nil {
		return errors.New("schedule not found")
	}

	isOwner := schedule.TeacherID == userID

	isAdmin := role == "admin"

	if !isOwner && !isAdmin {
		return errors.New("forbidden")
	}

	schedule.Title = title
	schedule.Description = description
	schedule.Date = date
	schedule.Time = time

	return repository.UpdateSchedule(schedule)
}

func DeleteSchedule(
	scheduleID uint,
	userID uint,
	role string,
) error {

	schedule, err := repository.GetScheduleByID(scheduleID)
	if err != nil {
		return errors.New("schedule not found")
	}

	isOwner := schedule.TeacherID == userID

	isAdmin := role == "admin"

	if !isOwner && !isAdmin {
		return errors.New("forbidden")
	}

	return repository.DeleteSchedule(scheduleID)
}