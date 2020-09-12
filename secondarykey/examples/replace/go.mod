module main

go 1.15

replace mypkgs => ./src

require mypkgs v0.0.0-00010101000000-000000000000 // indirect
