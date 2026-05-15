import api from "./axios";

export const getPosts = () => api.get("/api/posts");
export const createPost = (data) => api.post("/api/posts", data);
export const deletePost = (id) => api.delete(`/api/posts/${id}`);
export const toggleLike = (postId) => api.post(`/api/posts/${postId}/like`);
export const getComments = (postId) => api.get(`/api/posts/${postId}/comments`);
export const createComment = (postId, data) =>
  api.post(`/api/posts/${postId}/comments`, data);