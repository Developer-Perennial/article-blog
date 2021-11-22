package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

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

func (a Article) MarshalJSON() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Article) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	srcBytes, ok := src.([]byte)
	if !ok {
		return errors.New("only support []byte type")
	}
	return json.Unmarshal(srcBytes, a)
}

func (a *Article) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	return json.Marshal(a)
}
