package activity

import (
	activityDomain "skyshi_gethired/domain/activity"
	activityRepository "skyshi_gethired/infrastructure/repository/mysql/activity"
)

type Service struct {
	ActivityRepository activityRepository.Repository
}

// GetAll is a function that returns all activitys
func (s *Service) GetAll() (todos []activityDomain.Activity, err error) {
	todos, err = s.ActivityRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// GetByActivity is a function that returns all todos
func (s *Service) GetByActivity(activityID string) (todos []activityDomain.Activity, err error) {
	todos, err = s.ActivityRepository.GetByActivity(activityID)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// GetByID is a function that returns a activity by id
func (s *Service) GetByID(id int) (*activityDomain.Activity, error) {
	return s.ActivityRepository.GetByID(id)
}

// Create is a function that creates a activity
func (s *Service) Create(activity *NewActivity) (*activityDomain.Activity, error) {
	todoModel := activity.toDomainMapper("very-high")
	return s.ActivityRepository.Create(todoModel)
}

// Delete is a function that deletes a activity by id
func (s *Service) Delete(id int) error {
	return s.ActivityRepository.Delete(id)
}

// Update is a function that updates a activity by id
func (s *Service) Update(id uint, medicineMap map[string]interface{}) (*activityDomain.Activity, error) {
	return s.ActivityRepository.Update(id, medicineMap)
}
