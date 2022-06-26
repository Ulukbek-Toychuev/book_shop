package entity

type User struct {
	Id           int    `json:"-" db:"id"`
	UserName     string `json:"username" binding:"required"`
	UserSurName  string `json:"usersurname" binding:"required"`
	UserPassword string `json:"password" binding:"required"`
}
