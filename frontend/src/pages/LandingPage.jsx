import HeroSection from "../components/HeroSection";
import NavbarUser from "../components/Navbar";

const LandingPage = () => {
    return (<>
        <div className="body mx-5">
            <NavbarUser />
            <HeroSection/>
        </div>
    </>);
}

export default LandingPage;