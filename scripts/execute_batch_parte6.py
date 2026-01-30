#!/usr/bin/env python3
"""
Executar Batch Parte 6 - SQL direto no banco
"""

import psycopg2

def main():
    # Conectar ao banco
    conn = psycopg2.connect(
        host="localhost",
        port=5432,
        database="plenya_db",
        user="plenya_user",
        password="plenya_password"
    )

    cursor = conn.cursor()

    # Ler arquivo SQL
    with open("/home/user/plenya/scripts/batch_parte6_complete.sql", "r", encoding="utf-8") as f:
        sql_script = f.read()

    try:
        # Executar script completo
        print("Executando script SQL...")
        cursor.execute(sql_script)
        conn.commit()

        # Obter resultados da query final (verificação)
        if cursor.description:
            results = cursor.fetchall()
            print("\n" + "="*80)
            print("RESULTADOS DA VERIFICAÇÃO")
            print("="*80)
            for row in results:
                print(f"{row[0]:40} {row[1]:15} clin:{row[2]:5} pat:{row[3]:5} cond:{row[4]:5} {row[5]}")

        print("\n✅ Batch Parte 6 executado com sucesso!")

    except Exception as e:
        print(f"\n❌ Erro: {e}")
        conn.rollback()
        raise

    finally:
        cursor.close()
        conn.close()

if __name__ == "__main__":
    main()
