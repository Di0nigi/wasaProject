package api

type ObjId struct {
	IDObj string `json:"idObj"`
}

type UserId struct {
	IDUser string `json:"idUser"`
}

type User struct {
	ID           UserId   `json:"id"`
	Followers    []UserId `json:"followers"`
	Follows      []UserId `json:"follows"`
	Posts        []ObjId  `json:"Posts"`
	BlockedArray []UserId `json:"blockedArray"`
}

type PostedImage struct {
	IDPhoto     ObjId     `json:"idPhoto"`
	Owner       UserId    `json:"owner"`
	Image       string    `json:"image"`
	Likes       int       `json:"likes"`
	Comments    []Comment `json:"comments"`
	NumComments int       `json:"numComments"`
	Date        string    `json:"Date"`
}

type Comment struct {
	IDComment ObjId  `json:"idComment"`
	Content   string `json:"content"`
	Owner     UserId `json:"owner"`
	Photo     ObjId  `json:"photo"`
}

type Like struct {
	IDLike  ObjId  `json:"idLike"`
	Owner   UserId `json:"owner"`
	ToPhoto ObjId  `json:"toPhoto"`
}
