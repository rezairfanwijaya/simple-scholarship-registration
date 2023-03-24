import HeroBanner from '../assets/image/hero.png'

const HeroSection = () => {
    return (<>
        {/* hero section small device */}
        <div className="herosection-small-device mt-9 flex flex-col gap-2 md:hidden">
            <div
                className="left flex flex-col"
                data-aos="fade-down"
                data-aos-delay="0"
                data-aos-duration="1300"
            >
                <img
                    src={HeroBanner}
                    alt="Hero Banner"
                />
            </div>

            <div
                className="right flex flex-col gap-5"
                data-aos="fade-up"
                data-aos-delay="0"
                data-aos-duration="1300"
            >
                <div className="title text-dark text-center font-black text-[1.7rem]">
                    Temukan <span className='text-primary'>Beasiswa</span> <br /> Impian <span className='text-primary'>Kalian</span>
                </div>

                <div className="desc text-mute text-xs text-center">
                    Mulai melangkah lebih maju dengan mengambil beassiwa
                    dan nikmati keuntungannya
                </div>

                <div className="button flex mt-2">
                    <a href="#isi-kuesioner" className='bg-primary hover:bg-primary-hover w-screen px-3 py-3 text-[0.7rem] text-white tracking-wider rounded-md text-center'>
                        Lihat Beasiswa
                    </a>
                </div>

            </div>
        </div>

        {/* hero section for non small device  */}
        <div className="herosection-non-small-device hidden md:mt-5 md:flex md:items-stretch md:justify-between lg:px-2 lg:mt-8">
            <div
                className="left flex flex-col gap-5 lg:gap-8 w-1/2 py-8 "
                data-aos="fade-right"
                data-aos-delay="0"
                data-aos-duration="1300"
            >
                <div className="title text-dark text-left font-bold md:text-4xl lg:text-6xl lg:font-extrabold lg:leading-tight">
                    Temukan <span className='text-primary'>Beasiswa</span> <br /> Impian <span className='text-primary'>Kalian</span>
                </div>

                <div className="desc text-mute md:text-xs text-justify md:w-4/5 lg:text-lg">
                    Mulai melangkah lebih maju dengan mengambil beassiwa
                    dan nikmati keuntungannya
                </div>

                <div className="button mt-2">
                    <a href="#isi-kuesioner" className='bg-primary hover:bg-primary-hover w-screen px-7 py-3 md:text-[0.7rem] text-white tracking-wider rounded-md text-center lg:px-9 lg:py-4 lg:text-sm'>
                        Lihat Beasiswa
                    </a>
                </div>

            </div>

            <div className="decoration md:w-36 md:h-36 lg:w-96 lg:h-96 rounded-full bg-primary-decoration absolute top-72 lg:top-80 lg:-left-24 -left-12 blur-2xl lg:blur-3xl -z-10"></div>

            <div
                className="right lg:w-1/2  flex justify-end"
                data-aos="fade-left"
                data-aos-delay="0"
                data-aos-duration="1300"
            >
                <img
                    src={HeroBanner}
                    alt="Hero Banner"
                    className='w-80 lg:w-[33rem]'
                />

            </div>
        </div>

    </>);
}

export default HeroSection;