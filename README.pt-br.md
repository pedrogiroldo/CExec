# CExec

Uma ferramenta de linha de comando para compilar e executar arquivos C/C++ com um único comando.

## Sobre

CExec é uma ferramenta simples desenvolvida em Go que automatiza o processo de compilação e execução de programas C/C++. Ela elimina a necessidade de digitar comandos separados para compilar e depois executar seu código.

## Pré-requisitos

- Go (para construir a partir do código-fonte)
- Um compilador C/C++ (como g++)

## Instalação

### A partir do código-fonte

1. Clone este repositório
2. Compile o programa:

```bash
go build -o build/CExec src/main.go
```

3. Adicione o executável compilado ao seu PATH para uso global (opcional)

## Uso

### Uso básico

```bash
CExec arquivo.(c/cpp)
```

O programa irá:

1. Compilar o arquivo C/C++ especificado usando o compilador configurado
2. Executar o programa resultante (se configurado)
3. Exibir a saída do programa

### Arquivo de configuração

O CExec pode ser configurado através de um arquivo JSON chamado `CExecConfig.json`. Este arquivo deve estar no mesmo diretório de onde o CExec é executado.

Exemplo de `CExecConfig.json`:

```json
{
  "compilerPath": "/usr/bin/g++",
  "compilerArgs": ["-Wall", "-std=c++17"],
  "outputName": "meu_programa",
  "runAfterCompile": true,
  "customRunCommand": "arg1 arg2",
  "sourceFile": "main.cpp"
}
```

#### Opções de configuração:

| Opção              | Descrição                                             | Obrigatório                                       |
| ------------------ | ----------------------------------------------------- | ------------------------------------------------- |
| `compilerPath`     | Caminho para o compilador (ex: g++)                   | Sim                                               |
| `compilerArgs`     | Lista de argumentos para o compilador                 | Não                                               |
| `outputName`       | Nome do arquivo executável gerado                     | Não (padrão: "output" ou "output.exe" no Windows) |
| `runAfterCompile`  | Se o programa deve ser executado após a compilação    | Não (padrão: false)                               |
| `customRunCommand` | Argumentos para passar ao programa durante a execução | Não                                               |
| `sourceFile`       | Arquivo fonte padrão a ser compilado                  | Não (pode ser sobrescrito via linha de comando)   |

**Nota:** Quando um arquivo é especificado via linha de comando, ele tem precedência sobre o arquivo definido na configuração.

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
