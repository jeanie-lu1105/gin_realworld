package storage

import "context"

func ListPopularTags(ctx context.Context) ([]string, error) {
	var res []string
	err := gormDb.Table("popular_tags").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
