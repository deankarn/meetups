package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/fr"
	"github.com/go-playground/pure"
	"github.com/go-playground/pure/examples/middleware/logging-recovery"
	"github.com/go-playground/universal-translator"
)

var (
	tmpls  *template.Template
	utrans *ut.UniversalTranslator
)

func main() {

	en := en.New()
	utrans = ut.New(en, en, fr.New())
	setup()

	tmpls, _ = template.ParseFiles("home.tmpl")

	r := pure.New()
	r.Use(middleware.LoggingAndRecovery(true))
	r.Get("/", home)

	http.ListenAndServe(":8080", r.Serve())
}

type homepage struct {
	Locale    string
	DaysLeft1 string
	DaysLeft2 string
	Date      string
	Time      string
	Currency  string
	Weekdays  []string
	Months    []string
}

func home(w http.ResponseWriter, r *http.Request) {

	// get locale translator
	t := getTranslator(r)

	s := new(homepage)
	s.Locale = t.Locale()
	s.DaysLeft1, _ = t.C("days-left", 1, 0, "1")
	s.DaysLeft2, _ = t.C("days-left", 2, 0, "2")
	s.Date = t.FmtDateLong(time.Now())
	s.Time = t.FmtTimeLong(time.Now())
	s.Weekdays = t.WeekdaysWide()
	s.Months = t.MonthsWide()
	s.Currency = t.FmtCurrency(1025045.45, 2, getCurrency(t))

	if err := tmpls.ExecuteTemplate(w, "home", s); err != nil {
		log.Fatal(err)
	}
}

func getCurrency(t ut.Translator) currency.Type {

	switch t.Locale() {
	case "en":
		return currency.USD
	case "fr":
		return currency.EUR
	default:
		return currency.USD
	}
}

func getTranslator(r *http.Request) (t ut.Translator) {

	params := r.URL.Query()

	locale := params.Get("locale")

	if len(locale) > 0 {

		var found bool

		if t, found = utrans.GetTranslator(locale); found {
			return
		}
	}

	// get and parse the "Accept-Language" http header and return an array
	languages := pure.AcceptedLanguages(r)
	t, _ = utrans.FindTranslator(languages...)

	return
}

func setup() {

	en, _ := utrans.FindTranslator("en")
	en.AddCardinal("days-left", "There is {0} day left", locales.PluralRuleOne, false)
	en.AddCardinal("days-left", "There are {0} days left", locales.PluralRuleOther, false)

	fr, _ := utrans.FindTranslator("fr")
	fr.AddCardinal("days-left", "Il reste {0} jour", locales.PluralRuleOne, false)
	fr.AddCardinal("days-left", "Il reste {0} jours", locales.PluralRuleOther, false)
}
