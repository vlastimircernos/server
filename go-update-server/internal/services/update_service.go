package services

import (
	"errors"
	"fmt"
)

// Update represents the structure of an update.
type Update struct {
	Version string
	Notes   string
}

// UpdateService provides methods for checking and applying updates.
type UpdateService struct {
	currentVersion string
	availableUpdates []Update
}

// NewUpdateService creates a new instance of UpdateService.
func NewUpdateService(currentVersion string, availableUpdates []Update) *UpdateService {
	return &UpdateService{
		currentVersion: currentVersion,
		availableUpdates: availableUpdates,
	}
}

// CheckForUpdates checks if there are any available updates.
func (us *UpdateService) CheckForUpdates() ([]Update, error) {
	if len(us.availableUpdates) == 0 {
		return nil, errors.New("no updates available")
	}

	var updates []Update
	for _, update := range us.availableUpdates {
		if update.Version != us.currentVersion {
			updates = append(updates, update)
		}
	}

	if len(updates) == 0 {
		return nil, errors.New("no new updates available")
	}

	return updates, nil
}

// ApplyUpdate applies the specified update.
func (us *UpdateService) ApplyUpdate(version string) error {
	for _, update := range us.availableUpdates {
		if update.Version == version {
			us.currentVersion = version
			return nil
		}
	}
	return fmt.Errorf("update version %s not found", version)
}