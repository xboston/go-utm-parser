package utm

import (
	"net/url"
	"reflect"
	"testing"
)

var (
	tests = []struct {
		url string
		utm Utm
	}{
		{"http://www.example.com/?utm_source=source",
			Utm{
				Source:   "source",
				Campaign: "",
				Medium:   "",
				Content:  "",
				Term:     "",
				Extra:    map[string]string{},
			},
		},
		{"http://www.example.com/?utm_source=source&utm_campaign=campaign",
			Utm{
				Source:   "source",
				Campaign: "campaign",
				Medium:   "",
				Content:  "",
				Term:     "",
				Extra:    map[string]string{},
			},
		},
		{"http://www.example.com/?utm_source=source&utm_campaign=campaign&utm_medium=medium",
			Utm{
				Source:   "source",
				Campaign: "campaign",
				Medium:   "medium",
				Content:  "",
				Term:     "",
				Extra:    map[string]string{},
			},
		},
		{"http://www.example.com/?utm_source=source&utm_campaign=campaign&utm_medium=medium&utm_content=content",
			Utm{
				Source:   "source",
				Campaign: "campaign",
				Medium:   "medium",
				Content:  "content",
				Term:     "",
				Extra:    map[string]string{},
			},
		},
		{"http://www.example.com/?utm_source=source&utm_campaign=campaign&utm_medium=medium&utm_content=content&utm_term=term",
			Utm{
				Source:   "source",
				Campaign: "campaign",
				Medium:   "medium",
				Content:  "content",
				Term:     "term",
				Extra:    map[string]string{},
			},
		},
		{"http://www.example.com/?utm_source=source&utm_campaign=campaign&utm_medium=medium&utm_content=content&utm_term=term&utm_extra1=extra1&utm_extra2=extra2",
			Utm{
				Source:   "source",
				Campaign: "campaign",
				Medium:   "medium",
				Content:  "content",
				Term:     "term",
				Extra: map[string]string{
					"utm_extra2": "extra2",
					"utm_extra1": "extra1",
				},
			},
		},
	}
)

func TestParseURL(t *testing.T) {

	for _, tt := range tests {

		utm, err := ParseURL(&tt.url)

		if err != nil {
			t.Error(err.Error())
		}

		if !reflect.DeepEqual(utm, tt.utm) {
			t.Errorf("New.Error(): got: %q, want %q", utm, tt.utm)
		}
	}
}

func TestParseValues(t *testing.T) {

	for _, tt := range tests {

		url, _ := url.Parse(tt.url)
		u := url.Query()
		utm, err := ParseValuesFull(&u)

		if err != nil {
			t.Error(err.Error())
		}

		if !reflect.DeepEqual(utm, tt.utm) {
			t.Errorf("New.Error(): got: %q, want %q", utm, tt.utm)
		}
	}
}

func BenchmarkParseURL(b *testing.B) {

	url := "http://www.example.com/?utm_source=source&utm_campaign=campaign&utm_medium=medium&utm_content=content&utm_term=term&utm_extra1=extra1&utm_extra2=extra2"
	for i := 0; i < b.N; i++ {
		ParseURL(&url)
	}
}

func BenchmarkParallelParseURL(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {

		url := "http://www.example.com/?utm_source=source&utm_campaign=campaign&utm_medium=medium&utm_content=content&utm_term=term&utm_extra1=extra1&utm_extra2=extra2"
		for pb.Next() {
			ParseURL(&url)
		}
	})
}

func BenchmarkParseQuery(b *testing.B) {

	rawurl := "http://www.example.com/?a=b&utm_source=source&utm_campaign=campaign&utm_medium=medium&utm_content=content&utm_term=term&utm_extra1=extra1&utm_extra2=extra2&utm_extra3=extra3&utm_extra5=extra5"
	url, _ := url.Parse(rawurl)
	u := url.Query()
	for i := 0; i < b.N; i++ {
		ParseValues(&u)
	}
}

func BenchmarkParallelParseQuery(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {

		rawurl := "http://www.example.com/?utm_source=source&utm_campaign=campaign&utm_medium=medium&utm_content=content&utm_term=term&utm_extra1=extra1&utm_extra2=extra2"
		url, _ := url.Parse(rawurl)
		u := url.Query()
		for pb.Next() {
			ParseValues(&u)
		}
	})
}

func BenchmarkParseQueryFull(b *testing.B) {

	rawurl := "http://www.example.com/?a=b&utm_source=source&utm_campaign=campaign&utm_medium=medium&utm_content=content&utm_term=term&utm_extra1=extra1&utm_extra2=extra2&utm_extra3=extra3&utm_extra5=extra5"
	url, _ := url.Parse(rawurl)
	u := url.Query()
	for i := 0; i < b.N; i++ {
		ParseValuesFull(&u)
	}
}

func BenchmarkParallelParseQueryFull(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {

		rawurl := "http://www.example.com/?utm_source=source&utm_campaign=campaign&utm_medium=medium&utm_content=content&utm_term=term&utm_extra1=extra1&utm_extra2=extra2"
		url, _ := url.Parse(rawurl)
		u := url.Query()
		for pb.Next() {
			ParseValuesFull(&u)
		}
	})
}
