package relysia

type Headers map[string]string

func (self *Client) POSTHeaders(additionalHeaders ...map[string]string) (d Headers) {
	d = Headers{
		"accept":       "*/*",
		"Content-Type": "application/json",
	}
	if len(self.authToken) > 0 {
		d["authToken"] = self.authToken
	}
	if len(self.serviceID) > 0 {
		d["serviceID"] = self.serviceID
	}
	if len(additionalHeaders) > 0 {
		for k, v := range additionalHeaders[0] {
			d[k] = v
		}
	}
	return d
}

func (self *Client) GETHeaders(additionalHeaders ...map[string]string) (d Headers) {
	d = Headers{
		"accept": "*/*",
	}
	if len(self.authToken) > 0 {
		d["authToken"] = self.authToken
	}
	if len(self.serviceID) > 0 {
		d["serviceID"] = self.serviceID
	}
	if len(additionalHeaders) > 0 {
		for k, v := range additionalHeaders[0] {
			d[k] = v
		}
	}
	return d
}
