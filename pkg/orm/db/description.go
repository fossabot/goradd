package db

// This file describes the export/import structure that allows a user to manually describe the structure of the
// database. In particular, this is required when working with a NoSQL database, since the structure cannot be
// inferred from the database.

type DatabaseDescription struct {
	// The database key corresponding to its key in the global database cluster
	Key string
	// Tables are the tables in the database
	Tables []TableDescription
	// MM are the many-to-many links between tables. In SQL databases, these are actual tables,
	// but in NoSQL, they become array fields on either side of the relationship.
	MM []ManyManyDescription

	// The prefix for related objects.
	AssociatedObjectPrefix string
}

type TableDescription struct {
	// Name is the name of the database table or collection.
	Name string
	// Columns is a list of ColumnDescriptions, one for each column in the table.
	// The first column(s) is the primary key
	Columns []ColumnDescription
	// Indexes are the indexes defined in the database. Unique indexes will result in LoadBy* functions.
	Indexes []IndexDescription
	// TypeData is the data of the type table if this is a type table. The data structure must match that of the columns.
	TypeData []map[string]interface{}

	// Code-generation overrides. Leave these as empty to get the default value inferred by the Name above

	// LiteralName is the name of the object when used in a context that can be presented to the user
	LiteralName string
	// LiteralPlural is the plural form of the object when used in a context that can be presented to the user
	LiteralPlural string
	// GoName is the name of the created ORM object that reflects what is in the database
	GoName string
	// GoPlural is the name that will be used to represent a collection of ORM objects
	GoPlural string
	// Comment is a comment about the table
	Comment string
	// Options are key-value settings that can be used to further describe code generation
	Options map[string]interface{}
}

type ColumnDescription struct {
	// Name is the name of the column in the database. This is blank if this is a "virtual" table for sql tables like an association or virtual attribute query.
	Name string
	// GoName is the name of the column in go code
	GoName string
	// NativeType is the type of the column as described by the database itself.
	NativeType string
	//  GoType is the goradd defined column type
	GoType string
	// MaxCharLength is the maximum length of characters to allow in the column if a string type column.
	// If the database has the ability to specify this, this will correspond to what is specified.
	// In any case, we will generate code to prevent fields from getting bigger than this. Zero indicates there is
	// no length checking or limiting.
	MaxCharLength uint64
	// DefaultValue is the default value as specified by the database. We will initialize new ORM objects
	// with this value. It will be cast to the corresponding GO type.
	DefaultValue interface{}
	// MaxValue is the maximum value allowed for numeric values. This can be used by UI objects to tell the user what the limits are.
	MaxValue interface{}
	// MinValue is the minimum value allowed for numeric values. This can be used by UI objects to tell the user what the limits are.
	MinValue interface{}
	// IsId is true if this column represents a unique identifier generated by the database
	IsId bool
	// IsPk is true if this is the primary key column. PK's do not necessarily need to be ID columns, and if not, we will need to do our own work to generate unique PKs.
	IsPk bool
	// IsNullable is true if the column can be given a NULL value
	IsNullable bool
	// IsUnique indicates that the field holds unique values
	IsUnique bool
	// IsTimestamp is true if the field is a timestamp. Timestamps represent a specific point in world time.
	IsTimestamp bool
	// IsAutoUpdateTimestamp is true if the database is updating the timestamp. Otherwise we will do it manually.
	IsAutoUpdateTimestamp bool
	// ForeignKey is additional information describing a foreign key relationship
	ForeignKey *ForeignKeyDescription
	// Comment is the contents of the comment associated with this column
	Comment string
	// Options are key-value settings that can be used to further describe code generation
	Options map[string]interface{}
}


type ForeignKeyDescription struct {
	//DbKey string	// We don't support cross database foreign keys yet. Someday maybe.
	// ReferencedTable is the name of the table on the other end of the foreign key
	ReferencedTable string
	// ReferencedColumn is the database column name in the linked table that matches this column. Often that is the primary key of the other table.
	ReferencedColumn string
	// UpdateAction indicates how the database will react when the referenced item's id changes.
	UpdateAction string
	// DeleteAction indicates how the database will react when the referenced item is deleted.
	DeleteAction string
	// IsUnique is true if the reference is one-to-one
	IsUnique bool

	// GoName is the name we should use to refer to the related object. Leave blank to get a computed value.
	GoName string
	// ReverseName is the name that the reverse reference should use to refer to the collection of objects pointing to it.
	// Leave blank to get a "ThisAsThat" type default name. The lower-case version of this name will be used as a column name
	// to store the values if using a NoSQL database.
	ReverseName string
}

// IndexDescription gives us information about how columns are indexed.
// If a column has a unique index, it will get a corresponding "LoadBy" function in its table's model.
// Otherwise, it will get a corresponding "LoadSliceBy" function.
type IndexDescription struct {
	// IsUnique indicates whether the index is unique
	IsUnique bool
	// ColumnNames are the columns that are part of the index
	ColumnNames []string
}

type ManyManyDescription struct {
	// Table1 is the name of the first table that is part of the relationship. The private key of that table will be referred to.
	Table1 string
	// Column1 is the database column name. For SQL databases, this is the name of the column in the assn table. For
	// NoSQL, this is the name of the column that will be used to store the ids of the other side. This is optional for
	// NoSQL, as one will be created based on the table names if left blank.
	Column1 string
	// GoName1 is the singular name of the object that Table2 will use to refer to Table1 objects.
	GoName1 string
	// GoPlural1 is the plural name of the object that Table2 will use to refer to Table1 objects.
	GoPlural1 string

	// Table2 is the name of the second table that is part of the relationship. The private key of that table will be referred to.
	Table2 string
	// Column2 is the database column name. For SQL databases, this is the name of the column in the assn table. For
	// NoSQL, this is the name of the column that will be used to store the ids of the other side. This is optional for
	// NoSQL, as one will be created based on the table names if left blank.
	Column2 string
	// GoName2 is the singular name of the object that Table1 will use to refer to Table2 objects.
	GoName2 string
	// GoPlural2 is the plural name of the object that Table1 will use to refer to Table2 objects.
	GoPlural2 string

	// AssnTableName is the name of the intermediate association table that will be used to create the relationship. This is
	// needed for SQL databases, but not for NoSQL, as NoSQL will create additional array columns on each side of the relationship.
	AssnTableName string
}


