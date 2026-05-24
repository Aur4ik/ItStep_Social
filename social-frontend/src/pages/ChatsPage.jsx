// src/pages/ChatsPage.jsx
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { getChats, createChat, addChatMember } from "../api/chats";
import { useAuth } from "../context/AuthContext";

export default function ChatsPage() {
  const [chats, setChats] = useState([]);
  const [name, setName] = useState("");
  const [showForm, setShowForm] = useState(false);
  const [error, setError] = useState("");
  const { user } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    getChats()
      .then((res) => setChats(res.data || []))
      .catch(() => setError("Не удалось загрузить чаты"));
  }, []);

  const handleCreate = async (e) => {
    e.preventDefault();
    if (!name.trim()) return;
    try {
      const res = await createChat({ name });
      await addChatMember(res.data.id, { user_id: user?.id }).catch(() => {});
      setChats((prev) => [...prev, res.data]);
      setName("");
      setShowForm(false);
    } catch {
      setError("Ошибка создания чата");
    }
  };

  return (
    <div style={styles.wrapper}>
      <div style={styles.header}>
        <h2>💬 Чаты</h2>
        <button onClick={() => setShowForm(!showForm)} style={styles.btn}>
          {showForm ? "Отмена" : "+ Создать чат"}
        </button>
      </div>

      {error && <div style={styles.error}>{error}</div>}

      {showForm && (
        <form onSubmit={handleCreate} style={styles.form}>
          <input
            style={styles.input}
            placeholder="Название чата"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
          <button type="submit" style={styles.btn}>Создать</button>
        </form>
      )}

      {chats.length === 0 ? (
        <div style={styles.empty}>Чатов пока нет</div>
      ) : (
        <div style={styles.list}>
          {chats.map((chat) => (
            <div
              key={chat.id}
              style={styles.card}
              onClick={() => navigate(`/chats/${chat.id}`)}
            >
              <div style={styles.avatar}>💬</div>
              <div>
                <strong>{chat.name}</strong>
                <p style={styles.sub}>Нажми чтобы открыть</p>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

const styles = {
  wrapper: { maxWidth: 640, margin: "24px auto", padding: "0 16px" },
  header: { display: "flex", justifyContent: "space-between", alignItems: "center", marginBottom: 20 },
  btn: { background: "#e94560", color: "#fff", border: "none", borderRadius: 8, padding: "8px 18px", cursor: "pointer" },
  form: { background: "#fff", borderRadius: 10, padding: 16, marginBottom: 16, display: "flex", gap: 10, boxShadow: "0 2px 8px rgba(0,0,0,0.08)" },
  input: { flex: 1, padding: "10px 14px", border: "1px solid #ddd", borderRadius: 8, fontSize: 14 },
  error: { color: "#c00", background: "#fee", padding: 10, borderRadius: 8, marginBottom: 12 },
  empty: { textAlign: "center", color: "#999", marginTop: 40 },
  list: { display: "flex", flexDirection: "column", gap: 10 },
  card: { background: "#fff", borderRadius: 10, padding: 16, display: "flex", alignItems: "center", gap: 14, boxShadow: "0 2px 8px rgba(0,0,0,0.08)", cursor: "pointer" },
  avatar: { fontSize: 28, width: 48, height: 48, background: "#f4f6fb", borderRadius: "50%", display: "flex", alignItems: "center", justifyContent: "center" },
  sub: { color: "#999", fontSize: 12, margin: "2px 0 0" },
};