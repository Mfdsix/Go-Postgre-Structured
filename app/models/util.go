package models

func MapPostToJSON(p *DPost) JPost {
	return JPost{
		ID:      p.ID,
		Title:   p.Title,
		Content: p.Content,
		Author:  p.Author,
	}
}
