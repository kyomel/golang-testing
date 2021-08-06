package repository

import "unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
