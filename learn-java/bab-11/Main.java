import java.util.Locale;

public class Main {
    enum Status { BARU, LUNAS, DIKIRIM, BATAL }
    record Hasil(Status status, boolean ok) {}

    static Hasil transisi(Status awal, String aksi) {
        // Status terminal tidak memiliki transisi keluar.
        return switch (awal) {
            case BARU -> switch (aksi) {
                case "bayar" -> new Hasil(Status.LUNAS, true);
                case "batal" -> new Hasil(Status.BATAL, true);
                default -> new Hasil(awal, false);
            };
            case LUNAS -> aksi.equals("kirim")
                ? new Hasil(Status.DIKIRIM, true) : new Hasil(awal, false);
            case DIKIRIM, BATAL -> new Hasil(awal, false);
        };
    }

    public static void main(String[] args) {
        Status status = Status.BARU;
        for (String aksi : new String[]{"kirim", "bayar", "batal", "kirim", "bayar"}) {
            var hasil = transisi(status, aksi);
            System.out.printf("%s %s -> %s ok=%b%n", status.name().toLowerCase(Locale.ROOT),
                aksi, hasil.status().name().toLowerCase(Locale.ROOT), hasil.ok());
            status = hasil.status();
        }
    }
}
