#!/bin/sh

until nc -z -v -w30 dauth_v2_db 3306; do
   echo "Waiting for database connection..."
   sleep 5
done

sleep 3

$1
