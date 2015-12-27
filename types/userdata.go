package types

////////////////////////////////////////////////////////////////////////////////
// Configuration for RiveScript user data.
////////////////////////////////////////////////////////////////////////////////

type UservarsConfig struct {
	Username string            `json:"username"`
	Data     map[string]string `json:"data"`
}