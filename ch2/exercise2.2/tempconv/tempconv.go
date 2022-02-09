package tempconv

import "fmt"

type Feet float64
type Miler float64

type Pound float64
type Kilogram float64

func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }
func (m Miler) String() string { return fmt.Sprintf("%gm", m) }

func (p Pound) String() string    { return fmt.Sprintf("%glb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
