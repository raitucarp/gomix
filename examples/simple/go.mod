module github.com/raitucarp/gomix/examples/simple

go 1.25.3

replace github.com/raitucarp/gomix => ../../

require (
	github.com/raitucarp/gomix v1.0.0
	github.com/raitucarp/gomix/addons/google-fonts v0.0.0-20251016084232-d0e88ae6cb53
	github.com/raitucarp/gomix/icons/ant-design-icons v0.0.0-20251016063651-b6909058c740
	github.com/raitucarp/gomix/icons/fontawesome-5 v0.0.0-20251016063651-b6909058c740
)

require (
	github.com/raitucarp/gomix/icons v0.0.0-20251016064055-aee4464f7e92 // indirect
	github.com/tdewolff/minify/v2 v2.24.3 // indirect
	github.com/tdewolff/parse/v2 v2.8.3 // indirect
	golang.org/x/net v0.46.0 // indirect
)
