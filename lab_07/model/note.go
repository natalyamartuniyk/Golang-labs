//go:generate easyjson -all
package model

type Note struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Phone   string `json:"phone"`
}

func (n Note) IsPhoneValid() bool {
	return len(n.Phone) == 10
}
