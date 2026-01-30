#!/usr/bin/env node
/**
 * Enriquecimento: Hemácias - Mulheres
 * Item ID: 501fd84a-a440-4c13-9b11-35e2f69017d1
 */

import Anthropic from '@anthropic-ai/sdk';
import pg from 'pg';
const { Client } = pg;

const ITEM_ID = '501fd84a-a440-4c13-9b11-35e2f69017d1';
const ITEM_NAME = 'Hemácias - Mulheres';

async function getDbConnection() {
  const client = new Client({
    host: process.env.DB_HOST || 'db',
    port: parseInt(process.env.DB_PORT || '5432'),
    database: process.env.DB_NAME || 'plenya_db',
    user: process.env.DB_USER || 'plenya_user',
    password: process.env.DB_PASSWORD || 'plenya_password',
  });
  await client.connect();
  return client;
}

async function searchArticles(anthropic) {
  console.log('\n' + '='.repeat(80));
  console.log(`BUSCANDO ARTIGOS CIENTÍFICOS: ${ITEM_NAME}`);
  console.log('='.repeat(80) + '\n');

  const searchPrompt = `Busque 3-4 artigos científicos sobre contagem de hemácias (eritrócitos) em mulheres:

TÓPICOS:
- Red blood cell count women (reference ranges 4.0-5.2 million/µL)
- Hemoglobin levels women vs men (sex differences)
- Iron deficiency anemia menstruation
- Menstrual blood loss and erythrocyte count
- Hormonal effects on RBC production

FONTES PRIORITÁRIAS: PubMed, Google Scholar, UpToDate

Para cada artigo, extrair:
1. Título completo
2. Autores principais
3. DOI/PMID
4. Ano
5. 2-3 insights clínicos principais

Liste os artigos com título, autores, DOI e principais achados de forma clara.`;

  const message = await anthropic.messages.create({
    model: 'claude-sonnet-4-5-20250929',
    max_tokens: 4000,
    messages: [{ role: 'user', content: searchPrompt }],
  });

  const articlesText = message.content
    .filter((block) => block.type === 'text')
    .map((block) => block.text)
    .join('\n');

  console.log(`\nARTIGOS ENCONTRADOS:\n${articlesText.slice(0, 500)}...\n`);
  return articlesText;
}

async function enrichContent(anthropic, articlesText) {
  console.log('\n' + '='.repeat(80));
  console.log('GERANDO CONTEÚDO CLÍNICO EM PT-BR');
  console.log('='.repeat(80) + '\n');

  const enrichmentPrompt = `Com base nos artigos científicos abaixo, gere conteúdo clínico para o item "${ITEM_NAME}":

ARTIGOS:
${articlesText}

CONTEXTO CLÍNICO:
- Hemácias (eritrócitos): células vermelhas que transportam oxigênio
- Valores de referência mulheres: 4.0-5.2 milhões/µL (vs homens 4.5-5.5)
- Diferenças por menstruação (perda ferro), hormônios, gravidez
- Baixo = anemia; alto = policitemia/desidratação

GERAR EM PT-BR (formato JSON):

{
  "clinical_relevance": "Texto médico 150-200 palavras. Por que hemácias são menores em mulheres? Causas de alterações (menstruação, ferro, B12). Quando investigar?",

  "patient_explanation": "Texto leigo 100-150 palavras. 'Suas hemácias são as células que levam oxigênio...'. Explicar valores, sintomas anemia (cansaço, palidez).",

  "conduct": "Texto objetivo 80-120 palavras. Se baixo (<4.0): ferritina, ferro, B12, reticulócitos. Se alto (>5.2): hidratação, EPO, policitemia vera. Quando referenciar hematologista."
}

REGRAS:
- PT-BR técnico mas compreensível
- Focar diferenças de gênero (menstruação, ferro)
- Incluir valores numéricos
- Ser objetivo e prático
- JSON válido (escapar aspas)`;

  const message = await anthropic.messages.create({
    model: 'claude-sonnet-4-5-20250929',
    max_tokens: 3000,
    messages: [{ role: 'user', content: enrichmentPrompt }],
  });

  const content = message.content
    .filter((block) => block.type === 'text')
    .map((block) => block.text)
    .join('\n');

  console.log(`CONTEÚDO GERADO:\n${content.slice(0, 400)}...\n`);

  // Extrair JSON
  const jsonMatch = content.match(/\{[\s\S]*\}/);
  if (jsonMatch) {
    return JSON.parse(jsonMatch[0]);
  }
  return null;
}

async function saveArticle(dbClient, article, articleIds) {
  if (!article.title) return;

  const url = article.doi
    ? `https://doi.org/${article.doi}`
    : `https://pubmed.ncbi.nlm.nih.gov/search?term=${encodeURIComponent(article.title.slice(0, 50))}`;

  const result = await dbClient.query(
    `INSERT INTO articles (
      title, url, content, article_type, created_at, updated_at
    ) VALUES ($1, $2, $3, $4, NOW(), NOW())
    RETURNING id`,
    [
      article.title.slice(0, 300),
      url.slice(0, 500),
      (article.content || '').slice(0, 1000),
      'scientific',
    ]
  );

  const articleId = result.rows[0].id;
  articleIds.push(articleId);

  await dbClient.query(
    `INSERT INTO score_item_articles (score_item_id, article_id, created_at)
     VALUES ($1, $2, NOW())
     ON CONFLICT DO NOTHING`,
    [ITEM_ID, articleId]
  );

  console.log(`✓ Artigo ${articleIds.length} salvo: ${articleId} - ${article.title.slice(0, 60)}...`);
}

