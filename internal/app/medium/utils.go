package medium

// MediumURL is root domain of medium.com
const MediumURL = "https://medium.com/"

// constructMediumAuthorURL returns url for author profile
func constructMediumAuthorURL(username string) string {
	return MediumURL + "@" + username + "/latest?format=json"
}

// constructMediumArticleURL returns url for article
func constructMediumArticleURL(username string, slug string) string {
	return MediumURL + "@" + username + "/" + slug + "?format=json"
}
