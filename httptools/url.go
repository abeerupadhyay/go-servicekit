package httptools

import "net/url"

func AddParamsToURL(u *url.URL, params map[string]string) *url.URL {
	existing := u.Query()
	for key, val := range params {
		existing.Set(key, val)
	}
	u.RawQuery = existing.Encode()
	return u
}

func AddParamsToURLString(rawURL string, params map[string]string) (*url.URL, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	return AddParamsToURL(u, params), nil
}
