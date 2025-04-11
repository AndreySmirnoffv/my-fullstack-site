"use client"
import styles from './styles.module.scss';
import { useState, useEffect } from 'react';
import Link from 'next/link';
import { usePathname, useRouter } from 'next/navigation';
import axios from 'axios';

export default function AuthForm() {
  const [username, setUsername] = useState<string>('');
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const path = usePathname();
  const router = useRouter();

  const [isLogin, setIsLogin] = useState(false);

  useEffect(() => {
    if (path.includes('login')) {
      setIsLogin(true);
    } else {
      setIsLogin(false);
    }
  }, [path]);

  const handleAuth = async (e: React.FormEvent) => {
    e.preventDefault();
    
    const authData = isLogin ? { email, password } : { username, email, password };
    const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8000";

    axios.post('http://localhost:8080/api/auth/login', {
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        username,
        password
      }),
    })
    .then(data => console.log(data))
    .catch(error => console.error('Error:', error));
    
  };

  return (
    <section className={styles.wrapper}>
      <h1 className={styles.title}>{isLogin ? "Login" : "Register"}</h1>
      <form className={styles.inputContainer} onSubmit={handleAuth}>
        {!isLogin && (
          <input
            className={styles.usernameInput}
            type="text"
            placeholder="Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
        )}
        <input
          className={styles.emailInput}
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
        <input
          className={styles.passwordInput}
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <button className={styles.authButton} type="submit">
          {isLogin ? "Login" : "Register"}
        </button>
        <p className={styles.switchText}>
          {isLogin ? "Don't have an account?" : "Already have an account?"}{' '}
          <Link href={isLogin ? "/auth/register" : "/auth/login"}>
            {isLogin ? "Register" : "Login"}
          </Link>
        </p>
      </form>
    </section>
  );
}
