package databases

type Metadata struct {
	Name   string `json:"name" bson:"name"`
	Artist string `json:"artist" bson:"artist"`
	Album  string `json:"album" bson:"album"`
}

type URLs struct {
	Thumbnail string `json:"thumbnail" bson:"thumbnail"`
	Audio     string `json:"audio" bson:"audio"`
}

type Song struct {
	Metadata Metadata `json:"metadata" bson:"metadata"`
	ID       string   `json:"id" bson:"id"`
	URLs     URLs     `json:"url" bson:"url"`
}

type User struct {
	ID       string `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	SongList []Song `json:"songlist" bson:"songlist"`
}
