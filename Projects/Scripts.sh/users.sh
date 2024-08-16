#!/bin/bash

#primeiro ira ver se inseriu os parametros
if [ $# -eq 0 ]; then
  echo "Porfavor forneca os parametros."
  exit 1
fi

contador=1

#aqui ira contar as palavras separadas por espaco
while [ $# -gt 0 ]; do
  echo "Parâmetro $contador: $1"
  contador=$((contador + 1))
  shift
done