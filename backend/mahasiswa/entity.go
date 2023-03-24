package mahasiswa

type Mahasiswa struct {
	ID        int     `json:"id" gomr:"primaryKey"`
	Nama      string  `json:"nama"`
	Email     string  `json:"email"`
	Telepon   string  `json:"telepon"`
	Semester  int     `json:"semester"`
	IPK       float64 `json:"ipk"`
	BerkasURL string  `json:"berkas_url"`
}
