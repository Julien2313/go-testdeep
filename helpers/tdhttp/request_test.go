// Copyright (c) 2019, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package tdhttp_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/maxatome/go-testdeep/helpers/tdhttp"
	"github.com/maxatome/go-testdeep/internal/dark"
	"github.com/maxatome/go-testdeep/td"
)

func TestBasicAuthHeader(t *testing.T) {
	td.Cmp(t,
		tdhttp.BasicAuthHeader("max", "5ecr3T"),
		http.Header{"Authorization": []string{"Basic bWF4OjVlY3IzVA=="}})
}

func TestNewRequest(tt *testing.T) {
	t := td.NewT(tt)

	t.Run("NewRequest", func(t *td.T) {
		req := tdhttp.NewRequest("GET", "/path", nil,
			"Foo", "Bar",
			"Zip", "Test")

		t.Cmp(req.Header, http.Header{
			"Foo": []string{"Bar"},
			"Zip": []string{"Test"},
		})
	})

	t.Run("NewRequest last header value less", func(t *td.T) {
		req := tdhttp.NewRequest("GET", "/path", nil,
			"Foo", "Bar",
			"Zip")

		t.Cmp(req.Header, http.Header{
			"Foo": []string{"Bar"},
			"Zip": []string{""},
		})
	})

	t.Run("NewRequest header http.Header", func(t *td.T) {
		req := tdhttp.NewRequest("GET", "/path", nil,
			http.Header{
				"Foo": []string{"Bar"},
				"Zip": []string{"Test"},
			})

		t.Cmp(req.Header, http.Header{
			"Foo": []string{"Bar"},
			"Zip": []string{"Test"},
		})
	})

	t.Run("NewRequest header http.Cookie", func(t *td.T) {
		req := tdhttp.NewRequest("GET", "/path", nil,
			&http.Cookie{Name: "cook1", Value: "val1"},
			http.Cookie{Name: "cook2", Value: "val2"},
		)

		t.Cmp(req.Header, http.Header{"Cookie": []string{"cook1=val1; cook2=val2"}})
	})

	t.Run("NewRequest header flattened", func(t *td.T) {
		req := tdhttp.NewRequest("GET", "/path", nil,
			td.Flatten([]string{
				"Foo", "Bar",
				"Zip", "Test",
			}),
			td.Flatten(map[string]string{
				"Pipo": "Bingo",
				"Hey":  "Yo",
			}),
		)

		t.Cmp(req.Header, http.Header{
			"Foo":  []string{"Bar"},
			"Zip":  []string{"Test"},
			"Pipo": []string{"Bingo"},
			"Hey":  []string{"Yo"},
		})
	})

	t.Run("NewRequest header combined", func(t *td.T) {
		req := tdhttp.NewRequest("GET", "/path", nil,
			"H1", "V1",
			http.Header{
				"H1": []string{"V2"},
				"H2": []string{"V1", "V2"},
			},
			"H2", "V3",
			td.Flatten([]string{
				"H2", "V4",
				"H3", "V1",
				"H3", "V2",
			}),
			td.Flatten(map[string]string{
				"H2": "V5",
				"H3": "V3",
			}),
		)

		t.Cmp(req.Header, http.Header{
			"H1": []string{"V1", "V2"},
			"H2": []string{"V1", "V2", "V3", "V4", "V5"},
			"H3": []string{"V1", "V2", "V3"},
		})
	})

	t.Run("NewRequest header panic", func(t *td.T) {
		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.NewRequest("GET", "/path", nil, "H", "V", true) },
			"headers... can only contains string and http.Header, not bool (@ headers[2])")

		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.NewRequest("GET", "/path", nil, "H1", true) },
			`header "H1" should have a string value, not a bool (@ headers[1])`)

		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.Get("/path", true) },
			"headers... can only contains string and http.Header, not bool (@ headers[0])")

		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.Head("/path", true) },
			"headers... can only contains string and http.Header, not bool (@ headers[0])")

		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.Post("/path", nil, true) },
			"headers... can only contains string and http.Header, not bool (@ headers[0])")

		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.PostForm("/path", nil, true) },
			"headers... can only contains string and http.Header, not bool (@ headers[0])")

		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.PostMultipartFormData("/path", &tdhttp.MultipartBody{}, true) },
			"headers... can only contains string and http.Header, not bool (@ headers[0])")

		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.Patch("/path", nil, true) },
			"headers... can only contains string and http.Header, not bool (@ headers[0])")

		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.Put("/path", nil, true) },
			"headers... can only contains string and http.Header, not bool (@ headers[0])")

		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.Delete("/path", nil, true) },
			"headers... can only contains string and http.Header, not bool (@ headers[0])")
	})

	// Get
	t.Cmp(tdhttp.Get("/path", "Foo", "Bar"),
		td.Struct(
			&http.Request{
				Method: "GET",
				Header: http.Header{"Foo": []string{"Bar"}},
			},
			td.StructFields{
				"URL": td.String("/path"),
			}))

	// Head
	t.Cmp(tdhttp.Head("/path", "Foo", "Bar"),
		td.Struct(
			&http.Request{
				Method: "HEAD",
				Header: http.Header{"Foo": []string{"Bar"}},
			},
			td.StructFields{
				"URL": td.String("/path"),
			}))

	// Post
	t.Cmp(tdhttp.Post("/path", nil, "Foo", "Bar"),
		td.Struct(
			&http.Request{
				Method: "POST",
				Header: http.Header{"Foo": []string{"Bar"}},
			},
			td.StructFields{
				"URL": td.String("/path"),
			}))

	// PostForm
	t.Cmp(
		tdhttp.PostForm("/path",
			url.Values{
				"param1": []string{"val1", "val2"},
				"param2": []string{"zip"},
			},
			"Foo", "Bar"),
		td.Struct(
			&http.Request{
				Method: "POST",
				Header: http.Header{
					"Content-Type": []string{"application/x-www-form-urlencoded"},
					"Foo":          []string{"Bar"},
				},
			},
			td.StructFields{
				"URL": td.String("/path"),
				"Body": td.Smuggle(
					ioutil.ReadAll,
					[]byte("param1=val1&param1=val2&param2=zip"),
				),
			}))

	// PostMultipartFormData
	req := tdhttp.PostMultipartFormData("/path",
		&tdhttp.MultipartBody{
			Boundary: "BoUnDaRy",
			Parts: []*tdhttp.MultipartPart{
				tdhttp.NewMultipartPartString("p1", "body1!"),
				tdhttp.NewMultipartPartString("p2", "body2!"),
			},
		},
		"Foo", "Bar")
	t.Cmp(req,
		td.Struct(
			&http.Request{
				Method: "POST",
				Header: http.Header{
					"Content-Type": []string{`multipart/form-data; boundary="BoUnDaRy"`},
					"Foo":          []string{"Bar"},
				},
			},
			td.StructFields{
				"URL": td.String("/path"),
			}))
	if t.CmpNoError(req.ParseMultipartForm(10000)) {
		t.Cmp(req.PostFormValue("p1"), "body1!")
		t.Cmp(req.PostFormValue("p2"), "body2!")
	}

	// Put
	t.Cmp(tdhttp.Put("/path", nil, "Foo", "Bar"),
		td.Struct(
			&http.Request{
				Method: "PUT",
				Header: http.Header{"Foo": []string{"Bar"}},
			},
			td.StructFields{
				"URL": td.String("/path"),
			}))

	// Patch
	t.Cmp(tdhttp.Patch("/path", nil, "Foo", "Bar"),
		td.Struct(
			&http.Request{
				Method: "PATCH",
				Header: http.Header{"Foo": []string{"Bar"}},
			},
			td.StructFields{
				"URL": td.String("/path"),
			}))

	// Delete
	t.Cmp(tdhttp.Delete("/path", nil, "Foo", "Bar"),
		td.Struct(
			&http.Request{
				Method: "DELETE",
				Header: http.Header{"Foo": []string{"Bar"}},
			},
			td.StructFields{
				"URL": td.String("/path"),
			}))
}

