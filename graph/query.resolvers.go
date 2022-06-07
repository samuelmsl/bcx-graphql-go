package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bcx/graph/generated"
	"bcx/graph/model"
	"context"
)

func (r *queryResolver) Items(ctx context.Context, input *model.FilterItem) ([]*model.Item, error) {
	var output []*model.Item
	var optionItem []*model.Option
	var categoryOption []*model.Category
	output = nil
	optionItem = nil
	categoryOption = nil

	//with id
	if input != nil {
		if input.ID != nil {
			//check is id equals
			for _, item := range r.items {
				if *input.ID == *item.ID {
					optionItem = nil
					item.Option = nil
					for _, option := range r.options {

						for _, itemMap := range r.itemMaps {
							if *itemMap.ItemID == *item.ID && *itemMap.OptionID == *option.ID {
								categoryOption = nil
								option.Category = nil
								for _, category := range r.categories {
									if *category.OptionID == *option.ID {
										categoryOption = append(categoryOption, category)
									}
								}
								option.Category = append(option.Category, categoryOption...)
								optionItem = append(optionItem, option)
							}
						}
					}
					item.Option = append(item.Option, optionItem...)
					output = append(output, item)
				}

			}
			return output, nil
		}
	}

	//general get
	for _, item := range r.items {
		for _, option := range r.options {

			for _, itemMap := range r.itemMaps {
				if *itemMap.ItemID == *item.ID && *itemMap.OptionID == *option.ID {
					categoryOption = nil
					option.Category = nil
					for _, category := range r.categories {
						if *category.OptionID == *option.ID {
							categoryOption = append(categoryOption, category)
						}
					}
					option.Category = append(option.Category, categoryOption...)
					optionItem = append(optionItem, option)
				}
			}
		}
		item.Option = append(item.Option, optionItem...)
		output = append(output, item)

		optionItem = nil
		item.Option = nil
	}

	return output, nil
}

func (r *queryResolver) Options(ctx context.Context, input *model.FilterOption) ([]*model.Option, error) {
	var output []*model.Option
	var categoryOption []*model.Category
	//with id
	if input != nil {
		if input.ID != nil {
			for _, option := range r.options {
				option.Category = nil
				categoryOption = nil
				if *input.ID == *option.ID {
					for _, category := range r.categories {
						if *category.OptionID == *option.ID {
							categoryOption = append(categoryOption, category)
						}
					}
					option.Category = append(option.Category, categoryOption...)
					output = append(output, option)
				}
			}
			return output, nil
		}
	}
	//general get
	for _, option := range r.options {
		option.Category = nil
		categoryOption = nil
		for _, category := range r.categories {
			if *category.OptionID == *option.ID {
				categoryOption = append(categoryOption, category)
			}
		}
		option.Category = append(option.Category, categoryOption...)
		output = append(output, option)
	}

	return output, nil
}

func (r *queryResolver) Categories(ctx context.Context, input *model.FilterCategory) ([]*model.Category, error) {
	var output []*model.Category
	//with id
	if input != nil {
		if input.ID != nil {
			for _, category := range r.categories {
				if *input.ID == *category.ID {
					output = append(output, category)
				}
			}
		}
	}

	return r.categories, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
