package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
    "fmt"
)

func TestServerHandlers(t *testing.T) {
    go StartExamples(false)
    time.Sleep(time.Second * 3)
    fmt.Printf("\n\nStarting tests...\n\n")

    testHomePathGET(t)
    testRoute1PathGET(t)
    testRoute2PathGET(t)
    testRoute3PathGET(t)
}

func testHomePathGET(t *testing.T) {
  // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
  // pass 'nil' as the third parameter.
  req, err := http.NewRequest("GET", "/", nil)
  if err != nil {
    t.Fatal(err)
  }

  // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(c.HandleFunc)

  // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
  // directly and pass in our Request and ResponseRecorder.
  handler.ServeHTTP(rr, req)

  t.Logf(
    "Testing /",
  )
  e := false

  // Check the status code is what we expect.
  if status := rr.Code; status != http.StatusOK {
    t.Errorf(
      "Error: Server handler returned wrong status code: got %v want %v",
      status,
      http.StatusOK,
    )
    e = true
  }

  // Check the response body is what we expect.
  expected := `Hello world!`
  if rr.Body.String() != expected {
    t.Errorf(
      "Error Server handler returned unexpected body: got %v want %v",
      rr.Body.String(),
      expected,
    )
    e = true
  }

  if e == false {
    t.Logf(
      "Success!",
    )
  }
}

func testRoute1PathGET(t *testing.T) {
  // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
  // pass 'nil' as the third parameter.
  req, err := http.NewRequest("GET", "/routeget1", nil)
  if err != nil {
    t.Fatal(err)
  }

  // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(c.HandleFunc)

  // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
  // directly and pass in our Request and ResponseRecorder.
  handler.ServeHTTP(rr, req)

  t.Logf(
    "Testing /routeget1",
  )
  e := false

  // Check the status code is what we expect.
  if status := rr.Code; status != http.StatusOK {
    t.Errorf(
      "Error: Server handler returned wrong status code: got %v want %v",
      status,
      http.StatusOK,
    )
    e = true
  }

  // Check the response body is what we expect.
  expected := `Hello route get 1!`
  if rr.Body.String() != expected {
    t.Errorf(
      "Error: Server handler returned unexpected body: got %v want %v",
      rr.Body.String(),
      expected,
    )
    e = true
  }

  if e == false {
    t.Logf(
      "Success!",
    )
  }
}

func testRoute2PathGET(t *testing.T) {
  // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
  // pass 'nil' as the third parameter.
  req, err := http.NewRequest("GET", "/routeget2/theparam1/and/theparam2", nil)
  if err != nil {
    t.Fatal(err)
  }

  // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(c.HandleFunc)

  // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
  // directly and pass in our Request and ResponseRecorder.
  handler.ServeHTTP(rr, req)

  t.Logf(
    "Testing /routeget2/theparam1/and/theparam2",
  )
  e := false

  // Check the status code is what we expect.
  if status := rr.Code; status != http.StatusOK {
    t.Errorf(
      "Error: Server handler returned wrong status code: got %v want %v",
      status,
      http.StatusOK,
    )
    e = true
  }

  // Check the response body is what we expect.
  expected := `Hello route get 1 with param1 = theparam1 and param2 = theparam2!`
  if rr.Body.String() != expected {
    t.Errorf(
      "Error: Server handler returned unexpected body: got %v want %v",
      rr.Body.String(),
      expected,
    )
    e = true
  }

  if e == false {
    t.Logf(
      "Success!",
    )
  }
}

func testRoute3PathGET(t *testing.T) {
  // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
  // pass 'nil' as the third parameter.
  req, err := http.NewRequest("GET", "/routeget3/theparam1/atheparam2b", nil)
  if err != nil {
    t.Fatal(err)
  }

  // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(c.HandleFunc)

  // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
  // directly and pass in our Request and ResponseRecorder.
  handler.ServeHTTP(rr, req)

  t.Logf(
    "Testing /routeget3/theparam1/atheparam2b",
  )
  e := false

  // Check the status code is what we expect.
  if status := rr.Code; status != http.StatusOK {
    t.Errorf(
      "Error: Server handler returned wrong status code: got %v want %v",
      status,
      http.StatusOK,
    )
    e = true
  }

  // Check the response body is what we expect.
  expected := `Hello! This route uses REGEX! Only URL that use parameters between 'a' and 'b'`
  if rr.Body.String() != expected {
    t.Errorf(
      "Error: Server handler returned unexpected body: got %v want %v",
      rr.Body.String(),
      expected,
    )
    e = true
  }

  if e == false {
    t.Logf(
      "Success!",
    )
  }
}

func testRenderPathGET(t *testing.T) {
  // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
  // pass 'nil' as the third parameter.
  req, err := http.NewRequest("GET", "/rendertest", nil)
  if err != nil {
    t.Fatal(err)
  }

  // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(c.HandleFunc)

  // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
  // directly and pass in our Request and ResponseRecorder.
  handler.ServeHTTP(rr, req)

  t.Logf(
    "Testing /rendertest",
  )
  e := false

  // Check the status code is what we expect.
  if status := rr.Code; status != http.StatusOK {
    t.Errorf(
      "Error: Server handler returned wrong status code: got %v want %v",
      status,
      http.StatusOK,
    )
    e = true
  }

  if e == false {
    t.Logf(
      "Success!",
    )
  }
}
