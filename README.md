## FamBirthdayBot

## 📔 Descrição

Este projeto se consiste em um bot no Telegram que será responsável por enviar uma mensagem para alguns usuários específicos sempre às 5 da manhã contendo um aviso do possível aniversariante do dia. Sempre às 5 da manhã o bot fará uma validação de uma lista de aniversariantes, se a data atual da validação bater com a data de alguma pessoa listada na lista de aniversariantes, então este nome será registrado em outra lista de aniversariantes. Posteriormente será enviada com uma mensagem padrão avisando que hoje é aniversário de X, Y e Z.

## 🎯 Objetivo

O objetivo deste projeto é ajudar, especialmente, meu pai a lembrar das datas de aniversário de pessoas da família ou amigos próximos dele. Inicialmente o projeto será apenas dedicado e cadastrado e usado pelo meu pai, mas futuramente poderei escalar o projeto para o mesmo poder ser usado por outros integrantes da família, tornando a aplicação em uma espécie de intra-net da família.

## ⚠️ Requisitos

### Funcionais

- [x] Deve ser possível fazer envio de mensagens para N contatos mediante um bot no Telegram;
- [x] O sistema deve funcionar independente do ano na hora da validação;
- [x] Deve ser possível lidar com casos onde há mais de um aniversariante no dia da validação;
- [x] Deve ser possível validar se alguém está fazendo aniversário no dia da validação;
- [x] O sistema deve listar o nome e complemento dos aniversariantes;

### Não funcionais

- [ ] Deve haver um QR code para acessar o Bot no Telegram;
- [x] O sistema deve constar com um subsistema de logs feitos a nível de linha de comando;
- [ ] Deve existir um Google Forms que seja capaz de coletar dados para serem usados no sistema;

### Regras de negócio

- [x] As mensagens só devem ser enviadas caso haja algum aniversariante no dia da validação;
- [x] As mensagens não podem ser enviadas mais de uma vez no mesmo dia;
- [x] As mensagens só devem ser enviadas às 5 da manhã;
- [x] As mensagens só devem ser enviadas para IDs cadastrados no sistema;

## ⚒️ Infraestrutura

### Fluxograma

![]()

### 🖿 Estrutura de pastas

```markdown
  - cmd/
  	- main.go
  - internal/
	  - config/
	  - types/
	  - utils/
  go.mod
  go.sum
```

### 🖥️ Tecnologias

|Biblioteca|Versão|Utilidade na aplicação|
|---|---|---|
|telegram-bot-api|5.5.1|Comunicação com a API do Telegram usando a linguagem Go|
|viper|1.19|Carregamento de variáveis de ambiente mediante um arquivo .env|
|go|1.22.6|Linguagem de programação usada no projeto|

## 🌐 Implementações futuras

- Deploy;
- Cadastro de novas datas de aniversário através do bot;
- Envio de mensagem de parabéns para o aniversariante;
- Envio de mensagens seletas para usuários específicos;
	- Eu só quero ser notificado das datas de fulano, sicrano e beltrano.
- Tratamento de erros na hora de enviar mensagens, adiando o envio até que seja enviada;
- Separar em dois micro serviços dependentes;
- Separação da aplicação em 2 micro serviços, um de mensageiria e outro para tratar das datas de aniversário;

## ✏️ O que eu aprendi com este projeto

- Envio de mensagens mediante um bot no Telegram com a linguagem Go;
- Criação de um bot de Telegram usando Go;
- Agendamento de tarefas feitas em Go;
- Deploy de aplicações Go;
- Tratamento de erros em Go com SLog;
- Calcular quanto tempo falta para determinada data em Go;
- Criação de uma data customizada em Go;
- Obtenção do diretório raiz do projeto em Go;
- Formatação de horas e minutos oriundos da diferença de tempo entre duas datas em Go;
- Testes unitários em Go;
- Trabalho com variáveis de ambiente usando a biblioteca Viper em Go;

## Meus contatos

- Instagram: https://www.instagram.com/antonioalmeida2003/
- LinkedIn: https://www.linkedin.com/in/antonio-mauricio-4645832b3/
- Email: antonioimportant@gmail.com