<?php

final class TokenBucket
{
    private float $token;
    private float $waktuTerakhir = 0.0;

    public function __construct(private float $kapasitas, private float $perDetik)
    {
        $this->token = $kapasitas;
    }

    public function minta(float $waktu, float $biaya = 1.0): bool
    {
        $berlalu = max(0.0, $waktu - $this->waktuTerakhir);
        $this->token = min($this->kapasitas, $this->token + $berlalu * $this->perDetik);
        $this->waktuTerakhir = max($this->waktuTerakhir, $waktu);
        if ($biaya > $this->token) {
            return false;
        }
        $this->token -= $biaya;
        return true;
    }

    public function token(): int
    {
        return (int) floor($this->token);
    }
}

$bucket = new TokenBucket(2, 1);
foreach ([0, 0, 0, 2] as $waktu) {
    echo $bucket->minta($waktu) ? "diizinkan\n" : "ditolak\n";
}
echo "token=" . $bucket->token() . "\n";
