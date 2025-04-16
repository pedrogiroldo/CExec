# CExec

Uma ferramenta de linha de comando para compilar e executar arquivos C/C++ com um único comando.

## Sobre

CExec é uma ferramenta leve desenvolvida em GO, projetada para simplificar a compilação e execução de programas C/C++. Com apenas um comando no terminal, ela automatiza todo o processo, tornando mais fácil para os desenvolvedores testarem e executarem seu código. Projetada para eficiência, o CExec simplifica os fluxos de trabalho combinando ambas as etapas em uma única ação intuitiva.

## Abordagem Tradicional vs CExec

### Abordagem Tradicional

O método tradicional requer múltiplos comandos para compilar e depois executar programas C/C++:

![Compilação e Execução Tradicional de C/C++](docs/assets/ex1.gif)

### Abordagem com CExec

Com o CExec, todo o processo é simplificado em um único comando:

![Compilação e Execução com CExec](docs/assets/ex2.gif)

## ✨ Modo Watch: Desenvolvimento Contínuo em Tempo Real! ✨

**Aumente sua produtividade com o poderoso modo Watch!** 🚀

O modo Watch do CExec transforma completamente sua experiência de desenvolvimento em C/C++. Com ele ativado, você pode:

- **Esquecer os comandos repetitivos** - enquanto você escreve código, o CExec observa suas alterações e recompila automaticamente.
- **Ver resultados instantaneamente** - cada vez que você salva um arquivo, seu código é recompilado e executado imediatamente.
- **Focar no que realmente importa: seu código** - nada de ficar alternando entre editor e terminal!

É como ter um assistente de programação que executa seu código sempre que você faz uma mudança. Perfeito para:

- Desenvolvimento iterativo e incremental
- Aprendizado de C/C++
- Debugging rápido
- Testes imediatos de pequenas alterações

### Como usar:

```bash
CExec -watch=true meu_programa.cpp
```

Pronto! Agora você pode editar seu arquivo com tranquilidade - a cada salvamento, o CExec irá recompilar e executar seu programa automaticamente, exibindo os resultados no terminal.

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

### Opções de linha de comando

Você também pode usar flags na linha de comando para personalizar o comportamento do CExec:

```bash
CExec -compiler=/caminho/para/g++ -args="-Wall,-std=c++17" -output=meu_programa -run=true -source=main.cpp -watch=true
```

## 🔧 Configurando o CExec

O CExec oferece **duas maneiras flexíveis de configuração**:

### 1️⃣ Arquivo de Configuração JSON

Você pode criar um arquivo `CExecConfig.json` no diretório de trabalho para configurações permanentes:

```json
{
  "compilerPath": "/usr/bin/g++",
  "compilerArgs": ["-Wall", "-std=c++17"],
  "outputName": "meu_programa",
  "runAfterCompile": true,
  "customRunCommand": "arg1 arg2",
  "sourceFile": "main.cpp",
  "watchChanges": false
}
```

### 2️⃣ Flags de Linha de Comando

Para uso rápido ou para substituir configurações do arquivo JSON, use flags de linha de comando:

```bash
CExec -compiler=/caminho/para/g++ -args="-Wall,-std=c++17" -output=meu_programa -run=true -source=main.cpp -watch=true
```

### Prioridade das Configurações

O CExec segue uma ordem clara para determinar qual configuração usar:

1. **Flags de linha de comando** (maior prioridade) - substituem qualquer outra configuração
2. **Arquivo CExecConfig.json** (prioridade média) - usado quando as flags não estão presentes
3. **Valores padrão** (menor prioridade) - usados quando nenhuma outra configuração é fornecida

Isso permite ter um arquivo de configuração para seus valores padrão, enquanto mantém a flexibilidade para substituições rápidas via linha de comando quando necessário.

### Opções de configuração:

| Opção                    | Descrição                                                | Flag relacionada | Config JSON        | Obrigatório                                         |
| ------------------------ | -------------------------------------------------------- | ---------------- | ------------------ | --------------------------------------------------- |
| Caminho do compilador    | Caminho para o compilador (ex: g++)                      | `-compiler`      | `compilerPath`     | Sim                                                 |
| Argumentos do compilador | Lista de argumentos para o compilador                    | `-args`          | `compilerArgs`     | Não                                                 |
| Nome do executável       | Nome do arquivo executável gerado                        | `-output`        | `outputName`       | Não (padrão: "output" ou "output.exe" no Windows)   |
| Executar após compilar   | Se o programa deve ser executado após a compilação       | `-run`           | `runAfterCompile`  | Não (padrão: false)                                 |
| Argumentos de execução   | Argumentos para passar ao programa durante a execução    | `-run-cmd`       | `customRunCommand` | Não                                                 |
| Arquivo fonte            | Arquivo fonte a ser compilado                            | `-source`        | `sourceFile`       | Não (pode ser especificado como primeiro argumento) |
| Modo watch               | Ativar modo de monitoramento para recompilar em mudanças | `-watch`         | `watchChanges`     | Não (padrão: false)                                 |

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

### Exemplo do modo de monitoramento

```bash
$ CExec -watch=true meu_programa.cpp
```

Isso iniciará o CExec no modo de monitoramento. Ele compilará o arquivo e continuará monitorando-o para alterações. Sempre que o arquivo for modificado e salvo, o CExec o recompilará automaticamente.

## Notas de compatibilidade

O executável detecta automaticamente o sistema operacional e ajusta o nome do arquivo de saída:

- Em sistemas Linux/Unix: `output`
- Em sistemas Windows: `output.exe`

## Licença

[Adicionar informações de licença]

## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.
