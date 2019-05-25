.DEFAULT_GOAL := all

.PHONY : all
all : clean build

.PHONY : clean
clean :
	@(./gradlew clean)

.PHONY : build
build :
	@(./gradlew build)

.PHONY : fix-gradlew
fix-gradlew :
	@(chmod +x gradlew)

.PHONY : watch
watch :
	@(./gradlew build --continuous)
