import styles from './about.module.scss'

export default async function About(){
    return (
        <section className={styles.wrapper}>
            <section className={styles.about_header_wrapper}>
                <h2 className={styles.about_header}>About Me</h2>
            </section>
            <section className={styles.about_content_wrapper}>
                <p className={styles.about_content_text}>I am a full-stack developer with experience in creating web and mobile applications. I specialize in building modern, user-friendly websites and apps that perform well across different devices. Whether it's developing from scratch, adding features to existing projects, or delivering custom solutions, I can help bring your ideas to life. In addition, I have a strong background in cybersecurity and can implement robust security measures to protect your applications and data. I ensure that your projects are not only functional but also secure, with the latest practices for safeguarding against potential threats. I also offer ongoing support and maintenance to ensure the continued success and security of your projects.</p>
            </section>
        </section>
    )
}
