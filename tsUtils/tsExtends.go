package tsUtils

import (
	"fmt"
	"reflect"
	"strings"
)

// PurgeString : supprime les "\n" et les espaces avant et après le string
//
//	entity : est un pointeur sur la structure à nettoyer
//
// ex purgeString(&movies)
func PurgeString(entity interface{}) {
	entityValue := reflect.ValueOf(entity).Elem()
	switch entityValue.Kind() {
	case reflect.String:
		// traitement du 'string' ....
		entityValue.SetString(strings.Trim(reflect.Value(entityValue).String(), "\n "))
	case reflect.Slice:
		// exploration du 'slice'
		for i := 0; i < entityValue.Len(); i++ {
			// **** appel récursif pour explorer chaque occurs du type 'slice'
			PurgeString(entityValue.Index(i).Addr().Interface()) // <== pointeur vers la structure pour permettre la modification par "reflection"
		}
	case reflect.Struct:
		// exploration du 'struct'
		for i := 0; i < entityValue.NumField(); i++ {
			field := entityValue.Field(i)
			// *** appel récursif pour explorer chaque element du type 'struct'
			PurgeString(field.Addr().Interface()) // <== pointeur vers la structure pour permettre la modification par "reflection"
		}
	}
}

// PrintAll : imprime l'intégralité d'une structure
//
//	entity	: la structure à imprimer
//	depth	: toujours 0 (usage interne pour indenter l'impression lors de struc ou de slice)
//	name	: "" ou nom externe de l'élément (usage interne)
//	indent	: toujours true (usage interne pour ne pas indenter après un slice)
//	idx		: -1 pour tout imprimer
//			: de 0 à n pour imprimer le n élément du 1er slice trouvé
func PrintAll(entity interface{}, depth int, name string, indent bool, idx int) string {
	var result string
	var tab string = ""
	if indent {
		tab = strings.Repeat("\t", depth)
	}

	entityType := reflect.TypeOf(entity)
	entityValue := reflect.ValueOf(entity)
	switch entityType.Kind() {
	case reflect.Struct:
		xName := name
		if entityType.Name() != "Name" && entityType.Name() != "Value" && entityType.Name() != "" {
			xName = entityType.Name()
		}
		result += fmt.Sprintln(tab, entityType.Kind(), xName)
		for i := 0; i < entityType.NumField(); i++ {
			fieldEntityType := entityType.Field(i)
			fieldEntityValue := entityValue.Field(i)
			result += PrintAll(fieldEntityValue.Interface(), depth+1, fieldEntityType.Name, true, idx)
		}
	case reflect.Slice:
		sliceValue := reflect.ValueOf(entityValue.Interface())
		if idx >= 0 {
			result += fmt.Sprint(tab, " ", name, "[", idx, "]")
			result += PrintAll(sliceValue.Index(idx).Interface(), depth, "", false, -1)
		} else {
			for i := 0; i < sliceValue.Len(); i++ {
				result += fmt.Sprint(tab, " ", name, "[", i, "]")
				result += PrintAll(sliceValue.Index(i).Interface(), depth, "", false, idx)
			}
		}
	default:
		if name == "" {
			result += fmt.Sprint(tab, entityType.Name(), ": '"+fmt.Sprint(entityValue)+"'"+"\n")
		} else {
			result += fmt.Sprint(tab, name, ": '"+fmt.Sprint(entityValue)+"'"+"\n")
		}
	}
	return result
}
