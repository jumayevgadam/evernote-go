// This package keeps abstract models, models with pagination and most used responses.

package abstract

// Res model is used when responsing data.
type Res struct {
	Error  error  `json:"error"`
	Res    any    `json:"response"`
	Msg    string `json:"msg"`
	Status string `json:"status"`
}
