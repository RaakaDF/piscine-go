 #! /bin/bash

 echo "raaka"
 curl -s https://academy.digifemmes.com/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"raaka\"}}){id}}"}
 