#!/usr/bin/env python3
"""
Article Upload Script for Plenya EMR
Uploads scientific PDFs and links them to relevant score items.
"""

import json
import os
import sys
from pathlib import Path
from typing import Dict, List, Optional, Tuple
import requests
from datetime import datetime

# Configuration
BASE_URL = "http://localhost:3001/api/v1"
AUTH_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIwM2EwY2EzMy1lMDJhLTRiOWItOGQ0My1mZjg2YWViMTU3MzAiLCJlbWFpbCI6ImFydGljbGV1cGxvYWRlckBwbGVueWEuY29tIiwicm9sZSI6ImRvY3RvciIsImlzcyI6InBsZW55YS1lbXIiLCJleHAiOjE3NjkzNzY5MjAsImlhdCI6MTc2OTM3NjAyMH0.DxIV6kWS6stVTcIPmLc2wRzP4tnuAn8D23zpu2V-Z-g"
UPLOAD_DIR = Path("/home/user/plenya/uploads/articles")

# Article definitions with metadata
ARTICLES = [
    {
        "filename": "ESC_Heart_Failure_Algorithms_2023.pdf",
        "title": "ESC Heart Failure Diagnostic Algorithms 2023",
        "description": "NT-proBNP diagnostic algorithms for heart failure from European Society of Cardiology",
        "keywords": ["NT-proBNP", "BNP", "heart failure", "cardiac biomarkers"],
        "search_terms": ["nt-probnp", "bnp", "natriuretic"],
    },
    {
        "filename": "HDL_Paradox_Frontiers_2025.pdf",
        "title": "HDL Cholesterol Paradox - Frontiers 2025",
        "description": "Research on the paradoxical relationship between HDL cholesterol and cardiovascular risk",
        "keywords": ["HDL", "cholesterol", "lipid panel", "cardiovascular"],
        "search_terms": ["hdl", "cholesterol", "lipid"],
    },
    {
        "filename": "hsCRP_HDL_Mortality_Frontiers_2025.pdf",
        "title": "hs-CRP/HDL Ratio and Mortality - Frontiers 2025",
        "description": "High-sensitivity CRP to HDL ratio as a biomarker for mortality prediction",
        "keywords": ["hs-CRP", "HDL", "inflammation", "mortality", "biomarker"],
        "search_terms": ["crp", "hdl", "inflammation", "proteina c"],
    },
    {
        "filename": "Selenium_Mortality_Frontiers_2025.pdf",
        "title": "Selenium and Mortality - Frontiers 2025",
        "description": "Association between selenium levels and all-cause mortality",
        "keywords": ["selenium", "mortality", "trace elements", "longevity"],
        "search_terms": ["selênio", "selenium", "selenio"],
    },
    {
        "filename": "Sodium_U_Shape_Aging_Frontiers_2025.pdf",
        "title": "Sodium U-Shape Curve and Biological Aging - Frontiers 2025",
        "description": "U-shaped relationship between sodium levels and biological aging markers",
        "keywords": ["sodium", "electrolytes", "aging", "longevity"],
        "search_terms": ["sódio", "sodio", "sodium", "eletrólito", "electrolyte"],
    },
    {
        "filename": "IL11_Lifespan_Extension_Nature_2024.pdf",
        "title": "IL-11 and Lifespan Extension - Nature 2024",
        "description": "Interleukin-11 inhibition extends lifespan and healthspan in mice",
        "keywords": ["IL-11", "interleukin", "inflammation", "longevity", "anti-aging"],
        "search_terms": ["il-11", "interleukin", "inflammation", "aging"],
    },
]


