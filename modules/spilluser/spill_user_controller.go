package spilluser

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/muhwyndhamhp/spill/db"
	"github.com/muhwyndhamhp/spill/utils/errs"
)

type SpillUserController struct {
	repo *SpillUserRepository
}

func NewSpillUserController(repo *SpillUserRepository) *SpillUserController {
	return &SpillUserController{
		repo: repo,
	}
}

func (c *SpillUserController) GetUserByServiceID(ctx context.Context, serviceID string) (*db.SpillUser, error) {
	usr, err := c.repo.GetUserByServiceID(ctx, serviceID)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	return usr, nil
}

func (c *SpillUserController) CreateSpillUser(ctx context.Context, alias, bio, serviceID string) (id int64, err error) {
	id, err = c.repo.CreateSpillUser(ctx, &db.SpillUser{
		Alias: alias,
		Bio: pgtype.Text{
			String: bio,
			Valid:  true,
		},
		ServiceID: serviceID,
	})
	if err != nil {
		return 0, errs.Wrap(err)
	}

	return id, nil
}
