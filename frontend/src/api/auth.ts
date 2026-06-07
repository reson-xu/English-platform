import { post } from "./client";
import type { LoginReq, AuthResp, RegisterReq } from "@/types/auth";

export const authApi = {
  login: (req: LoginReq) => post<AuthResp>("/api/auth/login", req),
  register: (req: RegisterReq) => post<AuthResp>("/api/auth/register", req),
};
