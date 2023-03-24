import { Label, TextInput, Button, Select, FileInput } from "flowbite-react";

const FormDaftar = () => {


    return (<>

        <div className="body flex flex-col gap-4 bg-white p-5 rounded-md">
            <form className="flex flex-col gap-4 bg-white rounded-md">
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
                    />
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
                        value={3}
                        readOnly={true}
                        required={true}
                    />
                </div>

                <div id="select">
                    <div className="mb-2 block">
                        <Label
                            htmlFor="semester"
                            value="Pilihan Beasiswa"
                        />
                    </div>
                    <Select
                        id="semester"
                        required={true}
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
                        helperText="Hanya format JPG atau PNG yang diperbolehkan"
                    />
                </div>


                <Button type="submit">
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