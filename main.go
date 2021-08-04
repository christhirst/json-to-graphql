package main

import (
	"fmt"
	"strings"

	"github.com/christhirst/json-to-graphql/pkg/jsonToGQL"
)

func main() {

	query := `
	
	[
		{
		  "userId": 1,
		  "id": 1,
		  "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		  "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
		},
		{
		  "userId": 1,
		  "id": 2,
		  "title": "qui est esse",
		  "body": "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla"
		} 
 
 
  ]
		`

	/* 	query3 = `type Person {
	   		id: ID!
	   		xid: String! @id
	   		name: String!
	   		age: Int @search
	   		friends: [Person] @hasInverse(field: "friends")
	   		ownsPets: [Animal] @hasInverse(field: "owner")
	   	  }

	   	  type Animal {
	   		id: ID!
	   		xid: String! @id
	   		name: String!
	   		owner: Person @hasInverse(field: "ownsPets")
	   	  }` */

	/* 	body := `
	   	type Task {
	   		id: ID!
	   		title: String!
	   		completed: Boolean!
	   	}` */
	parsedSchema := jsonToGQL.Schema{Typ: "custom"}
	dataMap := jsonToGQL.ListOrMapSwitch(query)
	parsedSchema = jsonToGQL.ExtractFields(dataMap)
	//fmt.Println(parsedSchema)

	var b strings.Builder

	// Write strings to the Buffer.

	// Convert to a string and print it.
	//fmt.Println(b.String())
	b.WriteString(parsedSchema.Typ)
	b.WriteString(" {")
	for i, element := range parsedSchema.FieldData {
		for _, ee := range element {
			b.WriteString("\n" + i + ": " + ee)
		}
	}
	b.WriteString("\n")
	b.WriteString("}")
	fmt.Println(b.String())

}
