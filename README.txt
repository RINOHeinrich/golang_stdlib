- Single Responsability: create an auth microservice
-It 'll be based on OAuth2 authentication
-It 'll use JWT with PKI (Private and public key)
-Private key will be used for signing token
-public key will be used  for verifying token
- Each microservice using this authentication microservice  will use public key for verification
- This microservice will be interfaces with HTTP REQUEST ("I 'll not call it REST API to respect Roy Fielding)
- We 'll use the standard library net/http and database/sql package for performance purpose
no ORM, no Web Framework (also for challenges)

OAuth2 principle:
-user will send username/password
-if username/password is correct:
-the server will send two kind of token:
authentication token and refresh token
-else:
- the server will send error
- the user will use the auth token to get access to service
- when auth token expired, user will send refresh_token to server
for getting new pair/token

JWT principle:
- JWT is the tool that we 'll use to generate token
- a JWT token have 3 part:
.hashing method
.payload
.Signature
- the important part is the Signature: it 'll verify that the token is from the server
- On the payload we'll just insert user information
- for signing the token, on JWT we have two method:
- HMAC method using SECRET_KEY
- PKI method using a key cut: private and public
- We'll use private key for signing token and public key for verifying it