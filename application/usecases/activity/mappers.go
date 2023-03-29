package activity

import (
	activityDomain "skyshi_gethired/domain/activity"
)

func (n *NewActivity) toDomainMapper() *activityDomain.Activity {
	return &activityDomain.Activity{
		Title: *n.Title,
		Email: *n.Email,
	}
}

func (n *UpdateActivity) toDomainMapper() *activityDomain.Activity {
	return &activityDomain.Activity{
		Title: *n.Title,
	}
}
