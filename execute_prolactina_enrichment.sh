#!/bin/bash
docker compose exec -T db psql -U plenya_user -d plenya_db -f - < /home/user/plenya/scripts/enrich_prolactina_homens.sql
