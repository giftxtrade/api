//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

// UseSchema sets a new schema name for all generated table SQL builder types. It is recommended to invoke
// this method only once at the beginning of the program.
func UseSchema(schema string) {
	Category = Category.FromSchema(schema)
	Draw = Draw.FromSchema(schema)
	Event = Event.FromSchema(schema)
	Link = Link.FromSchema(schema)
	Migration = Migration.FromSchema(schema)
	Participant = Participant.FromSchema(schema)
	Product = Product.FromSchema(schema)
	User = User.FromSchema(schema)
	Wish = Wish.FromSchema(schema)
}