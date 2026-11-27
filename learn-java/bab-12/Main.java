public class Main {
    static class Alamat {
        String kota = "Bandung";
        int versi = 1;

        String ubah(int dibaca, String tujuan) {
            if (versi != dibaca) return "konflik";
            // Versi bertambah satu kali untuk setiap penulisan sukses.
            kota = tujuan;
            versi++;
            return "tersimpan";
        }
    }

    static void lapor(Alamat a, String pelaku, int versi, String kota) {
        String hasil = a.ubah(versi, kota);
        System.out.printf("%s %s kota=%s versi=%d%n", pelaku, hasil, a.kota, a.versi);
    }

    public static void main(String[] args) {
        var a = new Alamat();
        int bacaA = a.versi;
        int bacaB = a.versi;
        lapor(a, "A", bacaA, "Solo");
        lapor(a, "B", bacaB, "Bogor");
        lapor(a, "B", a.versi, "Bogor"); // Setelah B membaca ulang dan mengonfirmasi.
    }
}
