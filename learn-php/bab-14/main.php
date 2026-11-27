<?php

$log = ["E1", "E2", "E3", "E4"];

function halamanBerikutnya(array $log, ?string $cursor, int $ukuran): array
{
    $mulai = 0;
    if ($cursor !== null) {
        $posisi = array_search($cursor, $log, true);
        if ($posisi === false) {
            throw new InvalidArgumentException("cursor_tidak_ditemukan");
        }
        $mulai = $posisi + 1;
    }
    return array_slice($log, $mulai, $ukuran);
}

$cursor = null;
for ($nomor = 0; $nomor < 3; $nomor++) {
    $halaman = halamanBerikutnya($log, $cursor, 2);
    if ($halaman === []) {
        break;
    }
    $cursor = $halaman[count($halaman) - 1];
    echo "halaman=" . implode(",", $halaman) . " cursor={$cursor}\n";
    if ($nomor === 0) {
        $log[] = "E5";
    }
}
