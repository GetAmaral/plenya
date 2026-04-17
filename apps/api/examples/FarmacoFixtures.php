<?php

namespace App\DataFixtures;

use App\Entity\Farmaco;
use Doctrine\Bundle\FixturesBundle\Fixture;
use Doctrine\Persistence\ObjectManager;

/**
 * Farmaco Fixtures - Loads pharmaceutical data from ANVISA CSV
 * Based on farmacosAnvisa_2025-07.csv file with complete drug database
 */
class FarmacoFixtures extends Fixture
{
    private const CSV_PATH = '/var/www/html/var/entityCsv/farmacosAnvisa_2025-07.csv';
    private const BATCH_SIZE = 1000;

    public function load(ObjectManager $manager): void
    {
        if (!file_exists(self::CSV_PATH)) {
            throw new \RuntimeException('CSV file not found: ' . self::CSV_PATH);
        }

        $csvFile = fopen(self::CSV_PATH, 'r');
        if (!$csvFile) {
            throw new \RuntimeException('Could not open CSV file: ' . self::CSV_PATH);
        }

        // Skip header row
        fgetcsv($csvFile, 0, ';');

        $count = 0;
        while (($row = fgetcsv($csvFile, 0, ';')) !== false) {
            if (count($row) < 8) {
                continue; // Skip malformed rows
            }

            $farmaco = $this->createFarmacoFromCsvRow($row);
            if ($farmaco) {
                $manager->persist($farmaco);
                $count++;

                // Batch processing for better performance
                if ($count % self::BATCH_SIZE === 0) {
                    $manager->flush();
                    $manager->clear();
                    echo "Processed {$count} farmacos...\n";
                }
            }
        }

        fclose($csvFile);
        $manager->flush();
        $manager->clear();

        echo "Finished loading {$count} farmacos from CSV.\n";
    }

    private function createFarmacoFromCsvRow(array $row): ?Farmaco
    {
        try {
            // Map CSV columns to entity properties
            $substancia = $this->cleanString($row[0] ?? '');
            $laboratorio = $this->cleanString($row[1] ?? '');
            $produto = $this->cleanString($row[2] ?? '');
            $apresentacao = $this->cleanString($row[3] ?? '');
            $classe = $this->cleanString($row[4] ?? '');
            $tipo = $this->cleanString($row[5] ?? '');
            $precoStr = $this->cleanString($row[6] ?? '');
            $tarja = $this->cleanString($row[7] ?? '');

            // Skip if essential fields are empty
            if (empty($substancia) || empty($produto)) {
                return null;
            }

            $farmaco = new Farmaco();
            $farmaco->setSubstancia($substancia);
            $farmaco->setLaboratorio($laboratorio);
            $farmaco->setProduto($produto);
            $farmaco->setApresentacao($apresentacao);
            $farmaco->setClasse($classe);
            $farmaco->setTipo($tipo);
            $farmaco->setTarja($tarja);

            // Parse price - handle Brazilian decimal separator
            $preco = $this->parsePrice($precoStr);
            if ($preco !== null) {
                $farmaco->setPreco($preco);
            }

            // The dose and nome will be automatically set by the PreFlush lifecycle callback

            return $farmaco;
        } catch (\Exception $e) {
            // Log error but continue processing
            error_log("Error processing CSV row: " . $e->getMessage());
            return null;
        }
    }

    private function cleanString(string $value): string
    {
        // Convert from ISO-8859-1 to UTF-8 if needed
        if (!mb_check_encoding($value, 'UTF-8')) {
            $value = mb_convert_encoding($value, 'UTF-8', 'ISO-8859-1');
        }

        // Remove extra quotes and trim
        $value = trim($value, '"');
        $value = trim($value);

        return $value;
    }

    private function parsePrice(string $priceStr): ?float
    {
        if (empty($priceStr) || $priceStr === '-' || strpos($priceStr, '(*)') !== false) {
            return null;
        }

        // Remove any non-numeric characters except comma and period
        $priceStr = preg_replace('/[^\d,.]/', '', $priceStr);
        
        if (empty($priceStr)) {
            return null;
        }

        // Handle Brazilian decimal format (comma as decimal separator)
        if (strpos($priceStr, ',') !== false) {
            // If there's both comma and period, assume comma is decimal separator
            if (strpos($priceStr, '.') !== false) {
                // Remove period (thousands separator) and replace comma with period
                $priceStr = str_replace('.', '', $priceStr);
            }
            $priceStr = str_replace(',', '.', $priceStr);
        }

        $price = (float) $priceStr;
        return $price > 0 ? $price : null;
    }
}