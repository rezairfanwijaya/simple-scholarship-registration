import Card from "../components/Card";
import HeroSection from "../components/HeroSection";
import NavbarUser from "../components/Navbar";
import beasiswaAkademik from "../assets/image/beasiswaAkademik.png"
import beasiswaNonAkademik from "../assets/image/beasiswaNonAkademik.png"

const LandingPage = () => {
    return (<>
        <div className="body px-5 mb-10">
            <NavbarUser />
            <HeroSection />

            <section id="list-beasiswa" className="mt-20">
                <div className="card flex flex-col gap-5 md:flex-row lg:px-44">
                    <Card
                        image={beasiswaAkademik}
                        title="Beasiswa Akademik"
                        idModal="akademik"
                        detail={
                            <div className="p-6 space-y-6">
                                <p className="text-base leading-relaxed text-gray-500 dark:text-gray-400">
                                    Ketentuan status mahasiswa yang bisa mengikuti program beasiswa ini adalah sebagai berikut :
                                    <br />
                                    <br />
                                    <ol type="1">
                                        <li>
                                            1. Calon penerima Beasiswa PPA adalah mahasiswa yang kuliah di perguruan tinggi Negeri;
                                        </li>
                                        <br />
                                        <li>
                                            2. Calon penerima Beasiswa PPA wajib terdaftar pada Pangkalan Data Pendidikan Tinggi (PD-Dikti);
                                        </li>
                                        <br />
                                        <li>
                                            3. Calon penerima Beasiswa PPA adalah mahasiswa aktif dalam jenjang pendidikan Sarjana atau Diploma.
                                        </li>
                                    </ol>
                                </p>
                            </div>
                        }
                    />
                    <Card
                        image={beasiswaNonAkademik}
                        title="Beasiswa Non Akademik"
                        idModal="non"
                        detail={
                            <div className="p-6 space-y-6">
                                <p className="text-base leading-relaxed text-gray-500 dark:text-gray-400">
                                    Ketentuan status mahasiswa yang bisa mengikuti program beasiswa ini adalah sebagai berikut :
                                    <br />
                                    <br />
                                    <ol type="1">
                                        <li>
                                            1. Calon penerima Beasiswa PPA adalah mahasiswa yang kuliah di perguruan tinggi Negeri;
                                        </li>
                                        <br />
                                        <li>
                                            2. Calon penerima Beasiswa PPA wajib terdaftar pada Pangkalan Data Pendidikan Tinggi (PD-Dikti);
                                        </li>
                                        <br />
                                        <li>
                                            3. Calon penerima Beasiswa PPA adalah mahasiswa aktif dalam jenjang pendidikan Sarjana atau Diploma.
                                        </li>
                                    </ol>
                                </p>
                            </div>
                        }
                    />
                </div>
            </section>

        </div >
    </>);
}

export default LandingPage;