package seeker

// Seeker is an interface for fetching Instagram stories
type Seeker interface {
	// Get retrieves Instagram stories for a given user
	// It takes cookies and the username of the Instagram account as parameters
	// It returns a slice of Story representing the stories and an error if any
	Get(cookies, username string) ([]Story, error)
}
