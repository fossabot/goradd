//** This file was code generated by got. DO NOT EDIT. ***

package orm

import (
	"bytes"
	"context"

	"github.com/goradd/goradd/web/examples/model"
)

func (ctrl *ObjectsPanel) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {

	buf.WriteString(`
<h1>The Code-generated GO Objects</h1>

<p>The Code Generator will create a GO object for most of the tables in your database in groups of two files.
One file is a base file, and contains code that is regenerated every time you run the code generator. This file
contains accessors to the various fields of the table, and functions to do queries, updates and deletes.
The other file is a companion object that embeds the base file, and that lets you override the functions
in the base file, as well as define your own functions to access the database. The companion file only gets
generated one time, and so you may edit it and change it as you like and your changes will be preserved.</p>

<p>You will see this idea of a base file that gets recreated every time the code generator runs, and an enclosing
object that embeds the base object, throughout GoRADD. We call this scaffolding...a way of layering the code
so that when you change the database, or when GoRADD itself is updated, you do not have to rewrite your program
to take advantage of the changes.</p>

<p>The example below shows how we can use the <strong>Load*()</strong> methods and the
    properties to view some of the data.  Be sure to click on the source links to view some of the code
    that made this page.</p>

`)
	// This "g" tag lets us drop in to go code whenever we want. Normally you would not write a lot of go code
	// inside a template, but rather you would put your go code in an object, and refer to variables or functions
	// from here. However, for purposes of simplifying this tutorial, we will access the database straight from here.

	// This loads the person that has an id, or primary key, of "1". Note that even though SQL can use integers as
	// primary keys, we always use strings to identify primary keys. Many other types of databases only use strings,
	// and this makes our code portable.
	person := model.LoadPerson(ctx, "1")
	project := model.LoadProject(ctx, "1")

	buf.WriteString(`<p>
<div>Person 1</div>
<div>`)

	buf.WriteString(person.FirstName())

	buf.WriteString(` `)

	buf.WriteString(person.LastName())

	buf.WriteString(`</div>
</p>
<p>
<div>Project 1</div>
<div>`)

	buf.WriteString(project.Name())

	buf.WriteString(`</div>
<div>`)

	buf.WriteString(project.Description())

	buf.WriteString(`</div>
<div>`)

	buf.WriteString(project.ProjectStatusType().String())

	buf.WriteString(`</div>
</p>

`)

	buf.WriteString(`
`)

	return
}
