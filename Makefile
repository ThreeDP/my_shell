################  COLORS	##################
L_RED					=			\033[0;31m
L_REDB					=			\033[1;31m
L_WHITE					=			\033[0;37m
L_WHITEB				=			\033[1;37m
L_YELLOW				=			\033[0;33m
L_YELLOWB				=			\033[1;33m
L_GREEN					=			\033[0;32m
L_GREENB				=			\033[1;32m
################ CONFIG		###################
CONTAINER_NAME			=			shell_go
CONTAINER_VERSION		=			latest
IMAGE_NAME				=			shell
PROD_PATH				=			$(pwd)/prod/
################ RULES		###################

all:
	@echo "$(L_GREENB)____ ___  ___  ___  _         _  _$(L_WHITE)"
	@echo "$(L_GREENB)<__ /| . \| . \/ __>| |_  ___ | || |$(L_WHITE)"
	@echo "$(L_GREENB) <_ \| | ||  _/\__ \| . |/ ._>| || |$(L_WHITE)"
	@echo "$(L_GREENB)<___/|___/|_|  <___/|_|_|\___.|_||_|$(L_WHITE)"
	@go run .

# Build the image and run in a container docker
bd:
	docker build -t $(CONTAINER_NAME):$(CONTAINER_VERSION) .
	@echo "$(L_GREENB)[built docker image]$(L_WHITE)"
	docker run --name $(IMAGE_NAME) -v $(PROD_PATH) -d $(CONTAINER_NAME):$(CONTAINER_VERSION)
	@echo "$(L_GREENB)[run docker container]$(L_WHITE)"

# Execute the bash in the container
rd:
	docker exec -it $(IMAGE_NAME) bash
	@echo "$(L_GREENB)[Finish work in container]$(L_WHITE)"


