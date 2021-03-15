package main

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/i18n"
)

func main() {
	i18n.MustLoadTranslationFile("en-us.yaml")
	//i18n.MustLoadTranslationFile("ar.toml")

	T, _ := i18n.Tfunc("en")

	bobMap := map[string]interface{}{"Person": "Bob"}
	bobStruct := struct{ Person string }{Person: "Bob"}

	fmt.Println(T("program_greeting", 2))
	fmt.Println(T("program_greeting", 1))
	fmt.Println(T("person_greeting", bobMap))
	fmt.Println(T("person_greeting", bobStruct))

	fmt.Println(T("my_height_in_meters", map[string]interface{}{
		"Name":  "jiebin",
		"Count": "1",
	}))
	fmt.Println(T("d_days", map[string]interface{}{
		"Count": 2,
	}))
	/*

		fmt.Println(T("person_unread_email_count", 0, bobMap))
		fmt.Println(T("person_unread_email_count", 1, bobMap))
		fmt.Println(T("person_unread_email_count", 2, bobMap))
		fmt.Println(T("person_unread_email_count", 0, bobStruct))
		fmt.Println(T("person_unread_email_count", 1, bobStruct))
		fmt.Println(T("person_unread_email_count", 2, bobStruct))

		type Count struct{ Count int }
		fmt.Println(T("your_unread_email_count", 1, Count{0}))
		fmt.Println(T("your_unread_email_count", 1, Count{1}))
		fmt.Println(T("your_unread_email_count", 2, Count{2}))

		fmt.Println(T("your_unread_email_count", map[string]interface{}{"Count": 0}))
		fmt.Println(T("your_unread_email_count", map[string]interface{}{"Count": "1"}))
		fmt.Println(T("your_unread_email_count", map[string]interface{}{"Count": 2}))
		fmt.Println(T("your_unread_email_count", map[string]interface{}{"Count": "3.14"}))

		fmt.Println(T("person_unread_email_count_timeframe", 3, map[string]interface{}{
			"Person":    "Bob",
			"Timeframe": T("d_days", 0),
		}))
		fmt.Println(T("person_unread_email_count_timeframe", 3, map[string]interface{}{
			"Person":    "Bob",
			"Timeframe": T("d_days", 1),
		}))
		fmt.Println(T("person_unread_email_count_timeframe", 3, map[string]interface{}{
			"Person":    "Bob",
			"Timeframe": T("d_days", 2),
		}))

		fmt.Println(T("person_unread_email_count_timeframe", 1, map[string]interface{}{
			"Count":     30,
			"Person":    "Bob",
			"Timeframe": T("d_days", 0),
		}))
		fmt.Println(T("person_unread_email_count_timeframe", 2, map[string]interface{}{
			"Count":     20,
			"Person":    "Bob",
			"Timeframe": T("d_days", 1),
		}))
		fmt.Println(T("person_unread_email_count_timeframe", 3, map[string]interface{}{
			"Count":     10,
			"Person":    "Bob",
			"Timeframe": T("d_days", 2),
		}))
		fmt.Println(T("jiebin_test", map[string]interface{}{
			"Name": "jiebin",
		}))
	*/
}
