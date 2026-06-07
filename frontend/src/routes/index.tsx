import { createBrowserRouter } from "react-router-dom";
import AuthPage from "@/features/auth";

export const router = createBrowserRouter([
  { path: "/", element: <AuthPage /> },
]);
