#include <stdio.h>
// Biblioteca para o system("cls");, que serve para limpar a tela
#include <windows.h>
// Biblioteca para o rand();, que gera um n�mero aleat�rio
#include <stdlib.h>
#include <time.h>

// Dificuldade e dimensoes do jogo
#define QTD_BOMBAS 30
#define DIMENSAO_X 20
#define DIMENSAO_Y 15
// Constantes
#define BOMBA -1
#define CONHECIDO 1
#define DESCONHECIDO 0

#define FRA 1

// Campo
int campo[DIMENSAO_X][DIMENSAO_Y][2];

// Fun��es a serem utilizadas
void inicializar_campo();
void visualizar_campo(int estado);
void preencher_espacos_livres(int pos_x, int pos_y);
int checar_mina(int pos_x, int pos_y);
int checar_vitoria();
int checar_pontuacao();

int main()
{

    int linha, coluna, inspecao;
    srand(time(NULL));
    inicializar_campo();

    while (1)
    {
        visualizar_campo(0);

        if (FRA)
            printf("Entrez les coordonnees de la position qui tu veux chequer\nPar example: b5\n-", checar_pontuacao());
        else
            printf("Digite as coordenadas da posicao que que voce quer checar:\nPor exemplo: b5\n-");

        scanf(" %c%d", &coluna, &linha);
        coluna = coluna - 'a';
        // Checando o que tem na posi��o inserida
        if (linha >= 0 && linha < DIMENSAO_X && coluna >= 0 && coluna < DIMENSAO_Y)
        {
            inspecao = checar_mina(linha, coluna);

            if (inspecao == BOMBA)
            {
                visualizar_campo(BOMBA);
                printf("\033[0;31m");
                if (FRA)
                    printf("Tu as perdu!\nTa score etait %d\n", checar_pontuacao());
                else
                    printf("Perdeu! Hihihi\nA tua pontuacao foi %d\n", checar_pontuacao());
                printf("\033[0m");
                system("pause");
                return 0;
            }
        }
        // Verificando se o jogador venceu
        if (checar_vitoria())
        {
            visualizar_campo(BOMBA);
            printf("\033[0;32m");
            if (FRA)
                printf("Tres bien!!\n");
            else
                printf("Ai sim, ganhou!!\n");
            printf("\033[0m");
            system("pause");
            return 0;
        }
    }
}

void inicializar_campo()
{
    int pos_bomba_x, pos_bomba_y, bombas_postas = 0, bombas_encontradas;
    // Inicializando o campo com zeros(nenhuma bomba) e todas as posi��es desconhecidas
    for (int i = 0; i < DIMENSAO_X; i++)
    {
        for (int j = 0; j < DIMENSAO_Y; j++)
        {
            campo[i][j][0] = 0;
            campo[i][j][1] = DESCONHECIDO;
        }
    }
    // Inserindo as bombas
    while (bombas_postas < QTD_BOMBAS)
    {
        pos_bomba_x = rand() % (DIMENSAO_X - 1);
        pos_bomba_y = rand() % (DIMENSAO_Y - 1);

        if (campo[pos_bomba_x][pos_bomba_y][0] != BOMBA)
        {
            campo[pos_bomba_x][pos_bomba_y][0] = BOMBA;
            bombas_postas++;
        }
    }
    // Colocando o n�mero referente � quantidade de bombas ao redor de uma casa
    for (int i = 0; i < DIMENSAO_X; i++)
    {
        for (int j = 0; j < DIMENSAO_Y; j++)
        {
            if (campo[i][j][0] != BOMBA)
            {
                bombas_encontradas = 0;

                for (int k = i - 1; k <= i + 1; k++)
                {
                    for (int l = j - 1; l <= j + 1; l++)
                    {
                        if (k >= 0 && k < DIMENSAO_X && l >= 0 && l < DIMENSAO_Y)
                        {
                            if (campo[k][l][0] == BOMBA)
                            {
                                bombas_encontradas++;
                            }
                        }
                    }
                }

                campo[i][j][0] = bombas_encontradas;
            }
        }
    }
}

