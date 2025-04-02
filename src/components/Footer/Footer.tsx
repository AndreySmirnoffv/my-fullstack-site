import styles from './footer.module.scss'
import Link from 'next/link'
import Image from 'next/image'
import tgIcon from '../../../public/images/tg.png'
import kworkIcon from '../../../public/images/kwork.png'

export default function Footer() {
    return (
        <footer className={styles.wrapper}>
            <section className={styles.footer_content}>
                <ul className={styles.footer_content_list}>
                    <li><Link href="/portfolio" passHref><span className={styles.footer_link}>Portfolio</span></Link></li>
                    <li><Link href="/about" passHref><span className={styles.footer_link}>About Me</span></Link></li>
                    <li><Link href="/technologies" passHref><span className={styles.footer_link}>Technologies</span></Link></li>
                    <li><Link href="/contact" passHref><span className={styles.footer_link}>Contacts</span></Link></li>
                </ul>
            </section>
            <section className={styles.footer_icons}>
                <Link href="https://t.me/asdwebnodejsts" passHref target="_blank">
                    <Image src={tgIcon} alt="Telegram Icon" height={100} width={75} style={{ objectFit: 'contain' }} />
                </Link>
                <Link href="https://kwork.ru/user/nikkykuznetsov" passHref target='_blank'>
                    <Image src={kworkIcon} alt="Kwork Icon" height={50} width={50} className={styles.inverted_icon} />
                </Link>
                <Link href="https://kwork.ru/user/nikkykuznetsov" passHref target='_blank'>
                    <Image src={kworkIcon} alt='Kwork Icon' height={50} width={50}/>
                </Link>
            </section>
        </footer>
    )
}