package domain


type Friendship struct {
	User1Id string `json:"user1_id"`
	User2Id string `json:"user2_id"`
}

type Mutuals struct {
	User1Id   string `json:"user1_id"`
	User2Name string `json:"user2_name"`
}


