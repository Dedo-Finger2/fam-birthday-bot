## FamBirthdayBot

## Descrição

Este projeto se consiste em um bot no Telegram que será responsável por enviar uma mensagem para alguns usuários específicos sempre às 5 da manhã contendo um aviso do possível aniversariante do dia. Sempre às 5 da manhã o bot fará uma validação de uma lista de aniversariantes, se a data atual da validação bater com a data de alguma pessoa listada na lista de aniversariantes, então este nome será registrado em outra lista de aniversariantes. Posteriormente será enviada com uma mensagem padrão avisando que hoje é aniversário de X, Y e Z.

## Objetivo

O objetivo deste projeto é ajudar, especialmente, meu pai a lembrar das datas de aniversário de pessoas da família ou amigos próximos dele. Inicialmente o projeto será apenas dedicado e cadastrado e usado pelo meu pai, mas futuramente poderei escalar o projeto para o mesmo poder ser usado por outros integrantes da família, tornando a aplicação em uma espécie de intra-net da família.

## Requisitos

### Funcionais
- [ ] Deve ser possível fazer envio de mensagens para N contatos mediante um bot no Telegram;
- [ ] O sistema deve funcionar independente do ano na hora da validação;
- [ ] Deve ser possível lidar com casos onde há mais de um aniversariante no dia da validação;
- [ ] Deve ser possível validar se alguém está fazendo aniversário no dia da validação;

### Não funcionais
- [ ] Deve haver um QR code para acessar o Bot no Telegram;

### Regras de negócio
- [ ] As mensagens só devem ser enviadas caso haja algum aniversariante no dia da validação;
- [ ] As mensagens só devem ser enviadas às 5 da manhã;
- [ ] As mensagens só devem ser enviadas para IDs cadastrados no sistema;
