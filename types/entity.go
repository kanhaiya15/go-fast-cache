package types

// Entity below here!
type Entity struct {
	ID   uint64 `form:"id" json:"id" xml:"id"  binding:"required"`
	Name string `form:"name" json:"name" xml:"name" binding:"required"`
}
