#!/bin/bash

#primeiro ira ver se inseriu as palavras necessarias
if [ $# -ne 2 ]; then
  echo "Porfavor insira duas palavras"
  exit 1
fi

palavra1="$1"
palavra2="$2"

#aqui ira ver se a primeira palavra esta contida na segunda
if echo "$palavra2" | grep -q "$palavra1"; then
  echo "$palavra1 está contida em $palavra2"
fi