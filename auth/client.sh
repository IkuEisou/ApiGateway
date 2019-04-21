#!/bin/bash

#register
#./src/tool/authentication-api-cli/authentication-api-cli register user --payload '{"email": "jamesbond@gmail.com","first_name": "John","last_name": "Doe","password": "abcd1234"}'

#Login
./src/tool/authentication-api-cli/authentication-api-cli login user --payload '{"email": "jamesbond@gmail.com", "password": "abcd1234"}'