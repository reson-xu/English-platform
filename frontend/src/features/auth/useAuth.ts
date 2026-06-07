import { useState } from "react";
import { authApi } from "@/api";
import type { LoginReq, RegisterReq } from "@/types/auth";

type AuthMode = "login" | "register";

export function useAuth() {
  const [mode, setMode] = useState<AuthMode>("login");
  const [message, setMessage] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);

  const isRegister = mode === "register";

  async function login(input: LoginReq) {
    setMessage("");
    setIsSubmitting(true);
    try {
      const result = await authApi.login(input);
      setMessage(`${result.user.nickname}, welcome back!`);
      return result;
    } catch (error) {
      setMessage(error instanceof Error ? error.message : "Request failed");
    } finally {
      setIsSubmitting(false);
    }
  }

  async function register(input: RegisterReq) {
    setMessage("");
    setIsSubmitting(true);
    try {
      const result = await authApi.register(input);
      setMessage(`${result.user.nickname}, welcome back!`);
      return result;
    } catch (error) {
      setMessage(error instanceof Error ? error.message : "Request failed");
    } finally {
      setIsSubmitting(false);
    }
  }

  function switchMode(nextMode: AuthMode) {
    setMode(nextMode);
    setMessage("");
  }

  return { mode, message, isSubmitting, isRegister, login, register, switchMode };
}
