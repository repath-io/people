// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewPerson struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Organization struct {
	ID     string    `json:"id"`
	People []*Person `json:"people"`
}

func (Organization) IsEntity() {}

type Person struct {
	ID           string        `json:"id"`
	FirstName    string        `json:"firstName"`
	LastName     string        `json:"lastName"`
	Organization *Organization `json:"organization"`
}

func (Person) IsEntity() {}
