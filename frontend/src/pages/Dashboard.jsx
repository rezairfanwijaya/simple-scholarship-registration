import { Outlet } from "react-router-dom";
import SideBar from "../components/SideBar";

const Dashboard = () => {
    return (<>
        <div className="dashboard">
            <div className="main md:flex justify-between gap-11">
                <section className="md:1/3 lg:w-1">
                    <SideBar />
                </section>

                <section className="md:w-4/6 lg:w-10/12 lg:px-24">
                    <main>
                        <Outlet />
                    </main>
                </section>
            </div>
        </div>
    </>);
}

export default Dashboard;