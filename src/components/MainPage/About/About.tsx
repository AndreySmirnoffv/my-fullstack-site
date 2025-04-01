import styles from './about.module.scss'

export default async function About(){
    return (
        <section className={styles.wrapper}>
            <section className={styles.about_header_wrapper}>
                <h2 className={styles.about_header}>About Me</h2>
            </section>
        </section>
    )
}