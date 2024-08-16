#!/bin/bash

#primeiro ira ver se inseriu os numeros necessarios
if [ $# -ne 2 ]; then
  echo "Porfavor insira os numeros"
  exit 1
fi

numero1=$1
numero2=$2

#aqui vai fazer a verificacao dos valores e printar a relacao
if [ $numero1 -lt $numero2 ]; then
  echo "$numero1 é menor que $numero2"
elif [ $numero1 -gt $numero2 ]; then
  echo "$numero1 é maior que $numero2"
else
  echo "$numero1 é igual a $numero2"
fi
