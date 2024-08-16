#!/bin/bash

#primeiro ira ver se inseriu os caractesres necessarios
if [ $# -eq 0 ]; then
  echo "Digite os caracteres"
  exit 1
fi

#aqui ira juntar todas as palavras
tudojunto=$(echo "$@" | tr -d ' ')

echo "$tudojunto"
