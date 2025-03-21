package handlers

import (
	"net/http"
	"github.com/vlastimircernos/go-update-server/internal/services"
)

// UpdateHandler handles HTTP requests for checking and applying updates.
func UpdateHandler(updateService *services.UpdateService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Check for updates
			updates, err := updateService.CheckForUpdates()
			if err != nil {
				http.Error(w, "Failed to check for updates", http.StatusInternalServerError)
				return
			}
			// Render update information (this part can be expanded with templates)
			w.Write([]byte("Updates available: " + updates))
		case http.MethodPost:
			// Apply update
			err := updateService.ApplyUpdate()
			if err != nil {
				http.Error(w, "Failed to apply update", http.StatusInternalServerError)
				return
			}
			w.Write([]byte("Update applied successfully"))
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}