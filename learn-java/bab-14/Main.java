import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

public class Main {
    static List<Integer> halaman(List<Integer> log, int setelah, int batas, int ukuran) {
        return log.stream().dropWhile(id -> id <= setelah)
            .takeWhile(id -> id <= batas).limit(ukuran).toList();
    }

    static String tampil(List<Integer> ids) {
        return ids.stream().map(id -> id.toString()).collect(Collectors.joining(","));
    }

    public static void main(String[] args) {
        var log = new ArrayList<>(List.of(101, 104, 107, 112));
        int batas = log.getLast();
        int cursor = 0;
        while (true) {
            var ids = halaman(log, cursor, batas, 2);
            if (ids.isEmpty()) break;
            System.out.println("halaman=" + tampil(ids));
            if (cursor == 0) log.add(115); // Tambahan ini tersedia pada sesi berikutnya.
            cursor = ids.getLast();
        }
        System.out.println("susulan=" + tampil(halaman(log, cursor, log.getLast(), 2)));
    }
}
