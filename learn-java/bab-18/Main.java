import java.util.ArrayDeque;
import java.util.HashSet;

public class Main {
    record Produsen(int stok, boolean dipesan, boolean tertunda) {
        Produsen pesan() {
            // Record baru mewakili commit yang mencatat stok dan outbox bersama.
            return dipesan ? this : new Produsen(stok - 1, true, true);
        }
    }
    record Hasil(int stok, int publikasi, int paket) {}

    static Hasil simulasi(boolean responsHilang, boolean relayPutus, boolean ackHilang) {
        var p = new Produsen(5, false, false).pesan();
        if (responsHilang) p = p.pesan();

        var antrean = new ArrayDeque<String>();
        int publikasi = 0;
        while (p.tertunda()) {
            antrean.addLast("A");
            publikasi++;
            if (relayPutus) {
                relayPutus = false;
                continue;
            }
            p = new Produsen(p.stok(), p.dipesan(), false);
        }

        var inbox = new HashSet<String>();
        int paket = 0;
        while (!antrean.isEmpty()) {
            if (inbox.add(antrean.getFirst())) paket++;
            if (ackHilang) {
                ackHilang = false;
                continue;
            }
            antrean.removeFirst();
        }
        return new Hasil(p.stok(), publikasi, paket);
    }

    public static void main(String[] args) {
        boolean[] opsi = {false, true};
        int lulus = 0;
        for (boolean respons : opsi) {
            for (boolean relay : opsi) {
                for (boolean ack : opsi) {
                    var hasil = simulasi(respons, relay, ack);
                    var target = new Hasil(4, relay ? 2 : 1, 1);
                    if (!hasil.equals(target)) throw new IllegalStateException("Invarian gagal");
                    lulus++;
                    if (respons && relay && ack) {
                        System.out.printf("stok=%d publikasi=%d paket=%d%n",
                            hasil.stok(), hasil.publikasi(), hasil.paket());
                    }
                }
            }
        }
        System.out.println("skenario_lulus=" + lulus);
    }
}
