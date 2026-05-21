import api from "./axios";

export const getCommunities = () => api.get("/api/teacher/communities/all");
export const createCommunity = (data) => api.post("/api/teacher/communities/create", data);

export const getCommunityById = (id) => api.get(`/api/communities/${id}`);
export const getCommunityPosts = (id) => api.get(`/api/communities/${id}/posts`);
export const getCommunityMembers = (id) => api.get(`/api/communities/${id}/members`);
export const joinCommunity = (id) => api.post(`/api/communities/${id}/join`);
export const leaveCommunity = (id) => api.post(`/api/communities/${id}/leave`);
export const createCommunityPost = (id, data) => api.post(`/api/communities/${id}/posts`, data);
export const deleteCommunity = (id) => api.delete(`/api/admin/communities/${id}`);

export const getCommunitySchedule = (id) => api.get(`/api/communities/${id}/schedule`);
export const createSchedule = (communityId, data) =>
  api.post(`/api/teacher/communities/${communityId}/schedule`, data);
export const updateSchedule = (scheduleId, data) =>
  api.put(`/api/teacher/schedule/${scheduleId}/update`, data);
export const deleteSchedule = (scheduleId) =>
  api.delete(`/api/teacher/schedule/${scheduleId}/delete`);