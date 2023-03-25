package mahasiswa

type formatMahasiswa struct {
	Nama      string  `json:"nama"`
	Email     string  `json:"email"`
	Telepon   string  `json:"telepon"`
	Semester  int     `json:"semester"`
	IPK       float64 `json:"ipk"`
	BerkasURL string  `json:"berkas_url"`
}

func FormatterMahasiswa(mahasiswa Mahasiswa) *formatMahasiswa {
	return &formatMahasiswa{
		Nama:      mahasiswa.Nama,
		Email:     mahasiswa.Email,
		Telepon:   mahasiswa.Telepon,
		Semester:  mahasiswa.Semester,
		IPK:       mahasiswa.IPK,
		BerkasURL: mahasiswa.BerkasURL,
	}
}
