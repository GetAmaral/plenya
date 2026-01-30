#!/usr/bin/env python3
"""
Upload IL-11 article (large file)
"""

import requests
from pathlib import Path

BASE_URL = "http://localhost:3001/api/v1"
AUTH_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIwM2EwY2EzMy1lMDJhLTRiOWItOGQ0My1mZjg2YWViMTU3MzAiLCJlbWFpbCI6ImFydGljbGV1cGxvYWRlckBwbGVueWEuY29tIiwicm9sZSI6ImRvY3RvciIsImlzcyI6InBsZW55YS1lbXIiLCJleHAiOjE3NjkzNzY5MjAsImlhdCI6MTc2OTM3NjAyMH0.DxIV6kWS6stVTcIPmLc2wRzP4tnuAn8D23zpu2V-Z-g"
UPLOAD_DIR = Path("/home/user/plenya/uploads/articles")

ARTICLE = {
    "filename": "IL11_Lifespan_Extension_Nature_2024.pdf",
    "title": "IL-11 and Lifespan Extension - Nature 2024",
    "description": "Interleukin-11 inhibition extends lifespan and healthspan in mice",
    "keywords": ["IL-11", "interleukin", "inflammation", "longevity", "anti-aging"],
    "search_terms": ["il-11", "interleukin", "interleucina", "inflamação", "inflammation"],
}


def upload_article():
    """Upload the IL-11 article."""
    filepath = UPLOAD_DIR / ARTICLE["filename"]

    if not filepath.exists():
        print(f"✗ File not found: {filepath}")
        return None

    print(f"📤 Uploading: {ARTICLE['filename']}")
    print(f"   Size: {filepath.stat().st_size / 1024 / 1024:.1f} MB")

    session = requests.Session()
    session.headers.update({"Authorization": f"Bearer {AUTH_TOKEN}"})

    try:
        with open(filepath, 'rb') as f:
            files = {'file': (ARTICLE['filename'], f, 'application/pdf')}
            data = {
                'title': ARTICLE['title'],
                'description': ARTICLE.get('description', ''),
            }

            response = session.post(
                f"{BASE_URL}/articles/upload",
                files=files,
                data=data,
                timeout=120  # Longer timeout for large file
            )

        if response.status_code in (200, 201):
            result = response.json()
            article_id = result.get('id') or result.get('data', {}).get('id')

            if article_id:
                print(f"   ✓ Uploaded successfully!")
                print(f"   Article ID: {article_id}")
                return article_id
            else:
                print(f"   ✗ No ID in response")
                return None
        else:
            print(f"   ✗ Upload failed: {response.status_code}")
            print(f"   Response: {response.text}")
            return None

    except Exception as e:
        print(f"   ✗ Exception: {e}")
        return None


def search_score_items(search_terms):
    """Search for score items."""
    session = requests.Session()
    session.headers.update({"Authorization": f"Bearer {AUTH_TOKEN}"})

    try:
        response = session.get(f"{BASE_URL}/score-groups/tree", timeout=30)
        if response.status_code != 200:
            return []

        data = response.json()
        groups = data.get('data', []) if isinstance(data, dict) else data
        matched_ids = []

        def search_in_tree(node, search_terms):
            if 'subgroups' in node:
                for subgroup in node['subgroups']:
                    if 'items' in subgroup:
                        for item in subgroup['items']:
                            item_name = item.get('name', '').lower()
                            item_desc = item.get('description', '').lower()

                            for term in search_terms:
                                if term.lower() in item_name or term.lower() in item_desc:
                                    item_id = item.get('id')
                                    if item_id and item_id not in matched_ids:
                                        matched_ids.append(item_id)
                                        print(f"   → {item.get('name')} ({item_id})")
                                        break
                    search_in_tree(subgroup, search_terms)

        for group in groups:
            search_in_tree(group, search_terms)

        return matched_ids

    except Exception as e:
        print(f"Error: {e}")
        return []


def link_items(article_id, score_item_ids):
    """Link score items to article."""
    session = requests.Session()
    session.headers.update({"Authorization": f"Bearer {AUTH_TOKEN}"})

    try:
        response = session.post(
            f"{BASE_URL}/articles/{article_id}/score-items",
            json={"scoreItemIds": score_item_ids},
            timeout=30
        )

        if response.status_code in (200, 201):
            print(f"   ✓ Linked {len(score_item_ids)} items")
            return True
        else:
            print(f"   ✗ Failed: {response.status_code}")
            return False

    except Exception as e:
        print(f"   ✗ Error: {e}")
        return False


def main():
    print("=" * 80)
    print("UPLOADING IL-11 ARTICLE")
    print("=" * 80 + "\n")

    article_id = upload_article()

    if article_id:
        print(f"\n🔍 Searching for related score items...")
        item_ids = search_score_items(ARTICLE['search_terms'])

        if item_ids:
            print(f"\n🔗 Linking {len(item_ids)} score item(s)...")
            link_items(article_id, item_ids)
        else:
            print("   ⚠ No matching score items found")

    print("\n" + "=" * 80)
    print("DONE")
    print("=" * 80)


if __name__ == "__main__":
    main()
