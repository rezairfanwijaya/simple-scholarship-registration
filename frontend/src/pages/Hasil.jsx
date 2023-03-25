import React, { useState, useEffect } from 'react';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import "primereact/resources/themes/lara-light-indigo/theme.css";
import "primereact/resources/primereact.min.css";
import "primeicons/primeicons.css";
import Swal from 'sweetalert2'
import withReactContent from 'sweetalert2-react-content'

const Hasil = () => {
    const [approvals, setapprovals] = useState([])
    const MySwal = withReactContent(Swal)

    useEffect(() => {
        fetch(`http://localhost:8989/mahasiswa/approval`, {
            method: 'GET'
        }).then(res => { return res.json() })
            .then(data => {
                if (data.meta.code === 200) {
                    setapprovals(data.data)
                } else {
                    MySwal.fire({
                        title: 'Terjadi kesalahan',
                        icon: 'error',
                        text: data.data
                    })
                }
            })
    }, [])

    const imageBodyTemplate = (approval) => {
        return <img src={approval.berkas_url} alt="berkas" className="w-56 shadow-2 border-round" />;
    };

    const statusBodyTemplate = (approval) => {
        if (approval.status === 0) {
            return <div className="p-2 bg-secondary rounded-full text-center text-sm">Belum di verifikasi</div>
        } else {
            return <div className="p-2 bg-primary text-white rounded-full text-center">Approved</div>
        }

    };

    return (<>
        <div className="content md:px-24 md:py-10">

            <a href="/" className='flex mb-5'>
                <div className='bg-primary py-2 px-5 text-center text-white rounded-md'>
                    Beranda
                </div>
            </a>
            <div className="card ">
                <DataTable value={approvals} tableStyle={{ minWidth: '50rem' }}>
                    <Column field="nama" sortable header="Nama"></Column>
                    <Column field="email" sortable header="Email"></Column>
                    <Column field="ipk" sortable header="IPK"></Column>
                    <Column field="semester" sortable header="Semester"></Column>
                    <Column field="telepon" sortable header="Telepon"></Column>
                    <Column field="jenis_beasiswa" sortable header="Beasiswa"></Column>
                    <Column header="Berkas" body={imageBodyTemplate}></Column>
                    <Column header="Status" body={statusBodyTemplate}></Column>
                </DataTable>
            </div>

        </div>
    </>);
}

export default Hasil;