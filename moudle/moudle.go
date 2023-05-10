package moudle

// Todo 定义结构体model
type Todo struct {
	//gorm.Model
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}
