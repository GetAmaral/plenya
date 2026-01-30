#!/usr/bin/env python3
"""
Script para popular lab_test_definitions a partir de score_items do grupo Exames
"""

import csv
import re
import psycopg2
from psycopg2.extras import RealDictCursor
from datetime import datetime
import uuid

# Configuração do banco
DB_CONFIG = {
    'host': 'localhost',
    'port': 5432,
    'database': 'plenya_db',
    'user': 'plenya_user',
    'password': 'plenya_dev_password'
}

# ID do grupo Exames
EXAMES_GROUP_ID = '5a897f43-eb64-47be-a9e9-0d9bbee6e152'

class LabTestMigration:
    def __init__(self):
        self.conn = None
        self.cursor = None
        self.tuss_map = {}
        self.stats = {
            'total_score_items': 0,
            'parent_tests': 0,
            'sub_tests': 0,
            'tests_created': 0,
            'tests_updated': 0,
            'tuss_matched': 0,
            'errors': []
        }

    def connect(self):
        """Conecta ao banco de dados"""
        self.conn = psycopg2.connect(**DB_CONFIG)
        self.cursor = self.conn.cursor(cursor_factory=RealDictCursor)
        print("✓ Conectado ao banco de dados")

    def load_tuss_codes(self):
        """Carrega códigos TUSS do arquivo CSV"""
        print("\n📋 Carregando códigos TUSS...")
        with open('/home/user/plenya/tuss.csv', 'r', encoding='utf-8', errors='ignore') as f:
            reader = csv.DictReader(f, delimiter=';')
            for row in reader:
                if row.get('EXAME') and row.get('CÓD. TUSS'):
                    exame = row['EXAME'].strip()
                    tuss = row['CÓD. TUSS'].strip()
                    if exame and tuss and tuss.isdigit():
                        self.tuss_map[exame.lower()] = tuss

        print(f"✓ Carregados {len(self.tuss_map)} códigos TUSS")

    def extract_tuss_from_name(self, name):
        """Extrai código TUSS do nome do exame (ex: '40901122 - USG abdome total')"""
        match = re.match(r'^(\d{8})\s*-\s*(.+)$', name)
        if match:
            return match.group(1), match.group(2).strip()
        return None, name

    def find_tuss_code(self, name, extracted_tuss=None):
        """Busca código TUSS pelo nome ou usa o extraído"""
        if extracted_tuss:
            return extracted_tuss

        # Busca no mapa TUSS
        name_clean = name.lower().strip()

        # Busca exata
        if name_clean in self.tuss_map:
            return self.tuss_map[name_clean]

        # Busca parcial (remove parênteses e conteúdo)
        name_without_parentheses = re.sub(r'\([^)]*\)', '', name_clean).strip()
        if name_without_parentheses in self.tuss_map:
            return self.tuss_map[name_without_parentheses]

        # Busca por similaridade (contém)
        for tuss_name, tuss_code in self.tuss_map.items():
            if tuss_name in name_clean or name_clean in tuss_name:
                return tuss_code

        return None

    def determine_result_type(self, unit):
        """Determina o tipo de resultado baseado na unidade"""
        if not unit:
            return 'text'

        unit_lower = unit.lower()

        # Tipos numéricos
        numeric_indicators = ['mg', 'dl', 'ml', 'cm', 'mm', 'ng', 'pg', 'µg', 'ui',
                             'mmol', 'µmol', '%', 'ratio', 'índice', 'score']

        for indicator in numeric_indicators:
            if indicator in unit_lower:
                return 'numeric'

        # Tipos categóricos
        categorical_indicators = ['grau', 'classificação', 'classe', 'positivo', 'negativo']

        for indicator in categorical_indicators:
            if indicator in unit_lower:
                return 'categorical'

        return 'text'

    def determine_category(self, subgroup_name):
        """Determina a categoria baseada no subgrupo"""
        category_map = {
            'imagem': 'imaging',
            'hemograma': 'hematology',
            'hormônio': 'endocrinology',
            'lipídio': 'biochemistry',
            'renal': 'biochemistry',
            'hepático': 'biochemistry',
            'mineral': 'biochemistry',
            'vitamina': 'biochemistry',
            'marcador': 'tumor_markers',
            'autoimune': 'immunology',
            'infecção': 'microbiology',
            'genética': 'genetics'
        }

        subgroup_lower = subgroup_name.lower()

        for key, value in category_map.items():
            if key in subgroup_lower:
                return value

        return 'other'

    def get_score_items(self):
        """Busca todos os score_items do grupo Exames"""
        print("\n🔍 Buscando score items do grupo Exames...")

        query = """
            SELECT
                si.id,
                si.name,
                si.code,
                si.unit,
                si.unit_conversion,
                si.parent_item_id,
                si.order,
                si.clinical_relevance,
                si.patient_explanation,
                si.conduct,
                sg.name as subgroup_name
            FROM score_items si
            JOIN score_subgroups sg ON si.subgroup_id = sg.id
            WHERE sg.group_id = %s
              AND si.deleted_at IS NULL
            ORDER BY sg.name, si.order
        """

        self.cursor.execute(query, (EXAMES_GROUP_ID,))
        items = self.cursor.fetchall()

        self.stats['total_score_items'] = len(items)
        print(f"✓ Encontrados {len(items)} score items")

        return items

    def create_or_update_lab_test(self, item, parent_id=None):
        """Cria ou atualiza um lab test definition"""

        # Extrai código TUSS do nome, se houver
        extracted_tuss, clean_name = self.extract_tuss_from_name(item['name'])

        # Busca código TUSS
        tuss_code = self.find_tuss_code(clean_name, extracted_tuss)

        if tuss_code:
            self.stats['tuss_matched'] += 1

        # Determina tipo de resultado
        result_type = self.determine_result_type(item['unit'])

        # Determina categoria
        category = self.determine_category(item['subgroup_name'])

        # Gera código único se não houver
        code = item['code'] if item['code'] else f"PLN{str(uuid.uuid4())[:8].upper()}"

        # Define se é requestable (exames principais são requestable)
        is_requestable = parent_id is None

        # Prepara dados
        lab_test_data = {
            'code': code,
            'name': clean_name,
            'short_name': clean_name[:100] if len(clean_name) > 100 else None,
            'tuss_code': tuss_code,
            'category': category,
            'is_requestable': is_requestable,
            'parent_test_id': parent_id,
            'unit': item['unit'],
            'unit_conversion': item['unit_conversion'],
            'result_type': result_type,
            'description': item['clinical_relevance'],
            'clinical_indications': item['conduct'],
            'display_order': item['order'] or 0,
            'is_active': True
        }

        # Verifica se já existe
        check_query = "SELECT id FROM lab_test_definitions WHERE code = %s AND deleted_at IS NULL"
        self.cursor.execute(check_query, (code,))
        existing = self.cursor.fetchone()

        if existing:
            # Atualiza
            update_query = """
                UPDATE lab_test_definitions SET
                    name = %(name)s,
                    short_name = %(short_name)s,
                    tuss_code = %(tuss_code)s,
                    category = %(category)s,
                    is_requestable = %(is_requestable)s,
                    parent_test_id = %(parent_test_id)s,
                    unit = %(unit)s,
                    unit_conversion = %(unit_conversion)s,
                    result_type = %(result_type)s,
                    description = %(description)s,
                    clinical_indications = %(clinical_indications)s,
                    display_order = %(display_order)s,
                    is_active = %(is_active)s,
                    updated_at = NOW()
                WHERE code = %(code)s
                RETURNING id
            """
            self.cursor.execute(update_query, lab_test_data)
            result = self.cursor.fetchone()
            self.stats['tests_updated'] += 1
            return result['id']
        else:
            # Insere
            insert_query = """
                INSERT INTO lab_test_definitions (
                    code, name, short_name, tuss_code, category,
                    is_requestable, parent_test_id, unit, unit_conversion,
                    result_type, description, clinical_indications,
                    display_order, is_active, created_at, updated_at
                ) VALUES (
                    %(code)s, %(name)s, %(short_name)s, %(tuss_code)s, %(category)s,
                    %(is_requestable)s, %(parent_test_id)s, %(unit)s, %(unit_conversion)s,
                    %(result_type)s, %(description)s, %(clinical_indications)s,
                    %(display_order)s, %(is_active)s, NOW(), NOW()
                )
                RETURNING id
            """
            self.cursor.execute(insert_query, lab_test_data)
            result = self.cursor.fetchone()
            self.stats['tests_created'] += 1
            return result['id']

    def process_items(self):
        """Processa todos os items"""
        print("\n⚙️  Processando items...")

        items = self.get_score_items()

        # Separa pais e filhos
        parent_items = [item for item in items if not item['parent_item_id']]
        child_items = [item for item in items if item['parent_item_id']]

        self.stats['parent_tests'] = len(parent_items)
        self.stats['sub_tests'] = len(child_items)

        print(f"  - Exames principais (requestable): {len(parent_items)}")
        print(f"  - Parâmetros/sub-testes: {len(child_items)}")

        # Mapa de score_item_id -> lab_test_id
        id_map = {}

        # Processa exames principais primeiro
        print("\n📝 Criando exames principais...")
        for item in parent_items:
            try:
                lab_test_id = self.create_or_update_lab_test(item)
                id_map[item['id']] = lab_test_id
                print(f"  ✓ {item['name'][:80]}")
            except Exception as e:
                error_msg = f"Erro ao processar '{item['name']}': {str(e)}"
                self.stats['errors'].append(error_msg)
                print(f"  ✗ {error_msg}")

        # Processa sub-testes
        print("\n📝 Criando parâmetros/sub-testes...")
        for item in child_items:
            try:
                parent_lab_test_id = id_map.get(item['parent_item_id'])
                if not parent_lab_test_id:
                    raise Exception(f"Parent test não encontrado para {item['parent_item_id']}")

                lab_test_id = self.create_or_update_lab_test(item, parent_lab_test_id)
                id_map[item['id']] = lab_test_id
                print(f"  ✓ {item['name'][:80]}")
            except Exception as e:
                error_msg = f"Erro ao processar '{item['name']}': {str(e)}"
                self.stats['errors'].append(error_msg)
                print(f"  ✗ {error_msg}")

        self.conn.commit()

    def generate_report(self):
        """Gera relatório final"""
        print("\n" + "="*80)
        print("📊 RELATÓRIO DE MIGRAÇÃO - LAB TEST DEFINITIONS")
        print("="*80)

        print(f"\n📦 SCORE ITEMS PROCESSADOS:")
        print(f"  Total de items: {self.stats['total_score_items']}")
        print(f"  Exames principais: {self.stats['parent_tests']}")
        print(f"  Parâmetros/sub-testes: {self.stats['sub_tests']}")

        print(f"\n✅ LAB TEST DEFINITIONS:")
        print(f"  Criados: {self.stats['tests_created']}")
        print(f"  Atualizados: {self.stats['tests_updated']}")
        print(f"  Total processado: {self.stats['tests_created'] + self.stats['tests_updated']}")

        print(f"\n🏥 CÓDIGOS TUSS:")
        print(f"  Códigos TUSS encontrados: {self.stats['tuss_matched']}")
        print(f"  Taxa de match: {self.stats['tuss_matched'] / self.stats['total_score_items'] * 100:.1f}%")

        if self.stats['errors']:
            print(f"\n⚠️  ERROS ({len(self.stats['errors'])}):")
            for error in self.stats['errors'][:10]:
                print(f"  - {error}")
            if len(self.stats['errors']) > 10:
                print(f"  ... e mais {len(self.stats['errors']) - 10} erros")
        else:
            print(f"\n✨ Sem erros!")

        # Estatísticas do banco
        self.cursor.execute("SELECT COUNT(*) as total FROM lab_test_definitions WHERE deleted_at IS NULL")
        total_in_db = self.cursor.fetchone()['total']

        self.cursor.execute("SELECT COUNT(*) as total FROM lab_test_definitions WHERE is_requestable = true AND deleted_at IS NULL")
        requestable = self.cursor.fetchone()['total']

        print(f"\n💾 BANCO DE DADOS:")
        print(f"  Total de lab test definitions: {total_in_db}")
        print(f"  Exames solicitáveis (requestable): {requestable}")
        print(f"  Parâmetros/sub-testes: {total_in_db - requestable}")

        print("\n" + "="*80)
        print(f"✅ Migração concluída em {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
        print("="*80)

    def run(self):
        """Executa a migração completa"""
        try:
            self.connect()
            self.load_tuss_codes()
            self.process_items()
            self.generate_report()
        except Exception as e:
            print(f"\n❌ Erro fatal: {str(e)}")
            import traceback
            traceback.print_exc()
        finally:
            if self.cursor:
                self.cursor.close()
            if self.conn:
                self.conn.close()
                print("\n✓ Conexão fechada")

if __name__ == '__main__':
    migration = LabTestMigration()
    migration.run()
