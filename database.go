package goutils

import (
	"fmt"
	"strings"
)

type Model interface {
	TableName() string
	TableColumns() []string
}

func ColumnNamesExclusive(model Model, exclude ...string) string {
	return strings.Join(RemoveExcludedFromSlice(model.TableColumns(), exclude), ",")
}

func ColumnNamesInclusive(model Model, include ...string) string {
	return strings.Join(KeepIncludedInSlice(model.TableColumns(), include), ",")
}

func PrepareBatchValues(paramLength int, valueLength int) string {
	var valString string
	for i := 0; i < paramLength; i++ {
		valString = valString + "?,"
	}
	valString = fmt.Sprintf("(%s)", strings.TrimSuffix(valString, ","))

	var values string
	for i := 0; i < valueLength; i++ {
		values = fmt.Sprintf("%s, %s", values, valString)
	}
	return strings.TrimPrefix(values, ", ")
}

func PrepareUpdateScript(updateColumns []string) string {
	updaterSQL := ""
	for _, columnName := range updateColumns {
		updaterSQL += fmt.Sprintf("%s = ?,", columnName)
	}
	updaterSQL = strings.TrimSuffix(updaterSQL, ",")
	return updaterSQL
}
