import React, { useState, useEffect } from 'react';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { Button } from 'primereact/button';
import "primereact/resources/themes/lara-light-indigo/theme.css";
import "primereact/resources/primereact.min.css";
import "primeicons/primeicons.css";
import Swal from 'sweetalert2'
import withReactContent from 'sweetalert2-react-content'

const Pengajuan = () => {
    // state
    const [approvals, setapprovals] = useState([])
    
    const MySwal = withReactContent(Swal)
    const APPROVE = 1
    const PENDING = 0

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

    const statusTemplate = (approval) => {
        if (approval.status === 0) {
            return <div className="p-2 bg-secondary rounded-full text-center text-sm">Belum di verifikasi</div>
        } else {
            return <div className="p-2 bg-primary text-white rounded-full text-center text-sm">Approved</div>
        }

    };

    const handleApprove = (data) => {
        const newStatus = APPROVE
        const payload = {
            status : newStatus
        }

        fetch(`http://localhost:8989/admin/status/${data.id}`, {
            method: 'PUT',
            body : JSON.stringify(payload)
        }).then(res => { return res.json() })
            .then(data => {
                if (data.meta.code === 200) {
                    window.location.reload()
                } else {
                    MySwal.fire({
                        title: 'Periksa Tindakan Anda',
                        icon: 'warning',
                        text: "Status Telah di Approve, silahkan pilih tindakan lain"
                    })
                }
            })
    }

    const handlePending = (data) => {
        const newStatus = PENDING
        const payload = {
            status : newStatus
        }

        fetch(`http://localhost:8989/admin/status/${data.id}`, {
            method: 'PUT',
            body : JSON.stringify(payload)
        }).then(res => { return res.json() })
            .then(data => {
                if (data.meta.code === 200) {
                    window.location.reload()
                } else {
                    MySwal.fire({
                        title: 'Periksa Tindakan Anda',
                        icon: 'warning',
                        text: "Status Telah di Unapprove, silahkan pilih tindakan lain"
                    })
                }
            })
    }

    const actionBodyTemplate = (rowData) => {
        return (
            <div className='flex gap-2'>
                <Button
                    icon="pi pi-check"
                    className="p-button-rounded p-button-success"
                    onClick={() => handleApprove(rowData)}
                />
                <Button
                    icon="pi pi-times"
                    className="p-button-rounded p-button-danger"
                    onClick={() => handlePending(rowData)}
                />
            </div>
        );
    }


    return (<>
        <div className="content md:px-0 md:py-10" >
            <DataTable value={approvals} tableStyle={{ minWidth: '50rem' }}>
                <Column field="nama" sortable header="Nama"></Column>
                <Column field="email" sortable header="Email"></Column>
                <Column field="ipk" sortable header="IPK"></Column>
                <Column field="semester" sortable header="Semester"></Column>
                <Column field="telepon" sortable header="Telepon"></Column>
                <Column field="jenis_beasiswa" sortable header="Beasiswa"></Column>
                <Column header="Berkas" body={imageBodyTemplate}></Column>
                <Column header="Status" body={statusTemplate} style={{ minWidth: '8rem' }}></Column>
                <Column body={actionBodyTemplate} exportable={false} style={{ minWidth: '8rem' }}></Column>
            </DataTable>
        </div>


    </>);
}

export default Pengajuan;