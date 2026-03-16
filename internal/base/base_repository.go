package base

import "gorm.io/gorm"

type BaseRepository[T any] struct {
	DB *gorm.DB
}

func (r *BaseRepository[T]) Create(data *T) error {
	return r.DB.Create(data).Error
}

func (r *BaseRepository[T]) FindAll(out *[]T) error {
	return r.DB.Find(out).Error
}

func (r *BaseRepository[T]) FindByID(id interface{}, out *T) error {
	return r.DB.First(out, id).Error
}

func (r *BaseRepository[T]) Update(data *T) error {
	return r.DB.Save(data).Error
}

func (r *BaseRepository[T]) Delete(id interface{}) error {
	return r.DB.Delete(new(T), id).Error
}

func (r *BaseRepository[T]) FindWhere(cond interface{}, out *[]T) error {
	return r.DB.Where(cond).Find(out).Error
}

func (r *BaseRepository[T]) Paginate(limit, offset int, out *[]T) error {
	return r.DB.Limit(limit).Offset(offset).Find(out).Error
}