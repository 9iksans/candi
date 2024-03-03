package candishared

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/golangid/candi/candihelper"
)

type partialUpdateOption struct {
	updateFields map[string]struct{}
	ignoreFields map[string]struct{}
}

// DBUpdateOptionFunc option func
type DBUpdateOptionFunc func(*partialUpdateOption)

// DBUpdateSetUpdatedFields option func
func DBUpdateSetUpdatedFields(fields ...string) DBUpdateOptionFunc {
	return func(o *partialUpdateOption) {
		o.updateFields = make(map[string]struct{})
		for _, field := range fields {
			o.updateFields[field] = struct{}{}
		}
	}
}

// DBUpdateSetIgnoredFields option func
func DBUpdateSetIgnoredFields(fields ...string) DBUpdateOptionFunc {
	return func(o *partialUpdateOption) {
		o.ignoreFields = make(map[string]struct{})
		for _, field := range fields {
			o.ignoreFields[field] = struct{}{}
		}
	}
}

// DBUpdateGORMExtractorKey struct field key extractor for gorm model
func DBUpdateGORMExtractorKey(structField reflect.StructField) string {
	gormTag := structField.Tag.Get("gorm")
	if strings.HasPrefix(gormTag, "column:") {
		return strings.Split(strings.TrimPrefix(gormTag, "column:"), ";")[0]
	}
	return candihelper.ToDelimited(structField.Name, '_')
}

// DBUpdateMongoExtractorKey struct field key extractor for mongo model
func DBUpdateMongoExtractorKey(structField reflect.StructField) string {
	if bsonTag := strings.TrimSuffix(structField.Tag.Get("bson"), ",omitempty"); bsonTag != "" {
		return bsonTag
	}
	return candihelper.ToDelimited(structField.Name, '_')
}

// DBUpdateTools for construct selected field to update
type DBUpdateTools struct {
	KeyExtractorFunc func(structTag reflect.StructField) string
	IgnoredFields    []string
}

func (d *DBUpdateTools) parseOption(opts ...DBUpdateOptionFunc) (o partialUpdateOption) {
	for _, opt := range opts {
		opt(&o)
	}
	return o
}

// ToMap method
func (d DBUpdateTools) ToMap(data interface{}, opts ...DBUpdateOptionFunc) map[string]interface{} {
	opt := d.parseOption(opts...)

	dataValue := reflect.ValueOf(data)
	dataType := reflect.TypeOf(data)
	if dataValue.Kind() == reflect.Ptr {
		dataValue = dataValue.Elem()
		dataType = dataType.Elem()
	}
	isPartial := len(opt.updateFields) > 0 || len(opt.ignoreFields) > 0

	updateFields := make(map[string]interface{}, 0)
	for i := 0; i < dataValue.NumField(); i++ {
		fieldValue := dataValue.Field(i)
		fieldType := dataType.Field(i)

		if fieldType.Anonymous {
			for k, v := range d.ToMap(fieldValue.Interface(), opts...) {
				updateFields[k] = v
			}
			continue
		}

		key := strings.TrimSuffix(fieldType.Tag.Get("json"), ",omitempty")
		if d.KeyExtractorFunc != nil {
			key = d.KeyExtractorFunc(fieldType)
		}
		isIgnore, _ := strconv.ParseBool(fieldType.Tag.Get("ignoreUpdate"))
		if key == "" || key == "-" || isIgnore {
			continue
		}

		val := fieldValue.Interface()
		if fieldValue.Kind() == reflect.Pointer && !fieldValue.IsNil() {
			val = fieldValue.Elem().Interface()
		}

		if !isPartial {
			updateFields[key] = val
			continue
		}

		_, isFieldUpdated := opt.updateFields[fieldType.Name]
		_, isFieldIgnored := opt.ignoreFields[fieldType.Name]
		if (isFieldUpdated && len(opt.updateFields) > 0) || (!isFieldIgnored && len(opt.ignoreFields) > 0) {
			updateFields[key] = val
		}
	}

	for _, ignored := range d.IgnoredFields {
		delete(updateFields, ignored)
	}
	return updateFields
}

// GetUpdatedFields method
func (d DBUpdateTools) GetFields(opts ...DBUpdateOptionFunc) (updates, ignores []string) {
	opt := d.parseOption(opts...)
	for k := range opt.updateFields {
		updates = append(updates, k)
	}
	for k := range opt.ignoreFields {
		ignores = append(ignores, k)
	}
	return
}
