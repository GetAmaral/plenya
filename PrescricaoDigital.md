# Requisitos para EMR com prescrição digital controlada no Brasil

Desenvolver um sistema de prontuário eletrônico (EMR) com prescrição digital de medicamentos controlados no Brasil exige conformidade com um arcabouço regulatório complexo envolvendo CFM, ANVISA e ICP-Brasil. **O certificado digital A1 é plenamente aceito para prescrições controladas de receita branca**, desde que seja ICP-Brasil. A principal mudança recente é a criação do **SNCR (Sistema Nacional de Controle de Receituários)**, cuja integração via API será obrigatória para prescrições eletrônicas controladas a partir de junho de 2026.

---

## Legislação federal estabelece base para prescrição digital

O marco legal brasileiro para prescrição digital está consolidado em três pilares: a **MP 2.200-2/2001** (institui ICP-Brasil e equivalência jurídica da assinatura digital), a **Lei 14.063/2020** (define tipos de assinatura eletrônica) e a **Lei 13.732/2018** (validade nacional de receitas controladas).

A Lei 14.063/2020 é especialmente relevante ao estabelecer três níveis de assinatura:
- **Simples**: identificação básica, insuficiente para receitas controladas
- **Avançada**: certificado não-ICP com autenticação forte, aceita apenas para receitas sujeitas à retenção
- **Qualificada**: certificado ICP-Brasil, **obrigatória para receitas de controle especial (Lista C1)**

Para que uma receita digital tenha validade jurídica equivalente à física, deve: (1) ser assinada digitalmente com certificado ICP-Brasil; (2) permitir verificação da assinatura via serviços do ITI; (3) conter todos os elementos obrigatórios da prescrição; (4) ser gerada digitalmente ("nata digital") — documentos escaneados não são válidos para controlados.

---

## Resoluções CFM regulamentam documentos médicos eletrônicos

O **Conselho Federal de Medicina** estabeleceu um conjunto robusto de normas através de resoluções específicas:

| Resolução | Data | Escopo |
|-----------|------|--------|
| **CFM 1.821/2007** | Nov/2007 | Base do prontuário eletrônico (PEP) |
| **CFM 2.299/2021** | Out/2021 | Documentos médicos eletrônicos |
| **CFM 2.314/2022** | Mai/2022 | Telemedicina |
| **CFM 2.381/2024** | Jul/2024 | Normalização de documentos médicos |

A **Resolução CFM 2.299/2021** é a principal norma sobre prescrição digital. Seu Artigo 4º determina que documentos médicos eletrônicos "deverão ser feitos mediante o uso de assinatura digital, gerada por meio de certificados e chaves emitidos pela ICP-Brasil, com **Nível de Garantia de Segurança 2 (NGS2)**".

Quanto aos requisitos obrigatórios das plataformas, o **Artigo 5º** exige registro junto ao CRM da jurisdição, com Diretor Técnico médico inscrito. O **Artigo 8º** determina verificação automatizada junto ao Cadastro Nacional de Médicos. O **Artigo 11** proíbe direcionamento de receitas para farmácias específicas.

A **Resolução CFM 1.821/2007** estabelece os níveis de segurança NGS1 e NGS2, sendo este último obrigatório para eliminação do papel. O NGS2 exige certificados ICP-Brasil para assinatura e autenticação, trilha de auditoria completa e conformidade com padrões PBAD (PAdES, CAdES ou XAdES).

---

## ANVISA regulamenta receitas controladas digitais

A regulamentação da ANVISA para prescrições digitais de medicamentos controlados passou por transformação significativa entre 2024-2025:

**Portaria SVS/MS 344/1998** continua como base normativa para substâncias controladas. A **Lista C1** ("Outras Substâncias Sujeitas a Controle Especial") inclui cerca de **165-208 substâncias**: antidepressivos, antiparkinsonianos, anticonvulsivantes, neurolépticos, anestésicos e ansiolíticos não-benzodiazepínicos. Exige Receita de Controle Especial (receita branca) em 2 vias.

A **RDC 471/2021** regulamenta antimicrobianos, que **não são Lista C1** mas também exigem retenção. Prescrições de antibióticos têm validade de **10 dias**. A RDC 973/2025 expandiu esta norma para incluir **agonistas GLP-1** (semaglutida, tirzepatida), com validade de 90 dias.

A **RDC 873/2024** criou o **SNCR** (Sistema Nacional de Controle de Receituários), sistema centralizado de numeração nacional que substitui os talonários estaduais. A partir de sua implementação completa, todas as prescrições controladas terão validade nacional e numeração única.

A **RDC 1.000/2025** (dezembro/2025, vigência 60 dias após publicação) é a norma-chave para prescrições eletrônicas controladas:

