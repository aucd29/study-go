module hw/learngo

go 1.17

replace hw/mydict => ./mydict

require (
	hw/mydict v0.0.0-00010101000000-000000000000
	hw/myrandstring v0.0.0-00010101000000-000000000000
)

replace hw/myrandstring => ./myrandstring
