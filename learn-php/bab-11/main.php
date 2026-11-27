<?php

final class Pesanan
{
    private string $status = "baru";

    public function bayar(): void
    {
        $this->transisi("baru", "dibayar");
    }

    public function kirim(): void
    {
        $this->transisi("dibayar", "dikirim");
    }

    public function status(): string
    {
        return $this->status;
    }

    private function transisi(string $dari, string $ke): void
    {
        if ($this->status !== $dari) {
            throw new DomainException("pesanan_harus_{$dari}");
        }
        $this->status = $ke;
    }
}

$pesanan = new Pesanan();
$sebelum = $pesanan->status();
$pesanan->bayar();
echo $sebelum . " -> " . $pesanan->status() . "\n";
try {
    $pesananLain = new Pesanan();
    $pesananLain->kirim();
} catch (DomainException $error) {
    echo "ditolak: " . $error->getMessage() . "\n";
}
$sebelum = $pesanan->status();
$pesanan->kirim();
echo $sebelum . " -> " . $pesanan->status() . "\n";
