import styles from './shortdescription.module.scss'

export default function ShortDescription() {
    return (
        <section className={styles.wrapper}>
            <section className={styles.card}>
                <section className={styles.head}>
                    <h1 className={styles.head_text}>Fullstack Development</h1>
                </section>
                <section className={styles.short_description}>
                    <span className={styles.short_description_text}>
                        End-to-end development of websites and applications from UI/UX design to fully functional web and mobile solutions.
                    </span>
                </section>
                    <section className={styles.order_contact}>
                        <section className={styles.order}>
                            <a href="order" className={styles.order_button_link}>
                                <button className={styles.order_button}>
                                    Order a Project
                                </button>
                            </a>
                        </section>
                        <section className={styles.contact}>
                            <a href="https://t.me//hashfuck" target="_blank" className={styles.contact_button_link}>
                                <button className={styles.contact_button}>
                                    Discuss the project on <span className={styles.contact_tg}>Telegram</span>
                                </button>
                            </a>
                        </section>
                    </section>
                </section>
            </section>
    )
}
