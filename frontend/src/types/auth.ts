// ── Request types ───────────────────────────────────────────────

export type LoginReq = {
  email: string;
  password: string;
};

export type RegisterReq = {
  email: string;
  password: string;
  nickname: string;
};

// ── Response types ──────────────────────────────────────────────

export type AuthUserResp = {
  id: string;
  email: string;
  nickname: string;
  role: string;
  status: string;
};

export type AuthResp = {
  user: AuthUserResp;
  token: string;
};
