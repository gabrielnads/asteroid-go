# Asteroid Go ☄️

Um clone do clássico jogo de arcade Asteroids, desenvolvido inteiramente em Go (Golang) utilizando a biblioteca gráfica Ebitengine.

![Gameplay do Asteroid Go](https://raw.githubusercontent.com/gabrielnads/asteroid-go/main/assets/gameplay.gif)


---
## 🚀 Sobre o Projeto

Este projeto foi criado como um exercício prático para explorar o desenvolvimento de jogos em Go. O objetivo era recriar a mecânica fundamental do Asteroids, incluindo movimento da nave, física de asteroides, tiros e sistema de pontuação, em uma arquitetura de código limpa e organizada.

---
## 🛠️ Tecnologias Utilizadas

* **Linguagem:** Go (Golang)
* **Motor Gráfico:** [Ebitengine](https://ebitengine.org/) - Um motor de jogo 2D de código aberto e muito performático para Go.

---
## ✨ Funcionalidades

* Controle da nave com rotação e propulsão.
* Tiros para destruir asteroides.
* Asteroides que se dividem em pedaços menores ao serem atingidos.
* Sistema de pontuação.
* Tela de "Game Over".

---
## ⚙️ Como Executar

**Pré-requisitos:**
* Go (versão 1.18+) instalado em sua máquina.
* As dependências do Ebitengine para o seu sistema operacional (como `gcc` no Windows ou `xcode` no macOS). Consulte a [documentação do Ebitengine](https://ebitengine.org/en/documents/install.html) para mais detalhes.

**Passos:**

1.  **Clone o repositório:**
    ```bash
    git clone [https://github.com/gabrielnads/asteroid-go.git](https://github.com/gabrielnads/asteroid-go.git)
    ```

2.  **Navegue até a pasta do projeto:**
    ```bash
    cd asteroid-go
    ```

3.  **Instale as dependências:**
    O Go Modules cuidará disso automaticamente, mas você pode forçar a instalação com:
    ```bash
    go mod tidy
    ```

4.  **Execute o jogo:**
    ```bash
    go run .
    ```

Uma janela com o jogo deverá abrir. Use as setas do teclado para mover e a barra de espaço para atirar!
