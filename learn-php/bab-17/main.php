<?php

final class AntrianKerja
{
    private int $generasi = 0;
    private ?int $pemilik = null;
    private int $leaseSampai = 0;
    private string $status = "menunggu";

    public function klaim(int $sekarang, int $durasi): int
    {
        if ($this->pemilik !== null && $sekarang < $this->leaseSampai) {
            throw new DomainException("lease_masih_aktif");
        }
        $this->generasi++;
        $this->pemilik = $this->generasi;
        $this->leaseSampai = $sekarang + $durasi;
        return $this->generasi;
    }

    public function selesai(int $token, int $sekarang): bool
    {
        if ($token !== $this->pemilik
            || $sekarang >= $this->leaseSampai
            || $this->status === "selesai") {
            return false;
        }
        $this->status = "selesai";
        return true;
    }

    public function status(): string
    {
        return $this->status;
    }
}

$antrian = new AntrianKerja();
$lama = $antrian->klaim(0, 5);
$baru = $antrian->klaim(6, 5);
echo "lease_lama={$lama} lease_baru={$baru}\n";
echo "hasil_lama=" . ($antrian->selesai($lama, 6) ? "diterima" : "ditolak") . "\n";
echo "hasil_baru=" . ($antrian->selesai($baru, 6) ? "diterima" : "ditolak") . "\n";
echo "status=" . $antrian->status() . "\n";
