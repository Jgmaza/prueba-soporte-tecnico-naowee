package category

import "errors"

// Category define el rango de años de nacimiento (inclusive) de una categoría.
type Category struct {
	Code    string
	MinYear int
	MaxYear int
}

// Catálogo de categorías del evento, ordenado de mayor a menor edad.
var categories = []Category{
	{Code: "JUV", MinYear: 2009, MaxYear: 2011},
	{Code: "PRE_JUV", MinYear: 2012, MaxYear: 2014},
	{Code: "INF", MinYear: 2015, MaxYear: 2017},
}

// DiscoverCategory devuelve la categoría que le corresponde a un deportista
// según su año de nacimiento.
func DiscoverCategory(birthYear int) (string, bool) {
	for _, c := range categories {
		if birthYear >= c.MinYear && birthYear <= c.MaxYear {
			return c.Code, true
		}
	}
	return "", false
}

// Athlete es un deportista inscrito.
type Athlete struct {
	Code      string
	BirthYear int
	// CategoryCode es la categoría con la que está inscrito actualmente.
	CategoryCode string
}

// Group es un grupo de inscripción de un deporte y una categoría.
type Group struct {
	Code         string
	SportCode    string
	CategoryCode string
}

var (
	ErrNoCategory = errors.New("el deportista no tiene una categoría válida para su edad")
	ErrNoGroups   = errors.New("no hay grupos disponibles para el deportista")
)

// EvaluateSportSwap decide en qué grupos del deporte destino puede quedar
// inscrito un deportista al cambiar de deporte.
//
// Regla de negocio: la categoría de un deportista se determina SIEMPRE por su
// año de nacimiento, y debe quedar en un grupo de esa categoría en el deporte
// destino.
func EvaluateSportSwap(a Athlete, targetSport string, groups []Group) ([]Group, error) {
	cat := resolveCategory(a)
	if cat == "" {
		return nil, ErrNoCategory
	}

	eligible := eligibleGroups(cat, targetSport, groups)
	if len(eligible) == 0 {
		return nil, ErrNoGroups
	}
	return eligible, nil
}

// resolveCategory determina la categoría del deportista según su año de nacimiento.
func resolveCategory(a Athlete) string {
	cat, ok := DiscoverCategory(a.BirthYear)
	if !ok {
		return ""
	}
	return cat
}

// eligibleGroups filtra los grupos del deporte destino que son de la categoría dada.
func eligibleGroups(categoryCode, sportCode string, groups []Group) []Group {
	var out []Group
	for _, g := range groups {
		if g.SportCode == sportCode && g.CategoryCode == categoryCode {
			out = append(out, g)
		}
	}
	return out
}