type TestStruct struct {
	Name string `json:"name" xml:"name"`
}

func TestNewJSONRequest(tt *testing.T) {
	t := td.NewT(tt)

	t.Run("NewJSONRequest", func(t *td.T) {
		req := tdhttp.NewJSONRequest("GET", "/path",
			TestStruct{
				Name: "Bob",
			},
			"Foo", "Bar",
			"Zip", "Test")

		t.String(req.Header.Get("Content-Type"), "application/json")
		t.String(req.Header.Get("Foo"), "Bar")
		t.String(req.Header.Get("Zip"), "Test")

		body, err := ioutil.ReadAll(req.Body)
		if t.CmpNoError(err, "read request body") {
			t.String(string(body), `{"name":"Bob"}`)
		}
	})

	t.Run("NewJSONRequest panic", func(t *td.T) {
		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.NewJSONRequest("GET", "/path", func() {}) },
			"json: unsupported type: func()")

		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.PostJSON("/path", func() {}) },
			"json: unsupported type: func()")

		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.PutJSON("/path", func() {}) },
			"json: unsupported type: func()")

		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.PatchJSON("/path", func() {}) },
			"json: unsupported type: func()")

		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.DeleteJSON("/path", func() {}) },
			"json: unsupported type: func()")

		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.NewJSONRequest("GET", "/path", td.JSONPointer("/a", 0)) },
			"JSON encoding failed: json: error calling MarshalJSON for type *td.tdJSONPointer: JSONPointer TestDeep operator cannot be json.Marshal'led")

		// Common user mistake
		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.NewJSONRequest("GET", "/path", td.JSON(`{}`)) },
			`JSON encoding failed: json: error calling MarshalJSON for type *td.tdJSON: JSON TestDeep operator cannot be json.Marshal'led, use json.RawMessage() instead`)
	})

	// Post
	t.Cmp(tdhttp.PostJSON("/path", 42, "Foo", "Bar"),
		td.Struct(
			&http.Request{
				Method: "POST",
				Header: http.Header{
					"Foo":          []string{"Bar"},
					"Content-Type": []string{"application/json"},
				},
			},
			td.StructFields{
				"URL": td.String("/path"),
			}))

	// Put
	t.Cmp(tdhttp.PutJSON("/path", 42, "Foo", "Bar"),
		td.Struct(
			&http.Request{
				Method: "PUT",
				Header: http.Header{
					"Foo":          []string{"Bar"},
					"Content-Type": []string{"application/json"},
				},
			},
			td.StructFields{
				"URL": td.String("/path"),
			}))

	// Patch
	t.Cmp(tdhttp.PatchJSON("/path", 42, "Foo", "Bar"),
		td.Struct(
			&http.Request{
				Method: "PATCH",
				Header: http.Header{
					"Foo":          []string{"Bar"},
					"Content-Type": []string{"application/json"},
				},
			},
			td.StructFields{
				"URL": td.String("/path"),
			}))

	// Delete
	t.Cmp(tdhttp.DeleteJSON("/path", 42, "Foo", "Bar"),
		td.Struct(
			&http.Request{
				Method: "DELETE",
				Header: http.Header{
					"Foo":          []string{"Bar"},
					"Content-Type": []string{"application/json"},
				},
			},
			td.StructFields{
				"URL": td.String("/path"),
			}))
}

