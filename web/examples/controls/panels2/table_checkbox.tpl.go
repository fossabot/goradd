//** This file was code generated by got. DO NOT EDIT. ***

package panels

import (
	"bytes"
	"context"
)

func (control *panels.TableCheckboxPanel) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {

	buf.WriteString(`
`)

	buf.WriteString(`
<style>
table {
  font-family: "Trebuchet MS", Arial, Helvetica, sans-serif;
  border-collapse: collapse;
  width: 100%;
}

table td, table th {
  border: 1px solid #ddd;
  padding: 8px;
}

table tr:nth-child(even){background-color: #f2f2f2;}

table tr:hover {background-color: #ddd;}

table th {
  padding-top: 12px;
  padding-bottom: 12px;
  text-align: left;
  background-color: #4CAF50;
  color: white;
}
</style>
<h1>Tables - Checkbox Columns</h1>
<p>
A CheckboxColumn displays a single checkbox in a column. When you create it, you designate a
CheckboxProvider, which will determine what the initial state of the checkboxes will be. Once set up,
the column will keep track of changes, and when you are ready to save the changes, you can call
Changes() on the column to get the state of the changed checkboxes. This is useful if you have a
Save button to finally record the changes, but you can also use the CheckboxColumnClick event to
record changes in real time through Javascript and Ajax.
</p>
`)

	buf.WriteString(`
`)

	{
		err := control.Pager1.Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)

	{
		err := control.Table1.Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`

`)

	buf.WriteString(`
`)

	{
		err := control.SubmitAjax.Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)

	{
		err := control.SubmitServer.Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`

`)

	buf.WriteString(`
`)

	return
}
