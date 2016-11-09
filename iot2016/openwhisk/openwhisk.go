package openwhisk

import "net/http"

type OpenWhisk struct {
	BaseURL   string
	Token     string
	Namespace string

	Client http.Client
}

func New(token string, namespace string) *OpenWhisk {
	ow := OpenWhisk{
		"https://openwhisk.ng.bluemix.net/api/v1",
		token,
		namespace,
		http.Client{},
	}
	return &ow
}

func (ow *OpenWhisk) TriggerAction(action string, args string) {
	url := "<base-url>/namespaces/<namespace>/actions/<action>"
	//TODO
	//addHeader(req, ow.Token)
	//client := ow.Client
	//doRequest(req, client)
}

func doRequest(req *http.Request, client http.Client) {
	/* TODO
	Do a Request
	Read the response body & print it to the console
	*/
}

func addHeader(req *http.Request, token string) {
	auth := "Basic " + token
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", auth)
}
