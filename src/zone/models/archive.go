package models

type ArticleItem struct {
	CreatedMouthDay string
	Title string
	URL string
}

type MouthItem struct {
	Grouper string
	Nums int64
	Article []ArticleItem
}

type Year struct {
	Grouper string
	Mouths []MouthItem
}

type ArchiveResp struct {
	ArchiveResp []Year
}