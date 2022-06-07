package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bcx/graph/generated"
	"bcx/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateItem(ctx context.Context, input model.NewItem) (*model.Item, error) {
	fmt.Println(input)

	item := &model.Item{
		ID:          &input.ID,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}

	for _, option := range input.Options.OptionIds {
		itemMap := &model.ItemMap{
			ItemID:   &input.ID,
			OptionID: option,
		}
		r.itemMaps = append(r.itemMaps, itemMap)
	}

	r.items = append(r.items, item)

	return item, nil
}

func (r *mutationResolver) CreateOption(ctx context.Context, input model.NewOption) (*model.Option, error) {
	option := &model.Option{
		OptionName: &input.OptionName,
		ID:         &input.ID,
	}

	r.options = append(r.options, option)
	return option, nil
}

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	category := &model.Category{
		ID:              &input.ID,
		CategoryName:    &input.CategoryName,
		AdditionalPrice: input.AdditionalPrice,
		OptionID:        input.OptionID,
	}

	r.categories = append(r.categories, category)
	return category, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
