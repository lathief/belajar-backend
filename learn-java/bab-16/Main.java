public class Main {
    static String kirim(int[] respons, int deadline) {
        int waktu = 0;
        int jeda = 1;
        for (int i = 0; i < respons.length; i++) {
            int kode = respons[i];
            System.out.printf("t=%d percobaan=%d kode=%d%n", waktu, i + 1, kode);
            if (kode == 200) return "sukses";
            if (kode != 503) return "permanen";
            if (i + 1 == respons.length) return "habis";
            // Uji batas sebelum memajukan jam simulasi.
            if (waktu + jeda >= deadline) return "deadline";
            waktu += jeda;
            jeda *= 2;
        }
        return "habis";
    }

    public static void main(String[] args) {
        System.out.println("hasil=" + kirim(new int[]{503, 503, 200}, 4));
        System.out.println("hasil=" + kirim(new int[]{503, 503, 200}, 3));
        System.out.println("hasil=" + kirim(new int[]{400, 200}, 4));
    }
}
