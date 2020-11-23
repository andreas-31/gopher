module example.org/gopher/hello-world

go 1.12

replace example.org/bar => ./bar

require example.org/bar v0.0.0-00010101000000-000000000000 // indirect
