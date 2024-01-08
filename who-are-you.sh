#!/bin/bash
jq_url="https://learn.zone01oujda.ma/assets/superhero/all.json"
subject_id=70
name=$(curl -s "$jq_url" | jq -r ".[] | select(.id == $subject_id) | .name")
echo "$name"
