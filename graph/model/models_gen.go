// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Category struct {
	ID              *int     `json:"id"`
	CategoryName    *string  `json:"categoryName"`
	OptionID        *int     `json:"optionId"`
	AdditionalPrice *float64 `json:"additionalPrice"`
}

type FilterCategory struct {
	ID           *int    `json:"id"`
	CategoryName *string `json:"CategoryName"`
}

type FilterItem struct {
	ID   *int    `json:"id"`
	Name *string `json:"Name"`
}

type FilterOption struct {
	ID         *int    `json:"id"`
	OptionName *string `json:"OptionName"`
}

type Item struct {
	ID          *int      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Option      []*Option `json:"option"`
}

type ItemMap struct {
	ID       *int `json:"id"`
	ItemID   *int `json:"itemId"`
	OptionID *int `json:"optionId"`
}

type NewCategory struct {
	ID              int      `json:"id"`
	CategoryName    string   `json:"categoryName"`
	AdditionalPrice *float64 `json:"additionalPrice"`
	OptionID        *int     `json:"optionId"`
}

type NewItem struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Options     *NewItemMap `json:"options"`
}

type NewItemMap struct {
	ItemID    *int   `json:"itemId"`
	OptionIds []*int `json:"optionIds"`
}

type NewOption struct {
	ID         int    `json:"id"`
	OptionName string `json:"optionName"`
}

type Option struct {
	ID         *int        `json:"id"`
	OptionName *string     `json:"optionName"`
	Category   []*Category `json:"category"`
}
