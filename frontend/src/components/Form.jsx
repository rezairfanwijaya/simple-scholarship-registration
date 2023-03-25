import { Label, TextInput, Button, Select, FileInput } from "flowbite-react";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import Swal from 'sweetalert2'
import withReactContent from 'sweetalert2-react-content'

const FormDaftar = () => {
    // state
    const [nama, setNama] = useState('')
    const [email, setEmail] = useState('')
    const [telepon, setTelepon] = useState('')
    const [semester, setSemester] = useState(1)
    const [beasiswa, setBeasiswa] = useState("Beasiswa Akademik")
    const [ipk, setIpk] = useState(3)
    const [berkas, setBerkas] = useState(null)
    const [emailNotFound, setEmailNotFound] = useState(false)
    const [isEligible, setIsEligible] = useState(true)

    const MySwal = withReactContent(Swal)
    const navigate = useNavigate();

    // hanlder berkas
    const handleChangeBerkas = (e) => {
        let file = e.target.files[0]
        setBerkas(file)
    }

    // get mahasiwa by email
    useEffect(() => {
        fetch(`http://localhost:8989/mahasiswa/${email}`, {
            method: 'GET'
        }).then(res => { return res.json() })
            .then(data => {
                if (data.meta.code === 404) {
                    setEmailNotFound(true)
                } else {
                    setSemester(data.data.semester)
                    setIpk(data.data.ipk)
                    setEmailNotFound(false)

                }
            })
    }, [email])


    // filtering ipk
    useEffect(() => {
        if (ipk < 3) {
            setIsEligible(false)
        } else {
            setIsEligible(true)
        }
    }, [ipk])


    // submit
    const handleDaftarBeasiswa = (e) => {
        e.preventDefault()

        // initial form data to send data
        let formData = new FormData()
        formData.append("nama", nama)
        formData.append("email", email)
        formData.append("telepon", telepon)
        formData.append("semester", semester)
        formData.append("ipk", ipk)
        formData.append("berkas", berkas)
        formData.append("beasiswa", beasiswa)

        fetch("http://localhost:8989/mahasiswa/daftar", {
            method : "POST",
            body : formData,
        })
        .then((res)=>{return res.json()})
        .then((data) => {
            if (data.meta.code === 200) {
                let timerInterval
                MySwal.fire({
                    title: 'Sukses daftar beasiswa',
                    icon: 'success',
                    timer: 2000,
                    timerProgressBar: false,
                    didOpen: () => {
                        Swal.showLoading()
                    },
                    willClose: () => {
                        clearInterval(timerInterval)
                    }
                }).then((result) => {
                    navigate('/hasil')
                })
            }else{
                alert('Error')
            }
        })    
    }

    return (<>

        <div className="body flex flex-col gap-4 bg-white p-5 rounded-md">
            <form className="flex flex-col gap-4 bg-white rounded-md" onSubmit={handleDaftarBeasiswa}>
                <div>
                    <div className="mb-2 block">
                        <Label
                            htmlFor="nama"
                            value="Masukan Nama"
                        />
                    </div>
                    <TextInput
                        id="nama"
                        type="text"
                        placeholder="Anggara"
                        required={true}
                        onChange={e => setNama(e.target.value)}
                    />
                </div>
                <div>
                    <div className="mb-2 block">
                        <Label
                            htmlFor="email"
                            value="Masukan Email"
                        />
                    </div>
                    <TextInput
                        id="email"
                        type="email"
                        placeholder="anggara@gmail.com"
                        required={true}
                        onChange={e => setEmail(e.target.value)}
                    />
                    {emailNotFound && <span className="text-red-700">email tidak terdaftar</span>}
                </div>
                <div>
                    <div className="mb-2 block">
                        <Label
                            htmlFor="hp"
                            value="Nomor HP"
                        />
                    </div>
                    <TextInput
                        id="hp"
                        type="number"
                        placeholder="087656787654"
                        required={true}
                        onChange={e => setTelepon(e.target.value)}
                        disabled={emailNotFound}
                    />
                </div>
                <div id="select">
                    <div className="mb-2 block">
                        <Label
                            htmlFor="semester"
                            value="Semester Saat Ini"
                        />
                    </div>
                    <Select
                        id="semester"
                        required={true}
                        onChange={e => setSemester(e.target.value)}
                        disabled={emailNotFound}
                    >
                        <option>
                            1
                        </option>
                        <option>
                            2
                        </option>
                        <option>
                            3
                        </option>
                        <option>
                            4
                        </option>
                        <option>
                            5
                        </option>
                        <option>
                            6
                        </option>
                        <option>
                            7
                        </option>
                        <option>
                            8
                        </option>
                    </Select>
                </div>

                <div>
                    <div className="mb-2 block">
                        <Label
                            htmlFor="ipk"
                            value="IPK Terakhir"
                        />
                    </div>
                    <TextInput
                        id="ipk"
                        type="number"
                        placeholder="087656787654"
                        value={ipk}
                        readOnly={true}
                        required={true}
                        disabled={emailNotFound}
                    />
                    {!isEligible && <span className="text-red-700">ipk minimal 3</span>}
                </div>

                <div id="select">
                    <div className="mb-2 block">
                        <Label
                            htmlFor="beasiswa"
                            value="Pilihan Beasiswa"
                        />
                    </div>
                    <Select
                        id="beasiswa"
                        required={true}
                        onChange={e => setBeasiswa(e.target.value)}
                        disabled={emailNotFound || !isEligible}
                    >
                        <option>
                            Beasiswa Akademik
                        </option>
                        <option>
                            Beasiswa Non Akademik
                        </option>

                    </Select>
                </div>

                <div id="fileUpload">
                    <div className="mb-2 block">
                        <Label
                            htmlFor="file"
                            value="Upload Berkas Syarat"
                        />
                    </div>
                    <FileInput
                        id="file"
                        helperText="Hanya format PNG yang diperbolehkan"
                        onChange={e => handleChangeBerkas(e)}
                        disabled={emailNotFound || !isEligible}
                    />
                </div>


                <Button type="submit" disabled={emailNotFound || !isEligible}>
                    Daftar
                </Button>
            </form>

            <div className="back border w-full">
                <a href="/" >
                    <button className="text-sm border border-primary p-2 rounded-md text-mute w-full">
                        Kembali
                    </button>
                </a>
            </div>
        </div>
    </>
    );
}

export default FormDaftar;