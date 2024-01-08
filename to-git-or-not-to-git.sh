#!/bin/bash
jq_url="https://learn.zone01oujda.ma/assets/superhero/all.json"
curl -s $jq_url | jq -r '.[] | select(.id==170) | .name'
curl -s $jq_url | jq -r '.[] | select(.id==170) | .powerstats.power'
curl -s $jq_url | jq -r '.[] | select(.id==170) | .appearance.gender'

