# API CARBI
### Sumário do uso dos métodos

- estoque se refere ao banco de dados que tem como colunas: (ID, Carro, Ano, Cor, Preço)
- hist se refere ao banco de dados que guarda informação das transações em estoque

GET

/{banco}/ -> retorna os elementos do banco (estoque, hist)

/{banco}/{col}={val} -> retorna as colunas que satifazem col=val

exemplo:
curl localhost:8080/estoque/Cor=Amarela/
retorna os carros no estoque que têm a cor amarela

PUT
/{ID}/{col1,col2...}/{lin1,lin2...}

atualiza o carro que tem o ID passado. As colunas devem ser separadas por ",".

exemplo:
curl -X PUT localhost:8080/5/Cor,Ano/Amarela,2001/
altera o carro 5

PUSH
/{col1,col2...}/{lin1,lin2...}

exemplo:
curl -X PUSH localhost:8080/5/Cor,Ano,Carro,Preço/Amarela,2001,Mustang-GT,130000.0/

DELETE
/{ID}/
