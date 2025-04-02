import React from 'react';
import styles from './skillcard.module.scss';

interface SkillCardProps {
  logo: string; 
  technology: string;
  skillLevel: number;
  color: string;
}

const SkillCard: React.FC<SkillCardProps> = ({ logo, technology, skillLevel, color }) => {
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
            style={{ width: `${skillLevel}%`, backgroundColor: color }}
          />
        </div>
      </div>
    </div>
  );
};

export default SkillCard;
