# API REST carbi

Este repositório guarda uma api simples que serve como modelo para uma empresa - a carbi - de revenda de carros.

Os programas estão feitos em golang e dependem de um servidor mysql e o drive em github.com/go-sql-driver/mysql para funcionamento.

- A request GET consulta o Estoque e Histórico.
- PUT atualiza os atributos de um carro (Ano, Cor, Carro, Preço).
- PUSH adiciona um carro ao Estoque.
- DELETE deleta um carro do Estoque - podendo representar uma venda, por exemplo.

## Sumário do uso dos métodos

- "estoque" se refere ao banco de dados que tem como colunas: (ID, Carro, Ano, Cor, Preço)
- "hist" se refere ao banco de dados que guarda informação das transações em estoque.

### GET

- /{banco}/  
retorna os elementos do banco (estoque, hist)

- /{banco}/{col}={val}  
retorna as colunas que satifazem col=val

exemplos:
curl localhost:8080/estoque/Cor=Amarela/
retorna os carros no estoque que têm a cor amarela

### PUT
- /{ID}/{col1,col2...}/{lin1,lin2...}  
atualiza o carro que tem o ID passado. As colunas devem ser separadas por ",".

exemplo:  
curl -X PUT localhost:8080/5/Cor,Ano/Amarela,2001/  
altera o carro 5

### PUSH
- /{col1,col2...}/{lin1,lin2...}

exemplo:  
curl -X PUSH localhost:8080/5/Cor,Ano,Carro,Preço/Amarela,2001,Mustang-GT,130000.0/  

### DELETE
- /{ID}/  
Deleta o carro com ID passado.

As requisições DELETE, PUT e PUSH são adicionadas ao histórico sempre que feitas. Afinal, uma empresa quer ter histórico de seus pedidos.

Isso é tudo, pessoal. Ficou bem simples, mas foi com carinho.
