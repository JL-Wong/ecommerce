package model

type User struct {
    User_ID  string `json:"user_id"`
    Name     string `json:"name"`
    Password string `json:"password"`
    Role     string `json:"role"`
}
