<h3>Esse projeto é uma resolução ao desafio proposto no site: https://www.devgym.com.br/ . </h3>
</br>
Desafio: </br> 
Um programa para popular um banco de dados com milhares de filmes

Nesse desafio você deve criar um programa de linha de comando (cli) que lê um arquivo csv de filmes e popula um banco de dados pensando em performance e esperando que o arquivo pode crescer muito. 


  Cada linha do csv contém colunas que devem ser salvas em colunas/campos separados no banco de dados:
    
     ID - número inteiro que identifica o filme encontrado na 1a coluna do csv
    
     title - título do filme encontrado na 2a coluna do csv. Com o valor "Jumanji (1995)" o title é Jumanji
    
     Year - ano do filme encontrado na 2a coluna do csv. Com o valor "Jumanji (1995)" o year é 1995
    
     Genres - múltiplos valores com os gêneros do filme separados por |. Encontrado na 3a coluna.
  O script deve pensar em performance e tirar proveito de concorrência/paralelismo para popular o banco de dados. 

link para o desafio:
https://app.devgym.com.br/challenges/ec36e7e2-6a2d-4406-98e1-3029f843b5c3


Tabela usada para salvar os dados do CSV:

![Alt text](https://i.ibb.co/mJvZ1Jk/Untitled-1.png "schema table")
