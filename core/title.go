package core

// todo add category and category rank
type Title struct {
	Id       uint   `json:"id" gorm:"primary_key"`
	BggId    uint   `json:"bggid" gorm:"unique"`
	Name     string `json:"name"`
	Designer string `json:"designer"`
	BggRank  uint   `json:"bggrank"`
}
