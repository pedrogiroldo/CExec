# CExec

Uma ferramenta de linha de comando para compilar e executar arquivos C/C++ com um único comando.

## Sobre

CExec é uma ferramenta simples desenvolvida em Go que automatiza o processo de compilação e execução de programas C/C++. Ela elimina a necessidade de digitar comandos separados para compilar e depois executar seu código.

## Pré-requisitos

- Go (para construir a partir do código-fonte)

## Instalação

### A partir do código-fonte

1. Clone este repositório
2. Compile o programa:

```bash
go build -o build/CExec src/main.go
```

3. Adicione o executável compilado ao seu PATH para uso global (opcional)

## Uso

```bash
CExec arquivo.(c/cpp)
```

O programa irá:

1. Compilar o arquivo C++ especificado usando g++
2. Executar o programa resultante
3. Exibir a saída do programa

## Exemplo

```bash
$ CExec meu_programa.cpp
```

Se o programa `meu_programa.cpp` contiver:

```cpp
#include <iostream>

int main() {
    std::cout << "Olá, mundo!" << std::endl;
    return 0;
}
```

A saída será:

```
Olá, mundo!
```

## Notas de compatibilidade

O executável detecta automaticamente o sistema operacional e ajusta o nome do arquivo de saída:

- Em sistemas Linux/Unix: `output`
- Em sistemas Windows: `output.exe`

## Licença

[Adicionar informações de licença]

## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.
