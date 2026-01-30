#!/usr/bin/env python3
"""
Script INTELIGENTE para popular lab_test_definitions
Analisa cada exame individualmente e cria estrutura hierárquica correta
"""

import csv
import re
import psycopg2
from psycopg2.extras import RealDictCursor
from datetime import datetime
import uuid
from difflib import SequenceMatcher

DB_CONFIG = {
    'host': 'localhost',
    'port': 5432,
    'database': 'plenya_db',
    'user': 'plenya_user',
    'password': 'plenya_dev_password'
}

EXAMES_GROUP_ID = '5a897f43-eb64-47be-a9e9-0d9bbee6e152'

class IntelligentLabTestMigration:
    def __init__(self):
        self.conn = None
        self.cursor = None
        self.tuss_data = []  # Lista de dicionários com nome e código
        self.created_tests = {}  # Mapeia score_item_id -> lab_test_id
        self.stats = {
            'total_processed': 0,
            'parents_created': 0,
            'children_created': 0,
            'tuss_matched': 0,
            'fuzzy_matches': 0,
            'errors': []
        }

    def connect(self):
        self.conn = psycopg2.connect(**DB_CONFIG)
        self.cursor = self.conn.cursor(cursor_factory=RealDictCursor)
        print("✓ Conectado ao banco")

    def load_tuss_intelligently(self):
        """Carrega TUSS e prepara para matching inteligente"""
        print("\n📋 Carregando códigos TUSS...")

        with open('/home/user/plenya/tuss.csv', 'r', encoding='latin-1') as f:
            content = f.read()
            lines = content.split('\n')

            for line in lines[1:]:  # Pula header
                if ';' in line:
                    parts = line.split(';')
                    if len(parts) >= 2:
                        nome = parts[0].strip()
                        codigo = parts[1].strip()

                        if nome and codigo and codigo.isdigit():
                            self.tuss_data.append({
                                'nome': nome.lower(),
                                'nome_original': nome,
                                'codigo': codigo
                            })

        print(f"✓ Carregados {len(self.tuss_data)} códigos TUSS")

    def similarity(self, a, b):
        """Calcula similaridade entre duas strings"""
        return SequenceMatcher(None, a.lower(), b.lower()).ratio()

    def find_tuss_fuzzy(self, exam_name, extracted_code=None):
        """Busca código TUSS usando matching fuzzy"""
        # Se já tem código extraído do nome, retorna direto
        if extracted_code:
            return extracted_code, 1.0, None

        clean_name = exam_name.lower()

        # Remove texto entre parênteses
        clean_name = re.sub(r'\([^)]*\)', '', clean_name).strip()

        # Remove conectores
        clean_name = re.sub(r'\s+(de|com|para|em|do|da|dos|das)\s+', ' ', clean_name)

        best_match = None
        best_score = 0.0

        for tuss_entry in self.tuss_data:
            # Tenta match exato primeiro
            if tuss_entry['nome'] in clean_name or clean_name in tuss_entry['nome']:
                return tuss_entry['codigo'], 1.0, tuss_entry['nome_original']

            # Calcula similaridade
            score = self.similarity(clean_name, tuss_entry['nome'])

            # Bonus para palavras-chave em comum
            exam_words = set(clean_name.split())
            tuss_words = set(tuss_entry['nome'].split())
            common_words = exam_words & tuss_words
            if len(common_words) >= 2:
                score += 0.2

            if score > best_score:
                best_score = score
                best_match = tuss_entry

        # Threshold: 0.7 para aceitar match
        if best_score >= 0.7 and best_match:
            return best_match['codigo'], best_score, best_match['nome_original']

        return None, 0.0, None

    def extract_tuss_from_name(self, name):
        """Extrai código TUSS se estiver no formato '40901122 - Nome'"""
        match = re.match(r'^(\d{8})\s*-\s*(.+)$', name.strip())
        if match:
            return match.group(1), match.group(2).strip()
        return None, name

    def is_parent_exam(self, exam_name):
        """Determina se o exame é um PAI (requestable) baseado no padrão do nome"""
        # Se tem código TUSS no nome, é PAI
        if re.match(r'^\d{8}\s*-', exam_name):
            return True

        # Nomes de painéis/conjuntos são pais
        panel_indicators = [
            'perfil', 'painel', 'lipidograma', 'hemograma completo',
            'gasometria', 'eletroforese', 'curva', 'rotina de urina',
            'urocultura', 'hepatite', 'genotipagem'
        ]

        name_lower = exam_name.lower()
        for indicator in panel_indicators:
            if indicator in name_lower:
                return True

        return False

    def categorize_exam(self, subgroup_name, exam_name):
        """
        Categoriza o exame usando APENAS as categorias permitidas pelo DB:
        'hematology', 'biochemistry', 'hormones', 'immunology', 'microbiology',
        'urine', 'imaging', 'functional', 'genetics', 'other'
        """
        name_lower = exam_name.lower()
        subgroup_lower = subgroup_name.lower()

        # Imagem
        if 'imagem' in subgroup_lower:
            return 'imaging'

        # Genética
        if any(x in name_lower for x in ['genotipagem', 'hla', 'jak2', 'dna', 'gene', 'polimorfismo', 'mutação']):
            return 'genetics'

        # Hematologia (inclui hemograma + coagulação)
        if any(x in name_lower for x in [
            'hemograma', 'hemoglobina', 'hematócrito', 'leucócito', 'plaqueta',
            'vhs', 'reticulócito', 'eritrócito', 'hemácia', 'vcm', 'hcm', 'chcm', 'rdw',
            'tap', 'inr', 'protrombina', 'fibrinogênio', 'd-dímero', 'ttpa', 'coagulação'
        ]):
            return 'hematology'

        # Hormônios (endocrinologia)
        if any(x in name_lower for x in [
            'tsh', 't3', 't4', 'cortisol', 'acth', 'testosterona', 'estradiol',
            'progesterona', 'fsh', 'lh', 'prolactina', 'dhea', 'shbg', 'hormônio',
            'dht', 'paratormônio', 'pth', 'gh', 'igf-1', 'aldosterona', 'renina'
        ]):
            return 'hormones'

        # Imunologia
        if any(x in name_lower for x in [
            'imunoglobulina', 'complemento', 'fan', 'anti-', 'pcr ultrassensível',
            'interleucina', 'il-6', 'fator reumatoide', 'anca', 'ena'
        ]):
            return 'immunology'

        # Microbiologia
        if any(x in name_lower for x in [
            'hepatite', 'hiv', 'vdrl', 'cultura', 'antibiograma', 'sorologia',
            'toxoplasma', 'rubéola', 'citomegalovírus', 'epstein'
        ]):
            return 'microbiology'

        # Urina específica
        if any(x in name_lower for x in [
            'eas ', 'eas-', 'elementos anormais sedimento', 'urina tipo 1',
            'urocultura', 'urina 24', 'microalbuminúria', 'proteinúria'
        ]):
            return 'urine'

        # Medicina Funcional / Especializados
        if any(x in name_lower for x in [
            'estresse oxidativo', 'permeabilidade intestinal', 'disbiose',
            'metilação', 'detoxificação', 'neurotransmissor', 'orgânico'
        ]):
            return 'functional'

        # Bioquímica (default amplo - inclui tudo que não se encaixa acima)
        # Vitaminas, minerais, lipídios, função renal, função hepática,
        # glicose, marcadores cardíacos, marcadores tumorais, etc.
        return 'biochemistry'

    def determine_result_type(self, unit):
        """Determina tipo de resultado"""
        if not unit:
            return 'text'

        unit_lower = unit.lower()

        # Categórico
        if any(x in unit_lower for x in ['grau', 'categoria', 'classificação', 'stage', 'classe']):
            return 'categorical'

        # Numérico
        numeric_patterns = ['mg', 'dl', 'ml', 'cm', 'mm', 'ng', 'pg', 'µg', 'ui', 'u/', '%', 'score', 'ratio', 'ppm', 'bpm', 'ms', 'nmol', 'μmol']
        if any(x in unit_lower for x in numeric_patterns):
            return 'numeric'

        # Boolean/qualitativo
        if any(x in unit_lower for x in ['positivo', 'negativo', 'reagente', 'não reagente']):
            return 'boolean'

        return 'text'

    def get_all_exams(self):
        """Busca TODOS os exames do grupo Exames em ordem"""
        query = """
            SELECT
                si.id,
                si.name,
                si.code,
                si.unit,
                si.unit_conversion,
                si.order,
                si.clinical_relevance,
                si.patient_explanation,
                si.conduct,
                sg.name as subgroup_name,
                sg.id as subgroup_id
            FROM score_items si
            JOIN score_subgroups sg ON si.subgroup_id = sg.id
            WHERE sg.group_id = %s
              AND si.deleted_at IS NULL
            ORDER BY sg.order, si.order
        """

        self.cursor.execute(query, (EXAMES_GROUP_ID,))
        return self.cursor.fetchall()

    def create_lab_test(self, exam, parent_id=None, is_requestable=True):
        """Cria um lab test definition"""

        # Extrai código TUSS
        extracted_tuss, clean_name = self.extract_tuss_from_name(exam['name'])

        # Busca TUSS
        tuss_code, match_score, tuss_match_name = self.find_tuss_fuzzy(clean_name, extracted_tuss)

        if tuss_code:
            self.stats['tuss_matched'] += 1
            if match_score < 1.0:
                self.stats['fuzzy_matches'] += 1
                print(f"      FUZZY MATCH ({match_score:.0%}): '{clean_name}' → '{tuss_match_name}' [{tuss_code}]")

        # Categoria
        category = self.categorize_exam(exam['subgroup_name'], clean_name)

        # Tipo de resultado
        result_type = self.determine_result_type(exam['unit'])

        # Código único
        code = exam['code'] if exam['code'] else f"PLN{str(uuid.uuid4())[:8].upper()}"

        # Insere
        insert_query = """
            INSERT INTO lab_test_definitions (
                code, name, tuss_code, category, is_requestable,
                parent_test_id, unit, unit_conversion, result_type,
                description, clinical_indications, display_order,
                is_active, created_at, updated_at
            ) VALUES (
                %(code)s, %(name)s, %(tuss_code)s, %(category)s, %(is_requestable)s,
                %(parent_test_id)s, %(unit)s, %(unit_conversion)s, %(result_type)s,
                %(description)s, %(clinical_indications)s, %(display_order)s,
                true, NOW(), NOW()
            )
            RETURNING id
        """

        data = {
            'code': code,
            'name': clean_name,
            'tuss_code': tuss_code,
            'category': category,
            'is_requestable': is_requestable,
            'parent_test_id': parent_id,
            'unit': exam['unit'],
            'unit_conversion': exam['unit_conversion'],
            'result_type': result_type,
            'description': exam['clinical_relevance'],
            'clinical_indications': exam['conduct'],
            'display_order': exam['order'] or 0
        }

        self.cursor.execute(insert_query, data)
        result = self.cursor.fetchone()

        # Registra
        self.created_tests[exam['id']] = result['id']

        return result['id']

    def process_intelligently(self):
        """Processa exames de forma inteligente"""
        print("\n🧠 PROCESSAMENTO INTELIGENTE\n")

        exams = self.get_all_exams()
        total = len(exams)
        print(f"Total de exames a processar: {total}\n")

        i = 0
        while i < total:
            exam = exams[i]

            # Verifica se é PAI
            if self.is_parent_exam(exam['name']):
                # É um exame PAI (requestable)
                print(f"\n{'='*80}")
                print(f"🔵 EXAME PAI (requestable): {exam['name']}")
                print(f"{'='*80}")

                parent_id = self.create_lab_test(exam, None, True)
                self.stats['parents_created'] += 1
                self.stats['total_processed'] += 1

                # Procura FILHOS (próximos items que não são pais)
                j = i + 1
                child_count = 0
                while j < total:
                    next_exam = exams[j]

                    # Se o próximo é PAI, para
                    if self.is_parent_exam(next_exam['name']):
                        break

                    # Se mudou de subgrupo, para
                    if next_exam['subgroup_id'] != exam['subgroup_id']:
                        break

                    # É FILHO deste PAI
                    print(f"  ↳ Parâmetro: {next_exam['name']} [{next_exam['unit'] or 'sem unidade'}]")
                    self.create_lab_test(next_exam, parent_id, False)
                    self.stats['children_created'] += 1
                    self.stats['total_processed'] += 1
                    child_count += 1
                    j += 1

                print(f"  Total de parâmetros: {child_count}")

                # Avança para o próximo PAI
                i = j
            else:
                # Exame independente (sem estrutura pai-filho clara)
                print(f"\n🟢 EXAME INDEPENDENTE: {exam['name']}")
                self.create_lab_test(exam, None, True)
                self.stats['parents_created'] += 1
                self.stats['total_processed'] += 1
                i += 1

        self.conn.commit()

    def generate_report(self):
        """Relatório final"""
        print("\n\n" + "="*100)
        print("📊 RELATÓRIO FINAL - MIGRAÇÃO INTELIGENTE")
        print("="*100)

        print(f"\n✅ EXAMES PROCESSADOS:")
        print(f"  Total: {self.stats['total_processed']}")
        print(f"  Exames principais (requestable=true): {self.stats['parents_created']}")
        print(f"  Parâmetros/sub-testes (requestable=false): {self.stats['children_created']}")

        print(f"\n🏥 CÓDIGOS TUSS:")
        print(f"  Total matched: {self.stats['tuss_matched']}")
        print(f"  Fuzzy matches (similaridade): {self.stats['fuzzy_matches']}")
        print(f"  Taxa de match: {self.stats['tuss_matched'] / self.stats['total_processed'] * 100:.1f}%")

        # Estatísticas do banco
        self.cursor.execute("""
            SELECT
                category,
                COUNT(*) as total,
                COUNT(CASE WHEN is_requestable THEN 1 END) as requestable,
                COUNT(CASE WHEN NOT is_requestable THEN 1 END) as parameters,
                COUNT(CASE WHEN tuss_code IS NOT NULL THEN 1 END) as with_tuss
            FROM lab_test_definitions
            WHERE deleted_at IS NULL
            GROUP BY category
            ORDER BY total DESC
        """)

        categories = self.cursor.fetchall()

        print(f"\n📊 DISTRIBUIÇÃO POR CATEGORIA:")
        print(f"{'Categoria':<20} {'Total':>8} {'Requestable':>12} {'Parâmetros':>12} {'Com TUSS':>10}")
        print("-" * 100)
        for cat in categories:
            print(f"{cat['category']:<20} {cat['total']:>8} {cat['requestable']:>12} {cat['parameters']:>12} {cat['with_tuss']:>10}")

        print("\n" + "="*100)
        print(f"✅ Migração concluída: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
        print("="*100)

    def run(self):
        try:
            self.connect()
            self.load_tuss_intelligently()
            self.process_intelligently()
            self.generate_report()
        except Exception as e:
            print(f"\n❌ Erro: {str(e)}")
            import traceback
            traceback.print_exc()
        finally:
            if self.cursor:
                self.cursor.close()
            if self.conn:
                self.conn.close()

if __name__ == '__main__':
    migration = IntelligentLabTestMigration()
    migration.run()
