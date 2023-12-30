package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("search for value in map", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}

		got, _ := dict.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}

		_, err := dict.Search("non-existing-key")
		assertError(t, err, ErrKeyNotFound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("add new key", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}

		dict.Add("another-test", "this is just another test")

		want := "this is just another test"

		assertContainsKey(t, dict, "another-test", want)
	})
	
	
	t.Run("add existing key", func(t *testing.T) {
		key 	:= "test"
		value 	:= "this is just a test"
		dict 	:= Dictionary{"test": "this is just a test"}

		err := dict.Add(key, "this is just another test")
		
		assertError(t, err, ErrKeyAlreadyExists)
		assertContainsKey(t, dict, key, value)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("update existing key", func(t *testing.T) {
		key	:= "test"
		value	:= "old value"
		newValue:= "new value"
		dict	:= Dictionary{key: value}

		dict.Update(key, newValue)

		assertContainsKey(t, dict, key, newValue)
	})


	t.Run("update non-existing key", func(t *testing.T) {
		key	:= "test"
		value	:= "value"
		dict	:= Dictionary{}

		err := dict.Update(key, value)

		assertError(t, err, ErrKeyNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing key", func(t *testing.T) {
		key 	:= "test"
		value	:= "value"
		dict	:= Dictionary{key: value}

		dict.Delete(key)

		_, err := dict.Search(key)
		want := ErrKeyNotFound

		assertError(t, err, want)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil || got != want {
		t.Fatalf("expected %#v but got %#v", want, got)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertContainsKey(t testing.TB, dict Dictionary, key, want string) {	
	got, _ := dict.Search(key)

	assertStrings(t, got, want)
}
