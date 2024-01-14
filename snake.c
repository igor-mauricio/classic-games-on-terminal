#include <stdio.h>

//Biblioteca para o rand();, que gera um inteiro aleatório
#include <stdlib.h>

//Biblioteca para o kbhit(); e getch();, que servem para verificar se alguma tecla foi pressionada e capturar essa tecla, respectivamente
#include <conio.h>

//Biblioteca para o system("cls");, que serve para limpar a tela e Sleep(int tempo);, que serve para pausar o programa por um dado tempo em ms
#include <Windows.h>

//Essa vai ser a dimensão do campo
#define DIMENSAO 11

//Esse vai ser o tempo de cada tick do jogo
#define TEMPO 500

//Definindo as 4 direções como inteiros quaisquer:
#define CIMA 1
#define BAIXO -1
#define DIREITA 2
#define ESQUERDA -2

//As informações relativas à "cabeça" da cobra:

typedef struct{
	int posicao_x;
	int posicao_y;
	int direcao;
	int tamanho;
}cobrinha;

int teclaPressionada();
void colocarFrutinha(int[DIMENSAO][DIMENSAO]);
void inicializarJogo(cobrinha*, int[DIMENSAO][DIMENSAO]);
void mostrarCampo(int[DIMENSAO][DIMENSAO]);
void mudarDirecao(cobrinha*);
int verificarVitoria(int[DIMENSAO][DIMENSAO]);
void decairPartesDaCobra(int[DIMENSAO][DIMENSAO]);
void colocarCobraNaPosicao(cobrinha*, int[DIMENSAO][DIMENSAO], int, int);
int rastejar(cobrinha*, int[DIMENSAO][DIMENSAO]);


int main(){
	
	/*Esse vetor vai armazenar o que tem em cada posição do campo. Inicializando todas as posições como 0 (Sem nada)
	-1 é onde há uma frutinha
	Algum número maior ou igual a 1 é uma parte da cobrinha. O inteiro indica a quantidade de ticks que esse pedaço vai continuar ali (tempo de vida da parte).
	A cada tick, as partes da cobrinha são decrementadas em 1 (exceto quando a cobrinha come uma fruta) e uma parte de cobrinha é adicionada na posição seguinte na direção da cobra
	de modo a parecer que ela está rastejando*/
	
	int campo[DIMENSAO][DIMENSAO] = {0}, foi_possivel_rastejar;
	cobrinha snake;
	
	inicializarJogo(&snake, campo);
	
	mostrarCampo(campo);
	printf("\nAperte qualquer tecla para iniciar!\nVoce controla a cobrinha com as setinhas do teclado!");
	getch();
	getch();
	
	for(int i=0;1;i++){
		
		//Tentando fazer a cobra rastejar e verificando o que houve ao fazer isso
		foi_possivel_rastejar = rastejar(&snake,campo);
		
		if(foi_possivel_rastejar==-1){
			//O \a faz um barulhinho
			printf("\a\nA cobrinha faleceu :c\nPontuacao: %d", snake.tamanho);
			return 0;
		}else if(foi_possivel_rastejar==1){
			printf("\nVoce venceu :o");
			return 0;
		} else {
			
			system("cls");
			mostrarCampo(campo);
			
			//Isso faz o programa pausar por um dado tempo
			Sleep(TEMPO);
			
			// Caso o usuário aperte uma tecla nesse meio tempo, isso poderá mudar a direção que a cobrinha irá andar
			mudarDirecao(&snake);
		}
	}
}


//Essa função vai colocar a cobrinha no centro do campo e apontada para cima, guardar as informações da cobrinha na variável snake e colocar uma frutinha no campo

void inicializarJogo(cobrinha *snake, int campo[DIMENSAO][DIMENSAO]){
	
	snake->posicao_x = (int) DIMENSAO/2;
	snake->posicao_y = (int) DIMENSAO/2;
	snake->direcao = CIMA;
	snake->tamanho = 1;
	campo[snake->posicao_y][snake->posicao_x] = 1;
	colocarFrutinha(campo);
}

/*Essa função vai gerar um número aleatório e verificar se nesse local da matriz há um espaço vazio (0) para colocar a frutinha
caso não ache, vai ficar verificando a posição seguinte da matriz até achar um espaço vazio e assim colocar a frutinha (-1)*/

void colocarFrutinha(int campo[DIMENSAO][DIMENSAO]){
	
	//Atribuindo um número aleatório entre 0 e (DIMENSAO - 1) para linha e coluna para colocar a frutinha na matriz
	int linha_frutinha = rand() % (DIMENSAO-1);
	int coluna_frutinha = rand() % (DIMENSAO-1);
	int frutinha_colocada = 0;
	
	
	for(int i=linha_frutinha;i<DIMENSAO;i++){
		for(int j=coluna_frutinha;j<DIMENSAO;j++){
			if(campo[i][j]==0){
				campo[i][j]=-1;
				frutinha_colocada=1;
				break;
			}
		}
		if(frutinha_colocada)
			break;
	}
	
	//Se chegou no fim da matriz e ainda não achou um espaço vazio, vai checando desde o início (0,0)
	if(frutinha_colocada==0){
		
		for(int i=0;i<DIMENSAO;i++){
			for(int j=0;j<DIMENSAO;j++){
				if(campo[i][j]==0){
					campo[i][j]=-1;
					frutinha_colocada=1;
					break;
				}
			}
			if(frutinha_colocada)
				break;
		}
	}
}

//Essa função printa a matriz do campo, de acordo com seus elementos

