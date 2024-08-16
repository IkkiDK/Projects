#!/bin/bash

if [ -e "/etc/passwd" ]; then
  cut -d: -f1,5 /etc/passwd
else
  echo "O arquivo /etc/passwd não foi encontrado."
  #No meu computador n achou o arquivo, estou usando o VScode com git bash, pode ser por isso
  exit 1
fi