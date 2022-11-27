package parser_test

//
//import (
//	"HomegrownDB/backend/new/internal/parser/internal/validator"
//	"HomegrownDB/common/tests/assert"
//	"testing"
//)
//
//func TestField_Parse_ShouldParse(t *testing.T) {
//	//given
//	queryParts := []string{
//		"t1.col1",
//		"t1.col1, t2.col2",
//		"t1.col1 FROM ",
//	}
//
//	tableAlias := "t1"
//	fieldName := "col1"
//	fieldAlias := "col1"
//
//	for _, queryPart := range queryParts {
//		source := newTestTokenSource(queryPart)
//		v := validator.NewValidator(source)
//		//when
//		result, err := Field.Parse(source, v)
//
//		if err != nil {
//			t.Error(err)
//		}
//
//		assert.Eq(result.TableAlias, tableAlias, t)
//		assert.Eq(result.FieldAlias, fieldAlias, t)
//		assert.Eq(result.FieldName, fieldName, t)
//	}
//}
//
//func TestField_Parse_ShouldReturnError(t *testing.T) {
//	//given
//	queryParts := []string{
//		"t1.2",
//		"t1,col1",
//		"t1 .col1",
//	}
//
//	for _, queryPart := range queryParts {
//		source := newTestTokenSource(queryPart)
//		v := validator.NewValidator(source)
//		//when
//		_, err := Field.Parse(source, v)
//
//		if err == nil {
//			t.Errorf(
//				"Field parser returned nil error for invalid token sequence: %s",
//				queryPart)
//		}
//	}
//}
