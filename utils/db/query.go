package db

import "fmt"

type SqlQuery struct {
	text string
}

func NewQuery() *SqlQuery {
	return &SqlQuery{text: ""}
}

func (query *SqlQuery) From(table string) *SqlQuery {
	query.text += "FROM " + table + "\n"

	return query
}

func (query *SqlQuery) Select(columns ...string) *SqlQuery {
	text := "SELECT "

	if len(columns) == 0 {
		text += "*"
	} else {
		for _, column := range columns {
			if column == columns[len(columns)-1] {
				text += column
			} else {
				text += fmt.Sprintf("%s,", column)
			}
		}
	}

	text += "\n"
	query.text += text

	return query
}
