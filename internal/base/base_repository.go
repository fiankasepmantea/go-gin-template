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

func (r *BaseRepository[T]) FindOneBy(column string, value interface{}, out *T) error {
	return r.DB.Where(column+" = ?", value).First(out).Error
}

func (r *BaseRepository[T]) FindByWhere(cond interface{}, out *[]T) error {
	return r.DB.Where(cond).Find(out).Error
}

func (r *BaseRepository[T]) FindOneWhere(cond interface{}, out *T) error {
	return r.DB.Where(cond).First(out).Error
}

func (r *BaseRepository[T]) Update(data *T) error {
	return r.DB.Save(data).Error
}

func (r *BaseRepository[T]) UpdatesByID(id interface{}, data interface{}, model *T) error {
	return r.DB.Model(model).Where("id = ?", id).Updates(data).Error
}

func (r *BaseRepository[T]) UpdatesWhere(cond interface{}, data interface{}, model *T) error {
	return r.DB.Model(model).Where(cond).Updates(data).Error
}

func (r *BaseRepository[T]) Delete(id interface{}) error {
	return r.DB.Delete(new(T), id).Error
}

func (r *BaseRepository[T]) DeleteWhere(cond interface{}, model *T) error {
	return r.DB.Where(cond).Delete(model).Error
}

func (r *BaseRepository[T]) Count(model *T) (int64, error) {
	var total int64
	err := r.DB.Model(model).Count(&total).Error
	return total, err
}

func (r *BaseRepository[T]) CountWhere(cond interface{}, model *T) (int64, error) {
	var total int64
	err := r.DB.Model(model).Where(cond).Count(&total).Error
	return total, err
}

func (r *BaseRepository[T]) Paginate(limit, offset int, out *[]T) error {
	return r.DB.Limit(limit).Offset(offset).Find(out).Error
}

func (r *BaseRepository[T]) InsertBatch(data *[]T) error {
	return r.DB.Create(data).Error
}