package habit

type Habit struct {
	Id   int    `bson:"_id,omitempty"`
	Name string `bson:"name"`
}
