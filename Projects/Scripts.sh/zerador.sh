#!/bin/bash

#primeiro ira ver se inseriu os numeros necessarios
if [ $# -ne 1 ]; then
  echo "Porfavor insira um numero"
  exit 1
fi

numero=$1

#aqui ira fazer a contagem ate o zero
while [ $numero -ge 0 ]; do
  echo -n "$numero "
  numero=$((numero - 1))
done
echo
