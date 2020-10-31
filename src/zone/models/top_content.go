package models

// 顶部文章内容
type TopContent struct {
	TopContentClass string     `json:"top_content_class"`
	TopContentName  string     `json:"top_content_name"`
	ContentNum      int64      `json:"content_num"`
	ContentText     string     `json:"content_text"`
	TopArticle      []*Article `json:"top_article"`
}
