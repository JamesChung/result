package result_test

import (
	"errors"
	"testing"

	"github.com/JamesChung/result"
)

func TestNew(t *testing.T) {
	value := 10
	r := result.New(value, nil)
	if r.IsError() {
		t.Fail()
	}
	if r.Unwrap() != value {
		t.Fail()
	}
}

func TestIsOk(t *testing.T) {
	value := 10
	r := result.New(value, nil)
	if r.Unwrap() != value {
		t.Fail()
	}
	if !r.IsOk() {
		t.Fail()
	}
	if r.IsError() {
		t.Fail()
	}
}

func TestIsErr(t *testing.T) {
	err := errors.New("test error")
	r := result.New[*int](nil, err)
	if !r.IsError() {
		t.Fail()
	}
}

func TestUnwrap(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		value := 10
		r := result.New(value, nil)
		if r.Unwrap() != value {
			t.Fail()
		}
	})
	t.Run("Successfully panic", func(t *testing.T) {
		defer func() {
			if err := recover(); err == nil {
				t.Fail()
			}
		}()
		r := result.New[*int](nil, errors.New("panic"))
		r.Unwrap()
	})
}

func TestUnwrapOr(t *testing.T) {
	t.Run("Successful Unwrap", func(t *testing.T) {
		value := 10
		r := result.New(value, nil)
		if r.UnwrapOr(0) != value {
			t.Fail()
		}
	})
	t.Run("Return or value", func(t *testing.T) {
		value := 10
		r := result.New(0, errors.New("no value"))
		if r.UnwrapOr(value) != value {
			t.Fail()
		}
	})
}

func TestUnwrapOrDefault(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		value := 10
		r := result.New(value, nil)
		if r.UnwrapOrDefault() != value {
			t.Fail()
		}
	})
	t.Run("Returns default", func(t *testing.T) {
		var defaultValue int
		r := result.New(10, errors.New("error"))
		if r.UnwrapOrDefault() != defaultValue {
			t.Fail()
		}
	})
}

func TestUnwrapErr(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		err := errors.New("test error")
		r := result.New[*int](nil, err)
		if r.UnwrapErr() != err {
			t.Fail()
		}
	})
	t.Run("Successfully panic", func(t *testing.T) {
		defer func() {
			if err := recover(); err == nil {
				t.Fail()
			}
		}()
		r := result.New(0, nil)
		r.UnwrapErr()
	})
}
