package dbcontext

import (
	"gorm.io/gorm"
)

func InitSearch[TReturn any, TSearchFields any]() func() *Search[TReturn, TSearchFields] {
	return func() *Search[TReturn, TSearchFields] {
		search := new(Search[TReturn, TSearchFields])
		search.size = 10

		search.globalOperands = &Operands{}

		searchFields := new(TSearchFields)
		search.fieldOperands = searchFields

		SetFieldCodes(searchFields)

		search.filters = make([]func(operands *Operands, fields *TSearchFields) *GLobalFilter, 0)
		return search
	}
}

type Search[TReturn any, TSearchFields any] struct {
	// Set filters.
	filters []func(operands *Operands, fields *TSearchFields) *GLobalFilter

	// Limit response.
	size int

	// Operands.
	globalOperands *Operands

	// Fields with operands.
	fieldOperands interface{}
}

// Models actions interface.
type ISearch[T interface{}, TReturnEntity interface{}] interface {
	// Use filters.
	Where(filterFunction func(operands *Operands, fields *T) *GLobalFilter) *Search[T, TReturnEntity]

	// Set querry limit. Default: 10.
	Size(count int) *Search[T, TReturnEntity]

	// Get all rows. (Limit by size, default 10)
	All() []TReturnEntity

	// Get first matched record.
	First() *TReturnEntity
}

// Use filters.
func (sb *Search[TReturn, TSearchFields]) Where(filterFunction func(operands *Operands, fields *TSearchFields) *GLobalFilter) *Search[TReturn, TSearchFields] {
	sb.filters = append(sb.filters, filterFunction)
	return sb
}

// Set querry limit. Default: 10.
func (sb *Search[T, TSearchFields]) Size(count int) *Search[T, TSearchFields] {
	sb.size = count
	return sb
}

// Get all rows. (Limit by size, default 10)
func (sb *Search[TReturn, TSearchFields]) All() ([]TReturn, error) {
	var err error
	var entities []TReturn

	db, err := sb.arrangeFilters()

	WithPreloads[TReturn](db)

	if err == nil {
		db.Limit(sb.size).Find(&entities)
	}

	return entities, err
}

// Get first matched record.
func (sb *Search[TReturn, TSearchFields]) First() (*TReturn, error) {
	var err error
	var entity *TReturn

	db, err := sb.arrangeFilters()

	WithPreloads[TReturn](db)

	if err == nil {
		result := db.First(&entity)
		if result.Error != nil {
			entity = nil
		}
	} else {
		entity = nil
	}

	return entity, err
}

// Arrange all filters declared by "Where" function chains to common filter.
func (sb *Search[TReturn, TSearchFields]) arrangeFilters() (*gorm.DB, error) {
	operands, _ := sb.fieldOperands.(*TSearchFields)
	var err error
	db := getContextModel().db

	if len(sb.filters) == 0 {
		return db, nil
	}

	for _, filterFunction := range sb.filters {
		db, err = filterFunction(sb.globalOperands, operands).arrangeFilters(db)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
