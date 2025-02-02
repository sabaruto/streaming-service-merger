services := $(subst ./services/,,$(wildcard ./services/*))
all: update-images

update-images: update-gateway-image update-service-images

update-gateway-image: gateway-image-build gateway-image-push gateway-image-pull

update-service-images: $(foreach wrd,$(services),$(wrd)-image-build $(wrd)-image-push $(wrd)-image-pull)

%-image-build:
	docker build -t sabaruto/${subst -image-build,,$@} --target ${subst -image-build,,$@} .

%-image-push:
	docker image push sabaruto/${subst -image-push,,$@}

%-image-pull:
	docker pull sabaruto/${subst -image-pull,,$@}

binaries:
	mkdir -p ./bin
	go build -o ./bin ./gateway/ ./services/*

clean:
	rm -rf ./bin
