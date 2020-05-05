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

func (store *ArticleStoreInmemory) Update(idconv int, title string, body string) error {

	var articles = store.ArticleMap
	articles[idconv-1] = Article{ID: idconv, Title: title, Body: body}
	return nil
}

func (store *ArticleStoreInmemory) Delete(idconv int) error {

	store.ArticleMap = append(store.ArticleMap[:idconv-1], store.ArticleMap[idconv:]...)
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
