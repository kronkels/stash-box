package api

import (
	"context"
	"time"

	"github.com/stashapp/stash-box/pkg/dataloader"
	"github.com/stashapp/stash-box/pkg/models"
)

type studioResolver struct{ *Resolver }

func (r *studioResolver) ID(ctx context.Context, obj *models.Studio) (string, error) {
	return obj.ID.String(), nil
}

func (r *studioResolver) Urls(ctx context.Context, obj *models.Studio) ([]*models.URL, error) {
	return dataloader.For(ctx).StudioUrlsByID.Load(obj.ID)
}

func (r *studioResolver) Parent(ctx context.Context, obj *models.Studio) (*models.Studio, error) {
	if !obj.ParentStudioID.Valid {
		return nil, nil
	}

	qb := r.getRepoFactory(ctx).Studio()
	parent, err := qb.Find(obj.ParentStudioID.UUID)

	if err != nil {
		return nil, err
	}

	return parent, nil
}

func (r *studioResolver) ChildStudios(ctx context.Context, obj *models.Studio) ([]*models.Studio, error) {
	qb := r.getRepoFactory(ctx).Studio()
	children, err := qb.FindByParentID(obj.ID)

	if err != nil {
		return nil, err
	}

	return children, nil
}
func (r *studioResolver) Images(ctx context.Context, obj *models.Studio) ([]*models.Image, error) {
	imageIDs, err := dataloader.For(ctx).StudioImageIDsByID.Load(obj.ID)
	if err != nil {
		return nil, err
	}
	images, errors := dataloader.For(ctx).ImageByID.LoadAll(imageIDs)
	for _, err := range errors {
		if err != nil {
			return nil, err
		}
	}
	return images, nil
}

func (r *studioResolver) IsFavorite(ctx context.Context, obj *models.Studio) (bool, error) {
	jqb := r.getRepoFactory(ctx).Joins()
	user := getCurrentUser(ctx)
	return jqb.IsStudioFavorite(models.StudioFavorite{StudioID: obj.ID, UserID: user.ID})
}

func (r *studioResolver) Created(ctx context.Context, obj *models.Studio) (*time.Time, error) {
	return &obj.CreatedAt.Timestamp, nil
}

func (r *studioResolver) Updated(ctx context.Context, obj *models.Studio) (*time.Time, error) {
	return &obj.UpdatedAt.Timestamp, nil
}
