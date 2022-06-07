package graph

import "bcx/graph/model"

type Resolver struct {
	items      []*model.Item
	options    []*model.Option
	categories []*model.Category
	itemMaps   []*model.ItemMap
}