class ArticleUploader:
    """Handles article uploads and score item linkages."""

    def __init__(self, base_url: str, auth_token: str):
        self.base_url = base_url
        self.session = requests.Session()
        self.session.headers.update({
            "Authorization": f"Bearer {auth_token}"
        })
        self.results = {
            "uploaded": [],
            "failed": [],
            "linked": [],
            "errors": [],
        }

    def test_connection(self) -> bool:
        """Test API connection and authentication."""
        try:
            # Try to get score groups tree to verify auth
            response = self.session.get(f"{self.base_url}/score-groups/tree", timeout=10)
            if response.status_code == 200:
                print("✓ API connection successful")
                return True
            elif response.status_code == 401:
                print("✗ Authentication failed - invalid token")
                return False
            else:
                print(f"✗ API returned status {response.status_code}")
                return False
        except requests.exceptions.RequestException as e:
            print(f"✗ Connection error: {e}")
            return False

    def upload_article(self, article: Dict) -> Optional[str]:
        """Upload a single article PDF and return the article ID."""
        filepath = UPLOAD_DIR / article["filename"]

        if not filepath.exists():
            error_msg = f"File not found: {filepath}"
            print(f"  ✗ {error_msg}")
            self.results["errors"].append(error_msg)
            return None

        try:
            print(f"\n📤 Uploading: {article['filename']}")
            print(f"   Size: {filepath.stat().st_size / 1024:.1f} KB")

            with open(filepath, 'rb') as f:
                files = {'file': (article['filename'], f, 'application/pdf')}
                data = {
                    'title': article['title'],
                    'description': article.get('description', ''),
                }

                response = self.session.post(
                    f"{self.base_url}/articles/upload",
                    files=files,
                    data=data,
                    timeout=60
                )

            if response.status_code in (200, 201):
                result = response.json()
                article_id = result.get('id') or result.get('data', {}).get('id')

                if article_id:
                    print(f"   ✓ Uploaded successfully - ID: {article_id}")
                    self.results["uploaded"].append({
                        "filename": article["filename"],
                        "article_id": article_id,
                        "title": article["title"],
                    })
                    return article_id
                else:
                    error_msg = f"No ID in response for {article['filename']}"
                    print(f"   ✗ {error_msg}")
                    self.results["errors"].append(error_msg)
                    return None
            else:
                error_msg = f"Upload failed for {article['filename']}: {response.status_code} - {response.text}"
                print(f"   ✗ {error_msg}")
                self.results["failed"].append(article["filename"])
                self.results["errors"].append(error_msg)
                return None

        except Exception as e:
            error_msg = f"Exception uploading {article['filename']}: {str(e)}"
            print(f"   ✗ {error_msg}")
            self.results["failed"].append(article["filename"])
            self.results["errors"].append(error_msg)
            return None

    def search_score_items(self, search_terms: List[str]) -> List[str]:
        """Search for score items matching the given terms."""
        try:
            # Get all score groups with their tree structure
            response = self.session.get(
                f"{self.base_url}/score-groups/tree",
                timeout=30
            )

            if response.status_code != 200:
                print(f"   ⚠ Failed to fetch score groups tree: {response.status_code}")
                return []

            data = response.json()
            groups = data.get('data', []) if isinstance(data, dict) else data

            # Recursively search through the tree structure
            matched_ids = []

            def search_in_tree(node, search_terms):
                """Recursively search through score groups tree."""
                # Check subgroups
                if 'subgroups' in node:
                    for subgroup in node['subgroups']:
                        # Check items in subgroup
                        if 'items' in subgroup:
                            for item in subgroup['items']:
                                item_name = item.get('name', '').lower()
                                item_desc = item.get('description', '').lower()

                                for term in search_terms:
                                    if term.lower() in item_name or term.lower() in item_desc:
                                        item_id = item.get('id')
                                        if item_id and item_id not in matched_ids:
                                            matched_ids.append(item_id)
                                            print(f"   → Found: {item.get('name')} (ID: {item_id})")
                                            break  # Found match, move to next item

                        # Recurse into nested subgroups if any
                        search_in_tree(subgroup, search_terms)

            # Search through all groups
            for group in groups:
                search_in_tree(group, search_terms)

            return matched_ids

        except Exception as e:
            print(f"   ⚠ Error searching score items: {e}")
            return []

    def link_article_to_items(self, article_id: str, score_item_ids: List[str], article_title: str) -> bool:
        """Link an article to score items."""
        if not score_item_ids:
            print(f"   ⚠ No score items to link")
            return False

        try:
            print(f"\n🔗 Linking article to {len(score_item_ids)} score item(s)")

            response = self.session.post(
                f"{self.base_url}/articles/{article_id}/score-items",
                json={"scoreItemIds": score_item_ids},
                timeout=30
            )

            if response.status_code in (200, 201):
                print(f"   ✓ Successfully linked {len(score_item_ids)} items")
                self.results["linked"].append({
                    "article_id": article_id,
                    "article_title": article_title,
                    "score_item_count": len(score_item_ids),
                    "score_item_ids": score_item_ids,
                })
                return True
            else:
                error_msg = f"Link failed for article {article_id}: {response.status_code} - {response.text}"
                print(f"   ✗ {error_msg}")
                self.results["errors"].append(error_msg)
                return False

        except Exception as e:
            error_msg = f"Exception linking article {article_id}: {str(e)}"
            print(f"   ✗ {error_msg}")
            self.results["errors"].append(error_msg)
            return False

    def process_all_articles(self):
        """Main workflow: upload all articles and create linkages."""
        print("=" * 80)
        print("PLENYA EMR - ARTICLE UPLOAD SCRIPT")
        print("=" * 80)

        if not self.test_connection():
            print("\n❌ Cannot proceed - API connection failed")
            return False

        print(f"\n📋 Processing {len(ARTICLES)} articles...")

        for article in ARTICLES:
            print("\n" + "-" * 80)

            # Step 1: Upload article
            article_id = self.upload_article(article)
            if not article_id:
                continue

            # Step 2: Search for related score items
            print(f"\n🔍 Searching for related score items...")
            score_item_ids = self.search_score_items(article.get('search_terms', []))

            # Step 3: Link article to score items
            if score_item_ids:
                self.link_article_to_items(article_id, score_item_ids, article['title'])
            else:
                print(f"   ⚠ No matching score items found")

        return True

    def generate_report(self, output_path: str):
        """Generate a detailed markdown report."""
        report = []
        report.append("# Plenya EMR - Article Upload Report")
        report.append(f"\n**Generated:** {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
        report.append(f"\n**Base URL:** {self.base_url}")

        # Summary
        report.append("\n## Summary\n")
        report.append(f"- **Total articles processed:** {len(ARTICLES)}")
        report.append(f"- **Successfully uploaded:** {len(self.results['uploaded'])}")
        report.append(f"- **Failed uploads:** {len(self.results['failed'])}")
        report.append(f"- **Articles with linkages:** {len(self.results['linked'])}")
        report.append(f"- **Total errors:** {len(self.results['errors'])}")

        # Uploaded Articles
        if self.results['uploaded']:
            report.append("\n## Successfully Uploaded Articles\n")
            for item in self.results['uploaded']:
                report.append(f"### {item['title']}")
                report.append(f"- **File:** `{item['filename']}`")
                report.append(f"- **Article ID:** `{item['article_id']}`")
                report.append("")

        # Linkages
        if self.results['linked']:
            report.append("\n## Score Item Linkages\n")
            for link in self.results['linked']:
                report.append(f"### {link['article_title']}")
                report.append(f"- **Article ID:** `{link['article_id']}`")
                report.append(f"- **Linked Items:** {link['score_item_count']}")
                report.append(f"- **Score Item IDs:**")
                for sid in link['score_item_ids']:
                    report.append(f"  - `{sid}`")
                report.append("")

        # Failed Uploads
        if self.results['failed']:
            report.append("\n## Failed Uploads\n")
            for filename in self.results['failed']:
                report.append(f"- `{filename}`")
            report.append("")

        # Errors
        if self.results['errors']:
            report.append("\n## Errors Encountered\n")
            report.append("```")
            for error in self.results['errors']:
                report.append(error)
            report.append("```")
            report.append("")

        # Statistics
        report.append("\n## Statistics\n")
        total_links = sum(link['score_item_count'] for link in self.results['linked'])
        report.append(f"- **Total score item linkages created:** {total_links}")

        if self.results['linked']:
            avg_links = total_links / len(self.results['linked'])
            report.append(f"- **Average links per article:** {avg_links:.1f}")

        report.append(f"\n---\n*Generated by Plenya EMR Article Upload Script*")

        # Write report
        report_content = "\n".join(report)
        with open(output_path, 'w') as f:
            f.write(report_content)

        print(f"\n📄 Report saved to: {output_path}")
        return report_content


def main():
    """Main entry point."""
    uploader = ArticleUploader(BASE_URL, AUTH_TOKEN)

    success = uploader.process_all_articles()

    # Generate report
    report_path = "/home/user/plenya/ARTICLE_UPLOAD_REPORT.md"
    uploader.generate_report(report_path)

    print("\n" + "=" * 80)
    if success:
        print("✅ UPLOAD PROCESS COMPLETED")
    else:
        print("❌ UPLOAD PROCESS COMPLETED WITH ERRORS")
    print("=" * 80)

    return 0 if success else 1


if __name__ == "__main__":
    sys.exit(main())
