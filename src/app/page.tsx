import About from "@/components/MainPage/About/About";
import NavBar from "@/components/MainPage/NavBar/NavBar";
import PortFolio from "@/components/MainPage/Portfolio/Portfolio";
import ShortDescription from "@/components/MainPage/ShortDescription/ShortDescription";
import Technologies from "@/components/MainPage/Technologies/Technologies";

export default function Home() {
  return (
    <>
      <NavBar/>
      {/* <ShortDescription/> */}
      <About/>
      <PortFolio/>
      <Technologies/>
    </>

  )
}
