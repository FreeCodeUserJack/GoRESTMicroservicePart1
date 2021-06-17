package oauth

import "time"

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	Expires     int64  `json:"expires"`
}

func (a *AccessToken) IsNotExpired() bool {
	return time.Unix(a.Expires, 0).UTC().Before(time.Now().UTC())
}