- Prescrições eletrônicas **devem ser geradas exclusivamente através de serviços integrados ao SNCR via API**
- Receitas de Controle Especial exigem **assinatura eletrônica qualificada** (ICP-Brasil)
- Receitas sujeitas à retenção (antimicrobianos, GLP-1) aceitam assinatura avançada ou qualificada
- Elimina necessidade de 2 vias para receitas eletrônicas
- Cada prescrição eletrônica só pode ser utilizada **uma única vez** (registrada no SNCR)
- **Prazo de implantação plena: 1º de junho de 2026**

---

## Certificado A1 é válido para prescrições controladas

Uma das questões mais frequentes está esclarecida: **tanto certificados A1 quanto A3 são legalmente válidos** para prescrições de medicamentos controlados, desde que sejam certificados ICP-Brasil.

| Característica | A1 | A3 | Cloud (A5) |
|---------------|-----|-----|------------|
| **Armazenamento** | Arquivo software (.PFX) | Token/smartcard físico | HSM em nuvem |
| **Validade** | 1 ano | 1-5 anos | Até 5 anos |
| **Geração de chaves** | Software no dispositivo | Hardware criptográfico | HSM remoto |
| **Segurança** | Menor | Maior | Equivalente ao A3 |
| **Praticidade** | Alta | Menor | Alta |
| **Custo** | Menor | Maior | Variável |

A plataforma oficial do CFM (prescricaoeletronica.cfm.org.br) aceita todos os tipos: A1, A3 e certificados em nuvem (BIRDID, VIDAAS, VAULTID, SAFEID, REMOTEID, NEOID). O CFM oferece **certificado digital gratuito em nuvem** para médicos regularmente inscritos.

**Requisitos técnicos ICP-Brasil:**
- Algoritmo RSA com chaves mínimas de **2048 bits**
- Hash da família **SHA-2** (mínimo SHA-256)
- Formatos de assinatura: **PAdES** (para PDF), **CAdES** ou **XAdES**
- Carimbo do tempo (timestamp): **não obrigatório**, mas recomendado para maior proteção jurídica

---

## Receita branca controlada: o que deve conter

A **Receita de Controle Especial** para medicamentos da Lista C1 deve incluir obrigatoriamente:

**Identificação do prescritor:**
- Nome completo
- CRM com estado (UF)
- RQE se especialista
- Endereço profissional
- Telefone/e-mail de contato

**Identificação do paciente:**
- Nome completo
- CPF (ou passaporte para estrangeiros)
- Endereço residencial
- Para antimicrobianos: idade, sexo, telefone

**Informações da prescrição:**
- Nome do medicamento (DCB - Denominação Comum Brasileira)
- Dosagem/concentração
- Forma farmacêutica
- Quantidade em algarismos e por extenso
- Posologia completa
- Data de emissão

**Requisitos formais:**
- **Máximo 3 substâncias C1** por receita
- Quantidade para até **60 dias** de tratamento (180 dias para antiparkinsonianos/anticonvulsivantes)
- Validade de **30 dias** a partir da emissão
- **Máximo 5 ampolas** para injetáveis

**Para receita digital, adicionar:**
- Código QR ou código de validação
- Link para verificação
- Identificação da plataforma emissora
- Assinatura digital ICP-Brasil verificável

O formato de arquivo deve ser **PDF assinado digitalmente** com padrão PAdES, nascido digital (não escaneado).

---

## SNCR é obrigatório, SNGPC é da farmácia

Uma distinção crítica para desenvolvedores de EMR: o **SNGPC** (Sistema Nacional de Gerenciamento de Produtos Controlados) é um sistema **do lado da farmácia**, não do prescritor. O EMR não precisa integrar com SNGPC.

Para prescrições eletrônicas controladas, a integração obrigatória será com o **SNCR**:

**Integrações obrigatórias (a partir de jun/2026):**
- **SNCR**: para obter numeração nacional única das receitas controladas e registro de utilização
- **ICP-Brasil**: para assinatura digital qualificada

**Integrações recomendadas:**
- **RNDS** (Rede Nacional de Dados em Saúde): Portaria 6100/2024 instituiu REPM (Registro Eletrônico da Prescrição) e REDFM (Registro de Dispensação), com padrão FHIR R4
- **CFM Prescrição Eletrônica**: plataforma gratuita com validação integrada

**Não é necessário:**
- Integração com Receita Federal
- Integrações estaduais específicas (SNCR centraliza nacionalmente)
- Registro ou certificação do sistema junto ao CFM (não é obrigatório, embora o SBIS ofereça certificação voluntária)

O **SBIS-CFM** oferece certificação voluntária para sistemas S-RES (Sistema de Registro Eletrônico em Saúde), mas não é legalmente obrigatória. Contudo, o sistema deve cumprir os requisitos funcionais do NGS2 para operar 100% digital.

