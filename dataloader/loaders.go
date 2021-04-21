package loaders

import (
	"time"

	"github.com/holyshared/go-graphql-server/graph/model"
	models "github.com/holyshared/go-graphql-server/model"
	"gorm.io/gorm"
)

const loadersKey = "dataloaders"

type Loaders struct {
	UserById UserLoader
}

func NewDefaultUserLoader(db *gorm.DB) *UserLoader {
	config := &UserLoaderConfig{
		Fetch: func(keys []string) ([]*model.User, []error) {
			users := make([]*model.User, len(keys))
			errors := make([]error, len(keys))

			// Fetch form data store
			macthedUsers := []*models.User{}
			db.Find(&macthedUsers, keys)

			macthedUserRegistry := make(map[string]*models.User, len(macthedUsers))
			for _, macthedUser := range macthedUsers {
				macthedUserRegistry[macthedUser.ID] = macthedUser
			}

			for i, key := range keys {
				matched, ok := macthedUserRegistry[key]
				if !ok {
					continue
				}
				users[i] = &model.User{ID: matched.ID, Name: matched.Name}
			}
			return users, errors
		},
		MaxBatch: 100,
		Wait:     2 * time.Millisecond,
	}

	return &UserLoader{
		fetch:    config.Fetch,
		wait:     config.Wait,
		maxBatch: config.MaxBatch,
	}
}
