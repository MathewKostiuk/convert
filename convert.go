// General purpose unit conversion. Temperatures, length, weight, etc.
package convert

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Meters float64
type Feet float64
type Kilograms float64
type Pounds float64
type Ounce float64
type Millilitres float64

func (c Celsius) String() string     { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string  { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string      { return fmt.Sprintf("%g°K", k) }
func (m Meters) String() string      { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string        { return fmt.Sprintf("%g°ft.", f) }
func (k Kilograms) String() string   { return fmt.Sprintf("%gkg", k) }
func (p Pounds) String() string      { return fmt.Sprintf("%glbs.", p) }
func (o Ounce) String() string       { return fmt.Sprintf("%goz", o) }
func (m Millilitres) String() string { return fmt.Sprintf("%gmm", m) }
