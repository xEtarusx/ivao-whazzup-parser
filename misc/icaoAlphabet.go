package misc

func GetIcaoCodeWordByLetter(letter string) string {
	m := make(map[string]string)

	m["A"] = "Alpha"
	m["B"] = "Bravo"
	m["C"] = "Charlie"
	m["D"] = "Delta"
	m["E"] = "Echo"
	m["F"] = "Foxtrot"
	m["G"] = "Golf"
	m["H"] = "Hotel"
	m["I"] = "India"
	m["J"] = "Juliett"
	m["K"] = "Kilo"
	m["L"] = "Lima"
	m["M"] = "Mike"
	m["N"] = "November"
	m["O"] = "Oscar"
	m["P"] = "Papa"
	m["Q"] = "Quebec"
	m["R"] = "Romeo"
	m["S"] = "Sierra"
	m["T"] = "Tango"
	m["U"] = "Uniform"
	m["V"] = "Victor"
	m["W"] = "Whiskey"
	m["X"] = "X-ray"
	m["Y"] = "Yankee"
	m["Z"] = "Zulu"

	return m[letter]
}