void mostrarCampo(int campo[DIMENSAO][DIMENSAO]){
	
	for(int i=0;i<DIMENSAO;i++){
		for(int j=0;j<DIMENSAO;j++){	
			switch(campo[i][j]){
				
				//Espaço vazio
				case 0:
					printf("[ ]");
				break;
				
				//Frutinha
				case -1:
					printf("[%c]", 155);
				break;
				
				//Cobrinha
				default:
					printf("[%c]", 254);
				break;
				
			}
		}
		printf("\n");
	}
}

//Essa função vai verificar se alguma tecla foi pressionada e retornar seu inteiro associado ao caractere ou -1 caso nenhuma tecla seja pressionada

int teclaPressionada(){
	
    int tecla;
    if(kbhit()){
    	getch();
    	tecla = getch();
        return tecla;
    }
    return -1;
}

// Muda a direção da cobrinha caso alguma tecla foi apertada, menos para a direção contrária à atual

void mudarDirecao(cobrinha *snake){
	
	switch(teclaPressionada()){
		case 75:
			if(snake->direcao!=DIREITA)
				snake->direcao=ESQUERDA;
			break;
		case 72:
			if(snake->direcao!=BAIXO)	
				snake->direcao=CIMA;
			break;
		case 77:
			if(snake->direcao!=ESQUERDA)	
				snake->direcao=DIREITA;
			break;
		case 80:
			if(snake->direcao!=CIMA)	
				snake->direcao=BAIXO;
			break;
	}
}

/*Essa função ia se chamar andar, mas cobra não anda!
Ela projeta a "cabeça" da cobra para a próxima posição na matriz de acordo com a direção que a cobrinha está apontando, e decrementa em 1 o tempo de vida das outras partes da cobrinha,
mas antes verificando se é possível fazer isso ( Se não for bater nela mesma ou numa parede)
Ou se vai comer a frutinha no processo, assim aumentando em 1 o seu tamanho! */

int rastejar(cobrinha *snake, int campo[DIMENSAO][DIMENSAO]){
	switch(snake->direcao){
		case CIMA:
			//Iria bater numa parede ou nela mesma
			if(snake->posicao_y -1 <0 ||campo[snake->posicao_y-1][snake->posicao_x]>0)
				return verificarVitoria(campo);
			//Vai comer uma frutinha
			else if(campo[snake->posicao_y - 1][snake->posicao_x]==-1){
				(snake->tamanho)++;
				colocarCobraNaPosicao(snake, campo, snake->posicao_y-1, snake->posicao_x);
				colocarFrutinha(campo);
			//Rastejar normal
			} else{
				decairPartesDaCobra(campo);
				colocarCobraNaPosicao(snake, campo, snake->posicao_y-1, snake->posicao_x);
			}
		break;
		case BAIXO:
			if(snake->posicao_y +1 ==DIMENSAO ||campo[snake->posicao_y+1][snake->posicao_x]>0)
				return verificarVitoria(campo);
			else if(campo[snake->posicao_y + 1][snake->posicao_x]==-1){
				(snake->tamanho)++;
				colocarCobraNaPosicao(snake, campo, snake->posicao_y+1, snake->posicao_x);
				colocarFrutinha(campo);
			} else{
				decairPartesDaCobra(campo);
				colocarCobraNaPosicao(snake, campo, snake->posicao_y+1, snake->posicao_x);
			}
		break;
		case DIREITA:
			if(snake->posicao_x +1 ==DIMENSAO ||campo[snake->posicao_y][snake->posicao_x+1]>0)
				return verificarVitoria(campo);
			else if(campo[snake->posicao_y][snake->posicao_x+1]==-1){
				(snake->tamanho)++;
				colocarCobraNaPosicao(snake, campo, snake->posicao_y, snake->posicao_x+1);
				colocarFrutinha(campo);
			} else{
				decairPartesDaCobra(campo);
				colocarCobraNaPosicao(snake, campo, snake->posicao_y, snake->posicao_x+1);
			}
		break;
		case ESQUERDA:
			if(snake->posicao_x -1 <0 ||campo[snake->posicao_y][snake->posicao_x-1]>0)
				return verificarVitoria(campo);
			else if(campo[snake->posicao_y][snake->posicao_x-1]==-1){
				(snake->tamanho)++;
				colocarCobraNaPosicao(snake, campo, snake->posicao_y, snake->posicao_x-1);
				colocarFrutinha(campo);
			} else{
				decairPartesDaCobra(campo);
				colocarCobraNaPosicao(snake, campo, snake->posicao_y, snake->posicao_x-1);
			}
		break;		
	}
	return 0;
}

/*Essa função vai verificar se há alguma posição na matriz que seja <=0  (Espaço vazio ou frutinha) e retornar -1 caso haja ou retornar 1 caso o usuario tenha vencido
(cobrinha ocupou todo o campo e não há mais espaços vazios ou frutinhas)*/

int verificarVitoria(int campo[DIMENSAO][DIMENSAO]){
	
	for(int i=0;i<DIMENSAO;i++)
		for(int j=0;j<DIMENSAO;j++)
			 if(campo[i][j]<=0)
			 	return -1;
			 
	return 1;	
		
}

//Coloca a cabeça da cobra em uma dada posição

void colocarCobraNaPosicao(cobrinha *snake, int campo[DIMENSAO][DIMENSAO], int linha, int coluna){
	snake->posicao_y = linha;
	snake->posicao_x = coluna;
	campo[snake->posicao_y][snake->posicao_x] = snake->tamanho;
}

//Decrementa em 1 o tempo de vida das partes da cobra (inteiros>0 da matriz)

void decairPartesDaCobra(int campo[DIMENSAO][DIMENSAO]){
	for(int i=0;i<DIMENSAO;i++)
		for(int j=0;j<DIMENSAO;j++)
			if(campo[i][j]>0)
				campo[i][j]--;
}

//Desenvolvido por Igor M. :P
