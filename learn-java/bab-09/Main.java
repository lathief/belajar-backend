public class Main {
    static int jumlahValid(String teks) {
        try {
            int n = Integer.parseInt(teks);
            return n >= 1 && n <= 5 ? n : 0;
        } catch (NumberFormatException e) {
            return 0;
        }
    }

    public static void main(String[] args) {
        long unit = 0;
        int ditolak = 0;
        for (String teks : args) {
            // Nilai penanda tidak ikut menambah total unit.
            int n = jumlahValid(teks);
            if (n == 0) ditolak++;
            else unit += n;
        }
        System.out.println("unit=" + unit + " ditolak=" + ditolak);
    }
}
