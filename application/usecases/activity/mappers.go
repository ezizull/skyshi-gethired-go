package activity

import (
	activityDomain "skyshi_gethired/domain/activity"
)

func (n *NewActivity) toDomainMapper(todoPriority string) *activityDomain.Activity {
	return &activityDomain.Activity{
		Title: *n.Title,
		Email: *n.Email,
	}
}
