.PHONY: new run tidy list

# make new name=auth
new:
	@test -n "$(name)" || (echo "Usage: make new name=folder_name" && exit 1)
	@test ! -d "$(name)" || (echo "Folder $(name) already exists, nothing overwritten" && exit 1)
	@mkdir -p $(name)
	@cp _template/main.go $(name)/main.go
	@echo "✓ $(name)/ created -> make run name=$(name)"

# make run name=auth   (default: okapi-basis)
run:
	go run ./$(or $(name),okapi-basis)

tidy:
	go mod tidy

list:
	@ls -d */ | grep -v '^_'