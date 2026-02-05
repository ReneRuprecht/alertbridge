FROM maven:3.9-eclipse-temurin-21 AS build

WORKDIR /app

COPY ./backend/pom.xml .
COPY ./backend/src ./src

RUN mvn -DskipTests clean package

FROM gcr.io/distroless/java21-debian12

WORKDIR /app
COPY --from=build --chown=nonroot:nonroot /app/target/*.jar ./app.jar

EXPOSE 8080
USER nonroot

CMD ["./app.jar", "--spring.profiles.active=prod"]
