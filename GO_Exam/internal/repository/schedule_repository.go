package repository

import (
	"project/itStep/internal/config"
	"project/itStep/internal/models"
)

func CreateSchedule(schedule *models.Schedule) error {
	return config.DB.Create(schedule).Error
}

func GetCommunitySchedule(
	communityID uint,
) ([]models.Schedule, error) {

	var schedules []models.Schedule

	err := config.DB.
		Preload("Teacher").
		Where("community_id = ?", communityID).
		Order("date asc").
		Find(&schedules).Error

	if err != nil {
		return nil, err
	}

	return schedules, nil
}
func GetScheduleByID(id uint) (*models.Schedule, error) {

	var schedule models.Schedule

	err := config.DB.
		First(&schedule, id).Error

	if err != nil {
		return nil, err
	}

	return &schedule, nil
}

func UpdateSchedule(schedule *models.Schedule) error {
	return config.DB.Save(schedule).Error
}
func DeleteSchedule(id uint) error {

	return config.DB.
		Delete(&models.Schedule{}, id).Error
}