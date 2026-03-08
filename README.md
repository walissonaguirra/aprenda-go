# Aprenda Go

Anotações e exemplos práticos do curso **"Aprenda Go"** do canal [Aprenda Go](https://www.youtube.com/@AprendaGo) no YouTube, ministrado por [Ellen Körbes](https://gotopia.tech/experts/476/ellen-koerbes).

---

## Grade do Curso

| # | Tópico | Diretório |
|---|--------|-----------|
| 02 | Variáveis, Valores & Tipos | [02_variaveis-valores-e-tipos](02_variaveis-valores-e-tipos/) |
| 03 | Exercícios Ninja Nível 1 | [03_exercicios-ninja-nivel-1](03_exercicios-ninja-nivel-1/) |
| 04 | Fundamentos de Programação | [04_fundamentos-de-programacao](04_fundamentos-de-programacao/) |
| 06 | Fluxo de Controle | [06_fluxo-de-controle](06_fluxo-de-controle/) |
| 08 | Agrupamentos de Dados | [08_agrupamentos-de-dados](08_agrupamentos-de-dados/) |
| 09 | Exercícios Ninja Nível 4 | [09_exercícios-ninja-nivel-4](09_exercícios-ninja-nivel-4/) |
| 10 | Structs | [10_structs](10_structs/) |
| 11 | Exercícios Ninja Nível 5 | [11_exercicios-ninja-nivel-5](11_exercicios-ninja-nivel-5/) |
| 12 | Funções | [12_funcoes](12_funcoes/) |
| 13 | Exercícios Ninja Nível 6 | [13_exercicios-ninja-nivel-6](13_exercicios-ninja-nivel-6/) |
| 14 | Ponteiros | [14_ponteiros](14_ponteiros/) |
| 15 | Exercícios Ninja Nível 7 | [15_exercicios-ninja-nivel-7](15_exercicios-ninja-nivel-7/) |
| 16 | Aplicações | [16_aplicacoes](16_aplicacoes/) |
| 18 | Concorrência | [18_concorrencia](18_concorrencia/) |
| 21 | Canais | [21_canais](21_canais/) |
| 23 | Tratamento de Erros | [23_tratamento-de-erros](23_tratamento-de-erros/) |

> **Nota:** Os capítulos 1, 5, 7, 17, 19, 20 e 22 não possuem código no repositório por serem aulas teóricas ou de configuração de ambiente.

---

## Estrutura dos Arquivos

```md
aprenda-go/
├── 02_variaveis-valores-e-tipos/
│   ├── 02_hello-world/
│   ├── 03_operador-curto-de-declaracao/
│   ├── 04_a-palavra-chave-var/
│   ├── 05_explorando-tipos/
│   ├── 06_valor_zero/
│   ├── 07_o-pacote-fmt/
│   ├── 08_criando-seu-primeiro-tipo/
│   └── 09_conversao-nao-coercao/
├── 03_exercicios-ninja-nivel-1/
│   ├── 01/ a 05/                        # 5 exercícios
├── 04_fundamentos-de-programacao/
│   ├── 01_tipo-booleano/
│   ├── 02_como-os-computadores-funcionam/
│   ├── 03_tipos-numericos/
│   ├── 04_overflow/
│   ├── 05_tipo-string/
│   ├── 06_sistemas-numericos/
│   ├── 07_constantes/
│   ├── 08_iota/
│   └── 09_deslocamento_de_bits/
├── 06_fluxo-de-controle/
│   └── 01_entendendo-fluxo-de-controle/
├── 08_agrupamentos-de-dados/
│   ├── 01_array/
│   ├── 02_slice-literal-composta/
│   ├── 03_slice-for-range/
│   ├── 04_slice-fatiando-ou-deletando-de-uma-fatia/
│   ├── 05_slice-anexando-a-uma-slice/
│   ├── 06_slice-make/
│   ├── 07_slice-slice-multi-dimensional/
│   ├── 08_slice-a-surpresa-do-array-subjacente/
│   ├── 09_maps-introducao/
│   └── 10_maps-range-&-deletando/
├── 09_exercícios-ninja-nivel-4/
│   ├── 01/ a 10/                        # 10 exercícios
├── 10_structs/
│   ├── 01_struct/
│   ├── 02_structs-embutidos/
│   ├── 03_lendo-a-documentacao/
│   └── 04_structs-anonimos/
├── 11_exercicios-ninja-nivel-5/
│   ├── 01/ a 04/                        # 4 exercícios
├── 12_funcoes/
│   ├── 01_sintaxe/
│   ├── 02_desenrolando-(enumerando)-uma-slice/
│   ├── 03_defer/
│   ├── 04_metodos/
│   ├── 05_interfaces-&-polomorfismo/
│   ├── 06_funcoes-anonimas/
│   ├── 07_func-como-expressao/
│   ├── 08_retornando-uma-funcao/
│   ├── 09_callback/
│   ├── 10_closure/
│   └── 11_recursividade/
├── 13_exercicios-ninja-nivel-6/
│   ├── 01/ a 10/                        # 10 exercícios
├── 14_ponteiros/
│   ├── 01_o-que-sao-ponteiros/
│   └── 02_quando-usar-ponteiros/
├── 15_exercicios-ninja-nivel-7/
│   ├── 01/ a 02/                        # 2 exercícios
├── 16_aplicacoes/
│   ├── 02_json-marshal-ordenacao/
│   ├── 03_json-unmarshal-desordenacao/
│   ├── 04_a-interface-writer/
│   ├── 05_o-pacote-sort/
│   ├── 06_customizando-o-sort/
│   └── 07_bcrypt/
├── 18_concorrencia/
│   ├── 01_goroutines-&-waitgroups/
│   ├── 02_condicao-de-corrida/
│   ├── 03_mutex/
│   └── 04_atomic/
├── 21_canais/
│   ├── 01_entendendo-canais/
│   ├── 02_canais-direcionais-&-ultilizando-canais/
│   └── 03_range-e-close/
├── 23_tratamento-de-erros/
│   ├── 02_print-e-logs/
│   ├── 04_recover/
│   └── 05_erros-com-informacoes-adicionais/
│       ├── 01/ a 05/                    # 5 sub-exemplos
└── README.md
```

---

## Referência

- **Curso:** [Aprenda Go (YouTube)](https://www.youtube.com/@AprendaGo)
- **Instrutora:** [Ellen Körbes](https://gotopia.tech/experts/476/ellen-koerbes)
