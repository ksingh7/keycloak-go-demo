## Introduction
This repository demostrates how to use Keycloak GO Lang Client [gocloak](https://github.com/Nerzal/gocloak)

### Features Demonstrated
-  Login with username and password and get access & refresh tokens from Keycloak
-  Validate tokens using GO gin-gonic middleware before executing the actual API
-  Get user info from keycloak using access token
-  Retrospect token using gocloak client
-  Logout user and invalidate any tokens


### Keycloak on Docker

> Note : You might want to use Keycloak v17.0.0-legacy or lower, because of [this bug in gocloak](https://github.com/Nerzal/gocloak/issues/346) 

```
docker run -d -p 8080:8080 -e KEYCLOAK_ADMIN=admin -e KEYCLOAK_ADMIN_PASSWORD=admin quay.io/keycloak/keycloak:17.0.0-legacy
export CONTAINER_ID=$(docker ps | cut -f 1 -d ' ' | tail -1)
docker exec $CONTAINER_ID /opt/jboss/keycloak/bin/add-user-keycloak.sh -u admin -p admin
docker restart $CONTAINER_ID
```
- Visit http://localhost:8080 and login to keycloak admin consule with `admin` as username and `admin` as password

### Keycloak on Kubernetes / OpenShift

### Local deployment
- Git Clone the repository
```
git clone https://github.com/ksingh7/keycloak-go-demo.git
cd keycloak-go-demo/keycloak-go-app
go build . && ./keycloak-go-app
```
- Open your API Client (postman) and hit the API endpoint
- Health : `http://localhost:8081/health` to check if the API is up and running

### Deployment on Kubernetes/OpenShift
- Create a new project `keycloak`
- Navigate to `operatorhub` search for `keycloak` and install it. Make sure the project is `keycloak`
- Once `keycloak` operator installation is done, to to `Installed Operator` > `keycloak` > `Create Instance` , name your keycloak instance and hit create
- Grab Keycloak username and password
```
oc get secret -n keycloak credential-keycloak-instance -o jsonpath="{.data.ADMIN_USERNAME}" | base64 --decode

oc get secret -n keyckoak credential-keycloak-instance -o jsonpath="{.data.ADMIN_PASSWORD}" | base64 --decode
```
- Get keycloak UI route
```
oc get route -n keycloak keycloak -o jsonpath="{.spec.host}"
```
### Keycloak Configuration
- Create a new Realm
- Create a user `user1` in the new realm
  - Set `user1` credential and disable temporary password
- Create a new Client `my-go-servie`
  - `client-protocol : openid-connect`
  - `Access Type : Confidential`
  - `Valid Redirect URIs : http://localhost:8081`
  - `Credentials > Secret > Regenerate Secret`

### Preparing Keycloak Go app
- edit `.env` and update environment variables for your setup
### Using the App

#### Login
- Provide username and password in the request body and hit `http://localhost:8081/login`
```
{
    "username":"user1",
    "password":"user1"
}
```
- Response
  - Access Token, Refresh Token, Status

#### GetQuote
- User Access Token as Auth > Bearer token in your next API call and hit `http://localhost:8081/auth/getQuote`

#### Status
- User Access Token as Auth > Bearer token in your next API call and hit `http://localhost:8081/auth/status`

#### Logout
- User Access Token as Auth > Bearer token in your next API call and hit `http://localhost:8081/auth/logout`

- Verify logout works by hitting `http://localhost:8081/auth/status` again. It should throw error