// Mostrar a matriz na tela de acordo com o que foi descoberto
void visualizar_campo(int estado = 1)
{
    system("cls");

    printf("      ");
    for (int i = 'a'; i < 'a' + DIMENSAO_Y; i++)
    {
        printf("%c  ", i);
    }
    printf("\n");

    for (int i = 0; i < DIMENSAO_X; i++)
    {
        if (i > 9)
        {
            printf(" %d -", i);
        }
        else
        {
            printf(" %d  -", i);
        }

        for (int j = 0; j < DIMENSAO_Y; j++)
        {
            if (estado == BOMBA)
            {
                if (campo[i][j][0] == BOMBA)
                {
                    printf("\033[0;31m");
                    printf("[*]");
                    printf("\033[0m");
                }
                else if (campo[i][j][0] == 0)
                {
                    printf("\033[0;32m");
                    printf("[ ]");
                    printf("\033[0m");
                }
                else
                {
                    printf("\033[0;36m");
                    printf("[%d]", campo[i][j][0]);
                    printf("\033[0m");
                }
            }
            else if (campo[i][j][1] == DESCONHECIDO)
            {
                printf("\033[0;33m");
                printf("[x]");
                printf("\033[0m");
            }
            else if (campo[i][j][0] == 0)
            {
                printf("\033[0;32m");
                printf("[ ]");
                printf("\033[0m");
            }
            else
            {
                printf("\033[0;36m");
                printf("[%d]", campo[i][j][0]);
                printf("\033[0m");
            }
        }
        printf("\n");
    }
}

// Verifica se h� mina em uma dada posi��o
int checar_mina(int pos_x, int pos_y)
{
    if (campo[pos_x][pos_y][0] != BOMBA)
    {
        campo[pos_x][pos_y][1] = CONHECIDO;
        preencher_espacos_livres(pos_x, pos_y);
        return 0;
    }
    else
    {
        return BOMBA;
    }
}

// Preenche todos os espa�os livres adjacentes (que n�o tem nenhuma bomba ao redor deles)
void preencher_espacos_livres(int pos_x, int pos_y)
{
    if (campo[pos_x][pos_y][0] == 0)
    {
        for (int linha = pos_x - 1; linha <= pos_x + 1; linha++)
        {
            for (int coluna = pos_y - 1; coluna <= pos_y + 1; coluna++)
            {
                if (linha >= 0 && linha < DIMENSAO_X && coluna >= 0 && coluna < DIMENSAO_Y)
                {
                    if (campo[linha][coluna][0] == 0 && campo[linha][coluna][1] == DESCONHECIDO)
                    {
                        campo[linha][coluna][1] = CONHECIDO;
                        preencher_espacos_livres(linha, coluna);
                    }
                    else
                    {
                        campo[linha][coluna][1] = CONHECIDO;
                    }
                }
            }
        }
    }
    else
    {
        campo[pos_x][pos_y][1] = CONHECIDO;
    }
}

// Verifica se o usu�rio checou todas as casas sem bombas
int checar_pontuacao()
{
    int qtd_conhecidos = 0;
    for (int i = 0; i < DIMENSAO_X; i++)
    {
        for (int j = 0; j < DIMENSAO_Y; j++)
        {
            if (campo[i][j][1] == CONHECIDO)
            {
                qtd_conhecidos++;
            }
        }
    }
    return qtd_conhecidos;
}

// Verifica se o usu�rio checou todas as casas sem bombas
int checar_vitoria()
{
    int qtd_desconhecidos = 0;
    for (int i = 0; i < DIMENSAO_X; i++)
    {
        for (int j = 0; j < DIMENSAO_Y; j++)
        {
            if (campo[i][j][1] == DESCONHECIDO)
            {
                qtd_desconhecidos++;
            }
        }
    }
    return qtd_desconhecidos == QTD_BOMBAS;
}

// Desenvolvido por Igor M. :P
