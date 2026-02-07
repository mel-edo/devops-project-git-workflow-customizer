IMAGE_NAME := seryn:dev

TEST_DIR := $(HOME)/Homework/seryn-files-test

.PHONY: build run clean

build:
	docker build -t $(IMAGE_NAME) -f infrastructure/docker/Dockerfile .

run: build
	docker run --rm \
			--user $$(id -u):$$(id -g) \
			-e GIT_AUTHOR_NAME="$$(git config --global user.name)" \
			-e GIT_AUTHOR_EMAIL="$$(git config --global user.email)" \
			-e GIT_COMMITTER_NAME="$$(git config --global user.name)" \
			-e GIT_COMMITTER_EMAIL="$$(git config --global user.email)" \
			-v "$(TEST_DIR)":/repo \
			$(IMAGE_NAME) --repo /repo $(ARGS)

clean:
	docker rmi $(IMAGE_NAME)