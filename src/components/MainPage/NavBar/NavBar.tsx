import Link from 'next/link'; 
import styles from './navbar.module.scss';

export default function NavBar() {
    return (
        <nav className={styles.navbar}>
            <section className={styles.navbarImage}></section>
            <section className={styles.navbarContent}>
                <ul className={styles.navbarContentList}>
                    <li><Link href="/portfolio" passHref><span className={styles.navLink}>Portfolio</span></Link></li>
                    <li><Link href="/about" passHref><span className={styles.navLink}>About Me</span></Link></li>
                    <li><Link href="/technologies" passHref><span className={styles.navLink}>Technologies</span></Link></li>
                    <li><Link href="/contact" passHref><span className={styles.navLink}>Contacts</span></Link></li>
                </ul>
            </section>
            <section className={styles.navbarIcons}></section>
        </nav>
    );
}
