public class Main {
    record Peristiwa(int nomor, int delta) {}

    static class Ringkasan {
        int posisi;
        int stok;

        String terapkan(Peristiwa e) {
            if (e.nomor() <= posisi) return "duplikat";
            if (e.nomor() != posisi + 1) return "celah";
            // Kedua field hanya bergerak ketika peristiwa berikutnya lengkap.
            stok += e.delta();
            posisi = e.nomor();
            return "diterapkan";
        }
    }

    public static void main(String[] args) {
        var r = new Ringkasan();
        var log = new Peristiwa[]{new Peristiwa(1, 5), new Peristiwa(3, -2),
            new Peristiwa(2, -1), new Peristiwa(2, -1), new Peristiwa(3, -2)};
        for (var e : log) {
            String hasil = r.terapkan(e);
            System.out.printf("%d %s posisi=%d stok=%d%n", e.nomor(), hasil, r.posisi, r.stok);
        }
    }
}
