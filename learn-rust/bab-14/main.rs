fn halaman(log: &[u32], setelah: u32, batas: u32, ukuran: usize) -> Vec<u32> {
    log.iter().copied()
        .skip_while(|&id| id <= setelah)
        .take_while(|&id| id <= batas)
        .take(ukuran)
        .collect()
}

fn tampil(ids: &[u32]) -> String {
    ids.iter().map(u32::to_string).collect::<Vec<_>>().join(",")
}

fn main() {
    let mut log = vec![101, 104, 107, 112];
    let batas = log.last().copied().unwrap_or(0);
    let mut cursor = 0;
    loop {
        let ids = halaman(&log, cursor, batas, 2);
        let Some(&terakhir) = ids.last() else { break };
        println!("halaman={}", tampil(&ids));
        if cursor == 0 { log.push(115); } // ID baru masuk setelah batas ditetapkan.
        cursor = terakhir;
    }
    let batas_baru = log.last().copied().unwrap_or(0);
    println!("susulan={}", tampil(&halaman(&log, cursor, batas_baru, 2)));
}
