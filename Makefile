MODULES := auth user
wire:
	go install github.com/google/wire/cmd/wire@latest
	for dir in $(MODULES); do \
		cd internal/$$dir/injection && wire; \
		cd -; \
	done