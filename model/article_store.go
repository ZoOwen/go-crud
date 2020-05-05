package model

type ArticleStoreInmemory struct {
	ArticleMap []Article
}

func NewArticleStoreInmemory() *ArticleStoreInmemory {
	return &ArticleStoreInmemory{
		ArticleMap: []Article{
			Article{ID: 1, Title: "ini judul", Body: "Hallo ini raif"},
		},
	}
}

func removeArticle(id int) *ArticleStoreInmemory {
	return &ArticleStoreInmemory{
		ArticleMap: []Article{
			Article{ID: 1, Title: "ini judul", Body: "Hallo ini raif"},
		},
	}
}

func (store *ArticleStoreInmemory) Update(article *Article) error {

	store.ArticleMap = append(store.ArticleMap, *article)
	return nil
}
func (store *ArticleStoreInmemory) Save(article *Article) error {
	lastID := len(store.ArticleMap)

	// set article id
	article.ID = lastID + 1

	// push to article map slices
	store.ArticleMap = append(store.ArticleMap, *article)

	return nil
}
