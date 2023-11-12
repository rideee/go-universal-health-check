ifeq ($(OS),Windows_NT)     		# is Windows_NT on XP, 2000, 7, Vista, 10...
    detected_OS := Windows
	binName = unihc.exe
else
    detected_OS := $(shell uname)  	# same as "uname -s"
	binName = unihc
endif


### Tasks ###

# Run application.
devrun:
	@go run ./src/

# Clean bin folder.
clean:
	$(clean)

# Build application.
build:
	go build -o bin/$(binName) ./src/

# Run the built application.
run:
	./bin/$(binName)
