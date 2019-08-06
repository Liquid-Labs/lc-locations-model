package locations_test

import (
  "database/sql"
  "reflect"
  "testing"

  "github.com/Liquid-Labs/lc-entities-model/go/entities"
  "github.com/stretchr/testify/assert"

  // the package we're testing
  . "github.com/Liquid-Labs/lc-locations-model/go/locations"
)

func TestLocationsClone(t *testing.T) {
  orig := &Location{
    entities.InternalID(1),
    `a`,
    `b`,
    `c`,
    `d`,
    `e`,
    sql.NullFloat64{2.0, true},
    sql.NullFloat64{3.0, true},
    []string{`f`, `g`},
  }
  clone := orig.Clone()
  assert.Equal(t, orig, clone, "Clone does not match.")

  clone.LocationID = entities.InternalID(2)
  clone.Address1 = `z`
  clone.Address2 = `y`
  clone.City = `x`
  clone.State = `w`
  clone.Zip = `u`
  clone.Lat = sql.NullFloat64{4.0, true}
  clone.Lng = sql.NullFloat64{5.0, true}
  clone.ChangeDesc = []string{`t`}

  oReflection := reflect.ValueOf(orig).Elem()
  cReflection := reflect.ValueOf(clone).Elem()
  for i := 0; i < oReflection.NumField(); i++ {
    assert.NotEqualf(
      t,
      oReflection.Field(i).Interface(),
      cReflection.Field(i).Interface(),
      `Fields '%s' unexpectedly match.`,
      oReflection.Type().Field(i),
    )
	}
}
