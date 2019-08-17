package control

import (
	"context"
	"github.com/goradd/goradd/pkg/page"
	"github.com/goradd/goradd/pkg/page/control/data"
)

type PagedTableI interface {
	TableEmbedder
	PagedControlI
}

type PagedTable struct {
	Table
	PagedControl
}

func NewPagedTable(parent page.ControlI, id string) *PagedTable {
	t := &PagedTable{}
	t.Init(t, parent, id)
	return t
}

func (t *PagedTable) Init(self page.ControlI, parent page.ControlI, id string) {
	t.Table.Init(self, parent, id)
	t.PagedControl.SetPageSize(0) // use the application default
}

// PagedTableCreator creates a table that can be paged
type PagedTableCreator struct {
	// ID is the control id
	ID               string
	// HasColTags will make the table render <col> tags
	HasColTags       bool
	// Caption is the content of the caption tag, and can either be a string, or a data pager
	Caption          interface{}
	// HideIfEmpty will hide the table completely if it has no data. Otherwise, the table and headers will be shown, but no data rows
	HideIfEmpty      bool
	// HeaderRowCount is the number of header rows. You must set this to at least 1 to show header rows.
	HeaderRowCount   int
	// FooterRowCount is the number of footer rows.
	FooterRowCount   int
	// RowStyler returns the attributes to be used in a cell.
	RowStyler        TableRowAttributer
	// RowStylerID is a control id for the control that will be the RowStyler of the table.
	RowStylerID      string
	// HeaderRowStyler returns the attributes to be used in a header cell.
	HeaderRowStyler  TableHeaderRowAttributer
	// HeaderRowStylerID is a control id for the control that will be the HeaderRowStyler of the table.
	HeaderRowStylerID  string
	// FooterRowStyler returns the attributes to be used in a footer cell. It can be either a control id or a TableFooterRowAttributer.
	FooterRowStyler  TableFooterRowAttributer
	// FooterRowStylerID is a control id for the control that will be the FooterRowStyler of the table.
	FooterRowStylerID  string
	// Columns are the column creators that will add columns to the table
	Columns          []ColumnCreator
	// DataProvider is the data binder for the table. It can be either a control id or a DataBinder
	DataProvider     data.DataBinder
	// DataProviderID is the control id of the data binder for the table.
	DataProviderID	string
	// Data is the actual data for the table, and should be a slice of objects
	Data             interface{}
	// Sortable will make the table sortable
	Sortable         bool
	// SortHistoryLimit will set how many columns deep we will remember the sorting for multi-level sorts
	SortHistoryLimit int
	page.ControlOptions
	// PageSize is the number of rows to include in a page
	PageSize         int
	// SaveState will cause the table to remember what page it was on
	SaveState bool

}



// Create is called by the framework to create a new control from the Creator. You
// do not normally need to call this.
func (c PagedTableCreator) Create(ctx context.Context, parent page.ControlI) page.ControlI {
	ctrl := NewPagedTable(parent, c.ID)
	c.Init(ctx, ctrl)
	return ctrl
}

// Init is called by implementations of Buttons to initialize a control with the
// creator. You do not normally need to call this.
func (c PagedTableCreator) Init(ctx context.Context, ctrl PagedTableI) {
	sub := TableCreator {
		ID:               c.ID,
		HasColTags:       c.HasColTags,
		Caption:          c.Caption,
		HideIfEmpty:      c.HideIfEmpty,
		HeaderRowCount:   c.HeaderRowCount,
		FooterRowCount:   c.FooterRowCount,
		RowStyler:        c.RowStyler,
		HeaderRowStyler:  c.HeaderRowStyler,
		FooterRowStyler:  c.FooterRowStyler,
		Columns:          c.Columns,
		DataProvider:     c.DataProvider,
		Data:             c.Data,
		Sortable:         c.Sortable,
		SortHistoryLimit: c.SortHistoryLimit,
		ControlOptions:   c.ControlOptions,
	}
	sub.Init(ctx, ctrl)
	if c.PageSize != 0 {
		ctrl.SetPageSize(c.PageSize)
	}
	if c.SaveState {
		ctrl.SaveState(ctx, true)
	}
}

// GetPagedTable is a convenience method to return the table with the given id from the page.
func GetPagedTable(c page.ControlI, id string) *PagedTable {
	return c.Page().GetControl(id).(*PagedTable)
}