package iconsdk

// IconService holds the URL to the ICON service.
type IconService struct {
	IconServiceURL string
}

// NewIconService creates a new instance of IconService with a default or specified URL.
func NewIconService(iconServiceURL *string) *IconService {
	defaultURL := "https://api.icon.community/api/v3"
	if iconServiceURL == nil {
		return &IconService{IconServiceURL: defaultURL}
	}
	return &IconService{IconServiceURL: *iconServiceURL}
}
