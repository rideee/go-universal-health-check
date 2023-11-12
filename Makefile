srcDir=./src
outDir=./out

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
	@go run $(srcDir)

# Clean bin folder.
clean:
	@echo -e "\n--> Cleaning ./out folder"
	rm -rf $(outDir)/* 2> /dev/nul || echo "clean: Nothing to do."

# Build application.
build:
	@echo -e "\n--> Building project"
	go build -o $(outDir)/$(binName) $(srcDir)

	@echo -e "\n--> Copying required files"
	test -d $(outDir)/config || mkdir $(outDir)/config
	cp -rf $(srcDir)/config/* $(outDir)/config/
# Run the built application.
run:
	@echo -e "\n--> Running application: $(outDir)/$(binName)\n"
	@$(outDir)/$(binName)
