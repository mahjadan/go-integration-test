package domain

type People struct {
	Name      string `json:"name" bson:"name"`
	Height    string `json:"height" bson:"height"`
	Mass      string `json:"mass" bson:"mass"`
	HairColor string `json:"hair_color" bson:"hair_color"`
	SkinColor string `json:"skin_color" bson:"skin_color"`
	EyeColor  string `json:"eye_color" bson:"eye_color"`
	BirthYear string `json:"birth_year" bson:"birth_year"`
	Gender    string `json:"gender" bson:"gender"`
}
