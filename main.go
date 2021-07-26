package main

import (
	"encoding/json"
	"fmt"
	
)

type schema struct {
	fields []string
	typ    string
}

func stringtomap(query string) map[string]interface{} {
	fmt.Println("query")
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(query), &data); err != nil {
	}
	return data
}

func jsonToList(json []map[string]interface{}) map[string]interface{} {
	return json[0]
}

func dataTypeSwitch(query string) map[string]interface{} {
	var data []map[string]interface{}
	var dataMap map[string]interface{}
	if err := json.Unmarshal([]byte(query), &data); err != nil {
		dataMap = stringtomap(query)
	} else {
		dataMap = jsonToList(data)
	}

	return dataMap
}

func extractFields(dataMap map[string]interface{}) schema {
	k := make([]string, len(dataMap))
	i := 0
	var ss schema

	// copy c's keys into k
	for s, _ := range dataMap {
		k[i] = s
		i++
	}
	ss.fields = k
	/* 	for key, value := range results {
	   		fmt.Println("key:", key, "value:", value)
	   	}
	   	fmt.Printf("%#v\n", ss) */
	return ss
}

func main() {

	query := `
	
   {
    "userId": 1,
    "id": 1,
    "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
    "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
  }
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

	dataMap := dataTypeSwitch(query)
pkg.parsedSchema := extractFields(dataMap)
	fmt.Println(parsedSchema)

	/* 	k := make([]string, len(data[0]))
	   	i := 0
	   	var ss schema
	   	gen("ss")
	   	// copy c's keys into k
	   	for s, _ := range data[0] {
	   		k[i] = s
	   		i++
	   	}ll
	   	ss.fields = k
	   	/* 	for key, value := range results {
	   		fmt.Println("key:", key, "value:", value)
	   	}
	   	fmt.Printf("%#v\n", ss) */

}
