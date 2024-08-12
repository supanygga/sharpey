package user

import "github.com/gofrs/uuid/v5"

// List.
type ListRequest struct {
	//Filters.
	Filters *Filters

	// Page.
	Page *Page

	// Sort
	Sort *Sort
}

// Filters.
type Filters struct {
	// ChannelID.
	ChannelID *uuid.UUID

	// ChannelName.
	ChannelName *string `validate:"omitempty,gt=0"`

	// SenderID.
	SenderID *uuid.UUID

	// SenderName.
	SenderName *string `validate:"omitempty,gt=0"`

	// ... .
}

// Page.
type Page struct {
	// Size.
	Size int `validate:"gt=0"`

	// Number
	Number int `validate:"gt=0"`
}

// Sort.
type Sort struct {
	// OrderBy.
	OrderBy string `validate:"required,oneof=id ..."`

	// Direction.
	Direction string `validate:"required,oneof=DESC ASC"`
}

// ListResponse.
type ListResponse struct {
	// Items.
	Items []*User

	// Total.
	Total int

	// Page.
	Page *Page
}
