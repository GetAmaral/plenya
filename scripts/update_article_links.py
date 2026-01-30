#!/usr/bin/env python3
"""
Update article linkages for Selenium and Sodium articles
"""

import requests
from typing import List

BASE_URL = "http://localhost:3001/api/v1"
AUTH_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIwM2EwY2EzMy1lMDJhLTRiOWItOGQ0My1mZjg2YWViMTU3MzAiLCJlbWFpbCI6ImFydGljbGV1cGxvYWRlckBwbGVueWEuY29tIiwicm9sZSI6ImRvY3RvciIsImlzcyI6InBsZW55YS1lbXIiLCJleHAiOjE3NjkzNzY5MjAsImlhdCI6MTc2OTM3NjAyMH0.DxIV6kWS6stVTcIPmLc2wRzP4tnuAn8D23zpu2V-Z-g"

# Article IDs from previous upload
ARTICLES = [
    {
        "id": "2a7b9e9f-004d-4bf2-86d6-b4f94be349ab",
        "title": "Selenium and Mortality",
        "search_terms": ["selênio", "selenium"],
    },
    {
        "id": "676e8556-a3dd-46d7-b1b2-63f1f5f75644",
        "title": "Sodium U-Shape and Aging",
        "search_terms": ["sódio", "sodio", "sodium", "eletrólito"],
    },
]


def search_score_items(search_terms: List[str]) -> List[str]:
    """Search for score items matching the given terms."""
    session = requests.Session()
    session.headers.update({"Authorization": f"Bearer {AUTH_TOKEN}"})

    try:
        response = session.get(f"{BASE_URL}/score-groups/tree", timeout=30)
        if response.status_code != 200:
            print(f"Failed to fetch score groups: {response.status_code}")
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


def link_items(article_id: str, score_item_ids: List[str]) -> bool:
    """Link score items to an article."""
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
            print(f"   ✗ Failed: {response.status_code} - {response.text}")
            return False

    except Exception as e:
        print(f"   ✗ Error: {e}")
        return False


def main():
    print("=" * 80)
    print("UPDATING ARTICLE LINKAGES")
    print("=" * 80)

    for article in ARTICLES:
        print(f"\n{article['title']}")
        print(f"Article ID: {article['id']}")
        print(f"Searching for: {', '.join(article['search_terms'])}")

        item_ids = search_score_items(article['search_terms'])

        if item_ids:
            print(f"\nFound {len(item_ids)} matching items. Linking...")
            link_items(article['id'], item_ids)
        else:
            print("   ⚠ No matching score items found")

    print("\n" + "=" * 80)
    print("DONE")
    print("=" * 80)


if __name__ == "__main__":
    main()
