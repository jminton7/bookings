package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T){
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a","b","c")
	if form.Valid(){
		t.Error("form shows valid when missing required fields")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a","b","c")
	if !form.Valid(){
		t.Error("form says it does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T){
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")
	if has{
		t.Error("form shows has field when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a","a")

	form = New(postedData)

	has = form.Has("a")
	if !has{
		t.Error("form shows doesnt have a field when it does")
	}
}


func TestForm_MinLength(t *testing.T){
	postedData := url.Values{}
	form := New(postedData)

	form.MinLength("asdasdasda", 5)

	if form.Valid() {
		t.Error("form shows minlength for non existant field")
	}

	isError := form.Errors.Get("X")
	if isError == "" {
		t.Error("should have an error but got one")
	}

	postedData = url.Values{}
	postedData.Add("asdasdasda","a")

	form = New(postedData)
	form.MinLength("asdasdasda", 500)

	if form.Valid() {
		t.Error("form shows minlength of 500 met when data is shorter")
	}

	postedData = url.Values{}
	postedData.Add("another_field","abc123")

	form = New(postedData)
	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("form shows moinlength is not met when it is")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should not have an error but got one")
	}
}

func TestForm_IsEmail(t *testing.T){
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("asd")
	if form.Valid() {
		t.Error("form shows that an invalid email is valid")
	}

	postedData = url.Values{}
	postedData.Add("email","a@a.com")

	form = New(postedData)
	
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("form shows that a valid email is invalid")
	}

	postedData = url.Values{}
	postedData.Add("email","x")

	form = New(postedData)
	
	form.IsEmail("email")
	if form.Valid() {
		t.Error("form shows that valid email when not valid")
	}
}