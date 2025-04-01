"use server"
import Link from 'next/link'
import styles from './portfolio.module.scss'

export default async function PortFolio() {
    
    const cardsInfo = [
        {
            name: "Devis",
            photo: "/images/first.jpg",   
        },
        {
            name: "Bike | SHOP",
            photo: "/images/second.jpg",   
        },       
        {
            name: "MarketPlace",
            photo: "/images/third.jpg",   
        },       
        {
            name: "SOCIAL MEDIA | PHOTO",
            photo: "/images/fourth.jpg",   
        },
        {
            name: "Raketka",
            photo: "/images/fifth.jpg",   
        },
        {
            name: "Prbenefit.com",
            photo: "/images/sixth.jpg",   
        },
    ]

    return (
        <section className={styles.wrapper}>
            <section className={styles.porfolio_header_wrapper}>
                <h2 className={styles.portfolio_header}>Portfolio</h2>
            </section>
            <section className={styles.content_wrapper}>
                {cardsInfo.map((element, index) => (
                    <section key={index} className={styles.card}>
                        <section className={styles.card_header}>
                            <p>{element.name}</p>
                        </section>
                        <section className={styles.cards_image}>
                            <img className={styles.card_image} src={element.photo} alt="" />
                        </section>
                    </section>
                ))}
                <Link href="/portfolio" className={styles.portfolio_link}>
                   Portfolio
                </Link>
            </section>
        </section>
    )
}
