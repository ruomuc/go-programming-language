package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c - Celsius(AbsoluteZeroC)) }

func KToC(k Kelvin) Celsius { return Celsius(k + Kelvin(AbsoluteZeroC)) }

func FToK(f Fahrenheit) Kelvin { return Kelvin((f-32)/1.8000 + Fahrenheit(AbsoluteZeroC)) }

func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-Kelvin(AbsoluteZeroC))*1.8000 + 32.00) }
