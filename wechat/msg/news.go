package msg

type NewsMsg struct {
	ArticleCount int        `xml:"ArticleCount"`
	Articles     []*Article `xml:"Articles>item,omitempty"`
}

type Article struct {
	Title       string `xml:"Title,omitempty"`
	Description string `xml:"Description,omitempty"`
	PicURL      string `xml:"PicUrl,omitempty"`
	URL         string `xml:"Url,omitempty"`
}

func NewArticle(t, d, p, u string) *Article {
	a := &Article{Title: t, Description: d, PicURL: p, URL: u}
	return a
}
