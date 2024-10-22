//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Event = newEventTable("public", "event", "")

type eventTable struct {
	postgres.Table

	// Columns
	ID                postgres.ColumnInteger
	Name              postgres.ColumnString
	Description       postgres.ColumnString
	Budget            postgres.ColumnString
	InvitationMessage postgres.ColumnString
	DrawAt            postgres.ColumnTimestampz
	CloseAt           postgres.ColumnTimestampz
	CreatedAt         postgres.ColumnTimestampz
	UpdatedAt         postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type EventTable struct {
	eventTable

	EXCLUDED eventTable
}

// AS creates new EventTable with assigned alias
func (a EventTable) AS(alias string) *EventTable {
	return newEventTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new EventTable with assigned schema name
func (a EventTable) FromSchema(schemaName string) *EventTable {
	return newEventTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new EventTable with assigned table prefix
func (a EventTable) WithPrefix(prefix string) *EventTable {
	return newEventTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new EventTable with assigned table suffix
func (a EventTable) WithSuffix(suffix string) *EventTable {
	return newEventTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newEventTable(schemaName, tableName, alias string) *EventTable {
	return &EventTable{
		eventTable: newEventTableImpl(schemaName, tableName, alias),
		EXCLUDED:   newEventTableImpl("", "excluded", ""),
	}
}

func newEventTableImpl(schemaName, tableName, alias string) eventTable {
	var (
		IDColumn                = postgres.IntegerColumn("id")
		NameColumn              = postgres.StringColumn("name")
		DescriptionColumn       = postgres.StringColumn("description")
		BudgetColumn            = postgres.StringColumn("budget")
		InvitationMessageColumn = postgres.StringColumn("invitation_message")
		DrawAtColumn            = postgres.TimestampzColumn("draw_at")
		CloseAtColumn           = postgres.TimestampzColumn("close_at")
		CreatedAtColumn         = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn         = postgres.TimestampzColumn("updated_at")
		allColumns              = postgres.ColumnList{IDColumn, NameColumn, DescriptionColumn, BudgetColumn, InvitationMessageColumn, DrawAtColumn, CloseAtColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns          = postgres.ColumnList{NameColumn, DescriptionColumn, BudgetColumn, InvitationMessageColumn, DrawAtColumn, CloseAtColumn, CreatedAtColumn, UpdatedAtColumn}
	)

	return eventTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                IDColumn,
		Name:              NameColumn,
		Description:       DescriptionColumn,
		Budget:            BudgetColumn,
		InvitationMessage: InvitationMessageColumn,
		DrawAt:            DrawAtColumn,
		CloseAt:           CloseAtColumn,
		CreatedAt:         CreatedAtColumn,
		UpdatedAt:         UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}