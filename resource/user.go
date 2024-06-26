package resource

// User implements the user object
type User struct {
	Username         string    `json:"username"`
	PresentationName string    `json:"presentation_name"`
	Origin           string    `json:"origin"`
	Metadata         *Metadata `json:"metadata"`
	Resource         `json:",inline"`
}

// UserCreate is used to create a new user in the Cloud Controller database
//
// Creating a user requires one value, a GUID. This creates a user in the Cloud Controller database.
// Generally, the GUID should match the GUID of an already-created
// user in the UAA database, though this is not required.
type UserCreate struct {
	GUID     string    `json:"guid"`
	Metadata *Metadata `json:"metadata,omitempty"`
}

type UserUpdate struct {
	Metadata *Metadata `json:"metadata,omitempty"`
}

type UserList struct {
	Pagination Pagination `json:"pagination"`
	Resources  []*User    `json:"resources"`
}
