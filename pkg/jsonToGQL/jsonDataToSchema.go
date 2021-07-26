package jsonDataToSchema

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

func DataTypeSwitch(query string) map[string]interface{} {
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