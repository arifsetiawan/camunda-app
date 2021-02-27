package camunda

// IdentityVerifyRequest ...
type IdentityVerifyRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// IdentityVerifyResponse ...
type IdentityVerifyResponse struct {
	AuthenticatedUser string `json:"authenticatedUser"`
	Authenticated     bool   `json:"authenticated"`
}

// ListTaskRequest ...
type ListTaskRequest struct {
	Assignee string `json:"assignee"`
}

// UserTask ...
type UserTask struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Assignee          string `json:"assignee"`
	Created           string `json:"created"`
	ProcessInstanceID string `json:"processInstanceId"`
}

// FetchAndLockTopic ...
type FetchAndLockTopic struct {
	Name         string `json:"topicName"`
	LockDuration int    `json:"lockDuration"`
}

// FetchAndLockRequest ...
type FetchAndLockRequest struct {
	WorkerID string              `json:"workerId"`
	MaxTasks int                 `json:"maxTasks"`
	Topics   []FetchAndLockTopic `json:"topics"`
}

// ExternalTask ...
type ExternalTask struct {
	ID        string              `json:"id"`
	TopicName string              `json:"topicName"`
	Variables map[string]Variable `json:"variables"`
}

// Variable ...
type Variable struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

// CompleteExternalTaskRequest ...
type CompleteExternalTaskRequest struct {
	WorkerID  string              `json:"workerId"`
	Variables map[string]Variable `json:"variables"`
}

// UserProfileResponse ...
type UserProfileResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}
