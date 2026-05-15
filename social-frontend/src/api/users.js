import api from "./axios";

export const updateUserRole = (userId, role) =>
  api.post(`/api/admin/users/${userId}/role`, { role });