package entity

type Article struct {
	Base
	Title   string `json:"title" validator:"required" gorm:"column:title"`
	Content string `json:"content" validator:"required" gorm:"column:content"`
	Author  string `json:"author" validator:"Required" gorm:"column:author"`
}

func NewArticle(title, content, author string) *Article {
	return &Article{
		Title:   title,
		Content: content,
		Author:  author,
	}
}