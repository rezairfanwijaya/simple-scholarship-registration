package mahasiswa

type formatMahasiswa struct {
	Nama      string  `json:"nama"`
	Email     string  `json:"email"`
	Telepon   string  `json:"telepon"`
	Semester  int     `json:"semester"`
	IPK       float64 `json:"ipk"`
	BerkasURL string  `json:"berkas_url"`
}

type registeredMahasiswaOnBeasiswa struct {
	ID            int     `json:"id"`
	Nama          string  `json:"nama"`
	Email         string  `json:"email"`
	Telepon       string  `json:"telepon"`
	Semester      int     `json:"semester"`
	IPK           float64 `json:"ipk"`
	JenisBeasiswa string  `json:"jenis_beasiswa"`
	BerkasURL     string  `json:"berkas_url"`
	Status        int     `json:"status"`
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
