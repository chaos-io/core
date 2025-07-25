package core

const UrlTypeName = "Url"
const UrlTypeFullName = "chaos.core.Url"

func NewUrl(url string) (*Url, error) {
	u, err := ParseUrl(url)
	if err != nil {
		return nil, err
	}
	return u, nil
}
