import { Label, TextInput, Button } from "flowbite-react";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import Swal from 'sweetalert2'
import withReactContent from 'sweetalert2-react-content'

const FormLogin = () => {
    // state
    const [email, setemail] = useState('')
    const [password, setpassword] = useState('')
    const MySwal = withReactContent(Swal)
    const navigate = useNavigate();

    const handleLogin = (e) => {
        e.preventDefault()
        const data = { email, password }

        try {
            fetch("http://localhost:8989/admin/login", {
                method: 'POST',
                body: JSON.stringify(data)
            })
                .then(res => { return res.json() })
                .then(data => {
                    if (data.meta.code === 404) {
                        MySwal.fire({
                            icon: 'error',
                            title: 'Login Gagal',
                            text: data.data,
                        })
                    } else if (data.meta.code === 200) {
                        // save jwt to local storage
                        localStorage.setItem('token', data.data.token)

                        let timerInterval
                        MySwal.fire({
                            title: 'Login Berhasil',
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
                            navigate('/dashboard')
                        })
                    } else {
                        MySwal.fire({
                            icon: 'error',
                            title: 'Gagal Login',
                            text: data.data,
                        })
                    }
                })
        } catch (error) {
            MySwal.fire({
                icon: 'error',
                title: 'Login Failed',
                text: error,
            })
        }
    }



    return (<>
        <div className="content">
            <form className="flex flex-col gap-4 bg-white p-12 rounded-md" onSubmit={handleLogin}>
                <div>
                    <div className="mb-2 block">
                        <Label
                            htmlFor="email"
                            value="Email"
                        />
                    </div>
                    <TextInput
                        id="email"
                        type="email"
                        placeholder="agus@gmail.com"
                        required={true}
                        onChange={e => setemail(e.target.value)}
                    />
                </div>
                <div>
                    <div className="mb-2 block">
                        <Label
                            htmlFor="password"
                            value="Password"
                        />
                    </div>
                    <TextInput
                        id="password"
                        type="password"
                        required={true}
                        onChange={e => setpassword(e.target.value)}
                    />
                </div>

                <Button type="submit">
                    Submit
                </Button>
                <a href="/">
                    <div className="border border-primary p-2 text-sm rounded-md text-center">
                        Kembali
                    </div>
                </a>
            </form>
        </div>
    </>);
}

export default FormLogin;