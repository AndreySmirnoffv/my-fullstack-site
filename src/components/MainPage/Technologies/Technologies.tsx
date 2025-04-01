import styles from './technologies.module.scss'
import SkillCard from '../SkillCard/SkillCard'

export default function Technologies() {
  const skills = [
    {
      logo: '/images/html.png',
      technology: 'HTML',
      skillLevel: 90,
      color: '#FF5722',
    },
    {
      logo: '/images/css.png',
      technology: 'CSS',
      skillLevel: 80,
      color: '#2196F3',
    },
    {
      logo: '/images/js.png',
      technology: 'JavaScript',
      skillLevel: 75,
      color: '#FFEB3B',
    },
    {
      logo: '/images/next.png',
      technology: 'Next',
      skillLevel: 70,
      color: '#61DAFB',
    },
    {
        logo: "/images/ts.png",
        technology: "Typescript",
        skillLevel: 85,
        color: "#0541ff"
    },
    {
        logo: "/images/go.png",
        technology: "Golang",
        skillLevel: 30,
        color: "#3e78ff"
    },
    {
        logo: "/images/swift.png",
        technology: "Swift",
        skillLevel: 40,
        color: "#FFA500"
    },    
    {
        logo: "/images/kotlin.png",
        technology: "Kotlin",
        skillLevel: 40,
        color: "#FFA500"
    },
    {
        logo: "/images/devops.png",
        technology: "DEVOPS",
        skillLevel: 40,
        color: "#FFA500"
    },
  ];

  return (
    <div className={styles.container}>
      <h1 className={styles.technologies_header}>Technologies</h1>
      <div className={styles.card_container}>
        {skills.map((skill, index) => (
          <SkillCard
            key={index}
            logo={skill.logo}
            technology={skill.technology}
            skillLevel={skill.skillLevel}
            color={skill.color}
          />
        ))}
      </div>
    </div>
  );
}
