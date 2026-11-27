import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Main {
    public static void main(String[] args) {
        List<String> kode = List.of("PS-1", "PS-2", "PS-1");
        Map<String, Integer> hitungan = new HashMap<>();
        for (String k : kode) {
            // Pada kemunculan pertama, hitungan sebelumnya nol.
            hitungan.put(k, hitungan.getOrDefault(k, 0) + 1);
        }
        System.out.println("PS-1=" + hitungan.get("PS-1"));
    }
}
