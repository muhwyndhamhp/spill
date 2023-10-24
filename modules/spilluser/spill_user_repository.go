package spilluser

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/muhwyndhamhp/spill/db"
	"github.com/muhwyndhamhp/spill/utils/errs"
)

var ErrAliasUsed = errors.New("ErrAliasUsed: alias has been used")

type SpillUserRepository struct {
	db *db.Queries
}

func NewSpillUserRepository(db *db.Queries) *SpillUserRepository {
	return &SpillUserRepository{
		db: db,
	}
}

func (r *SpillUserRepository) CreateSpillUser(ctx context.Context, value *db.SpillUser) (id int64, err error) {
	now := time.Now()
	exists, _ := r.db.SpillUserByAliasExist(ctx, value.Alias)
	if exists {
		return 0, ErrAliasUsed
	}

	id, err = r.db.CreateSpillUser(ctx, db.CreateSpillUserParams{
		Alias:     value.Alias,
		ServiceID: value.ServiceID,
		Bio:       value.Bio,
		CreatedAt: pgtype.Timestamptz{
			Time:             now,
			InfinityModifier: pgtype.Finite,
			Valid:            true,
		},
		UpdatedAt: pgtype.Timestamptz{
			Time:             now,
			InfinityModifier: pgtype.Finite,
			Valid:            true,
		},
	})
	if err != nil {
		return 0, errs.Wrap(err)
	}

	return id, nil
}

func (r *SpillUserRepository) GetUserByID(ctx context.Context, id int64) (*db.SpillUser, error) {
	usr, err := r.db.GetSpillUserByID(ctx, id)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	return &usr, nil
}

func (r *SpillUserRepository) GetUserByServiceID(ctx context.Context, serviceID string) (*db.SpillUser, error) {
	usr, err := r.db.GetSpillUserByServiceID(ctx, serviceID)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	return &usr, nil
}
