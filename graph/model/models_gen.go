// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Node interface {
	IsNode()
}

type OrganizationResource interface {
	IsOrganizationResource()
}

type AttachPersonToOrganization struct {
	PersonID       string `json:"personId"`
	OrganizationID string `json:"organizationId"`
}

type CreatePerson struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Title     *string `json:"title"`
}

type DeletePerson struct {
	ID string `json:"id"`
}

type Organization struct {
	ID     string    `json:"id"`
	People []*Person `json:"people"`
}

func (Organization) IsEntity() {}

type Person struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Title     string `json:"title"`
}

func (Person) IsNode()                 {}
func (Person) IsEntity()               {}
func (Person) IsOrganizationResource() {}

type PersonDelete struct {
	Record  *Person `json:"record"`
	Success bool    `json:"success"`
}

type PersonUpdate struct {
	Previous *Person `json:"previous"`
	Current  *Person `json:"current"`
	Success  bool    `json:"success"`
}

type Team struct {
	ID     string    `json:"id"`
	People []*Person `json:"people"`
}

func (Team) IsEntity() {}

type UpdatePerson struct {
	ID        string  `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Title     *string `json:"title"`
}
