#!/bin/sh
set -e


echo "================================================================"
echo "===========================DEPLOYMENT==========================="
echo "================================================================"

echo "Update codebase..."
cd /home/fanz/project/synapsis-test
git fetch origin main
git reset --hard origin/main

echo "Installing dependencies 🛠"
go mod tidy

echo "Restart pm2 service backend 🔥"
pm2 restart running.json

echo "Deploying Backend Application Successfully Yeayyyy ......."