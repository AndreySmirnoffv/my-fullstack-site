import About from "@/components/MainPage/About/About";
import NavBar from "@/components/NavBar/NavBar";
import PortFolio from "@/components/MainPage/Portfolio/Portfolio";
import Technologies from "@/components/MainPage/Technologies/Technologies";
import Footer from "@/components/Footer/Footer";

export default function Home() {
  return (
    <>
      <NavBar/>
      <About/>
      <PortFolio/>
      <Technologies/>
      <Footer/>
    </>

  )
}
