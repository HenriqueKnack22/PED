Para rodar o código:

1.Abra um terminal e rode o servidor:
go run server.go    

2.Abra outro terminal e rode a agency
go run agency.go create localhost 5000 nomedesuaescolha 300.0
go run agency.go consult localhost 5000 nomedesuaescolha
go run agency.go delete localhost 5000 nomedesuaescolha

3.Abra outro terminal com o teller
go run teller.go consult localhost 5000 nomedesuaescolha
go run teller.go deposit localhost 5000 nomedesuaescolha 100.0
go run teller.go withdraw localhost 5000 nomedesuaescolha 100.0


Vale ressaltar que podemos criar, remover, consultar nesse codigo. No entanto alguns agentes não consegue realizar certas ações!