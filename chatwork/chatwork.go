// Package facebook provides constants for using OAuth2 to access Facebook.
package chatwork // import "github.com/eiel/golang-oauth2/chatwork"

import (
	"github.com/eiel/golang-oauth2"
)

// Endpoint is Chatwork's OAuth 2.0 endpoint
var Endpoint = oauth2.Endpoint{
	AuthURL:   "https://www.chatwork.com/packages/oauth2/login.php",
	TokenURL:  "https://oauth.chatwork.com/token",
	AuthStyle: oauth2.AuthStyleInHeader,
}
