package tags

type TagController struct {
	Collection string
}

func NewTagController() *TagController {
	return &TagController{
		Collection: "tags",
	}
}
