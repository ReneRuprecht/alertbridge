
.PHONY: unit-test integration-test clean-install

clean-install:
	./mvnw clean
	./mvnw install -DskipTests

unit-test:
	./mvnw clean
	./mvnw test

integration-test:
	${MAKE} clean-install
	./mvnw failsafe:integration-test
