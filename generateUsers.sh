#!/usr/bin/env bash

truncate -s 0 users.htpasswd
echo "server: 'example.com'" >> config.yaml
echo "cookieName: 'ccc'" >> config.yaml
echo "credentials:" >> config.yaml

for ID in {1..100}; do
	USERNAME="student$ID"
	PASSWORD="student$ID"

	echo " - username: $USERNAME" >> config.yaml
	echo "   password: $PASSWORD" >> config.yaml

	htpasswd -b users.htpasswd $USERNAME $PASSWORD
done


