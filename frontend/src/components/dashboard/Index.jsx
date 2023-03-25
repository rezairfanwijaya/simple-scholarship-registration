import BarChart from "../BarChart";

const Index = () => {
    return (<>
        <section>
            <div className="title mb-2 text-dark font-bold w-10/12 lg:text-2xl lg:mt-10 ">
                Statistik
            </div>
            <div className="bg-white p-2 rounded-md">
                <BarChart />
            </div>
        </section>
    </>);
}

export default Index;