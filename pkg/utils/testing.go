package utils

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/domain"
)

func AssertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("expected no error but got error: %v", err)
	}
}

func AssertUserId(t testing.TB, got, want uint64) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertError(t testing.TB, err error) {
	t.Helper()

	if err == nil {
		t.Error("expected error but didn't get one")
	}
}

func AssertNoApplicationError(t testing.TB, err *ApplicationError) {
	t.Helper()

	if err != nil {
		t.Errorf("expected no error but got error: %v", err)
	}
}

func AssertUserFound(t testing.TB, got, want *domain.User) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertApplicationError(t testing.TB, err *ApplicationError, expectedStatusCode int, expectedCode string) {
	t.Helper()

	if err == nil || err.StatusCode != expectedStatusCode || err.Code != expectedCode {
		t.Error("expected error but got no error")
	}
}

func GenerateIntSlice(size int) []int {
	res := make([]int, size)

	for i := 0; i < size; i++ {
		res = append(res, rand.Int())
	}

	return res
}

func AssertEqualInstance(t testing.TB, got, want interface{}) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}