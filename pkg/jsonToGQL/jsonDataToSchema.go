package jsonToGQL

import (
	"encoding/json"
)

type Schema struct {
	Typ       string
	FieldData map[string][]string
}

func NewSchema() *Schema {
	var g Schema
	g.Typ = "custom"
	g.FieldData = make(map[string][]string)
	return &g
}

func stringtomap(query string) map[string]interface{} {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(query), &data); err != nil {
	}
	return data
}

func DataTypesToSchema(i interface{}) string {

	switch i.(type) {
	case float64:
		if i.(float64) == float64(int(i.(float64))) {
			return "int"
		}
		return "float"
	case string:
		return "string"
	default:
		return "string"

	}
}

func jsonToList(json []map[string]interface{}) map[string]interface{} {
	return json[0]
}

func ListOrMapSwitch(query string) map[string]interface{} {
	var dataList []map[string]interface{}
	var dataMap map[string]interface{}
	if err := json.Unmarshal([]byte(query), &dataList); err != nil {
		dataMap = stringtomap(query)
	} else {
		dataMap = jsonToList(dataList)
	}

	return dataMap
}

func ExtractFields(dataMap map[string]interface{}) Schema {
	inputData := make([]string, len(dataMap))

	//FieldMap := make(map[string][]string)

	ss := *NewSchema()

	i := 0
	// copy c's keys into k
	for s, w := range dataMap {
		var sl []string
		sl = append(sl, DataTypesToSchema(w))

		//FieldMap[inputData[i]] = sl
		inputData[i] = s
		ss.FieldData[inputData[i]] = sl
		i++
	}
	/* 	for key, value := range results {
	   		fmt.Println("key:", key, "value:", value)
	   	}
	   	fmt.Printf("%#v\n", ss) */
	return ss
}
