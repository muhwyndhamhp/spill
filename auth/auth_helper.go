package auth

import (
	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/muhwyndhamhp/spill/utils/errs"
)

type Session struct {
	Claim *clerk.SessionClaims
}

func (sc *Session) GetUser(client clerk.Client) (*clerk.User, error) {
	s, err := client.Sessions().Read(sc.Claim.SessionID)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	u, err := client.Users().Read(s.UserID)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	return u, nil
}
