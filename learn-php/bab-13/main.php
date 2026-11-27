<?php

$stok = 10;
$urutanBerikutnya = 1;
$diproses = [];

function terapkan(array $event, int &$stok, int &$urutanBerikutnya, array &$diproses): string
{
    if (isset($diproses[$event["id"]])) {
        return "duplikat=" . $event["id"];
    }
    if ($event["urutan"] !== $urutanBerikutnya) {
        return "urutan_terlewat=" . $event["urutan"] . " harapan={$urutanBerikutnya}";
    }

    $stok += $event["perubahan"];
    $diproses[$event["id"]] = true;
    $urutanBerikutnya++;
    return "diterapkan";
}

$eventLog = [
    ["id" => "E1", "urutan" => 1, "perubahan" => -2],
    ["id" => "E2", "urutan" => 2, "perubahan" => -1],
    ["id" => "E2", "urutan" => 2, "perubahan" => -1],
    ["id" => "E4", "urutan" => 4, "perubahan" => 3],
];
$hasil = [];
foreach ($eventLog as $event) {
    $hasil[] = terapkan($event, $stok, $urutanBerikutnya, $diproses);
}

echo "stok={$stok} diterapkan=" . count($diproses) . "\n";
echo $hasil[2] . "\n";
echo $hasil[3] . "\n";
