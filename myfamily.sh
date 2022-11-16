#! \bin\bash

 curl https://academy.digifemmes.com/assets/superhero/all.json | jq --argjson id "$HERO_ID" '. [] | select(.id == $id) | .connections.relatives' | sed 's/^"\(.*\)"$/\1/'