async function saveArticles(dbClient, articlesText) {
  console.log('\n' + '='.repeat(80));
  console.log('SALVANDO ARTIGOS NO BANCO');
  console.log('='.repeat(80) + '\n');

  const articleIds = [];

  // Parse simples - buscar padrões de título e DOI
  const lines = articlesText.split('\n');
  let currentArticle = {};

  for (let i = 0; i < lines.length; i++) {
    const line = lines[i].trim();

    // Detectar título
    if (line.match(/^(Title|Título|Article|[0-9]\.|###)/i)) {
      if (currentArticle.title) {
        // Salvar artigo anterior
        await saveArticle(dbClient, currentArticle, articleIds);
        currentArticle = {};
      }
      currentArticle.title = line.replace(/^(Title|Título|Article|[0-9]\.+|###):?\s*/i, '');
    }

    // Detectar DOI/PMID
    if (line.match(/DOI|PMID|doi\.org/i)) {
      currentArticle.doi = line.match(/10\.\d{4,}\/[^\s]+|PMID:?\s*\d+/i)?.[0] || '';
    }

    // Acumular conteúdo
    if (line && !line.match(/^(Title|DOI|PMID|Authors?|Year)/i)) {
      currentArticle.content = (currentArticle.content || '') + line + '\n';
    }
  }

  // Salvar último artigo
  if (currentArticle.title) {
    await saveArticle(dbClient, currentArticle, articleIds);
  }

  // Se nenhum artigo foi parseado, criar um genérico
  if (articleIds.length === 0) {
    const result = await dbClient.query(
      `INSERT INTO articles (
        title, url, content, article_type, created_at, updated_at
      ) VALUES ($1, $2, $3, $4, NOW(), NOW())
      RETURNING id`,
      [
        `${ITEM_NAME} - Evidências Científicas`,
        'https://pubmed.ncbi.nlm.nih.gov/search?term=RBC+women+reference+ranges',
        articlesText.slice(0, 1000),
        'scientific',
      ]
    );
    const articleId = result.rows[0].id;
    articleIds.push(articleId);

    await dbClient.query(
      `INSERT INTO score_item_articles (score_item_id, article_id, created_at)
       VALUES ($1, $2, NOW())
       ON CONFLICT DO NOTHING`,
      [ITEM_ID, articleId]
    );

    console.log(`✓ Artigo genérico salvo: ${articleId}`);
  }

  console.log(`\n✓ ${articleIds.length} artigo(s) salvo(s) e vinculado(s)`);
  return articleIds;
}

async function updateScoreItem(dbClient, enrichedData) {
  console.log('\n' + '='.repeat(80));
  console.log(`ATUALIZANDO SCORE ITEM: ${ITEM_ID}`);
  console.log('='.repeat(80) + '\n');

  // Buscar dados atuais
  const current = await dbClient.query(
    'SELECT name, clinical_relevance, patient_explanation, conduct FROM score_items WHERE id = $1',
    [ITEM_ID]
  );

  if (current.rows.length > 0) {
    const row = current.rows[0];
    console.log('Dados atuais:');
    console.log(`  Name: ${row.name}`);
    console.log(`  Clinical: ${(row.clinical_relevance || '').length} chars`);
    console.log(`  Patient: ${(row.patient_explanation || '').length} chars`);
    console.log(`  Conduct: ${(row.conduct || '').length} chars`);
  }

  // Atualizar
  await dbClient.query(
    `UPDATE score_items
     SET clinical_relevance = $1,
         patient_explanation = $2,
         conduct = $3,
         updated_at = NOW()
     WHERE id = $4`,
    [
      enrichedData.clinical_relevance,
      enrichedData.patient_explanation,
      enrichedData.conduct,
      ITEM_ID,
    ]
  );

  console.log('\n✓ Score item atualizado com sucesso!');
  console.log(`  Clinical: ${enrichedData.clinical_relevance.length} chars`);
  console.log(`  Patient: ${enrichedData.patient_explanation.length} chars`);
  console.log(`  Conduct: ${enrichedData.conduct.length} chars`);
}

async function main() {
  console.log('\n' + '#'.repeat(80));
  console.log(`# ENRIQUECIMENTO: ${ITEM_NAME}`);
  console.log(`# ID: ${ITEM_ID}`);
  console.log(`# Timestamp: ${new Date().toISOString()}`);
  console.log('#'.repeat(80));

  // Verificar API key
  const apiKey = process.env.ANTHROPIC_API_KEY;
  if (!apiKey) {
    console.error('\n❌ ERRO: ANTHROPIC_API_KEY não configurada');
    process.exit(1);
  }

  let dbClient = null;

  try {
    // Inicializar Claude
    const anthropic = new Anthropic({ apiKey });

    // 1. Buscar artigos
    const articlesText = await searchArticles(anthropic);

    // 2. Gerar conteúdo enriquecido
    const enrichedData = await enrichContent(anthropic, articlesText);

    if (!enrichedData) {
      console.error('\n❌ ERRO: Falha ao gerar conteúdo enriquecido');
      process.exit(1);
    }

    // 3. Conectar ao banco
    dbClient = await getDbConnection();

    // 4. Salvar artigos
    const articleIds = await saveArticles(dbClient, articlesText);

    // 5. Atualizar score_item
    await updateScoreItem(dbClient, enrichedData);

    console.log('\n' + '#'.repeat(80));
    console.log('# ✓ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO');
    console.log(`# Artigos salvos: ${articleIds.length}`);
    console.log(`# Item atualizado: ${ITEM_ID}`);
    console.log('#'.repeat(80) + '\n');

    process.exit(0);
  } catch (error) {
    console.error('\n❌ ERRO:', error);
    console.error(error.stack);
    process.exit(1);
  } finally {
    if (dbClient) {
      await dbClient.end();
    }
  }
}

main();
