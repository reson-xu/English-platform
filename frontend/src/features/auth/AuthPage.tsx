import { FormEvent, useState } from "react";
import { useAuth } from "./useAuth";
import "./AuthPage.css";

export default function AuthPage() {
  const { mode, message, isSubmitting, isRegister, login, register, switchMode } = useAuth();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [nickname, setNickname] = useState("");

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    if (isRegister) {
      await register({ email, password, nickname });
    } else {
      await login({ email, password });
    }
  }

  return (
    <main className="auth-page">
      <section className="auth-card" aria-labelledby="auth-title">
        <div className="auth-visual" aria-hidden="true">
          <img
            src="/images/login.png"
            alt="Learning Now"
            className="auth-visual-img"
          />
        </div>

        <div className="auth-panel">
          <h1 id="auth-title" className="auth-title">
            {isRegister ? "Create Account" : "Welcome Back"}
          </h1>

          <form className="auth-form" onSubmit={handleSubmit}>
            <label
              className={`auth-field${isRegister ? "" : " auth-field--hidden"}`}
              htmlFor="nickname"
            >
              <span>Nickname</span>
              <input
                id="nickname"
                name="nickname"
                type="text"
                value={nickname}
                onChange={(event) => setNickname(event.target.value)}
                autoComplete="nickname"
                required={isRegister}
                disabled={!isRegister}
                tabIndex={isRegister ? 0 : -1}
              />
            </label>

            <label className="auth-field" htmlFor="email">
              <span>Email</span>
              <input
                id="email"
                name="email"
                type="email"
                value={email}
                onChange={(event) => setEmail(event.target.value)}
                autoComplete="email"
                required
              />
            </label>

            <label className="auth-field" htmlFor="password">
              <span>Password</span>
              <input
                id="password"
                name="password"
                type="password"
                value={password}
                onChange={(event) => setPassword(event.target.value)}
                autoComplete={isRegister ? "new-password" : "current-password"}
                minLength={8}
                required
              />
            </label>

            <button className="auth-submit" type="submit" disabled={isSubmitting}>
              {isSubmitting ? "Processing…" : isRegister ? "Sign Up" : "Log In"}
            </button>

            {message ? <p className="auth-message">{message}</p> : null}
          </form>

          <div className="auth-switch" role="group" aria-label="Account">
            <span className="auth-switch-hint">
              {isRegister ? "Already have an account?" : "Don't have an account yet?"}
            </span>
            {isRegister ? (
              <button
                className="auth-switch-btn"
                type="button"
                onClick={() => switchMode("login")}
              >
                Log In
              </button>
            ) : (
              <button
                className="auth-switch-btn"
                type="button"
                onClick={() => switchMode("register")}
              >
                Sign Up
              </button>
            )}
          </div>
        </div>
      </section>
    </main>
  );
}
