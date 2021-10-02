package tempconv

func FToM(f Feet) Miler { return Miler(f * 0.3048) }

func MToF(m Miler) Feet { return Feet(m / 0.3048) }

func PToK(p Pound) Kilogram { return Kilogram(p / 2.2046) }

func KToP(k Kilogram) Pound { return Pound(k * 2.2046) }
