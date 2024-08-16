#include <stdio.h>
#include <string.h>
#include <stdlib.h>

// Registradores
char R1[10];
char R2[10];
char RT[10];
char OPCODE[10];

// Contador de Programa
int PC = 0;

// Registradores $S0 a $S4
int S0 = 8, S1 = -9, S2 = 5, S3 = 6, S4 = 1;

// Memória simulada
int memoria[100] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20};

void printProcessorState() {
    printf("PC: %d\n", PC);
    printf("$S0 = %d , $S1 = %d , $S2 = %d , $S3 = %d , $S4 = %d\n", S0, S1, S2, S3, S4);
}

int main() {
    // Loop principal do programa
    while (1) {
        // Variável para armazenar a instrução MIPS digitada pelo usuário
        char codigo[50];

        // Exibe informações atuais
        printProcessorState();
        printf("Digite a linha do MIPS em Assembly com espaço entre cada comando.\nEscreva com letras MAIÚSCULAS\n\nex: ""ADD $S0 $S2 $S3""\n");

        // Solicitação de entrada do usuário
        printf("-> ");
        
        // Usa fgets para ler a entrada de forma segura
        fgets(codigo, sizeof(codigo), stdin);

        // Verifica se o usuário digitou 0 para encerrar o programa
        if (codigo[0] == '0') {
            printf("Programa encerrado.\n");
            break;
        }

        // Extrai as partes da instrução
        sscanf(codigo, " %s %s %s %s", OPCODE, RT, R1, R2);

        // Exibe o OPCODE da instrução
        printf("\nOPCODE: %s\n", OPCODE);

        // Verifica se a instrução é uma operação aritmética ou a instrução LW
        if (strcmp(OPCODE, "ADD") == 0 || strcmp(OPCODE, "SUB") == 0 || strcmp(OPCODE, "MUL") == 0 || strcmp(OPCODE, "DIV") == 0 || strcmp(OPCODE, "ADDU") == 0 || strcmp(OPCODE, "ADDI") == 0 || strcmp(OPCODE, "LW") == 0) {
            // Variáveis para armazenar os valores dos registradores ou imediatos
            int r1 = 0, r2 = 0;

            // Exibe os registradores e seus valores
            printf("RT: %s\n", RT);
            printf("R1: %s\n", R1);
            printf("R2: %s\n", R2);

            // Obtém os valores dos registradores ou imediatos
            if (strcmp(R1, "$S0") == 0) {
                r1 = S0;
            } else if (strcmp(R1, "$S1") == 0) {
                r1 = S1;
            } else if (strcmp(R1, "$S2") == 0) {
                r1 = S2;
            } else if (strcmp(R1, "$S3") == 0) {
                r1 = S3;
            } else if (strcmp(R1, "$S4") == 0) {
                r1 = S4;
            } else {
                // Manipula caso R1 seja um valor imediato ou não corresponda aos registradores
                printf("R1 é um valor imediato ou não corresponde aos registradores\n");
            }

            // Obtém os valores dos registradores ou imediatos
            if (strcmp(R2, "$S0") == 0) {
                r2 = S0;
            } else if (strcmp(R2, "$S1") == 0) {
                r2 = S1;
            } else if (strcmp(R2, "$S2") == 0) {
                r2 = S2;
            } else if (strcmp(R2, "$S3") == 0) {
                r2 = S3;
            } else if (strcmp(R2, "$S4") == 0) {
                r2 = S4;
            } else {
                // Manipula caso R2 seja um valor imediato ou não corresponda aos registradores
                printf("R2 é um valor imediato ou não corresponde aos registradores\n");
            }

            // Executa a operação aritmética ou a instrução LW baseada no OPCODE
            if (strcmp(OPCODE, "ADD") == 0) {
                // Realiza a adição e armazena o resultado no registrador de destino
                if (strcmp(RT, "$S0") == 0) {
                    S0 = r1 + r2;
                } else if (strcmp(RT, "$S1") == 0) {
                    S1 = r1 + r2;
                } else if (strcmp(RT, "$S2") == 0) {
                    S2 = r1 + r2;
                } else if (strcmp(RT, "$S3") == 0) {
                    S3 = r1 + r2;
                } else if (strcmp(RT, "$S4") == 0) {
                    S4 = r1 + r2;
                } else {
                    printf("RT não encontrado");
                }
            }

            if (strcmp(OPCODE, "ADDI") == 0) {
                if (strcmp(RT, "$S0") == 0) {
                    S0 = r1 + atoi(R2);
                } else if (strcmp(RT, "$S1") == 0) {
                    S1 = r1 + atoi(R2);
                } else if (strcmp(RT, "$S2") == 0) {
                    S2 = r1 + atoi(R2);
                } else if (strcmp(RT, "$S3") == 0) {
                    S3 = r1 + atoi(R2);
                } else if (strcmp(RT, "$S4") == 0) {
                    S4 = r1 + atoi(R2);
                } else {
                    printf("RT não encontrado");
                }
            }

            if (strcmp(OPCODE, "ADDU") == 0) {
                if (strcmp(RT, "$S0") == 0) {
                    S0 = (abs(r1) + abs(r2));
                } else if (strcmp(RT, "$S1") == 0) {
                    S1 = (abs(r1) + abs(r2));
                } else if (strcmp(RT, "$S2") == 0) {
                    S2 = (abs(r1) + abs(r2));
                } else if (strcmp(RT, "$S3") == 0) {
                    S3 = (abs(r1) + abs(r2));
                } else if (strcmp(RT, "$S4") == 0) {
                    S4 = (abs(r1) + abs(r2));
                } else {
                    printf("RT não encontrado");
                }
            }

            if (strcmp(OPCODE, "SUB") == 0) {
                if (strcmp(RT, "$S0") == 0) {
                    S0 = r1 - r2;
                } else if (strcmp(RT, "$S1") == 0) {
                    S1 = r1 - r2;
                } else if (strcmp(RT, "$S2") == 0) {
                    S2 = r1 - r2;
                } else if (strcmp(RT, "$S3") == 0) {
                    S3 = r1 - r2;
                } else if (strcmp(RT, "$S4") == 0) {
                    S4 = r1 - r2;
                } else {
                    printf("RT não encontrado");
                }
            }

            if (strcmp(OPCODE, "MUL") == 0) {
                if (strcmp(RT, "$S0") == 0) {
                    S0 = r1 * r2;
                } else if (strcmp(RT, "$S1") == 0) {
                    S1 = r1 * r2;
                } else if (strcmp(RT, "$S2") == 0) {
                    S2 = r1 * r2;
                } else if (strcmp(RT, "$S3") == 0) {
                    S3 = r1 * r2;
                } else if (strcmp(RT, "$S4") == 0) {
                    S4 = r1 * r2;
                } else {
                    printf("RT não encontrado");
                }
            }

            if (strcmp(OPCODE, "DIV") == 0) {
                if (strcmp(RT, "$S0") == 0) {
                    S0 = r1 / r2;
                } else if (strcmp(RT, "$S1") == 0) {
                    S1 = r1 / r2;
                } else if (strcmp(RT, "$S2") == 0) {
                    S2 = r1 / r2;
                } else if (strcmp(RT, "$S3") == 0) {
                    S3 = r1 / r2;
                } else if (strcmp(RT, "$S4") == 0) {
                    S4 = r1 / r2;
                } else {
                    printf("RT não encontrado");
                }
            }

            if (strcmp(OPCODE, "LW") == 0) {
    // Obtém o registrador de destino
    char registradorDestino[10];
    strcpy(registradorDestino, strtok(RT, "$"));

    // Obtém o deslocamento da memória
    int deslocamento = atoi(R1);

    // Verifica se o registrador de destino é válido
    if (strcmp(registradorDestino, "S0") == 0 || strcmp(registradorDestino, "S1") == 0 || strcmp(registradorDestino, "S2") == 0 || strcmp(registradorDestino, "S3") == 0 || strcmp(registradorDestino, "S4") == 0) {
        // Calcula o endereço na memória
        int endereco = deslocamento;

        // Verifica se o endereço está dentro dos limites da memória
        if (endereco >= 0 && endereco < 100) {
            // Carrega o valor da memória para o registrador de destino
            if (strcmp(RT, "$S0") == 0) {
                S0 = memoria[endereco];
            } else if (strcmp(RT, "$S1") == 0) {
                S1 = memoria[endereco];
            } else if (strcmp(RT, "$S2") == 0) {
                S2 = memoria[endereco];
            } else if (strcmp(RT, "$S3") == 0) {
                S3 = memoria[endereco];
            } else if (strcmp(RT, "$S4") == 0) {
                S4 = memoria[endereco];
            } else {
                printf("RT não encontrado");
            }
        } else {
            printf("Endereço de memória inválido");
        }
    } else {
        printf("Registrador de destino inválido");
    }
}
            // Incrementa o PC em 4 após a execução da instrução
            PC += 4;
        }
    }
    return 0;
}
