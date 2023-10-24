package company

import (
	"context"
	"strings"

	"github.com/muhwyndhamhp/spill/db"
	"github.com/muhwyndhamhp/spill/utils/errs"
)

type CompanyRepository struct {
	db *db.Queries
}

func NewCompanyRepository(db *db.Queries) *CompanyRepository {
	return &CompanyRepository{
		db: db,
	}
}

func (r *CompanyRepository) CompanyWithServiceIDExists(ctx context.Context, serviceID string) (bool, error) {
	exists, err := r.db.CompanyByServiceIDExist(ctx, serviceID)
	if err != nil {
		return false, errs.Wrap(err)
	}

	return exists, nil
}

func (r *CompanyRepository) UpsertCompany(ctx context.Context, name string, userID int64) error {
	existing, _ := r.db.GetCompanyByName(ctx, strings.ToLower(name))

	companyID := int64(0)
	if existing.ID == 0 {
		cid, err := r.db.CreateCompany(ctx, name)
		if err != nil {
			return errs.Wrap(err)
		}

		companyID = cid
	} else {
		companyID = existing.ID
	}

	err := r.db.CreateCompanySpillUser(ctx, db.CreateCompanySpillUserParams{
		SpillUserID: userID,
		CompanyID:   companyID,
	})
	if err != nil {
		return errs.Wrap(err)
	}

	return nil
}
