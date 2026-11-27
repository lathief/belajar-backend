<?php

function simulasi(bool $responsHilang, bool $relayPutus, bool $ackHilang): array
{
    $produsen = ["stok" => 5, "dipesan" => false, "tertunda" => false];
    $pesan = static function () use (&$produsen): void {
        if (!$produsen["dipesan"]) {
            $produsen = [
                "stok" => $produsen["stok"] - 1,
                "dipesan" => true,
                "tertunda" => true,
            ];
        }
    };

    $pesan();
    if ($responsHilang) {
        $pesan();
    }

    $antrean = [];
    $publikasi = 0;
    while ($produsen["tertunda"]) {
        $antrean[] = "A";
        $publikasi++;
        if ($relayPutus) {
            $relayPutus = false;
            continue;
        }
        $produsen["tertunda"] = false;
    }

    $inbox = [];
    $paket = 0;
    while ($antrean !== []) {
        $id = $antrean[0];
        if (!isset($inbox[$id])) {
            $inbox[$id] = true;
            $paket++;
        }
        if ($ackHilang) {
            $ackHilang = false;
            continue;
        }
        array_shift($antrean);
    }

    return ["stok" => $produsen["stok"], "publikasi" => $publikasi, "paket" => $paket];
}

$lulus = 0;
foreach ([false, true] as $respons) {
    foreach ([false, true] as $relay) {
        foreach ([false, true] as $ack) {
            $hasil = simulasi($respons, $relay, $ack);
            $publikasiHarapan = $relay ? 2 : 1;
            if ($hasil["stok"] !== 4
                || $hasil["publikasi"] !== $publikasiHarapan
                || $hasil["paket"] !== 1) {
                throw new LogicException("invarian_gagal");
            }
            $lulus++;
            if ($respons && $relay && $ack) {
                echo "stok=" . $hasil["stok"]
                    . " publikasi=" . $hasil["publikasi"]
                    . " paket=" . $hasil["paket"] . "\n";
            }
        }
    }
}

echo "skenario_lulus={$lulus}\n";
