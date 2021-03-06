//** This file was code generated by got. DO NOT EDIT. ***

package orm

import (
	"bytes"
	"context"

	"github.com/goradd/goradd/web/examples/gen/goradd/model"
	"github.com/goradd/goradd/web/examples/gen/goradd/model/node"
)

func (ctrl *OneManyPanel) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {

	buf.WriteString(`
<h1>One-to-Many Relationships</h1>
<p>
A typical foreign key will create a one-to-many relationship. The field with the foreign key points to a single record
elsewhere in the database, but that other record may have many such foreign keys pointing to it. Goradd will create functions that let
you load a single record from the "one" side of the relationship, and a slice of records from the "many" side.
</p>
<p>
In the example below, we load person 3 from the database, and then examine the addresses associated with that person.
Be sure to click on the View Source button above to see the source code.
`)
	person := model.LoadPerson(ctx, "3")
	addresses := person.LoadAddresses(ctx)

	buf.WriteString(`</p>
<p>
    Person: `)

	buf.WriteString(person.FirstName())

	buf.WriteString(` `)

	buf.WriteString(person.LastName())

	buf.WriteString(` <br>
    Addresses: <br>
    `)
	for _, addr := range addresses {
		buf.WriteString(`        `)

		buf.WriteString(addr.Street())

		buf.WriteString(`, `)

		buf.WriteString(addr.City())

		buf.WriteString(` <br>
    `)
	}

	buf.WriteString(`</p>

<h2>Creating One-to-Many Linked Records</h2>
<p>
One way to create linked records is to save a record, get its ID, and then set the foreign key in the other record to this value.
You will need to enclose the process in a transaction so that you know that the multiple steps of the save will be sure to complete.
</p>
`)
	newId1 := ctrl.addRecord(ctx) // See addRecord in 6-onetomany.go
	address1 := model.LoadAddress(ctx, newId1, node.Address().Person())

	buf.WriteString(`<p>
New Person: `)

	buf.WriteString(address1.Person().FirstName())

	buf.WriteString(` `)

	buf.WriteString(address1.Person().LastName())

	buf.WriteString(`, `)

	buf.WriteString(address1.Street())

	buf.WriteString(`, `)

	buf.WriteString(address1.City())

	buf.WriteString(`
</p>
<p>
An easier way to do this is to create the two objects, and then use the Set* function to link the two objects together.
When you do the Save() this way, internally the two database operations are wrapped in a transaction automatically for you.
</p>
`)
	newId2 := ctrl.addRecordSimpler(ctx) // See addRecordSimpler in 6-onetomany.go
	address2 := model.LoadAddress(ctx, newId2, node.Address().Person())

	buf.WriteString(`<p>
New Person: `)

	buf.WriteString(address2.Person().FirstName())

	buf.WriteString(` `)

	buf.WriteString(address2.Person().LastName())

	buf.WriteString(`, `)

	buf.WriteString(address2.Street())

	buf.WriteString(`, `)

	buf.WriteString(address2.City())

	buf.WriteString(`
</p>

<p>
Saving on the "many" side is simply a matter of setting the slice of records, and then calling Save(). As above, the transaction is
built in to the Save() process in this situation.
</p>
`)
	manyPersonId := ctrl.addMany(ctx)
	manyPerson := model.LoadPerson(ctx, manyPersonId, node.Person().Addresses())

	buf.WriteString(`<p>
    New Person: `)

	buf.WriteString(manyPerson.FirstName())

	buf.WriteString(` `)

	buf.WriteString(manyPerson.LastName())

	buf.WriteString(` <br>
    Addresses: <br>
    `)
	for _, addr := range manyPerson.Addresses() {
		buf.WriteString(`        `)

		buf.WriteString(addr.Street())

		buf.WriteString(`, `)

		buf.WriteString(addr.City())

		buf.WriteString(` <br>
    `)
	}

	buf.WriteString(`</p>
<h2>Updating One-to-Many Linked Records</h2>
<p>
    To modify a linked record, simply call the appropriate Set* function, and then call Save().
</p>
`)
	address1.Person().SetFirstName("Bob") // Here we change a name in the person
	address1.Save(ctx)                    // Note that even though we are saving the address, it will also call Save on all the linked objects, including the Person.

	linkedPerson := model.LoadPerson(ctx, address1.PersonID())

	buf.WriteString(`<p>
`)

	buf.WriteString(`    The linked person&#39;s name was changed to `)

	buf.WriteString(linkedPerson.FirstName())

	buf.WriteString(`
`)

	buf.WriteString(`
</p>
<p>
    You can also change the link to point to a new record. However, be aware that the old linked record may still exist.
    This will depend on whether
        you are changing the link from the "one" side or "many" side, how you set up the foreign key, and whether a NULL value
        is allowed in the foreign key field.
</p>
`)
	address3 := model.NewAddress()
	address3.SetStreet("Main Street")
	address3.SetCity("San Diego")
	linkedPerson.SetAddresses([]*model.Address{address3})
	linkedPerson.Save(ctx) // this will save the new address and join it to the linked person in one step.
	// And, since the address's Person constraint is set to CASCADE on DELETE, the old addresses will be deleted

	buf.WriteString(`<p>
`)

	buf.WriteString(`    The linked person&#39;s address was changed to `)

	buf.WriteString(linkedPerson.Addresses()[0].Street())

	buf.WriteString(`, `)

	buf.WriteString(linkedPerson.Addresses()[0].City())

	buf.WriteString(`
`)

	buf.WriteString(`
</p>


`)
	// Cleanup by deleting whatever we created above
	address1.Person().Delete(ctx)
	address2.Person().Delete(ctx)
	manyPerson.Delete(ctx)

	buf.WriteString(`
`)

	buf.WriteString(`
`)

	return
}
