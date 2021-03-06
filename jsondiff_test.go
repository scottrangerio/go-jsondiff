package jsondiff

import "testing"
import "strings"
import "reflect"

func TestDecode(t *testing.T) {
	r := strings.NewReader(`{ "foo": "bar" }`)
	e := map[string]interface{}{
		"foo": "bar",
	}

	g, err := Decode(r)

	if err != nil {
		t.Error(err)
	}

	eq := reflect.DeepEqual(g, e)

	if !eq {
		t.Errorf("Wrong result:\ngot: %v\nwant: %v\n", g, e)
	}
}

func TestAlignKeys(t *testing.T) {
	keysA := []string{"foo", "bar"}
	keysB := []string{"bar", "baz"}
	expected := []string{"foo", "bar", "baz"}
	ch := make(chan string)

	go AlignKeys(keysA, keysB, ch)

	got := make([]string, 0, 3)
	for key := range ch {
		got = append(got, key)
	}

	if l := len(got); l != 3 {
		t.Errorf("Wrong number of elements:\ngot: %v, want: %v\n", l, 3)
	}

	eq := reflect.DeepEqual(got, expected)

	if !eq {
		t.Errorf("Wrong result:\ngot: %v\nwant: %v\n", got, expected)
	}
}

func TestKeysFromMap(t *testing.T) {
	m := map[string]interface{}{
		"foo": struct{}{},
		"bar": struct{}{},
		"baz": struct{}{},
	}

	expected := []string{"foo", "bar", "baz"}

	keys := KeysFromMap(m)

	eq := reflect.DeepEqual(keys, expected)

	if !eq {
		t.Errorf("Wrong result:\ngot: %v\nwant: %v\n", keys, expected)
	}
}
