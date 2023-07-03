# Weather CLI

Uma ferramenta de linha de comando (CLI) para verificar a previsão do tempo em qualquer cidade do mundo.

## Instalação

Para executar o Weather CLI na sua CLI a partir de qualquer diretório, siga estes passos:

1. Clone o repositório:

```git clone https://github.com/seu-nome-de-usuario/weather-cli.git```

2. Navegue até o diretório do projeto:
   
```cd weather-cli```

3. Compile o arquivo executável:

```go build -o weather```

4. Mova o arquivo executável para a pasta `/usr/local/bin` utilizando o comando `sudo mv`:
   
```sudo mv weather /usr/local/bin```

Esse passo requer privilégios de administrador e pode solicitar que você insira sua senha.

Ao mover o arquivo executável para o diretório `/usr/local/bin`, ele se torna acessível de qualquer lugar na sua CLI.

## Uso

Após completar as etapas de instalação, você pode utilizar o Weather CLI digitando `weather` na sua CLI a partir de qualquer diretório. Essa ferramenta permite verificar a previsão do tempo para qualquer cidade do mundo.

Para obter a previsão do tempo de uma cidade específica, utilize o seguinte comando:

```weather <cidade>```

Substitua `<cidade>` pelo nome da cidade para a qual você deseja verificar a previsão do tempo.

## Exemplo

Aqui está um exemplo de como utilizar o Weather CLI:

```weather São Paulo```

Esse comando irá buscar e exibir a previsão do tempo para a cidade de São Paulo.

## Contribuições

Contribuições são bem-vindas! Se você deseja contribuir com o projeto Weather CLI, siga os passos abaixo:

1. Faça um fork do repositório.
2. Crie um novo branch para a sua funcionalidade ou correção de bug.
3. Faça as alterações necessárias e faça o commit delas.
4. Faça o push das suas alterações para o seu repositório fork.
5. Envie um pull request descrevendo as alterações que você fez.

## Licença

O projeto Weather CLI está licenciado sob a [Licença MIT](LICENSE). Sinta-se à vontade para utilizar, modificar e distribuir essa ferramenta CLI conforme os termos da licença.

## Agradecimentos

Este projeto utiliza os dados de previsão do tempo fornecidos pelo [Weather API](https://www.weatherapi.com/). Agradecemos pelos seus serviços.

Se você tiver alguma dúvida ou precisar de ajuda adicional, não hesite em entrar em contato. Aproveite a verificação do tempo com o Weather CLI!
