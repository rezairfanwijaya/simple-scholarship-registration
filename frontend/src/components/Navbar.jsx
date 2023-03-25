import { Navbar, Button } from "flowbite-react";
import Logo from "../assets/image/logo.png"

const NavbarUser = () => {
    return (<>

        <Navbar
            fluid={true}
            rounded={true}
            className="lg:px-12"
        >
            <Navbar.Brand href="/">
                <img
                    src={Logo}
                    className="mr-3 h-7 sm:h-9"
                    alt="CDC Logo"
                />

            </Navbar.Brand>
            <div className="flex md:order-2 ">
                <a href="/login" >
                    <Button className="bg-primary hover:bg-primary-hover px-0 py-0 md:px-1 lg:px-2 ">
                        <span className="text-xs tracking-wide lg:text-sm">Masuk</span>
                    </Button>
                </a>
                <Navbar.Toggle />
            </div>
            <Navbar.Collapse>
                <Navbar.Link
                    href="#list-beasiswa"
                >
                    <span className="hover:text-primary tracking-wide">Lihat Beasiswa</span>
                </Navbar.Link>
                <Navbar.Link href="/daftar" className="hover:text-primary">
                    <span className="hover:text-primary tracking-wide">Daftar</span>
                </Navbar.Link>
                <Navbar.Link href="/hasil">
                    <span className="hover:text-primary tracking-wide">Hasil</span>
                </Navbar.Link>
            </Navbar.Collapse>
        </Navbar>


    </>);
}

export default NavbarUser;
