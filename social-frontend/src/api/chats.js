import api from "./axios";

export const getChats = () => api.get("/api/chats");
export const createChat = (data) => api.post("/api/chats", data);
export const updateChat = (id, data) => api.put(`/api/chats/${id}`, data);
export const deleteChat = (id) => api.delete(`/api/chats/${id}`);

export const getChatMembers = (id) => api.get(`/api/chats/${id}/members`);
export const addChatMember = (id, data) => api.post(`/api/chats/${id}/members`, data);

export const getChatMessages = (id) => api.get(`/api/chats/${id}/messages`);
export const sendMessage = (id, data) => api.post(`/api/chats/${id}/messages`, data);

export const createDM = (data) => api.post("/api/dm", data);