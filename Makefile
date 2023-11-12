appName = UniversalHealthCheck

ifeq ($(OS),Windows_NT)     		# is Windows_NT on XP, 2000, 7, Vista, 10...
    detected_OS := Windows
	binName = unihc.exe
else
    detected_OS := $(shell uname)  	# same as "uname -s"
	binName = unihc
endif


### Definitions ###

runApp = go run ./src/$(appName).go
build = go build -o bin/$(binName) ./src/.
clean = rm ./bin/*


### Tasks ###

# Run application.
run:
	@$(runApp)

# Clean bin folder.
clean:
	$(clean)

# Build application.
build:
	$(build)
