package axios

// RequestConfig represents axios [RequestConfig](https://github.com/axios/axios#request-config)
type RequestConfig struct {
	URL              *string                 `json:"url,omitempty"`
	Method           *string                 `json:"method,omitempty"`
	BaseURL          *string                 `json:"baseURL,omitempty"`
	Headers          *map[string]string      `json:"headers,omitempty"`
	Params           *map[string]interface{} `json:"params,omitempty"`
	Data             *map[string]interface{} `json:"data,omitempty"`
	Timeout          *int                    `json:"timeout,omitempty"`
	WithCredentials  *bool                   `json:"withCredentials,omitempty"`
	ResponseType     *string                 `json:"responseType,omitempty"`
	XSRFCookieName   *string                 `json:"xsrfCookieName,omitempty"`
	XSRFHeaderName   *string                 `json:"xsrfHeaderName,omitempty"`
	MaxContentLength *int                    `json:"maxContentLength,omitempty"`
	MaxRedirects     *int                    `json:"maxRedirects,omitempty"`

	Auth *struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	} `json:"auth,omitempty"`
}

// Response represents axios [Response](https://github.com/axios/axios#response-schema)
type Response struct {
	Data       interface{}       `json:"data"`
	Status     int               `json:"status"`
	StatusText string            `json:"statusText"`
	Headers    map[string]string `json:"headers"`
	Config     map[string]string `json:"config"`
}

type lambdaError struct {
	ErrorMessage string `json:"errorMessage"`
	ErrorType    string `json:"errorType"`
}
