import { useEffect, useState } from "react";
import { getPosts, createPost } from "../api/posts";
import PostCard from "../components/PostCard";

export default function FeedPage() {
  const [posts, setPosts] = useState([]);
  const [content, setContent] = useState("");
  const [error, setError] = useState("");

  useEffect(() => {
  getPosts()
    .then((res) => setPosts(res.data || []))
    .catch(() => setError("Не удалось загрузить посты"));
}, []);

  const handleCreate = async (e) => {
    e.preventDefault();
    if (!content.trim()) return;
    try {
      const res = await createPost({ content });
      setPosts((prev) => [res.data, ...prev]);
      setContent("");
    } catch {
      setError("Не удалось создать пост");
    }
  };

  const handleDelete = (id) => {
    setPosts((prev) => prev.filter((p) => p.id !== id));
  };

  return (
    <div style={styles.wrapper}>
      <form onSubmit={handleCreate} style={styles.createForm}>
        <textarea
          value={content}
          onChange={(e) => setContent(e.target.value)}
          placeholder="Что у вас нового?"
          style={styles.textarea}
        />
        {error && <div style={styles.error}>{error}</div>}
        <button type="submit" style={styles.btn}>Опубликовать</button>
      </form>
      {posts.map((post) => (
        <PostCard key={post.id} post={post} onDelete={handleDelete} />
      ))}
    </div>
  );
}

const styles = {
  wrapper: { maxWidth: 640, margin: "24px auto", padding: "0 16px" },
  createForm: { background: "#fff", borderRadius: 10, padding: 16, marginBottom: 20, boxShadow: "0 2px 8px rgba(0,0,0,0.08)", display: "flex", flexDirection: "column", gap: 10 },
  textarea: { padding: 10, border: "1px solid #ddd", borderRadius: 8, resize: "vertical", minHeight: 80, fontSize: 15 },
  btn: { alignSelf: "flex-end", background: "#e94560", color: "#fff", border: "none", borderRadius: 8, padding: "8px 20px", cursor: "pointer" },
  error: { color: "#c00" },
};