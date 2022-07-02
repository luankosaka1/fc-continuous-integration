# Continuous Integration

É o processo de integrar modificações do codebase de forma contínua e automatizada, evitando assim erros humanos de verificação , garantindo mais agilidade e segurança no processo de desenvolvimento de um software.

## Principais Pocessos

- Execução de testes
- Linter
- Verificações de qualidade de código
- Verificações de segurança
- Geração de artefatos prontos para o processo de deploy
- Identificação da próxima versão a ser gerada no software
- Geração de tags e releases

## Ferramentas Populares

- Jenkins
- Github Actions
- Circle CI
- AWS Code Build
- Azure DevOps
- Google Cloud Build
- GitLab Pipelines/CI

## Dinamica

- Workflow
-- Evento (on: push)
-- Filtros (branches: master)
-- Ambiente (runs-on: ubuntu)
-- Ações (uses: action/run-composer; run: npm run prod)

## Criando Workflow

1. Criar arquivo .github/workflows/ci.yaml
2. Adicionar o código no arquivo ci.yaml
```
name: ci-golang-workflow
on: [pull_request]
jobs:
    check-application:
        runs-on: ubuntu-latest
        steps:
            - uses: actions/checkout@v2
            - uses: actions/setup-go@v2
                with:
                    go-version: 1.18
            - run: go test
            - run: go run math.go
```
