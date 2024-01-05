package relationship

type objectRegistrationStatus int

const (
	statusUnregistered objectRegistrationStatus = iota
	statusArchived
	statusActive
)
