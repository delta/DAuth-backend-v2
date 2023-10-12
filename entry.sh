#!/bin/sh

until nc -z -v -w30 dauth_v2_db 3306; do
   echo "Waiting for database connection..."
   sleep 5
done

sleep 3

if [[ "$1" == "DEV" ]]
then
   reflex -s -r '\.go$$' go run main.go
else
   $1
fi
