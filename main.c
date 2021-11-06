#include <stdio.h>
#include <windows.h>
#include <stdlib.h>
#include <conio.h>
#include <time.h>
#include <math.h>

//Dificuldade e configurações do jogo
#define DIMENSAO_X 6
#define DIMENSAO_Y 70
#define QTD_CACTOS 2
#define CACTOS_RANDOM 50
#define ALTURA_MAX_CACTO 2
#define VELO_INICIAL_PULO -0.9
#define ACELERACAO_GRAV 0.14
#define VELO_INICIAL_CACTOS 1
#define INCREMENTO_VELO_CACTOS 0.1
//Estruturas
typedef struct{
    float pos_x;
    int pos_y;
    float velocidade;
} Dino;

typedef struct{
    int altura;
    float pos_y;
} Cacto;

void visualizar_jogo(Dino dino, Cacto cactos[]);
int tecla_pressionada();
void tamanho_terminal();
void resetar_cor();
void texto_amarelo();
void texto_verde();
void texto_vermelho();
void mostrar_dino(Dino dino);
void mover_dino(Dino *dino);
void resetar_cacto(Cacto *cacto);

//Função principal
int main(){
    Dino dino = {.pos_x = DIMENSAO_X-1, .pos_y=1, .velocidade=0};
    Cacto cactos[QTD_CACTOS];
    float velocidade_cacto = 1;

    srand(time(NULL));
    
    tamanho_terminal();

    system("cls");
    texto_amarelo();
    printf("Jogo do dinossauro pelo terminal em C!\n\nAperte espaco para pular!\n\n");
    texto_verde();
    system("pause");
    resetar_cor();

    for(int i=0; i<QTD_CACTOS; i++)
        resetar_cacto(&cactos[i]);

    for(int iteracao=0; 1; iteracao++){
        if(iteracao%100 == 0 && iteracao != 0)
            velocidade_cacto += INCREMENTO_VELO_CACTOS;
            
        mover_dino(&dino);
        
        for(int i=0; i<QTD_CACTOS; i++){
            cactos[i].pos_y -= velocidade_cacto;
            
            if(cactos[i].pos_y <= dino.pos_y && dino.pos_x > DIMENSAO_X - cactos[i].altura - 1){
                visualizar_jogo(dino, cactos);
                texto_vermelho();
                printf("\nPerdeu! hihihi\nTua pontuacao foi %d\n\a", iteracao/10);
                resetar_cor();
                system("pause");
                return 0;
            }

            if(cactos[i].pos_y < 0){
                cactos[i].pos_y = DIMENSAO_Y + rand() % CACTOS_RANDOM;
                cactos[i].altura = 1 + rand()%2;
            }
        }
        
        visualizar_jogo(dino, cactos);
        printf("\nPontuacao: %d", iteracao/10);
    }
}

//Função principal para mostrar os elementos do jogo na tela
void visualizar_jogo(Dino dino, Cacto cactos[]){
    int printou_cacto = 0;
    system("cls");
    
    for(int i=0; i<DIMENSAO_X-ALTURA_MAX_CACTO;i++){
        if(i==(int)dino.pos_x){
            printf(" ");
            mostrar_dino(dino);
        }
        printf("\n");
    }

    for(int i=DIMENSAO_X-ALTURA_MAX_CACTO; i<DIMENSAO_X;i++){
        for(int j=0; j<DIMENSAO_Y;j++){
            printou_cacto = 0;
            for(int k=0; k<QTD_CACTOS; k++){
                if(j == floor(cactos[k].pos_y) && i > floor(DIMENSAO_X - cactos[k].altura-1)){
                    texto_verde();
                    printf("%c", 178);
                    resetar_cor();
                    printou_cacto = 1;
                }
            }
            if(i == (int)dino.pos_x && j == dino.pos_y)
                mostrar_dino(dino);

            else{
                if(printou_cacto)
                    printou_cacto=0;
                else
                    printf(" ");
            }
        }
        printf("\n");
    }

     for(int i=0; i<DIMENSAO_Y; i++)
         printf("%c", 219);
}

//Mostra na tela o dino baseado na sua velocidade
void mostrar_dino(Dino dino){
    texto_amarelo();

    if(dino.velocidade>0)
        printf("%c", 219);
    else if(dino.velocidade<0)
        printf("%c", 223);
    else
        printf("%c", 220);
    
    resetar_cor();
}

//Gera uma notva posição e altura para o cacto
void resetar_cacto(Cacto *cacto){
    cacto->pos_y = DIMENSAO_Y + rand() % CACTOS_RANDOM;
    cacto->altura = 1 + rand() % ALTURA_MAX_CACTO;
}

//Move o dino a partir da sua velocidade e aceleração se o usuário apertou a tecla de pulo
void mover_dino(Dino *dino){
    dino->velocidade += ACELERACAO_GRAV;
    
    if(dino->pos_x >= DIMENSAO_X-1){
        dino->pos_x = DIMENSAO_X-1;
        dino->velocidade = 0;
    }

    if(tecla_pressionada() == ' ' && dino->pos_x >= DIMENSAO_X-1)
        dino->velocidade = VELO_INICIAL_PULO;
    
    dino->pos_x += dino->velocidade;
    
}

//Capta uma tecla pressionada e retorna ela
int tecla_pressionada(){
    int tecla;
    if(kbhit()){
    	tecla = getch();
        return tecla;
    }
    return -1;
}

//Mudam a cor do texto exibido no terminal
void texto_verde(){
    printf("\033[0;92m");
}

void texto_amarelo(){
    printf("\033[0;33m");
}

void texto_vermelho(){
    printf("\033[0;31m");
}

void resetar_cor(){
    printf("\033[0m");  
}

//Seta o tamanho do terminal para quanto vamos usar
void tamanho_terminal(){
    SMALL_RECT windowSize = {0 , 0 , DIMENSAO_Y+1 , DIMENSAO_X+1};
    SetConsoleWindowInfo(GetStdHandle(STD_OUTPUT_HANDLE), TRUE, &windowSize);
}
//Desenvolvido por Igor M. :P
