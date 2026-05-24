import { useState } from "react";
import { toggleLike, getComments, createComment, deletePost } from "../api/posts";
import { useAuth } from "../context/AuthContext";

export default function PostCard({ post, onDelete }) {
  const { user } = useAuth();
  const [likes, setLikes] = useState(post.likes_count || 0);
  const [showComments, setShowComments] = useState(false);
  const [comments, setComments] = useState([]);
  const [commentText, setCommentText] = useState("");

  const handleLike = async () => {
    const res = await toggleLike(post.id);
    setLikes(res.data.like_count ?? likes);
  };

  const handleToggleComments = async () => {
    if (!showComments) {
      const res = await getComments(post.id);
      setComments(res.data || []);
    }
    setShowComments(!showComments);
  };

  const handleComment = async (e) => {
    e.preventDefault();
    if (!commentText.trim()) return;
    const res = await createComment(post.id, { content: commentText });
    setComments((prev) => [...prev, res.data]);
    setCommentText("");
  };

  const handleDelete = async () => {
    if (!window.confirm("Удалить пост?")) return;
    await deletePost(post.id);
    onDelete(post.id);
  };

  const avatarUrl = post.author?.avatar
    ? `http://localhost:8080${post.author.avatar}`
    : null;

  return (
    <div style={styles.card}>
      <div style={styles.header}>
        {avatarUrl ? (
          <img src={avatarUrl} alt="avatar" style={styles.avatar} />
        ) : (
          <div style={styles.avatarPlaceholder}>👤</div>
        )}
        <div>
          <strong>{post.author?.first_name} {post.author?.last_name}</strong>
          <div style={styles.date}>{new Date(post.created_at).toLocaleString("ru")}</div>
        </div>
        {(user?.id === post.author_id || user?.role === "admin") && (
          <button onClick={handleDelete} style={styles.deleteBtn}>🗑️</button>
        )}
      </div>

      <p style={styles.content}>{post.content}</p>
      {post.image && (
        <img src={post.image} alt="post" style={styles.postImage} />
      )}

      <div style={styles.actions}>
        <button onClick={handleLike} style={styles.actionBtn}>❤️ {likes}</button>
        <button onClick={handleToggleComments} style={styles.actionBtn}>
          💬 Комменты
        </button>
      </div>

      {showComments && (
        <div style={styles.comments}>
          {comments.map((c) => (
            <div key={c.id} style={styles.comment}>
              <strong>{c.author?.first_name || "Аноним"}: </strong>
              {c.content}
            </div>
          ))}
          <form onSubmit={handleComment} style={styles.commentForm}>
            <input
              value={commentText}
              onChange={(e) => setCommentText(e.target.value)}
              placeholder="Написать комментарий..."
              style={styles.input}
            />
            <button type="submit" style={styles.submitBtn}>Отправить</button>
          </form>
        </div>
      )}
    </div>
  );
}

const styles = {
  card: { background: "#fff", borderRadius: 10, padding: 16, marginBottom: 16, boxShadow: "0 2px 8px rgba(0,0,0,0.08)" },
  header: { display: "flex", alignItems: "center", gap: 10, marginBottom: 10 },
  avatar: { width: 40, height: 40, borderRadius: "50%", objectFit: "cover" },
  avatarPlaceholder: { width: 40, height: 40, borderRadius: "50%", background: "#eee", display: "flex", alignItems: "center", justifyContent: "center", fontSize: 20 },
  date: { color: "#999", fontSize: 12 },
  deleteBtn: { marginLeft: "auto", background: "none", border: "none", cursor: "pointer", fontSize: 18 },
  content: { margin: "8px 0" },
  postImage: { width: "100%", borderRadius: 8, marginBottom: 8 },
  actions: { display: "flex", gap: 12 },
  actionBtn: { background: "none", border: "1px solid #ddd", borderRadius: 6, padding: "4px 12px", cursor: "pointer" },
  comments: { marginTop: 12, borderTop: "1px solid #eee", paddingTop: 10 },
  comment: { marginBottom: 6, fontSize: 14 },
  commentForm: { display: "flex", gap: 8, marginTop: 8 },
  input: { flex: 1, padding: "6px 10px", border: "1px solid #ddd", borderRadius: 6 },
  submitBtn: { background: "#e94560", color: "#fff", border: "none", borderRadius: 6, padding: "6px 14px", cursor: "pointer" },
};