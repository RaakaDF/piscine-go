curl -s https://academy.digifemmes.com/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"raaka\"}}){id}}"}'|jq '.data.user[0].id'
