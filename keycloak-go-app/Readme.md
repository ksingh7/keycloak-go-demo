```
## use old version of keycloak because of bug https://github.com/Nerzal/gocloak/issues/346

docker run -d -p 8080:8080 -e KEYCLOAK_ADMIN=admin -e KEYCLOAK_ADMIN_PASSWORD=admin quay.io/keycloak/keycloak:17.0.0-legacy
export CONTAINER_ID=$(docker ps | cut -f 1 -d ' ' | tail -1)
docker exec $CONTAINER_ID /opt/jboss/keycloak/bin/add-user-keycloak.sh -u admin -p admin
docker restart $CONTAINER_ID

```