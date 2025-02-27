package api

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"

	"github.com/stashapp/stash-box/pkg/manager/config"
	"github.com/stashapp/stash-box/pkg/models"
)

func (r *queryResolver) FindScene(ctx context.Context, id uuid.UUID) (*models.Scene, error) {
	fac := r.getRepoFactory(ctx)
	qb := fac.Scene()

	return qb.Find(id)
}

func (r *queryResolver) FindSceneByFingerprint(ctx context.Context, fingerprint models.FingerprintQueryInput) ([]*models.Scene, error) {
	fac := r.getRepoFactory(ctx)
	qb := fac.Scene()

	return qb.FindByFingerprint(fingerprint.Algorithm, fingerprint.Hash)
}

func (r *queryResolver) FindScenesByFingerprints(ctx context.Context, fingerprints []string) ([]*models.Scene, error) {
	if len(fingerprints) > 100 {
		return nil, errors.New("Too many fingerprints")
	}

	fac := r.getRepoFactory(ctx)
	qb := fac.Scene()

	return qb.FindByFingerprints(fingerprints)
}

func (r *queryResolver) FindScenesByFullFingerprints(ctx context.Context, fingerprints []*models.FingerprintQueryInput) ([]*models.Scene, error) {
	if len(fingerprints) > 100 {
		return nil, errors.New("Too many fingerprints")
	}

	fac := r.getRepoFactory(ctx)
	qb := fac.Scene()

	if config.GetPHashDistance() == 0 {
		var hashes []string
		for _, fp := range fingerprints {
			hashes = append(hashes, fp.Hash)
		}
		return qb.FindByFingerprints(hashes)
	}

	return qb.FindByFullFingerprints(fingerprints)
}

func (r *queryResolver) QueryScenes(ctx context.Context, input models.SceneQueryInput) (*models.SceneQuery, error) {
	return &models.SceneQuery{
		Filter: input,
	}, nil
}

type querySceneResolver struct{ *Resolver }

func (r *querySceneResolver) Count(ctx context.Context, obj *models.SceneQuery) (int, error) {
	fac := r.getRepoFactory(ctx)
	qb := fac.Scene()
	return qb.QueryCount(obj.Filter)
}

func (r *querySceneResolver) Scenes(ctx context.Context, obj *models.SceneQuery) ([]*models.Scene, error) {
	fac := r.getRepoFactory(ctx)
	qb := fac.Scene()
	return qb.QueryScenes(obj.Filter)
}
