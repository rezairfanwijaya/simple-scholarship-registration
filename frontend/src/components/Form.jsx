import { Label, TextInput, Button, Select } from "flowbite-react";

const FormDaftar = () => {
    return (<>

        <form className="flex flex-col gap-4 bg-white p-5 rounded-md">
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

            <Button type="submit">
                Daftar
            </Button>
        </form>

    </>
    );
}

export default FormDaftar;