import api from "./axios";

export const register = (data) => api.post("/auth/register", data);
export const login = (data) => api.post("/auth/login", data);
export const getMe = () => api.get("/api/me");
export const uploadAvatar = (formData) =>
  api.post("/api/avatar", formData, {
    headers: { "Content-Type": "multipart/form-data" },
  });