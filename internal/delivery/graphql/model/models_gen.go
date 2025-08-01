// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type Order struct {
	ID        string  `json:"id"`
	UserID    string  `json:"userId"`
	Amount    float64 `json:"amount"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

type Query struct {
}

type Subscription struct {
}

type Time struct {
	UnixTime  int32  `json:"unixTime"`
	Timestamp string `json:"timestamp"`
}

type User struct {
	ID        string   `json:"id"`
	Email     string   `json:"email"`
	Age       int32    `json:"age"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
	Orders    []*Order `json:"orders"`
}

type UserCreateInput struct {
	Email string `json:"email"`
	Age   int32  `json:"age"`
}
