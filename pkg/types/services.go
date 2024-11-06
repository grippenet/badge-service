package types

type BadgeServices struct {
	Pioneer PioneerService
}

type PioneerService interface {
	Check(instanceID string, studyKey string, postalCode string) (bool, error)
}