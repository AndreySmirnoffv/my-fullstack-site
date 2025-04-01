import React from 'react';
import styles from './skillcard.module.scss';

const SkillCard = ({ logo, technology, skillLevel, color }) => {
  return (
    <div className={styles.card}>
      <div className={styles.cardHeader}>
        <img src={logo} alt={technology} className={styles.logo} />
        <span className={styles.technology}>{technology}</span>
      </div>
      <div className={styles.progressWrapper}>
        <span className={styles.skillLabel}>Skill</span>
        <div className={styles.skillBar}>
          <div
            className={styles.skillFill}
            style={{ width: `${skillLevel}%`, backgroundColor: color }} // Используем переданный цвет
          />
        </div>
      </div>
    </div>
  );
};

export default SkillCard;
