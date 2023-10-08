Id de usuário sempre contendo 16 bytes. A mensagem pode ter tamanho indefinido. Perceba que a estrutura termina com um byte nulo,
isso pode gerar alguns problemas, mas concertamos no futuro.

```
FFFFFFFFFFFFFFFFAAAAAAAAAAAAAAAAAAAAAAAAAAAAA0
|   User ID    |      Message               |
```

Suponhamos o cliente1 e cliente2.

Quando cliente1 deseja enviar uma mensage ao cliente2, é necessário enviar na seguinte estrutura. É esse dado que o server deve receber
para enviar a mensagem ao cliente2.

```
FFFFFFFFFFFFFFFFAAAAAAAAAAAAAAAAAAAAAAAAAAAAA0
| ID cliente2  |      Message               |
```

O server irá repassar a mensagem na seguinte estrutura, tal qual será recebida pelo cliente2.

```
FFFFFFFFFFFFFFFFAAAAAAAAAAAAAAAAAAAAAAAAAAAAA0
| ID cliente1  |      Message               |
```



