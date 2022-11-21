package reflection

import (
	"fmt"
	"log"
	"reflect"
)

type S struct {
	Symbol      string `json:"symbol"`
	Position    string `persistence:"position"`
	ProfitStats string `persistence:"profit_stats"`
	TradeStats  string `persistence:"trade_stats"`
}

type Stash map[string]interface{}

func filter_tag(conf S, tag string) (Stash, error) {
	s := Stash{}
	v := reflect.ValueOf(conf)

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		if tag_name, ok := field.Tag.Lookup("persistence"); ok {
			if tag_name != "" {
				// fmt.Println("name", field.Name)
				// fmt.Println("value", v.Field(i).Interface())
				s[field.Name] = v.Field(i).Interface()
			}
		}
	}
	return s, nil
}
func Exec_filter_tag() {
	s := S{Symbol: "USDT", Position: "1", ProfitStats: "2", TradeStats: "3"}
	tag_value, err := filter_tag(s, "persistence")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tag_value", tag_value)
}