func TestNewXMLRequest(tt *testing.T) {
	t := td.NewT(tt)

	t.Run("NewXMLRequest", func(t *td.T) {
		req := tdhttp.NewXMLRequest("GET", "/path",
			TestStruct{
				Name: "Bob",
			},
			"Foo", "Bar",
			"Zip", "Test")

		t.String(req.Header.Get("Content-Type"), "application/xml")
		t.String(req.Header.Get("Foo"), "Bar")
		t.String(req.Header.Get("Zip"), "Test")

		body, err := ioutil.ReadAll(req.Body)
		if t.CmpNoError(err, "read request body") {
			t.String(string(body), `<TestStruct><name>Bob</name></TestStruct>`)
		}
	})

	t.Run("NewXMLRequest panic", func(t *td.T) {
		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.NewXMLRequest("GET", "/path", func() {}) },
			"XML encoding failed")

		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.PostXML("/path", func() {}) },
			"XML encoding failed")

		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.PutXML("/path", func() {}) },
			"XML encoding failed")

		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.PatchXML("/path", func() {}) },
			"XML encoding failed")

		dark.CheckFatalizerBarrierErr(t,
			func() { tdhttp.DeleteXML("/path", func() {}) },
			"XML encoding failed")
	})

	// Post
	t.Cmp(tdhttp.PostXML("/path", 42, "Foo", "Bar"),
		td.Struct(
			&http.Request{
				Method: "POST",
				Header: http.Header{
					"Foo":          []string{"Bar"},
					"Content-Type": []string{"application/xml"},
				},
			},
			td.StructFields{
				"URL": td.String("/path"),
			}))

	// Put
	t.Cmp(tdhttp.PutXML("/path", 42, "Foo", "Bar"),
		td.Struct(
			&http.Request{
				Method: "PUT",
				Header: http.Header{
					"Foo":          []string{"Bar"},
					"Content-Type": []string{"application/xml"},
				},
			},
			td.StructFields{
				"URL": td.String("/path"),
			}))

	// Patch
	t.Cmp(tdhttp.PatchXML("/path", 42, "Foo", "Bar"),
		td.Struct(
			&http.Request{
				Method: "PATCH",
				Header: http.Header{
					"Foo":          []string{"Bar"},
					"Content-Type": []string{"application/xml"},
				},
			},
			td.StructFields{
				"URL": td.String("/path"),
			}))

	// Delete
	t.Cmp(tdhttp.DeleteXML("/path", 42, "Foo", "Bar"),
		td.Struct(
			&http.Request{
				Method: "DELETE",
				Header: http.Header{
					"Foo":          []string{"Bar"},
					"Content-Type": []string{"application/xml"},
				},
			},
			td.StructFields{
				"URL": td.String("/path"),
			}))
}
