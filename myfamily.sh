#!bin/bash
Api_curl="https://learn.zone01oujda.ma/assets/superhero/all.json"
curl -s "$Api_curl" | jq ".[] | select(.id == $HERO_ID) | .connections .relatives" | tr -d "\""
