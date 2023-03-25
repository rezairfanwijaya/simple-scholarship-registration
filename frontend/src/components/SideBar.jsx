// import library
import { useNavigate } from "react-router-dom";
import Swal from 'sweetalert2'
import withReactContent from 'sweetalert2-react-content'

const SideBar = () => {
    const MySwal = withReactContent(Swal)
    const navigate = useNavigate();

    const handleLogout = () => {
        MySwal.fire({
            title: 'Apakah anda yakin?',
            text: "Lanjutkan untuk keluar",
            icon: 'warning',
            showCancelButton: true,
            confirmButtonColor: '#3085d6',
            cancelButtonColor: '#d33',
            cancelButtonText: 'Batal',
            confirmButtonText: 'Keluar'
        }).then((result) => {
            if (result.isConfirmed) {
                localStorage.clear('token')
                navigate('/')
            }
        })
    }

    return (<>
        {/* small screen */}
        <div
            className="nav md:hidden fixed bottom-5 ml-auto mr-auto left-5 right-5 py-4 px-7 bg-white drop-shadow-md rounded-xl z-50"
        >
            <div className="menus flex gap-0 text-xl justify-center items-center">
                <div className="dashboard p-2 px-4 text-mute">
                    <a href="/dashboard" className="text-[1.3rem]">
                        <ion-icon name="grid-outline"></ion-icon>
                    </a>
                </div>
                <div className="alumni p-2 px-4 text-mute">
                    <a href="/dashboard/pengajuan" className="text-[1.5rem]">
                        <ion-icon name="reader-outline"></ion-icon>
                    </a>
                </div>
                <div className="logout p-2 px-4 text-mute">
                    <button className="text-[1.5rem]" onClick={handleLogout}>
                        <ion-icon name="log-out-outline"></ion-icon>
                    </button>
                </div>
            </div>
        </div>


        {/* non small screen */}
        <div className="menus bg-white hidden md:flex flex-col mt-7 fixed gap-5 p-6 left-5 top-5 drop-shadow-sm rounded-md z-50">
            <div className="group dashboard p-1 px-2 rounded-md hover:bg-primary">
                <a href="/dashboard" className="text-mute flex gap-4 items-start group-hover:text-white">
                    <div className="icon text-2xl">
                        <ion-icon name="grid-outline"></ion-icon>
                    </div>
                    <div className="title text-md text-mute group-hover:text-white">Dashboard</div>
                </a>
            </div>
            <div className="group alumni p-1 px-2 rounded-md hover:bg-primary">
                <a href="/dashboard/pengajuan" className="text-mute flex gap-4 items-start group-hover:text-white">
                    <div className="icon text-2xl">
                        <ion-icon name="school-outline"></ion-icon>
                    </div>
                    <div className="title text-md text-mute group-hover:text-white">Pengajuan</div>
                </a>
            </div>
            <div className="group logout p-1 px-2 rounded-md hover:bg-primary ">
                <button
                    className="text-mute flex gap-4 items-start group-hover:text-white"
                    onClick={handleLogout}
                >
                    <div className="icon text-2xl">
                        <ion-icon name="log-out-outline"></ion-icon>
                    </div>
                    <div className="title text-md text-mute group-hover:text-white">Keluar</div>
                </button>
            </div>
        </div>

    </>);
}

export default SideBar;