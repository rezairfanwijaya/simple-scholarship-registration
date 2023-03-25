import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';
import { Bar } from 'react-chartjs-2';
import React, { useState, useEffect } from 'react';
import Swal from 'sweetalert2'
import withReactContent from 'sweetalert2-react-content'

ChartJS.register(
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend
);

export const options = {
    responsive: true,
    plugins: {
        legend: {
            position: 'top',
            display: false,
        },
        title: {
            display: false,
            text: 'Grafik Status Approval',
        },
    },
};

const labels = [
    'Menunggu Konfirmasi', 
    'Telah dikonfirmasi', 
];

const BarChart = () => {
    // state
    const [approve, setapprove] = useState(0)
    const [pending, setpending] = useState(0)
    const MySwal = withReactContent(Swal)
    useEffect(() => {
        fetch(`http://localhost:8989/admin/status/show`, {
            method: 'GET'
        }).then(res => { return res.json() })
            .then(data => {
                if (data.meta.code === 200) {
                    setapprove(data.data.approve)
                    setpending(data.data.pending)
                } else {
                    MySwal.fire({
                        title: 'Terjadi kesalahan',
                        icon: 'error',
                        text: data.data
                    })
                }
            })
    }, []) 

    // data pada table
    const data = {
        labels,
        datasets: [
            {
                data: [pending, approve],
                backgroundColor: [
                    'rgba(255, 214, 0, 0.5)',
                    'rgba(0, 115, 255, 0.5)',
                ],
                borderWidth: 0,
            }
        ],
    };
    return (<>
        <Bar options={options} data={data} />
    </>);
}

export default BarChart;