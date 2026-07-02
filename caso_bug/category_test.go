package category

import "testing"

// Grupos disponibles en el evento para las pruebas.
func sampleGroups() []Group {
	return []Group{
		{Code: "G-BASKET-JUV", SportCode: "BALONCESTO", CategoryCode: "JUV"},
		{Code: "G-FUTSAL-PREJUV", SportCode: "FUTSAL", CategoryCode: "PRE_JUV"},
		{Code: "G-FUTSAL-JUV", SportCode: "FUTSAL", CategoryCode: "JUV"},
		{Code: "G-NATACION-PREJUV", SportCode: "NATACION", CategoryCode: "PRE_JUV"},
	}
}

func codes(groups []Group) []string {
	out := make([]string, 0, len(groups))
	for _, g := range groups {
		out = append(out, g.Code)
	}
	return out
}

func TestDiscoverCategory(t *testing.T) {
	cases := map[int]string{2011: "JUV", 2013: "PRE_JUV", 2016: "INF"}
	for year, want := range cases {
		got, ok := DiscoverCategory(year)
		if !ok || got != want {
			t.Errorf("DiscoverCategory(%d) = %q,%v; want %q,true", year, got, ok, want)
		}
	}
}

// Deportista sin categoría inscrita: se determina por su año de nacimiento
// (2013 -> PRE_JUV) y se cambia a FUTSAL.
func TestSportSwap_CategoryFromBirthYear(t *testing.T) {
	a := Athlete{Code: "A1", BirthYear: 2013}

	got, err := EvaluateSportSwap(a, "FUTSAL", sampleGroups())
	if err != nil || len(got) != 1 || got[0].Code != "G-FUTSAL-PREJUV" {
		t.Fatalf("esperaba [G-FUTSAL-PREJUV] sin error; got %v err=%v", codes(got), err)
	}
}

// Deportista nacido en 2011, inscrito como PRE_JUV, que se cambia a BALONCESTO
// (deporte que solo tiene grupo JUV).
func TestSportSwap_InconsistentCategory(t *testing.T) {
	a := Athlete{Code: "A2", BirthYear: 2011, CategoryCode: "PRE_JUV"}

	got, err := EvaluateSportSwap(a, "BALONCESTO", sampleGroups())
	if err != nil || len(got) != 1 || got[0].Code != "G-BASKET-JUV" {
		t.Fatalf("esperaba [G-BASKET-JUV] sin error; got %v err=%v", codes(got), err)
	}
}

// Deportista de 2016 (INF) que se cambia a NATACION, donde no hay grupos de su
// categoría: debe devolver ErrNoGroups.
func TestSportSwap_NoGroupsInTargetSport(t *testing.T) {
	a := Athlete{Code: "A3", BirthYear: 2016}

	_, err := EvaluateSportSwap(a, "NATACION", sampleGroups())
	if err != ErrNoGroups {
		t.Fatalf("esperaba ErrNoGroups; got %v", err)
	}
}
