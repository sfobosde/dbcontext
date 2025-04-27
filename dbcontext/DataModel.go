package dbcontext

import "github.com/gofrs/uuid"

// Data model entitiy structure.
type DataModel[T any, TSearchFields any] struct {
	entity *T

	// Using ORM. Contains filter, limit and fetchind actions.
	Search func() *Search[T, TSearchFields]

	// Create new element.
	Create func() *T
}

// DataModel constructor.
func InitDataModel[TDataModel DataModel[TModel, TSearchFields], TSearchFields any, TModel interface{}]() *DataModel[TModel, TSearchFields] {
	dataModel := &DataModel[TModel, TSearchFields]{
		entity: new(TModel),
		Search: InitSearch[TModel, TSearchFields](),
		Create: InitCreate[TModel](),
	}

	contextModel := getContextModel()
	contextModel.Entities = append(contextModel.Entities, dataModel.entity)

	return dataModel
}

// Initialize Create function.
func InitCreate[T any]() func() *T {
	return func() *T {
		entity := new(T)
		id, _ := uuid.NewV4()

		SetObjectProperty(entity, "Id", id.String())
		return entity
	}
}

// Saving data.
func (d *DataModel[T, TSearchFields]) Save(entity *T) {
	getContextModel().db.Save(entity)
}

// Delete element.
func (d *DataModel[T, TSearchFields]) Delete(entity *T) {
	getContextModel().db.Delete(entity)
}

// Get object by Id.
func (d *DataModel[T, TSearchFields]) Fetch(id string) (*T, error) {
	baseSearch := &BaseEntitySearch{Id: &StringFieldOperands{FieldValueOperandsParams: FieldValueOperandsParams{FieldName: "id"}}}

	return d.Search().Where(func(operands *Operands, fields *TSearchFields) *GLobalFilter {
		return operands.And(*baseSearch.Id.Equals(id))
	}).First()
}

// Get objects by ids.
func (d *DataModel[T, TSearchFields]) FetchAll(id []string) ([]T, error) {
	baseSearch := &BaseEntitySearch{Id: &StringFieldOperands{FieldValueOperandsParams: FieldValueOperandsParams{FieldName: "id"}}}

	return d.Search().Where(func(operands *Operands, fields *TSearchFields) *GLobalFilter {
		return operands.And(*baseSearch.Id.In(id))
	}).All()
}
