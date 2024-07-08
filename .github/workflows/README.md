***Projeto de exemplo de pipeline com o Git hub actions***


# SONARQUBE
docker run \
    --rm \
    -v ".:/usr/src" \
    --network="host" \
    -e SONAR_HOST_URL="http://localhost:9000" \
    -e SONAR_SCANNER_OPTS="-Dsonar.projectKey=go-project" \
    -e SONAR_TOKEN="sqp_f2fcf74ebc34ccb7e9b524cf88916763fd97c741" \
    sonarsource/sonar-scanner-cli

/Applications/sonar-scanner-6.1.0.4477-macosx-x64

docker run -rm --network="host" sonarsource/sonar-scanner-cli


# Command to overload properties files
docker run \
    --rm \
    -v ".:/usr/src" \
    --network="host" \
    -v $PWD/sonar-project.properties:/usr/lib/sonar-scanner/conf/sonar-scanner.properties \
    sonarsource/sonar-scanner-cli

