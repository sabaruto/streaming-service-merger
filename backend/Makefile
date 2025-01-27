services := $(subst ./services/,,$(wildcard ./services/*))
all: services

images: service-images

service-images: $(foreach wrd,$(services),$(wrd)-image)

%-image:
	docker build -t sabaruto/${subst -image,,$@} --target ${subst -image,,$@} .

binaries: service-binaries

service-binaries: $(foreach wrd,$(services),$(wrd)-binary)

%-binary:
	go build -o ./bin/ "./services/${subst -binary,,$@}"

clean:
	rm -rf ./bin
