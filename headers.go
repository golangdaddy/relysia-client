package relysia

type Headers map[string]string

func (self *Client) defaultHeaders(additionalHeaders ...map[string]string) (d Headers) {
	d = map[string]string{
		"accept":       "*/*",
		"Content-Type": "application/json",
	}
	if len(additionalHeaders) > 0 {
		if len(self.accessToken) > 0 {
			d["accessToken"] = self.accessToken
		}
		for k, v := range additionalHeaders[0] {
			d[k] = v
		}
	}
	return d
}
