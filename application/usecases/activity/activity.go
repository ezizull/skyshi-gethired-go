package activity

import (
	activityDomain "skyshi_gethired/domain/activity"
	activityRepository "skyshi_gethired/infrastructure/repository/mysql/activity"
)

type Service struct {
	ActivityRepository activityRepository.Repository
}

// GetAll is a function that returns all todos
func (s *Service) GetAll() (todos []activityDomain.Activity, err error) {
	todos, err = s.ActivityRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// GetByID is a function that returns a activity by id
func (s *Service) GetByID(id string) (*activityDomain.Activity, error) {
	return s.ActivityRepository.GetByID(id)
}

// GetActivity is a function that returns a activity by id
func (s *Service) GetActivity(id string) (*activityDomain.Activity, error) {
	return s.ActivityRepository.GetActivity(id)
}

// Create is a function that creates a activity
func (s *Service) Create(activity *NewActivity) (*activityDomain.Activity, error) {
	todoModel := activity.toDomainMapper()
	return s.ActivityRepository.Create(todoModel)
}

// Update is a function that updates a activity by id
func (s *Service) Update(id string, activity *UpdateActivity) (*activityDomain.Activity, error) {
	todoModel := activity.toDomainMapper()
	_, err := s.ActivityRepository.Update(id, todoModel)
	if err != nil {
		return nil, err
	}

	return s.ActivityRepository.GetByID(id)
}

// Delete is a function that deletes a activity by id
func (s *Service) Delete(id string) error {
	_, err := s.ActivityRepository.GetByID(id)
	if err != nil {
		return err
	}

	return s.ActivityRepository.Delete(id)
}
