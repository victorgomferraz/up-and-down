# UP and DOWN

Gerar um artefato que sobe um servidor dumb e morre após o delay configurado.

## Motivação
Existe o caso quando trabalhamos com containers que um dos containers vai subir, realizar um trabalho e morrer,
entretando nos precisamos que outro container suba após a morte do primeiro.

Para isso ser possível precisamos de um artefato que suba um servidor, após realizar seu trabalho, ocupando uma porta que vai ser observada pelo wait-for-it.sh para inicio
do container que esta aguardando o processo do primeiro container finalizar.

Esse container precisa ficar com o servidor online tempo suficiente para ser detectado pelo wait-for-it.sh, e após isso pode morrer.

## Como usar
Para gerar o artefato, basta executar o comando abaixo:
```bash
$ docker compose up
```
O vai ser gerado um artefato chamado "up-and-down" na arquitetura linux.
Esse artefato utiliza 2 variáveis de ambiente:
- UP_AND_DOWN_PORT: Porta que o servidor vai subir e ficar online (Obrigatório)
- UP_AND_DOWN_DELAY: Tempo em segundos que o servidor vai ficar online. (Opcional, padrão 3 segundos)