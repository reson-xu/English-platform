import type { ApiErrorResp } from "@/types/api";

const API_BASE_URL = "http://127.0.0.1:8080";

export async function get<T>(path: string): Promise<T> {
  return request<T>("GET", path);
}

export async function post<T>(path: string, body?: unknown): Promise<T> {
  return request<T>("POST", path, body);
}

export async function put<T>(path: string, body?: unknown): Promise<T> {
  return request<T>("PUT", path, body);
}

export async function del<T>(path: string): Promise<T> {
  return request<T>("DELETE", path);
}

async function request<T>(
  method: string,
  path: string,
  body?: unknown,
): Promise<T> {
  const response = await fetch(`${API_BASE_URL}${path}`, {
    method,
    headers: { "Content-Type": "application/json" },
    body: body ? JSON.stringify(body) : undefined,
  });

  const data = (await response.json()) as T | ApiErrorResp;

  if (!response.ok) {
    throw new Error((data as ApiErrorResp).message);
  }

  return data as T;
}
