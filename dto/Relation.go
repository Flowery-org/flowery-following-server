package dto

type CreateRelation struct {
	FollowerId  string `json:"followerId"`
	FollowingId string `json:"followingId"`
	CreatedAt   string
	//* No updated at, since only 'create' and 'delete' exists in relationship operation.
}

type DeleteRelation struct {
	FollowerId  string `json:"followerId"`
	FollowingId string `json:"followingId"`
}