---

## Farmácias validam via portal ITI e SNCR

O processo de validação nas farmácias para receitas digitais controladas funciona assim:

**Ferramentas de validação oficiais:**
- **Portal ITI**: https://validar.iti.gov.br/ — verifica autenticidade da assinatura ICP-Brasil
- **Plataforma CFM**: https://prescricao.cfm.org.br/documento — valida receitas emitidas pela plataforma
- **App VALIDAR QR CODE**: aplicativo móvel para validação via QR Code
- **SNCR**: verificação da numeração e status de utilização (a partir de jun/2026)

**Fluxo de dispensação com receita digital controlada:**
1. Farmácia recebe receita (arquivo PDF por e-mail, WhatsApp ou na tela do celular do paciente)
2. Valida assinatura no portal ITI
3. Confirma numeração no SNCR
4. Verifica conformidade com Portaria 344/98
5. Registra utilização no SNCR (receita torna-se inutilizável)
6. Dispensa medicamento
7. Registra dispensação no SNGPC (transmissão XML em até 7 dias)
8. Arquiva registro eletrônico por **2 anos**

**Importante:** A farmácia **não precisa de software especial** para aceitar receitas digitais — basta acesso à internet e aos portais de validação. Porém, para controlados, precisa de certificado digital do farmacêutico para registro da dispensação e conformidade com SNGPC.

---

## Aspectos práticos da emissão e apresentação

**Emissão pelo médico:**
1. Cria receita no EMR
2. Sistema solicita numeração ao SNCR (para controlados)
3. Médico assina com certificado ICP-Brasil
4. Sistema envia dados ao RNDS (se integrado)
5. Receita disponibilizada ao paciente (PDF com QR code)

**Entrega ao paciente:**
- E-mail com PDF anexo
- WhatsApp/SMS com link para download
- QR Code para acesso direto
- Disponibilização no portal do paciente do EMR

**Apresentação na farmácia:**
- **Pode apresentar no celular** — não há obrigatoriedade de impressão
- A versão impressa serve apenas como referência ao documento eletrônico original
- O documento legal é o arquivo digital, não a cópia impressa
- Receita digitalizada (foto/scan de papel) **não é válida** para controlados

**Prazos de validade (idênticos à receita física):**

| Tipo de receita | Validade |
|-----------------|----------|
| Receita simples | 30 dias |
| Antimicrobianos | 10 dias |
| Controle especial (C1, C5) | 30 dias |
| GLP-1 (semaglutida, etc.) | 90 dias |
| Uso contínuo (não controlados) | Até 180 dias |

---

## Cronograma de implementação e requisitos técnicos do EMR

**Requisitos técnicos mínimos para o EMR:**
- Suporte a assinatura digital ICP-Brasil (A1, A3 ou nuvem)
- Geração de PDF assinado padrão PAdES
- Geração de QR Code para validação
- Verificação automática do registro profissional (integração com cadastro CFM)
- Conformidade com LGPD
- Trilha de auditoria completa
- Armazenamento seguro por 20 anos (conforme CFM 1.821/2007)
- **API de integração com SNCR** (obrigatório a partir de jun/2026)

**Cronograma regulatório:**

| Data | Marco |
|------|-------|
| Jul/2024 | RDC 873/2024 em vigor (SNCR instituído) |
| Jan/2025 | SNCR obrigatório para Vigilâncias Sanitárias |
| Fev/2026 | RDC 1.000/2025 em vigor (prescrições eletrônicas) |
| **Jun/2026** | SNCR disponível para integração com plataformas de prescrição |
| 30 dias após SNCR | Fim do período de transição |

**Recomendações para desenvolvedores:**
1. Implementar suporte a certificados ICP-Brasil (A1, A3, nuvem) imediatamente
2. Preparar arquitetura para integração SNCR via API (especificações a serem publicadas pela ANVISA)
3. Considerar integração voluntária com RNDS (padrão FHIR R4)
4. Manter documentação de conformidade com requisitos NGS2 do Manual SBIS
5. Implementar validação de CRM automatizada via cadastro nacional

---

## Conclusão

O desenvolvimento de EMR com prescrição digital controlada no Brasil é plenamente viável com certificado A1, exigindo conformidade com Resolução CFM 2.299/2021, Portaria 344/98 e, crucialmente, integração com o SNCR prevista pela RDC 1.000/2025. A receita branca controlada (Lista C1) pode ser emitida digitalmente desde que assinada com ICP-Brasil qualificado e, a partir de junho de 2026, integrada ao SNCR. O sistema regulatório brasileiro favorece a digitalização, mantendo equivalência de prazos e regras com receitas físicas, enquanto adiciona camadas de segurança como uso único verificado e rastreabilidade nacional.