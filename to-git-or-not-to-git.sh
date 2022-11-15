 #! /bin/bash

 echo "7719"
 curl -s https://academy.digifemmes.com/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"raaka\"}}){id}}"}'
