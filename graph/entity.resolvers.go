package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"repath.io/data"
	"repath.io/graph/generated"
	"repath.io/graph/model"
)

func (r *entityResolver) FindOrganizationByID(ctx context.Context, id string) (*model.Organization, error) {
	people, err := data.Store().FindByOrganizationId(id)
	if err != nil {
		return nil, err
	}
	out := make([]*model.Person, 0)
	for _, v := range people {
		out = append(out, &model.Person{
			ID:        v.Id,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Title:     *v.Title,
		})
	}
	return &model.Organization{ID: "id", People: out}, nil
}

func (r *entityResolver) FindPersonByID(ctx context.Context, id string) (*model.Person, error) {
	people, err := data.Store().FindById(id)
	if err != nil {
		return nil, err
	}
	out := make([]*model.Person, 0)
	for _, v := range people {
		out = append(out, &model.Person{
			ID:        v.Id,
			FirstName: v.FirstName,
			LastName:  v.LastName,
		})
	}
	return out[0], nil
}

func (r *entityResolver) FindTeamByID(ctx context.Context, id string) (*model.Team, error) {
	people, err := data.Store().FindByTeamId(id)
	if err != nil {
		return nil, err
	}
	out := make([]*model.Person, 0)
	for _, v := range people {
		out = append(out, &model.Person{
			ID:        v.Id,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Title:     *v.Title,
		})
	}
	return &model.Team{ID: "id", People: out}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
