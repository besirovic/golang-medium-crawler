package medium

// Scrape function is responsible for crawling data from medium.com
func Scrape() {
	authors := []string{
		"dan_abramov",
		"alexmngn",
		"francescod_ales",
		"nicolascole77",
		"richienorton",
		"JoubranJad",
	}

	slugC := make(chan slugResponse)
	for _, a := range authors {
		go getAuthorProfile(a, slugC)
	}

	for s := range slugC {
		go getArticle(s.author, s.slug)
	}
}
