package article_service

import "errors"

var db map[int]*Article = make(map[int]*Article)

type Article struct {
	ID            int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
}

func (a *Article) Add() {
	if a.ID == 0 {
		a.ID = a.Count() + 1
	}
	db[a.ID] = a
}

func (a *Article) Edit() error {
	if _, ok := db[a.ID]; !ok {
		return errors.New("article not found")
	}
	e := db[a.ID]
	if a.Title != "" {
		e.Title = a.Title
	}
	if a.Desc != "" {
		e.Desc = a.Desc
	}
	if a.Content != "" {
		e.Content = a.Content
	}
	if a.CoverImageUrl != "" {
		e.CoverImageUrl = a.CoverImageUrl
	}
	return nil
}

func (a *Article) Get() (*Article, error) {
	f, ok := db[a.ID]
	if !ok {
		return nil, errors.New("article not found")
	}
	return f, nil
}

func (a *Article) GetAll() []*Article {
	list := make([]*Article, 0, len(db))
	for _, v := range db {
		list = append(list, v)
	}
	return list
}

func (a *Article) Delete() error {
	if _, ok := db[a.ID]; !ok {
		return errors.New("article not found")
	}
	delete(db, a.ID)
	return nil
}

func (a *Article) ExistByID() bool {
	_, ok := db[a.ID]
	return ok
}

func (a *Article) Count() int {
	return len(db)
}
