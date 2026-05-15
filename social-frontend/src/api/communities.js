import api from "./axios";

export const getCommunities = () => api.get("/api/communities/all");
export const getCommunityById = (id) => api.get(`/api/communities/${id}`);
export const getCommunityPosts = (id) => api.get(`/api/communities/${id}/posts`);
export const getCommunityMembers = (id) => api.get(`/api/communities/${id}/members`);
export const createCommunity = (data) => api.post("/api/communities/create", data);
export const joinCommunity = (id) => api.post(`/api/communities/${id}/join`);
export const leaveCommunity = (id) => api.post(`/api/communities/${id}/leave`);
export const createCommunityPost = (id, data) =>
  api.post(`/api/communities/${id}/posts`, data);
export const deleteCommunity = (id) => api.delete(`/api/admin/communities/${id}`);